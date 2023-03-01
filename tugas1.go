package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Job string `json:"job"`
	Reason string `json:"reason"`
}

func getStudentDataById(id string) (Student, error) {
	var s Student

	for _, student := range students {
		if student.ID == id {
			s = student
		}
	}

	var err error = nil

	if s == (Student{}) {
		err = errors.New("Data of id " + id + " could not be found.")
	}

	return s, err
}

func (studentData Student) showStudentData(){
	fmt.Println("ID:\t", studentData.ID)
	fmt.Println("Name:\t", studentData.Name)
	fmt.Println("Address:", studentData.Address)
	fmt.Println("Job:\t", studentData.Job)
	fmt.Println("Reason:\t", studentData.Reason)
}

var students []Student = []Student{
	{
		ID: "S1",
		Name: "Fayyadh",
		Address: "Keputih Utara",
		Job: "Pelajar",
		Reason: "Jatuh hati pada Golang",
	},
	{
		ID: "S2",
		Name: "Hafizh",
		Address: "Keputih Selatan",
		Job: "Pelukis",
		Reason: "Lelah melukis banting stir ke Golang",
	},
	{
		ID: "S3",
		Name: "Thomas",
		Address: "Pacetan Timur",
		Job: "Komika",
		Reason: "Nyoba-nyoba ternyata seru",
	},
	{
		ID: "S4",
		Name: "Balmond",
		Address: "Land of Dawn",
		Job: "Kuli",
		Reason: "Bukan era gua lagi di Legenda Seluler ",
	},
}

func main() {
	args := os.Args

	studentData, err := getStudentDataById(args[1])

	if err != nil {
		log.Fatal(err)
	}

	studentData.showStudentData()
}