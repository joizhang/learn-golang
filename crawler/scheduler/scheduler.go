package scheduler

import "imooc.com/joizhang/learn-golang/crawler/types"

type Scheduler interface {
	Submit(types.Request)
	ConfigureMasterWorkerChan(chan types.Request)
}
