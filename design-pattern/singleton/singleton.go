package singleton

import (
	"fmt"
	"sync"
)

//单例模式保证每次调用的都是同一个对象
//lazy加载，使用sync.Once双重检验加锁保证并发安全

//
// Singleton
// @Description: 导出的单例模式接口
//
type Singleton interface {
	foo()
}

var (
	instance singleton
	once     sync.Once
)

type singleton struct {
}

func (s singleton) foo() {
	fmt.Printf("i've been implanted")
}

func GetInstance() Singleton {
	once.Do(func() {
		instance = singleton{}
	})
	return instance
}
