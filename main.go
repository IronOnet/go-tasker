package main 

import (
	"context" 
	"log" 
	"os"
	
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo" 
	"go.mongodb.org/mongo-driver/mongo/options"
)


var collection *mongo.Collection 
var ctx = context.TODO()

func init(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27018")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil{
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil) 
	if err != nil{
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks");
}

func main(){
	app := &cli.App{
		Name: "tasker", 
		Usage: "A simple CLI Program to manager your task", 
		Commands: []*cli.Command{},
	}

	err := app.Run(os.Args) 
	if err != nil{
		log.Fatal(err)
	}
}