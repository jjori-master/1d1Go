package unit36_reflection

type Data struct {
	a, b int
}

//noinspection ALL
type Person struct {
	name string `tag1:"이름"tag2:"Name"`
	age  int    `tag1:"나이"tag2:"Age"`
}
