package factory

import "fmt"

//抽象工厂模式将建造不同产品的方法抽象出来，对于不同的产品，
//有着不同的接口，适合建造同族产品

type CreateFactory interface {
	CreateProduct() Product
}

type Product interface {
	Describe()
}

type Car struct {
	Type string
}

func (c Car) Describe() {
	fmt.Printf("i am a car,my type is %s", c.Type)
}

type CarFactory struct{}

func (c CarFactory) CreateProduct() Product {
	return Car{Type: "bmw"}
}
