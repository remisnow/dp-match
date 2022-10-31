package utils

type ServiceGId interface {
	GetGId() string
}
type Obj interface {
}

type EvalFunc func(Obj) Obj

// BuildLazyEvaluator
/*惰性生产器的工厂函数
//例 获取一个从0累加的整数接口：
定义
func BuildIntEvaluator() func() int {
	refunc := utils.BuildLazyEvaluator(func(obj utils.Obj) utils.Obj {
		return obj.(int) + 1
	}, 0)
	return func() int {
		return refunc().(int)
	}
}
调用
	GetIntValue := BuildIntEvaluator()
	GetIntValue()
*/

func buildLazyEvaluator(evalFunc EvalFunc, initData Obj) func() Obj {
	reChan := make(chan Obj)
	var result = initData
	loopFunc := func() {
		for true {
			result = evalFunc(result)
			reChan <- result
		}
	}
	reFunc := func() Obj {
		return <-reChan
	}
	go loopFunc()
	return reFunc
}

func GetIntEvaluator() func() int {
	reFunc := buildLazyEvaluator(func(obj Obj) Obj {
		return obj.(int) + 1
	}, 0)
	return func() int {
		return reFunc().(int)
	}
}
