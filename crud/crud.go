package crud

import(

    "fmt"
    "log"
    "context"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/x/bsonx"
    "go-junior/models"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func Connect(dbUrl string, dbName string, collectionName string, constraintName string) (*mongo.Collection, error){
    client, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
    if err != nil {
        return nil, err
    }
    // Create connect
    err = client.Connect(context.TODO())
    if err != nil {
        return nil, err
    }
    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        return nil, err
    }
    fmt.Println("Connected to MongoDB!")
    collection := client.Database(dbName).Collection(collectionName)
    _, err = collection.Indexes().CreateOne(
        context.Background(),
        mongo.IndexModel{
            Keys:     bsonx.Doc{{constraintName, bsonx.Int32(1)}},
            Options:  options.Index().SetUnique(true),
        },
    )
    if err != nil{
        return nil,err
    }
    return collection, nil
}
func SearchByEmail(email string, coll *mongo.Collection) (*models.User, error){

    options := options.FindOne()
    filter := bson.D{{"email",email}}
    var user models.User

    err := coll.FindOne(context.TODO(), filter, options).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user,nil
}
func UpdateUser(stringId string, user *models.User, coll *mongo.Collection) error{
    id, _ := primitive.ObjectIDFromHex(stringId)
    _, err := coll.UpdateOne(
        context.TODO(),
        bson.M{"_id": id},
        bson.D{
            {"$set", bson.D{{"email",user.Email},{"lastname",user.LastName},
                {"birth_date",user.BirthDate},{"country",user.Country},{"gender",user.Gender}}},
        },
    )
    return err
}

func InsertUser(user models.User, coll *mongo.Collection) {
    insertResult, err := coll.InsertOne(context.TODO(), user)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)
    return
}
func GetUsers(coll *mongo.Collection, pageSize int ,pageNum int) []*models.User{
    limit :=  int64(pageSize)
    skip := int64(pageSize * (pageNum ))
    findOptions := options.Find()
    findOptions.SetLimit(limit)
    findOptions.SetSkip(skip)
    cur, err := coll.Find(context.TODO(), bson.D{{}}, findOptions)
    if err != nil{
        log.Fatal(err)
    }
    var results[] *models.User
    for cur.Next(context.TODO()) {
        // create a value into which the single document can be decoded
        var elem models.User
        err := cur.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }

        results = append(results, &elem)
    }

        if err := cur.Err(); err != nil {
            log.Fatal(err)
        }
        // Close the cursor once finished
        cur.Close(context.TODO())
        return results
}