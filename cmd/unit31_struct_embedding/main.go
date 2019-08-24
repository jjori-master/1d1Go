package unit31_struct_embedding

type Person struct {
	name string
	age  int
}

func (p *Person) greeting() string {
	return "Say Hello!!"
}

type Students struct {
	p      Person
	school string
	grade  int
}

type Man struct {
	Person
	penisSize int
}

type Bitch struct {
	Person
	ssagage int
}

func (b *Bitch) greeting() string {
	return "It's none your business"
}
