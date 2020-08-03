package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Employee struct {
	employeeName string
	employeeId   string
	employeeAge  int
	employeeAddress string
}

func (emp *Employee) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	 emp.CreateNewEmployee()
	 e := &Employee{
		 employeeName:    "dakkkfl",
		 employeeId:      "fnnsf",
		 employeeAge:     1223,
		 employeeAddress: "fasfdfs",
	 }
	 responseString,err := json.Marshal(&e)
	 if err!=nil{

	 }
	 fmt.Println(e)
	 fmt.Println(responseString)
	switch request.URL.Path {
	case "/Get":
		 io.WriteString(writer,"Bokachoda")
	}
}
func (emp *Employee) CreateNewEmployee(){
	emp = new(Employee)
	emp.employeeAddress = "nkasa"
	emp.employeeAge = 20
	emp.employeeId = "89e892892r"
	emp.employeeName = "Pal"
}

func main(){
	var employee *Employee
	http.ListenAndServe(":8080",employee)
}