package test

import (
	threadPool2 "commonArchitecture/threadPool"
	"fmt"
	"testing"
	"time"
)

//测试任务
type Add struct{
	FirstOp int
	SecondOp int
	Result int
}
func(self *Add)Do()int{
	self.Result=self.FirstOp+self.SecondOp
	return 0
}
/**
	测试线程池
 */
func TestThreadPool(t *testing.T){
	threadPool:=threadPool2.NewThreadPool(3)//容量为100的线程池
	threadPool.Run()
	add:=&Add{FirstOp: 1,SecondOp: 2}
	add1:=&Add{FirstOp: 1,SecondOp: 2}
	add3:=&Add{FirstOp: 1,SecondOp: 2}
	add4:=&Add{FirstOp: 1,SecondOp: 2}
	threadPool.AddJob(add)
	threadPool.AddJob(add1)
	threadPool.AddJob(add3)
	threadPool.AddJob(add4)
	time.Sleep(1e8) //如果没有sleep,那么可能还没完成done就已经threadPool.End()
	fmt.Println(add.Result)
	fmt.Println(add1.Result)
	fmt.Println(add3.Result)
	fmt.Println(add4.Result)
	threadPool.End()
}