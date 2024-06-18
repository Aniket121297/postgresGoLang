package main

import
	"fmt"

type Student struct {
	ID      int    
	Name    string 
}

func main{
	host := "some-host-value"
	port := "5432"
	user := "postgres"
	password := "password"
	dbname := "sampleStudentDB"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
  		log.Fatalf("Error while connecting to DB: %s", err)
	}

	sql := "select id,name FROM student;"
	data, err := DB.Query(sql)
	if err != nil {
		saveError := fmt.Sprintf("Error while executing the get query, and %s", err)
	}

	type Students []Student

	for data.Next() {
		var Student Student
		err = data.Scan(&Student.ID, &Student.Name)
		if err != nil {
			saveError := fmt.Sprintf("Error Looping data, and %s", err)
		}
		Students = append(Students, Student)
	}

	fmt.Println("List of Students: ", Student)
}