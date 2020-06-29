package main
import(
	"testing"
    "go-junior/crud"
    "go-junior/parsejson"
    "go.mongodb.org/mongo-driver/mongo"
    "go-junior/models"
)

func setup() (*mongo.Collection,models.Objects,error){
	jsonPath := "users_go.json"
    users,_ := parsejson.ParseJson(jsonPath)
    collection,err := crud.Connect("mongodb://localhost:27017", "go-junior", "users", "email")
    return collection, *users, err
}

// db should be dropped prior to all tests
func TestMultipleInsert(t *testing.T) {
    collection,users,err := setup()
    if err!= nil{
         t.Error("No errors expected, instead got", err)
    }
    err = crud.AddUsers(users.Objects,collection)
    if err != nil{
        t.Error("No errors expected, instead got", err)
    }
}
func TestSearchByEmail(t *testing.T){
    collection,_,err := setup()
    if err != nil{
        t.Error("No errors expected, instead got", err)
    }
    //err = crud.AddUsers(users.Objects,collection)
    usr, err := crud.SearchByEmail("Valerie_Gavin9167@nimogy.biz", collection)
    if (usr.LastName != "Gavin") || (usr.Email != "Valerie_Gavin9167@nimogy.biz") ||
    	(usr.Country != "Kazakhstan") || (usr.Gender != "Female"){
    		t.Error("Fields don`t match")
    	}
}
func DuplicateEmailThrowsError(t *testing.T){
	collection,_,err := setup()
	usr, err := crud.SearchByEmail("Valerie_Gavin9167@nimogy.biz", collection)
	err = crud.InsertUser(*usr, collection)
	if err == nil{
		t.Error("Error expected")
	}
}
func GetUsersPaginates(t *testing.T){
	collection, _, _ := setup()
	users := crud.GetUsers(collection, 5, 1)
	if len(users) != 5{
		t.Error("Expected 5 records, got ", len(users))
	}
}

func InsertUserUpdatesTable(t *testing.T){
	collection, _, _ := setup()
	user := models.User{
		Email: "david.rudenko.2001@ukr.net",
		BirthDate: "Monday, March 28, 8546 2:32 AM",
		LastName: "Rudenko",
		Country: "Ukraine",
		Gender: "Male",
	}
	err := crud.InsertUser(user,collection)
	if err != nil{
		t.Error("Expected no errors, got ", err)
	}
	usr, err := crud.SearchByEmail("Valerie_Gavin9167@nimogy.biz", collection)
	if (usr.LastName != "Rudenko") || (usr.Email != "david.rudenko.2001@ukr.net") ||
    	(usr.Country != "Ukraine") || (usr.Gender != "Male"){
    		t.Error("Fields don`t match")
    	}
}