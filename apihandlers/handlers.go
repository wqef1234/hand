package apihandlers

import (
	"encoding/json"
	"github.com/course_spec/newApi/apimodels"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var employees apimodels.Employees

func init(){
	employees = apimodels.Employees{
		apimodels.Employee{
			Id:        "1",
			FirstName: "Bob",
			LastName:  "Jack",
		},
		apimodels.Employee{
			Id:        "2",
			FirstName: "Alice",
			LastName:  "Tompson",
		},
		apimodels.Employee{
			Id:        "3",
			FirstName: "George",
			LastName:  "Lighter",
		},
	}
}

func GetEmployees(w http.ResponseWriter,r *http.Request){
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	for _,employee := range employees {
		if employee.Id == id {
			if err := json.NewEncoder(w).Encode(employee); err != nil {
				log.Println("Error getting employee by id",err)
			}
		}
	}


}
func AddEmployee(w http.ResponseWriter,r *http.Request){
	employee := apimodels.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Println("error while parsing POST-body",err)
		return
	}
	log.Println("Post body succeeded parsing")

	employees = append(employees,apimodels.Employee{
		Id:       employee.Id,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
	})
	json.NewEncoder(w).Encode(employees)

}

func UpdateEmployee(w http.ResponseWriter,r *http.Request) {
	employee := apimodels.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Println("error while parsing PUT method")
		return
	}
	var isUpsert = true
	for i, empl := range employees{
		if empl.Id == employee.Id{
			isUpsert = false
			log.Println("Found that Id")
			employees[i].FirstName = employee.FirstName
			employees[i].LastName = employee.LastName
			break
		}
	}

	if isUpsert {
		employees = append(employees,apimodels.Employee{
			Id:        employee.Id,
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
		})
	}

}


func DeleteEmployee(w http.ResponseWriter,r *http.Request) {
	employee := apimodels.Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Println("error while parsing Delete method")
		return
	}

	for i, empl := range employees{
		if empl.Id == employee.Id{
			log.Println("Found that Id")
			employees = append(employees[:i],employees[i+1:]...)

			log.Println(w,"employee with that ID successfully deleted")
			return
		}
	}
	log.Println(w,"employee with that ID does not exist")
}