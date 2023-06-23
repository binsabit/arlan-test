package main

import (
	"bufio"
	"io"
	"os"
	"path"

	storage "github.com/binsabit/arlan-test/proto"
)

type Storage interface {
	StoreImage(*storage.StoreImageRequest) error
	GetImage(*storage.GetImageRequest) (*storage.Image, error)
}

type FileStorage struct {
	Dir string
}

func (f FileStorage) StoreImage(req *storage.StoreImageRequest) error {
	//create directory with name
	filepath := path.Join(f.Dir, req.Image.Name)
	_ = os.Mkdir(filepath, os.ModePerm)
	//create file
	file, err := os.OpenFile(path.Join(filepath, req.Image.Name+".webp"), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	//write to file
	file.Write([]byte(req.Image.Content))
	defer file.Close()
	return nil
}

func (f FileStorage) GetImage(req *storage.GetImageRequest) (*storage.Image, error) {
	filepath := path.Join(f.Dir, req.Name, req.Name+".webp")
	//open file
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	//get file stat to convert file to bytes
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// Read the file into a byte slice
	bs := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bs)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return &storage.Image{
		Uid:     "1",
		Content: bs,
		Name:    req.Name,
	}, nil
}
