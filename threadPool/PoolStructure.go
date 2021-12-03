package threadPool

import "fmt"

//线程池中的封装"线程"
type worker struct{
	jobChan chan Job
	endChan chan bool //结束线程通道
}
//线程池中线程的"job",
type Job interface {
	Do() int
}

//运行线程
func(self*worker)RunWorker(){
	go func(){
		for{
			select{
			case job:=<-self.jobChan:
				job.Do()
			case <-self.endChan:
				fmt.Println("worker end!")
				return //结束该worker线程
			}
		}
	}()
}
//结束线程
func(self*worker)EndWorker(){
	self.endChan<-true
}

//创建线程
func NewWorker(jobChan chan Job)worker{
	return worker{
		jobChan: jobChan,
		endChan: make(chan bool),
	}
}