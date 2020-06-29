package validation
import(
    "errors"
    "regexp"
    "time"
    "go-junior/models"
)
func  ValidateBDay(userString string) error{
	layout_string := "Monday, January 2, 2006 3:04 PM"
    _, err := time.Parse(layout_string, userString)
    return err
}
func ValidateEmail(userEmail string) error{
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    if !re.MatchString(userEmail){
        return errors.New("Invalid Email")
    }
    return nil
}
func ValidateNameOrCountry(userString string) error{
	if userString == ""{
		return errors.New("Empty Name/Country")
	}else{
		return nil
	}
}
func ValidateGender(userGender string) error{
	if !((userGender == "Male") || (userGender == "Female")){
		return errors.New("Invalid Gender")
	}else{ 
	return nil}
}
func ValidateUser(user *models.User) error{
	if  err := ValidateBDay(user.BirthDate); err != nil{
        return err 
    }
    if err := ValidateEmail(user.Email); err != nil{
        return err
    }
    if err := ValidateBDay(user.BirthDate); err != nil{
        return err
    }
    if err := ValidateNameOrCountry(user.LastName); err != nil{
        return err
    }
    if err := ValidateNameOrCountry(user.Country); err != nil{
        return err
    }
    if err := ValidateGender(user.Gender); err != nil{
        return err
    }
    return nil

}