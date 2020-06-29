package parsejson

import(
    "encoding/json"
    "go-junior/models"
    "os"
    "io/ioutil"
)

func ParseJson(filePath string) (*models.Objects, error){

    jsonFile, err := os.Open(filePath)

    if err != nil {
        return nil, err
    }
    defer jsonFile.Close()

    byteValue, err := ioutil.ReadAll(jsonFile)
    if err != nil{
        return nil, err
    }

    var users models.Objects


    json.Unmarshal(byteValue, &users)
    return &users, nil

}