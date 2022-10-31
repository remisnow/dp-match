package protocol

import (
	"github.com/astaxie/beego/context"
	"github.com/golang/protobuf/proto"
	"match/lib/utils/messageHandler"
)

// ===================
// client - match
// ===================

// MatchResp -----------------------------------------------------------------------------------------------------------

type MatchReqHandle func(ctx *context.Context, req *MatchReq) (rsp *MatchResp)

type MatchReqHandler struct {
	Handle MatchReqHandle
}

func (h *MatchReqHandler) HttpDo(ctx *context.Context, in proto.Message) (out proto.Message) {
	req := in.(*MatchReq)
	out = h.Handle(ctx, req)
	return
}

func InitMatchReqHandler(container *messageHandler.MessageContainer, handle MatchReqHandle) {
	svr := &MatchReqHandler{Handle: handle}
	container.HttpRegister(&MatchReq{}, svr)
}

// ---------------------------------------------------------------------------------------------------------------------

// ===================
// redis - match
// ===================

// UpdateRoomService ---------------------------------------------------------------------------------------------------

type UpdateRoomServiceHandle func(req *UpdateRoomService) (out *UpdateRoomService)

type UpdateRoomServiceHandler struct {
	Handle UpdateRoomServiceHandle
}

func (h *UpdateRoomServiceHandler) Do(in proto.Message) (out proto.Message) {
	req := in.(*UpdateRoomService)
	out = h.Handle(req)
	return
}

func InitUpdateRoomServiceHandler(container *messageHandler.MessageContainer, handle UpdateRoomServiceHandle) {
	svr := &UpdateRoomServiceHandler{Handle: handle}
	container.Register(&UpdateRoomService{}, svr)
}

// UserEnterRoom -------------------------------------------------------------------------------------------------------

type UserEnterRoomHandle func(req *UserEnterRoom) (out *UserEnterRoom)

type UserEnterRoomHandler struct {
	Handle UserEnterRoomHandle
}

func (h *UserEnterRoomHandler) Do(in proto.Message) (out proto.Message) {
	req := in.(*UserEnterRoom)
	out = h.Handle(req)
	return
}

func InitUserEnterRoomHandler(container *messageHandler.MessageContainer, handle UserEnterRoomHandle) {
	svr := &UserEnterRoomHandler{Handle: handle}
	container.Register(&UserEnterRoom{}, svr)
}

// UserLeaveRoom -------------------------------------------------------------------------------------------------------

type UserLeaveRoomHandle func(req *UserLeaveRoom) (out *UserLeaveRoom)

type UserLeaveRoomHandler struct {
	Handle UserLeaveRoomHandle
}

func (h *UserLeaveRoomHandler) Do(in proto.Message) (out proto.Message) {
	req := in.(*UserLeaveRoom)
	out = h.Handle(req)
	return
}

func InitUserLeaveRoomHandler(container *messageHandler.MessageContainer, handle UserLeaveRoomHandle) {
	svr := &UserLeaveRoomHandler{Handle: handle}
	container.Register(&UserLeaveRoom{}, svr)
}

// CreateRoomBack ------------------------------------------------------------------------------------------------------

type CreateRoomBackHandle func(req *CreateRoomBack) (out *CreateRoomBack)

type CreateRoomBackHandler struct {
	Handle CreateRoomBackHandle
}

func (h *CreateRoomBackHandler) Do(in proto.Message) (out proto.Message) {
	req := in.(*CreateRoomBack)
	out = h.Handle(req)
	return
}

func InitCreateRoomBackHandler(container *messageHandler.MessageContainer, handle CreateRoomBackHandle) {
	svr := &CreateRoomBackHandler{Handle: handle}
	container.Register(&CreateRoomBack{}, svr)
}

// ---------------------------------------------------------------------------------------------------------------------

// UpdateRoom ------------------------------------------------------------------------------------------------------

type UpdateRoomHandle func(req *UpdateRoom) (out *UpdateRoom)

type UpdateRoomHandler struct {
	Handle UpdateRoomHandle
}

func (h *UpdateRoomHandler) Do(in proto.Message) (out proto.Message) {
	req := in.(*UpdateRoom)
	out = h.Handle(req)
	return
}

func InitUpdateRoomHandler(container *messageHandler.MessageContainer, handle UpdateRoomHandle) {
	svr := &UpdateRoomHandler{Handle: handle}
	container.Register(&UpdateRoom{}, svr)
}
