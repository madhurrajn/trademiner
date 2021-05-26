package scribledb

import (
	"log"

	scribble "github.com/nanobox-io/golang-scribble"
)

var globalDbObj *scribble.Driver

var dir = "/Users/madhurrajn/nseobject"

func InitScribbleDb() (*scribble.Driver, error) {
	dbObj, err := scribble.New(dir, nil)
	if err != nil {
		log.Println("Error ", err)
	}
	globalDbObj = dbObj
	return dbObj, err
}

func Write(collection interface{}, key interface{}, obj interface{}) error {
	err := globalDbObj.Write(collection.(string), key.(string), obj)
	if err != nil {
		log.Printf("Unable to Write to Db")
	}
	return err
}

func Read(collection interface{}, key interface{}, newObj interface{}) (interface{}, error) {
	err := globalDbObj.Read(collection.(string), key.(string), newObj)
	if err != nil {
		log.Printf("Unable to Write to Db")
	}
	return newObj, err
}

func ReadAll(collection interface{}) ([]string, error) {
	records, err := globalDbObj.ReadAll(collection.(string))
	if err != nil {
		log.Printf("Unable to Write to Db")
		return nil, err
	}
	return records, nil
}
