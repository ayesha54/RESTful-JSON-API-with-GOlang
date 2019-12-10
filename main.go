package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type student struct {
	ID         int    `json:"ID"`
	Name       string `json:"Name"`
	Enrollment string `json:"Enrollment"`
}
type allStudents []student

var students = allStudents{
	{
		ID:         1,
		Name:       "Ayesha",
		Enrollment: "2015/COMP/BSCS/18575",
	},
	{
		ID:         2,
		Name:       "Faiza",
		Enrollment: "2015/comp/bscs/1367674",
	},
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the student name and enrollment only in order to update")
	}
	json.Unmarshal(reqBody, &newStudent)
	newStudent.ID = students[len(students)-1].ID + 1
	students = append(students, newStudent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)

}
func getOneStudent(w http.ResponseWriter, r *http.Request) {
	studentID := mux.Vars(r)["id"]
	for _, singleStudent := range students {
		if strconv.Itoa(singleStudent.ID) == studentID {
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}
func getAllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(students)

}
func updateStudent(w http.ResponseWriter, r *http.Request) {
	studentID := mux.Vars(r)["id"]
	var updatedStudent student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedStudent)
	for i, singleStudent := range students {
		if strconv.Itoa(singleStudent.ID) == studentID {
			singleStudent.Name = updatedStudent.Name
			singleStudent.Enrollment = updatedStudent.Enrollment
			students[i] = singleStudent
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	studentID := mux.Vars(r)["id"]
	for i, singleStudent := range students {
		if strconv.Itoa(singleStudent.ID) == studentID {
			students = append(students[:i], students[i+1:]...)
			fmt.Fprintf(w, "The student with ID %v has been deleted successfully", studentID)
		}
	}
}
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/student", createStudent).Methods("POST")
	router.HandleFunc("/students/{id:[0-9]+}", getOneStudent).Methods("GET")
	router.HandleFunc("/students", getAllStudents).Methods("GET")
	router.HandleFunc("/students/{id}", updateStudent).Methods("PATCH")
	router.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
