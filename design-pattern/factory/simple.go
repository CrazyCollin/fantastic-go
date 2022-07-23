package factory

//简单工厂模式，直接通过工厂方法，在内部根据类别生产不同的产品，
//缺点就是对于需要新增产品就要修改源代码，不符合开闭原则

type Api interface {
	Say(name string) string
}

func NewApi(key int) Api {
	switch key {
	case 1:
		return ImplApi01{}
	case 2:
		return ImplApi02{}
	default:
		return nil
	}
}

type ImplApi01 struct{}

func (i ImplApi01) Say(name string) string {
	return "01:" + name
}

type ImplApi02 struct{}

func (i ImplApi02) Say(name string) string {
	return "02:" + name
}
