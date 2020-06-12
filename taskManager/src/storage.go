package src

import bolt "go.etcd.io/bbolt"

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

const defaultFolder string = "/tmp"
const defaultFileName string = "data.db"

const bucketName string = "TASKS"

var fileLoc string

func InitDb(fileLocParam string) {
	if fileLocParam == "" {

		fileLoc = defaultFolder + "/" + defaultFileName
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

func GetAll() []ListItemEntry {

	db := openDb()
	defer db.Close()

	var itemList []ListItemEntry
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))

		b.ForEach(func(k, v []byte) error {

			var thisTask ListItemEntry
			err := json.Unmarshal(v, &thisTask.Data)
			thisTask.Index = btoi(k)

			if err != nil {
				fmt.Println(err)
				return err
			}

			itemList = append(itemList, thisTask)

			return nil
		})

		//eventually maybe return an error if one of forEach fails
		return nil
	})

	return itemList
}

func Add(itemData ListItem) error {

	db := openDb()
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		id, _ := b.NextSequence()

		buf, err := json.Marshal(itemData)
		if err != nil {
			return err
		}

		// Persist bytes to users bucket.
		return b.Put(itob(int(id)), buf)
	})

	return nil
}

//not sure why docs convert to bigendian for bbolt. Need to figure out.
// ^ This made sense when I wasn't so sleppy.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	i := int(binary.BigEndian.Uint64(b))
	return i
}

//Update Task
//Returns nil if no error, error type if problem removing Task.
func Update(itemEntry ListItemEntry) error {

	db := openDb()
	defer db.Close()

	res := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		buf, err := json.Marshal(itemEntry.Data)
		if err != nil {
			return err
		}

		return b.Put(itob(itemEntry.Index), buf)
	})

	return res

}
