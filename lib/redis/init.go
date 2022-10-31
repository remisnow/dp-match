package redis

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
	"match/lib/file"
	"match/lib/log"
	"time"
)

/*
{
	"default": {
		"addrs":["127.0.0.1:1111","127.0.0.1:2222"],
		"auth":"123",
		"pool_size":5,
		"is_cluster":false
	},
	"test": {
		"addrs":["127.0.0.1:3333","127.0.0.1:4444"],
		"auth":"321",
		"pool_size":15,
		"is_cluster":true
	}
}
*/

const (
	StreamMQMaxLen        = 500000                  //消息队列最大长度
	StreamConsumeMaxCount = 10                      //单次消费的最大长度
	StreamBlockTime       = 5000 * time.Millisecond //阻塞读取的时间小于redis超时时间的一半

	Common  = 0
	Cluster = 1
	PubSub  = 2
)

// Config 配置结构
type Config struct {
	Addrs    []string `json:"addrs"`
	Auth     string   `json:"auth"`
	PoolSize int      `json:"pool_size"`
	ConnType int      `json:"connType"`
}

var clients map[string]radix.Client = make(map[string]radix.Client)
var pubSubConn map[string]radix.PubSubConn = make(map[string]radix.PubSubConn)
var pubSubChan map[string]chan<- radix.PubSubMessage = make(map[string]chan<- radix.PubSubMessage)
var ErrCh chan error = make(chan error, 1)

func GetClient(name string) radix.Client {
	if _, ok := clients[name]; ok {
		return clients[name]
	}
	return nil
}

func GetPubMassageChan(name string) chan<- radix.PubSubMessage {
	if _, ok := pubSubChan[name]; ok {
		return pubSubChan[name]
	}
	return nil
}
func GetPubSubConn(name string) radix.PubSubConn {
	if _, ok := pubSubConn[name]; ok {
		return pubSubConn[name]
	}
	return nil
}

func GetSubMessage(name string, channel string) chan radix.PubSubMessage {
	msgCh := make(chan radix.PubSubMessage)
	if err := GetPubSubConn(name).Subscribe(msgCh, channel); err != nil {
		log.Println(err)
	}
	return msgCh
}

func PubMessage(name string, m *radix.PubSubMessage) {
	GetPubMassageChan(name) <- *m
}

//--------------------------------------------------------------------------------------

// StreamSendMsg 添加消息
func StreamSendMsg(clientName string, streamKey string, msgKey string, msgValue []byte) (strMsgId string, err error) {
	conn := GetClient(clientName)
	//*表示由Redis自己生成消息ID，设置MAXLEN可以保证消息队列的长度不会一直累加
	err = conn.Do(radix.FlatCmd(&strMsgId, "XADD", streamKey, "MAXLEN", "=", StreamMQMaxLen, "*", msgKey, msgValue))
	if err != nil {
		fmt.Println("XADD failed, err: ", err)
		return "", err
	}
	log.Println("PutMsg", strMsgId, string(msgValue))
	//fmt.Println("Reply Msg Id:", strMsgId)
	return strMsgId, nil
}

// CreateConsumerGroup 创建消费者组
func CreateConsumerGroup(clientName string, streamKey string, groupName string, beginMsgId string) (err error) {
	conn := GetClient(clientName)

	var out string
	err = conn.Do(radix.FlatCmd(&out, "XGROUP", "CREATE", streamKey, groupName, beginMsgId, "MKSTREAM"))
	if err != nil {
		fmt.Println("XGROUP CREATE Failed. err:", err)
		//return err
	}
	log.Println("CreateConsumerGroup out = ", out)
	return err
}

func CreateStreamReaderOpts(stream map[string]*radix.StreamEntryID, groupName string, consumerName string, isBlock bool, count int) radix.StreamReaderOpts {
	return radix.StreamReaderOpts{
		Streams:               stream,
		FallbackToUndelivered: true,
		Group:                 groupName,
		Consumer:              consumerName,
		NoAck:                 true,
		Block:                 StreamBlockTime,
		NoBlock:               !isBlock,
		Count:                 count,
	}
}

// CreateStreamBlockByGroupConsumer 组内消息分配操作，组内每个消费者消费多少消息
func CreateStreamBlockByGroupConsumer(clientName string, streamOpts radix.StreamReaderOpts, outChan chan []radix.StreamEntry) {
	conn := GetClient(clientName)
	sReader := radix.NewStreamReader(conn, streamOpts)
	go func() {
		for true {
			_, entries, ok := sReader.Next()
			if ok && len(entries) > 0 {
				outChan <- entries
			}
		}
	}()
}

// --------------------------------------------------------------------------------------

// 在redis 获取率自增的key
func GetINCRValue(clientName string, keyName string) int {
	var res int
	GetClient(clientName).Do(radix.FlatCmd(&res, "incr", keyName))
	return res
}

//--------------------------------------------------------------------------------------

func Initial(filename string) error {
	var cfgs = make(map[string]*Config)
	var e error
	e = file.LoadJsonToObject(filename, &cfgs)
	if e != nil {
		return e
	}

	for k, cfg := range cfgs {
		switch cfg.ConnType {
		case Common:
			var cli radix.Client

			//连接的额外配置
			customConnFunc := func(network, addr string) (radix.Conn, error) {
				return radix.Dial(network, addr,
					radix.DialTimeout(15*time.Second), // 默认超时的时间是10秒
				)
			}
			cli, e = radix.NewPool("tcp", cfg.Addrs[0], cfg.PoolSize, radix.PoolConnFunc(customConnFunc))
			clients[k] = cli
		case Cluster:
			var cli radix.Client
			cli, e = radix.NewCluster(cfg.Addrs)
			clients[k] = cli
		case PubSub:
			stub, stubCh := radix.PubSubStub("tcp", "127.0.0.1:6379", func([]string) interface{} {
				fmt.Println("radix.PubSubStub")
				return nil
			})
			pubSubConn[k] = radix.PubSub(stub)
			pubSubChan[k] = stubCh
		default:
			log.Println("redis initial error type = " + string(cfg.ConnType))
		}
	}

	return nil
}
