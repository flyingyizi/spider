package spider

import (
	"encoding/json"
	"os"
)

//DocCard 存放内容的卡片
type DocCard struct {
	OrigHtml []byte
	Content  []byte
	Title    string
}

// Save saves model in the file
func (f *DocCard) Save(path string) error {

	file, err := os.Create(path)
	if err == nil {
		encoder := json.NewEncoder(file)
		err = encoder.Encode(f)
	}
	file.Close()
	return err
}

// Load loads from the file
func (f *DocCard) Load(path string) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&f)
	}
	file.Close()

	return err
}
