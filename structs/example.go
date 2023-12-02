package structs

type Address struct {
	City  string
	State string
}

type Person struct {
	Name string
	Age  int
	Address
}
