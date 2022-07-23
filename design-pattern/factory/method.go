package factory

import "fmt"

//根据简单工厂模式改进而来，产品实现工厂的生产方法，
//这样就将生产产品的责任交给子类，实现代码的解耦，
//缺点就是对于每次新增一个产品都要新增代码，略显繁琐

type AnimalFactory interface {
	Create() Animal
}

type Animal interface {
	Call()
}

type Cat struct{}

func (c Cat) Create() Animal {
	return Cat{}
}

func (c Cat) Call() {
	fmt.Printf("i am a cat")
}

type Dog struct{}

func (d Dog) Create() Animal {
	return Dog{}
}

func (d Dog) Call() {
	fmt.Printf("i am a dog")
}

func Create(animal AnimalFactory) Animal {
	return animal.Create()
}
