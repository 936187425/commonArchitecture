package threadPool
//线程池中的封装"线程"
type worker struct{
	jobChan chan job
	endChan chan bool //结束线程通道
}
//线程池中线程的"job",
type job interface {
	Do()
}

//运行线程
func(self*worker)RunWorker(){
	go func(){
		for{
			select{
			case job:=<-self.jobChan:
				job.Do()
			case <-self.endChan:
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
func NewWorker(jobChan chan job)worker{
	return worker{
		jobChan: jobChan,
		endChan: make(chan bool),
	}
}