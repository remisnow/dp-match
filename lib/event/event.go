package event

import "match/lib/log"

type _eventBus map[uint32]func(interface{})
type Event struct {
	_events map[string]_eventBus
}

func (e *Event) WatchEvent(topicId uint32, eventName string, callback func(interface{})) {

	if e._events == nil {
		e._events = make(map[string]_eventBus)
	}
	curEventBus := e._events[eventName] // 获取总事件列表中某一个事件的列表
	if curEventBus == nil {
		curEventBus = make(_eventBus)
	}
	curEventBus[topicId] = callback    // 某一个事件的列表添加新事件
	e._events[eventName] = curEventBus // 把某一个事件列表放回到总事件列表中

}

func (e *Event) CallEvent(eventName string, param interface{}) {
	curEventBus := e._events[eventName] // 获取总事件列表中某一个事件的列表
	num := 0
	for _, event := range curEventBus { // 迭代事件列表，然后调用
		event(param)
		num++
	}
	log.Println("CallEvent = ", num)
}

func (e *Event) UnWatchEvent(topicId uint32, eventName string) {
	if e._events == nil {
		return
	}
	curEventBus := e._events[eventName] // 获取总事件列表中某一个事件的列表
	if curEventBus == nil {
		return
	}
	delete(curEventBus, topicId)
}

/*

var eventList = make(map[string][]func(interface{}), 10) // 所有事件的列表
type event struct {
}

func Event() *event { // 实例化对象
	return &event{}
}

func (e event) RegisterEvent(name string, callback func(interface{})) {
	oneEventList := eventList[name]               // 获取总事件列表中某一个事件的列表
	oneEventList = append(oneEventList, callback) // 某一个事件的列表添加新事件
	eventList[name] = oneEventList                // 把某一个事件列表放回到总事件列表中
}

func (e event) CallEvent(name string, param interface{}) {
	oneEventList := eventList[name]      // 根据名字获取一个事件列表
	for _, event := range oneEventList { // 迭代事件列表，然后调用
		event(param)
	}
	delete(eventList, name)
}

/*
func networkRegister(i interface{}) {
	user, ok := i.(string) // 取出接口中的数据
	if ok {
		fmt.Println("user: ", user)
	}
}

func handleNetwork(i interface{}) {
	par := i.(map[string]interface{}) // 取出接口中的数据
	if id, ok := par["id"].([]int); ok {
		fmt.Println("handle id: ", id)
	}
	if user, ok := par["user"].([]string); ok {
		fmt.Println("handle user: ", user)
	}
}
func main() {
	eventObj := eventClass()                         // 调用实例化对象函数，并获取对象
	eventObj.registerEvent("net_1", networkRegister) // 注册一个事件
	eventObj.registerEvent("net_2", networkRegister) // 再次注册该事件
	eventObj.registerEvent("net_3", handleNetwork)   // 注册另一个处理事件
	eventObj.callEvent("net_1", "zhong")             // 调用事件
	eventObj.callEvent("net_2", "xiaohang")
	// 传递多种类型数据的参数
	pa := make(map[string]interface{})
	pa["id"] = []int{18, 22}
	pa["user"] = []string{"zhong", "hang"}
	eventObj.callEvent("net_3", pa)
}*/
