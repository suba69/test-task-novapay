package storage

import (
	"encoding/json"
	"io"
	"os"
	"sync"
)

type FileStorage struct {
	filename string
	mu       sync.Mutex
	data     map[int]string
}

func NewFileStorage(filename string) *FileStorage {
	fs := &FileStorage{
		filename: filename,
		data:     make(map[int]string),
	}
	fs.load()
	return fs
}

func (fs *FileStorage) Get(id int) (string, bool) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	result, found := fs.data[id]
	return result, found
}

func (fs *FileStorage) Save(id int, result string) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.data[id] = result
	fs.save()
}

func (fs *FileStorage) load() {
	file, err := os.Open(fs.filename)
	if err != nil {
		if !os.IsNotExist(err) {
			println("Error opening file:", err.Error())
		}
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		println("Error reading file:", err.Error())
		return
	}

	if len(data) == 0 {
		return
	}

	if err := json.Unmarshal(data, &fs.data); err != nil {
		println("Error unmarshaling JSON:", err.Error())
	}
}

func (fs *FileStorage) save() {
	data, err := json.Marshal(fs.data)
	if err != nil {
		println("Error marshaling JSON:", err.Error())
		return
	}

	if err := os.WriteFile(fs.filename, data, 0644); err != nil {
		println("Error writing file:", err.Error())
	}
}
