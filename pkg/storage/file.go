package storage

import (
	"encoding/json"
	"os"
)

type Storage interface {
	Read(data any) error
	Write(data any) error
}

type Type string

const (
	FileType Type = "file"
)

type FileStorage struct {
	FileName string
}

func (fs *FileStorage) Read(data any) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}

func (fs *FileStorage) Write(data any) error {
	normalizedData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.FileName, normalizedData, 0644)
}

func NewFileStorage(storage Type, fileName string) Storage {
	switch storage {
	case FileType:
		return &FileStorage{FileName: fileName}
	}

	return nil
}
