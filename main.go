package main
import(
    "log"
    "go-junior/crud"
    "go-junior/parsejson"
)

func main() {
    jsonPath := "users_go.json"
    users,_ := parsejson.ParseJson(jsonPath)
    collection,err := crud.Connect("mongodb://localhost:27017", "go-junior", "users", "email")

    if err!= nil{
        log.Fatal(err)
    }

    err = crud.AddUsers(users.Objects,collection)
    if err != nil{
        log.Fatal(err)
    }
}
