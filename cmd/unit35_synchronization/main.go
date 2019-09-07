package unit35_synchronization

type Hello struct {
	messages []string
}

func (hello *Hello) sayHello() {
	hello.messages = append(hello.messages, "hello world")
}
