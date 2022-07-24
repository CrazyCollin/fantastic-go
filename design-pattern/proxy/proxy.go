package proxy

//静态代理，实现相同方法做方法增强，缺点就是修改方法的时候修改代码繁琐

type Subject interface {
	Do() string
}

type RealSubject struct {
}

func (r RealSubject) Do() string {
	return "i am a real subject"
}

type Proxy struct {
	subject RealSubject
}

func (p Proxy) Do() string {
	var res string
	res += "pre action+"
	res += p.subject.Do()
	res += "+after action"
	return res
}
