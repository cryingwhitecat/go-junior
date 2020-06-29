package main
import(
    "fmt"
    "log"
    "context"
    //"go.mongodb.org/mongo-driver/mongo/readpref"
    //"time"
    //"bufio"
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

    users1 := make([]interface{}, len(users.Objects))
    for i := 0; i < len(users.Objects); i++ {
        users1[i]=users.Objects[i]
    }
    
    _,err = collection.InsertMany(context.TODO(), users1)
    for i:=0; i < 5; i++ {
        rslts := crud.GetUsers(collection, 5, i)
        fmt.Printf("----------------------Page # %d----------------------\n",(i+1)) 
        for i := range rslts{
            fmt.Println(rslts[i].LastName)
        }
    }
}
