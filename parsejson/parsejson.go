package parsejson

import(
    "encoding/json"
    "go-junior/models"
    "os"
    "io/ioutil"
)

func ParseJson(filePath string) (*models.Objects, error){
    // Open our jsonFile
    jsonFile, err := os.Open(filePath)
    // if we os.Open returns an error then handle it
    if err != nil {
        return nil, err
    }
    defer jsonFile.Close()
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()
    // read our opened jsonFile as a byte array.
    byteValue, err := ioutil.ReadAll(jsonFile)
    if err != nil{
        return nil, err
    }
    // we initialize our Users array
    var users models.Objects

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
    json.Unmarshal(byteValue, &users)
    return &users, nil

}