package singleton

import (
	"fmt"
	"sync"
)

//单例模式保证每次调用的都是同一个对象
//lazy加载，使用sync.Once双重检验加锁保证并发安全

var (
	instance *singleton
	once     sync.Once
)

type singleton struct {
}

func GetInstance() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})
	return instance
}

func (s *singleton) GetString() {
	fmt.Println("i am a singleton")
}
