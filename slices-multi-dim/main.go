package main

import "fmt"

func main() {
	records := make([][]string, 0)
	student1 := make([]string, 3)
	student1[0] = "Teri"
	student1[1] = "Deewani"
	student1[2] = "100.00"

	records = append(records, student1)

	student2 := make([]string, 3)
	student2[0] = "Laxmi"
	student2[1] = "Ghadse"
	student2[2] = "32"

	records = append(records, student2)

	fmt.Println(records)

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}

	fmt.Println(transactions)

	employee := make([]string, 35)
	employees := make([][]string, 72)
	employee[0] = "Temple"
	employee = append(employee, "Run")
	fmt.Println(employee)
	fmt.Println(employees)
	fmt.Println(employee == nil)
	fmt.Println(employees == nil)

	var contractor []string
	var contractors [][]string
	//contractor[0] = "Sudhee"
	contractor = append(contractor, "Peram")
	fmt.Println(contractor)
	fmt.Println(contractors)
	fmt.Println(contractor == nil)
	fmt.Println(contractors == nil)

	manager := []string{}
	managers := [][]string{}
	//manager[0] = "Sudhee"
	manager = append(manager, "Peram")
	fmt.Println(manager)
	fmt.Println(managers)
	fmt.Println(manager == nil)
	fmt.Println(managers == nil)
}
