package commonArchitecture

import (
	"commonArchitecture/threadPool"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct{
	pool threadPool.WorkerPool
	jobQueue chan threadPool.Job
}

//服务器
func(self *Server)Service(){
	self.pool=threadPool.NewThreadPool(10)
	self.pool.Run()
	time.Sleep(1e9) //等候1s 线程池完成创建过程
	defer self.pool.End()
	rpc.Register(self)
	rpc.HandleHTTP()
	l,err:=net.Listen("tcp","localhost:8080")
	if err!=nil{
		fmt.Println("start server failed!")
		return
	}
	go self.recvJob()
	go http.Serve(l,nil)
}

//接收Job
func (self*Server)recvJob(){
	for{
		select {
			case job:=<-self.jobQueue:
				job.Do()
		}
	}
}


type RequestArgs struct{
	Job threadPool.Job
}

type ResponseReply struct{

}
//供客户端调用
func (self *Server)DoJob(args RequestArgs,reply *ResponseReply)error{
	self.jobQueue<-args.Job
	return nil
}