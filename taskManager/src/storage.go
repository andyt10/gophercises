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

func GetAll() []ListItem {

	db := openDb()
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(bucketName))

		b.ForEach(func(k, v []byte) error {
			kInt := binary.BigEndian.Uint64(k)
			fmt.Printf("key=%v, value=%s\n", kInt, v)
			return nil
		})
		return nil
	})

	return nil
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

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func Remove(itemIndex int) {
}
