package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func IsCourseExists(courseName, newCourseName string) bool {
	return strings.EqualFold(courseName, newCourseName)
}

func main() {
	mux := mux.NewRouter()

	courses = append(courses, Course{CourseId: "2", CourseName: "React", Price: 199, Author: &Author{FullName: "HiteshChoudhary", Website: "Lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "NODEJS", Price: 299, Author: &Author{FullName: "AkshaySaini", Website: "go.dev"}})

	mux.HandleFunc("/", HomeServe)
	mux.HandleFunc("/courses", getAllCourses).Methods("GET")
	mux.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	mux.HandleFunc("/course", createCourses).Methods("POST")
	mux.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	mux.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	fmt.Println("Server is listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func HomeServe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcom to Abhinav's Server </h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(courses)
	if err != nil {
		http.Error(
			w,
			"No course found",
			http.StatusNotFound,
		)
		return
	}
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pathVar := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == pathVar["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
}

func createCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCourse Course

	err := json.NewDecoder(r.Body).Decode(&newCourse)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	if newCourse.CourseName == "" {
		http.Error(w, "CourseName cannot be empty", http.StatusBadRequest)
	}

	for _, course := range courses {
		if IsCourseExists(course.CourseName, newCourse.CourseName) {
			http.Error(w, "Course with Same name already exists", http.StatusFound)
			return
		}
	}

	newCourse.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var newCourse Course

	err := json.NewDecoder(r.Body).Decode(&newCourse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			newCourse.CourseId = params["id"]
			courses = append(courses, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Course with given id is not found")
}
