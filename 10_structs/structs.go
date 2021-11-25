package main

func main() {
	type person struct {
		name string
		age  int
	}
	p := person{
		name: "John",
		age:  30,
	}
	p.age = 31
	println(p.age)
}
