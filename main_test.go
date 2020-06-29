package main
import(
	"testing"
    "go-junior/crud"
    "go-junior/parsejson"
    "go.mongodb.org/mongo-driver/mongo"
    "go-junior/models"
    "go-junior/validation"
    "fmt"
)
func TestInvalidateGender(t *testing.T){
	fmt.Println("testing gender")
	invalidGender := "Pig"
	err := validation.ValidateGender(invalidGender)
	if err == nil{
		t.Error("error expected")
	}
}
func TestValidateEmail(t *testing.T){
	validEmail := "Valerie_Gavin9167@nimogy.biz"
	err := validation.ValidateEmail(validEmail)
	if err != nil{
		t.Error("no errors expected")
	}
}
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
func TestDuplicateEmailThrowsError(t *testing.T){
	collection,_,err := setup()
	usr, err := crud.SearchByEmail("Valerie_Gavin9167@nimogy.biz", collection)
	err = crud.InsertUser(usr, collection)
	if err == nil{
		t.Error("Error expected")
	}
}
func TestGetUsersPaginates(t *testing.T){
	collection, _, _ := setup()
	users,_ := crud.GetUsers(collection, 5, 1)
	if len(users) != 5{
		t.Error("Expected 5 records, got ", len(users))
	}
}

func TestInsertUserUpdatesTable(t *testing.T){
	collection, _, _ := setup()
	user := models.User{
		Email: "david.rudenko.2001@ukr.net",
		BirthDate: "Monday, March 28, 8546 2:32 AM",
		LastName: "Rudenko",
		Country: "Ukraine",
		Gender: "Male",
	}
	err := crud.InsertUser(&user,collection)
	if err != nil{
		t.Error("Expected no errors, got ", err)
	}
	usr, err := crud.SearchByEmail("david.rudenko.2001@ukr.net", collection)
	if (usr.LastName != "Rudenko") || (usr.Email != "david.rudenko.2001@ukr.net") ||
    	(usr.Country != "Ukraine") || (usr.Gender != "Male"){
    		t.Error("Fields don`t match")
    	}
}
func TestInvalidateName(t *testing.T){
	invalidName := ""
	err := validation.ValidateNameOrCountry(invalidName)
	if err == nil{
		t.Error("error expected")
	}
}
func TestInvalidateCountry(t *testing.T){
	invalidCountry := ""
	err := validation.ValidateNameOrCountry(invalidCountry)
	if err == nil{
		t.Error("error expected")
	}
}
func TestValidateName(t *testing.T){
	validName := "David"
	err := validation.ValidateNameOrCountry(validName)
	if err != nil{
		t.Error("no errors expected, got",err)
	}
}
func TestInvalidateBDate(t *testing.T){
	invalidDate := "Monday, March 35, 8546 2:32 AM"
	err := validation.ValidateBDay(invalidDate)
	if err == nil{
		t.Error("error expected")
	}
}
func TestValidateBDay(t *testing.T){
	validDate := "Monday, March 28, 8546 2:32 AM"
	err := validation.ValidateBDay(validDate)
	if err != nil{
		t.Error("no errors expected")
	}
}
func TestInvalidateEmail(t *testing.T){
	invalidEmail := "Valerie_Gavin9167nimogy.biz"
	err := validation.ValidateEmail(invalidEmail)
	if err == nil{
		t.Error("error expected")
	}
}
