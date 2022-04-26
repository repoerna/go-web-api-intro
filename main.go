package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"main/utils"
	"net/http"
	"utils"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
    DB_USER     = "lwzbzejw"
    DB_PASSWORD = "i0xAI776rL4ETUAljCkEhAxI3hH3Oh39"
    DB_NAME     = "lwzbzejw"
    DB_HOST     = "arjuna.db.elephantsql.com"
    DB_PORT     = 5432
)

// define model
type Student struct {
    ID   int `json:"id"`
    Name string `json:"name"`
    Gender string `json:"gender"`
}

type JsonResponse struct {
    Type    string `json:"type"`
    Message string `json:"message"`
    Data    []Student `json:"data"`
}

type JsonResponseOther struct {
    Type    string `json:"type"`
    Message string `json:"message"`
    Data    interface{} `json:"data"`
}

// Function for handling errors
// func checkErr(err error) {
//     if err != nil {
//         panic(err)
//     }
// }

// DB set up
func setupDB() *sql.DB {
    dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)

    utils.checkErr(err)

    return db
}

// Main function
func main() {

    // Init the mux router
    router := mux.NewRouter()

// Route handles & endpoints

    // Get all students
    router.HandleFunc("/", homePage).Methods("GET")
  
    router.HandleFunc("/students", GetStudents).Methods("GET")

  // to insert data
    router.HandleFunc("/students", CreateStudents).Methods("POST")


    // serve the app
    fmt.Println("Server at 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}

// util help logging
func printMessage(message string) {
    fmt.Println("")
    fmt.Println(message)
    fmt.Println("")
}


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "OK")
    fmt.Println("Endpoint Hit: homePage")
}

// Get all students
// response and request handlers
func GetStudents(w http.ResponseWriter, r *http.Request) {
    db := setupDB()

    printMessage("Getting students...")

    // Get all movies from movies table that don't have movieID = "1"
    rows, err := db.Query("SELECT * FROM students")

    // check errors
    utils.CheckErr(err)

    // var response []JsonResponse
    var students []Student

    // Foreach movie
    for rows.Next() {
        var id int
        var name string
        var gender string

        err = rows.Scan(&id, &name, &gender)

        // check errors
        checkErr(err)

        students = append(students, Student{ID: id, Name: name, Gender: gender})
    }

    var response = JsonResponse{Type: "success", Data: students}

    json.NewEncoder(w).Encode(response)
}

func CreateStudents(w http.ResponseWriter, r *http.Request) {
  
    id := r.FormValue("id")
    name := r.FormValue("name")
    gender := r.FormValue("gender")

    var response = JsonResponseOther{}

    if id == "" || name == "" {
        response = JsonResponseOther{Type: "error", Message: "You are missing id or name parameter."}
    } else {
        db := setupDB()

        printMessage("Inserting student into DB")

        fmt.Println("Inserting new student with ID: " + id + " and name: " + name)

        var lastInsertID int
    err := db.QueryRow("INSERT INTO students(id, name, gender) VALUES($1, $2, $3) returning id;", id, name, gender).Scan(&lastInsertID)

      fmt.Println(lastInsertID)

    // check errors
    utils.CheckErr(err)

    response = JsonResponseOther{Type: "success", Message: "The student has been inserted successfully!", Data: lastInsertID}
    }

    json.NewEncoder(w).Encode(response)
}