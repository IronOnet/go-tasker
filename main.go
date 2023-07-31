package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gookit/color.v1"
) 

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://admin:password123@localhost:6000/?authSource=admin")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("tasker").Collection("tasks")
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAT time.Time          `bson:"created_at"`
	UpdatedAT time.Time          `bson:"updated_at"`
	Text      string             `bson:"text"`
	Completed bool               `bson:"completed"`
}

func createTask(task *Task) error{
	_, err := collection.InsertOne(ctx, task) 
	return err 
}

func getAll() ([]*Task, error){
	// passing bson.D{{}} matches all documents in the collection 
	filter := bson.D{{}}
	return filterTasks(filter)
}

func filterTasks(filter interface{}) ([]*Task, error){
	var tasks []*Task 

	cur, err := collection.Find(ctx, filter) 
	if err != nil{
		return tasks, err 
	}

	for cur.Next(ctx){
		var t Task 
		err := cur.Decode(&t) 
		if err != nil{
			return tasks, err 
		}

		tasks = append(tasks, &t)
	}

	if err := cur.Err(); err != nil{
		return tasks, err 
	}

	cur.Close(ctx) 

	if len(tasks) == 0{
		return tasks, mongo.ErrNoDocuments
	}

	return tasks, nil 
}

func completeTask(text string) error{
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	t := &Task{} 
	return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
}

func printTasks(tasks []*Task){
	for i, v := range tasks{
		if v.Completed{
			color.Green.Printf("%d: %s\n", i+1, v.Text)
		} else{
			color.Yellow.Printf("%d: %s\n", i+1, v.Text)
		}
	}
}

func getPending()([]*Task, error){
	filter := bson.D{
		primitive.E{Key: "completed", Value: false}, 
	}
	return filterTasks(filter)
}

func getFinished()([]*Task, error){
	filter := bson.D{
		primitive.E{Key: "completed", Value: true},
	}

	return filterTasks(filter)
}


func deleteTask(text string) error{
	filter := bson.D{primitive.E{Key: "text", Value: text}} 

	res, err := collection.DeleteOne(ctx, filter) 
	if err != nil{
		return err 
	}

	if res.DeletedCount == 0{
		return errors.New("No tasks were deleted")
	}

	return nil 
}

func main() {
	app := &cli.App{
		Name:     "tasker",
		Usage:    "A simple CLI Program to manager your task",
		Action: func(c *cli.Context) error{
			tasks, err := getPending() 
			if err != nil{
				if err == mongo.ErrNoDocuments{
					fmt.Print("Nothing to see here\nRun `add 'task'`to add a task")
					return nil 
				}
				return err 
			}
			printTasks(tasks) 
			return nil 
		},
		Commands: []*cli.Command{
			{
				Name: "add", 
				Aliases: []string{"a"}, 
				Usage: "add a task to the list", 
				Action: func(c *cli.Context) error{
					str := c.Args().First() 
					if str == ""{
						return errors.New("cannot pass an empty parameter")
					}

					task := &Task{
						ID: primitive.NewObjectID(), 
						CreatedAT: time.Now(), 
						UpdatedAT: time.Now(),
						Text: str, 
						Completed: false,
					}

					return createTask(task)
				},
			},
			{
				Name: "all", 
				Aliases: []string{"l"},
				Usage: "list all tasks",
				Action: func(c *cli.Context) error{
					tasks, err := getAll() 
					if err != nil{
						if err == mongo.ErrNoDocuments{
							fmt.Printf("Nothing to see heare this database hasn't been updated yet")
							return nil 
						}
						return err 
					}
					printTasks(tasks)
					return nil 
				},
			},
			{
				Name: "done", 
				Aliases: []string{"d"}, 
				Usage: "completed a task on the list", 
				Action: func(c *cli.Context) error{
					text := c.Args().First() 
					return completeTask(text)
				},
			},
			{
				Name: "finished", 
				Aliases: []string{"f"}, 
				Usage: "list completed tasks", 
				Action: func(c *cli.Context) error{
					tasks, err := getFinished() 
					if err != nil{
						if err == mongo.ErrNoDocuments{
							fmt.Print("Nothing to show") 
							return nil 
						}
						return err 
					}
					printTasks(tasks) 
					return nil 
				},
			},
			{
				Name:"rm", 
				Usage: "deletes a task on the list", 
				Action: func(c *cli.Context) error{
					text := c.Args().First() 
					err := deleteTask(text) 
					if err != nil{
						return err 
					}
					return nil 
				},
			},


		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
