package messageHandler

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/golang/protobuf/proto"
	"match/lib/log"
	"match/lib/redis"
	pc "match/lib/utils/protocol"
	"strconv"
	"sync"
)

var messageContainer *MessageContainer

type HttpHandler interface {
	HttpDo(ctx *context.Context, in proto.Message) (out proto.Message)
}

func NewHttpHandler(reqMessage proto.Message, handler HttpHandler) *messageHttpHandler {
	return &messageHttpHandler{
		ReqMessage: reqMessage,
		Handler:    handler,
	}
}

type messageHttpHandler struct {
	ReqMessage proto.Message
	Handler    HttpHandler
}

func (mh *messageHttpHandler) HttpDo(ctx *context.Context, date []byte) (out proto.Message, err error) {
	msg := proto.Clone(mh.ReqMessage)
	err = proto.Unmarshal(date, msg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal ReqCall failed, %v", err)
	}

	log.WithField("msg", msg.String()).Debug("recv")
	return mh.Handler.HttpDo(ctx, msg), nil
}

type Handler interface {
	Do(in proto.Message) (out proto.Message)
}

func NewHandler(reqMessage proto.Message, handler Handler) *messageHandler {
	return &messageHandler{
		ReqMessage: reqMessage,
		Handler:    handler,
	}
}

type messageHandler struct {
	ReqMessage proto.Message
	Handler    Handler
}

func (mh *messageHandler) Do(date []byte) (out proto.Message, err error) {
	msg := proto.Clone(mh.ReqMessage)
	err = proto.Unmarshal(date, msg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal ReqCall failed, %v", err)
	}

	log.WithField("msg", msg.String()).Debug("recv")
	return mh.Handler.Do(msg), nil
}

func InitMessageContainer() {
	messageContainer = &MessageContainer{}
	messageContainer.initial()

}
func GetMessageContainer() *MessageContainer {
	return messageContainer

}

type MessageContainer struct {
	Handlers     map[int32]*messageHandler
	HttpHandlers map[int32]*messageHttpHandler
	lock         sync.RWMutex
}

func (mc *MessageContainer) initial() {
	mc.Handlers = make(map[int32]*messageHandler)
	mc.HttpHandlers = make(map[int32]*messageHttpHandler)
}

func (mc *MessageContainer) HttpRegister(message proto.Message, handler HttpHandler) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	key := pc.MessageId(message)
	_, ok := mc.HttpHandlers[key]
	if ok {
		return errors.New("the key already exists, cannot be registered again")
	}

	mc.HttpHandlers[key] = NewHttpHandler(message, handler)
	log.WithField("key", key).Info("register rpc")

	return nil
}

func (mc *MessageContainer) Register(message proto.Message, handler Handler) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	key := pc.MessageId(message)
	_, ok := mc.Handlers[key]
	if ok {
		return errors.New("the key already exists, cannot be registered again")
	}

	mc.Handlers[key] = NewHandler(message, handler)
	log.WithField("key", key).Info("register rpc")

	return nil
}

func (mc *MessageContainer) Unregister(message proto.Message) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	key := pc.MessageId(message)

	_, ok := mc.Handlers[key]
	if ok {
		delete(mc.Handlers, key)
	}
	_, ok1 := mc.HttpHandlers[key]
	if ok1 {
		delete(mc.HttpHandlers, key)
	}
	log.WithField("key", key).Info("unregister rpc")

	return nil
}

// HttpCall http消息路由
func (mc *MessageContainer) HttpCall(ctx *context.Context, messageId int32, buffer []byte) (int32, []byte, int32, error) {
	log.WithFields(map[string]interface{}{
		"Key": messageId,
	}).Info("ReqCall")

	mc.lock.RLock()
	srv, ok := mc.HttpHandlers[messageId]
	mc.lock.RUnlock()
	if !ok {
		return 0, nil, 0, nil
	}

	out, err := srv.HttpDo(ctx, buffer)
	if err != nil {
		return 0, nil, 0, err
	}

	var data []byte
	var outMessageId int32
	var dataLen = 0
	if out != nil {
		data, err = proto.Marshal(out)
		if err != nil {
			return 0, nil, 0, err
		}
		dataLen = len(data)
		outMessageId = pc.MessageId(out)
		log.WithField("msg", out.String()).Debug("send msg")
	}

	return outMessageId, data, int32(dataLen), nil
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToInt 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

// RedisCallByte 收到房间服消息
func (mc *MessageContainer) RedisCallByte(buffer []byte) ([]byte, error) {
	return mc.RedisCall(int32(BytesToInt(buffer[:4])), buffer[4:])
}

// RedisCall 收到房间服消息
func (mc *MessageContainer) RedisCall(messageId int32, dataBuffer []byte) ([]byte, error) {
	log.WithFields(map[string]interface{}{
		"Key": messageId,
	}).Info("ReqCall")

	mc.lock.RLock()
	srv, ok := mc.Handlers[messageId]
	mc.lock.RUnlock()
	if !ok {
		return nil, nil
	}
	_, err := srv.Do(dataBuffer)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// CallRedis  给redis消息队列发消息
func (mc *MessageContainer) CallRedis(clientName string, streamName string, message proto.Message) error {
	mID := strconv.Itoa(int(pc.MessageId(message)))
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	_, err = redis.StreamSendMsg(clientName, streamName, mID, data)
	return err
}
