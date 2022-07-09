package db

import (
	"encoding/binary"
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"os"
	"path/filepath"
	"time"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
	Done  bool
}

func Init() error {
	dbPath, err := getPath()
	if err != nil {
		return err
	}

	var err2 error
	db, err2 = bolt.Open(dbPath, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err2 != nil {
		return err2
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func getPath() (string, error) {
	dirname, err := os.UserHomeDir()

	return filepath.Join(dirname, "/task_manager.db"), err
}

func CreateTask(taskName string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)

		task := Task{
			Key:   id,
			Value: taskName,
			Done:  false,
		}
		encoded, _ := json.Marshal(task)

		return b.Put(key, encoded)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func DoneTask(key int) (int, error) {
	id := make([]byte, 8)
	id = itob(key)

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		result := b.Get(id)

		var task Task
		json.Unmarshal(result, &task)

		task.Done = true

		encoded, _ := json.Marshal(task)

		return b.Put(id, encoded)
	})
	if err != nil {
		return -1, err
	}
	return btoi(id), nil
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var uncodedTask Task
			json.Unmarshal(v, &uncodedTask)

			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: uncodedTask.Value,
				Done:  uncodedTask.Done,
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
