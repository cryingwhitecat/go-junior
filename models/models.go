package models

import("fmt")

type Objects struct {
	Objects []User `json:"objects"`
}
type User struct {
	Email      string `json:"email"`
	LastName   string `json:"last_name"`
	Country    string `json:"country"`
	Gender     string `json:"gender"`
	BirthDate  string `json:"birth_date"`
}
func (usr *User) PrintString(){
    fmt.Println("User`s Email: " + usr.Email)
    fmt.Println("User`s Last Name: " + usr.LastName)
    fmt.Println("User`s DOB: " + usr.BirthDate)
    fmt.Println("User`s Country: " + usr.Country)
    fmt.Println("Users`s Gender: " + usr.Gender)
}