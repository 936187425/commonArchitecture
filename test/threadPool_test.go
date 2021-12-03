package test

import (
	threadPool2 "commonArchitecture/threadPool"
	"testing"
)

/**
	测试线程池
 */
func TestTreadPool(t *testing.T){
	threadPool:=threadPool2.NewThreadPool(100)//容量为100的线程池
	threadPool.Run()
	threadPool.AddJob()//
	threadPool.End()
}