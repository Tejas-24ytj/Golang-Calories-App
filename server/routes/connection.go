package routes

import (
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options")

func DBinstance() *mongo.Client
{
   MongoDb="mongodb://localhost:27017/caloriesdb"

   client,err:=mongo.NewClient(options.Client().ApplyURL(MongoDb))
   if(err!=nil){
	log.Fatal(err)
}
  ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
  defer cancel()
  err!=client.Connect(ctx)
  if err!=nil{
	log.Fatal(err)
  }
  fmt.Println("Connected to MongoDB")
  return
}

var Client *mongo.Client=DBinstance()

func openCollection(client *mongo.Client, collectionname string) *mongo.Collection{
	  var collection *mongo.Collection =client.Database("caloriesdb").Collection(collectionName)
	  return collection 
}