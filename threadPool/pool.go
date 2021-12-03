package threadPool

import "fmt"

/**
	实现线程池,代替"消息队列+缓冲"解决高并发的问题
	线程池有分为有界队列的线程池和无界队列的线程池
	采用RPC协议作为客户端和服务器端之间的通信
	每个线程work可以实现各种功能
*/
//有界队列的"线程池"
type WorkerPool struct{
	workSize int
	endChan chan bool //结束线程池信号的通道
	jobQueue chan Job //内部队列,接收来自外部的job:此处利用了chan的缓冲特性,来分配任务
	wokerQueue []worker //线程池中的线程集合
}

func NewThreadPool(worerSize int) WorkerPool{
	return WorkerPool{
		workSize: worerSize,
		jobQueue: make(chan Job,worerSize),
		endChan: make(chan bool),
		wokerQueue: make([]worker,0,worerSize),
	}
}

//运行线程池
func (self *WorkerPool)Run(){
	go func(){
		//根据线程池中workerSize开启线程
		for i:=0;i<self.workSize;i++{
			fmt.Println("This is the thread: ",i," created!")
			worker:=NewWorker(self.jobQueue)
			self.wokerQueue=append(self.wokerQueue,worker)
			worker.RunWorker()//运行常存协程
		}

		//等候结束run
		for{
			select {
			case <-self.endChan:
				return
			}
		}
	}()

}

//结束线程池
func (self *WorkerPool)End(){
	//结束线程池中的所有线程
	for ix,_:=range self.wokerQueue{ //切片的遍历是会重新创建一个空间
		self.wokerQueue[ix].EndWorker()
	}
	//结束workerPool
	self.endChan<-true
}

//接收来自外部的job数据
func (self*WorkerPool)AddJob(job Job){
	self.jobQueue<-job
}