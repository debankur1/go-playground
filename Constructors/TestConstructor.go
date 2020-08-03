package Constructors

type employee struct {
	Name string
	Age  string
}

func NewEmployee(name string,age string) *employee{
	emp := new(employee)
	emp.Name = name
	emp.Age  = age
	return emp
}