package jsondb

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/pnkj-kmr/looking-glass/repos"
	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
)

type db struct {
	mu   sync.Mutex
	name string
	path string
}

// New - initilizer for a db model
func New(f string) (repos.Model, error) {
	path := filepath.Join(storagePath, f)
	dir, err := getOrCreateDir(path)
	if err != nil {
		utils.L.Error("new table path", zap.String("path", path), zap.String("err", err.Error()))
		return nil, err
	}
	if !dir.IsDir() {
		utils.L.Warn("it's not a directory", zap.String("path", path))
		return nil, fmt.Errorf("not a directory")
	}

	return &db{name: f, path: path}, err
}

// GetAll - returns all record of model
func (db *db) GetAll() [][]byte {
	var list [][]byte
	records, err := ioutil.ReadDir(db.path)
	if err != nil {
		utils.L.Error("no data available", zap.String("err", err.Error()))
		return list
	}
	for _, r := range records {
		if !r.IsDir() {
			fPath := filepath.Join(db.path, r.Name())
			data, err := os.ReadFile(fPath)
			if err != nil {
				utils.L.Error("unable to read the data file", zap.String("err", err.Error()), zap.String("f", fPath))
				continue
			}
			list = append(list, data)
		}
	}
	return list
}

// Get help to retrive key based record
func (db *db) Get(key string) ([]byte, error) {
	record := key + fileExtension
	fPath := filepath.Join(db.path, record)
	r, err := os.Stat(fPath)
	if err != nil {
		utils.L.Error("no data available", zap.String("err", err.Error()), zap.String("record", record))
		return nil, err
	}
	if r.IsDir() {
		utils.L.Warn("invalid record", zap.String("record", r.Name()))
		return nil, fmt.Errorf("invalid record key")
	}
	data, err := os.ReadFile(fPath)
	if err != nil {
		utils.L.Error("unable to read the record", zap.String("err", err.Error()), zap.String("f", fPath))
		return nil, err
	}
	return data, nil
}

// Insert - helps to save data into model dir
func (db *db) Insert(key string, data []byte) error {
	db.mu.Lock()
	record := key + fileExtension
	err := ioutil.WriteFile(filepath.Join(db.path, record), data, os.ModePerm)
	if err != nil {
		utils.L.Error("insertion error", zap.String("err", err.Error()))
	}
	db.mu.Unlock()
	return err
}

// Delete - helps to delete model dir record
func (db *db) Delete(key string) error {
	db.mu.Lock()
	record := key + fileExtension
	err := os.Remove(filepath.Join(db.path, record))
	if err != nil {
		utils.L.Error("deletion error", zap.String("err", err.Error()))
	}
	db.mu.Unlock()
	return err
}
