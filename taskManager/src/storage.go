package src

import bolt "go.etcd.io/bbolt"

import (
	"fmt"
	"os"
)

const defaultLoc string = "/tmp"

const bucketName string = "TASKS"

var fileLoc string

func InitDb(fileLocParam string) {
	if fileLocParam == "" {
		fileLoc = defaultLoc
		return
	}

	fileLoc = fileLocParam
}

func openDb() *bolt.DB {
	db, err := bolt.Open(fileLoc, 0666, nil)
	if err != nil {
		fmt.Println("Error opening DB")
		fmt.Println(err)
		os.Exit(1)
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			fmt.Println(fmt.Errorf("create bucket: %s", err))
			os.Exit(1)
		}
		return nil
	})

	return db
}

func GetAll() []ListItem {

	db := openDb()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		//figure out what the type of a bucket is, will probably need to serailise the ListItem struct somehow. Time for sleep.
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		return err
	})

	return nil
}

func Add(itemData string) {
}

func Remove(itemIndex int) {
}
