package main

import (
	"fmt"
)

type StorageType int
type Store struct{}

func (store *Store) Open(file string) (string, error) { return "", nil }

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func NewStore(t StorageType) *Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage("")
	case DiskStorage:
		return newDiskStorage("")
	default:
		return newTempStorage("")
	}
}

func main() {
	s := NewStore(MemoryStorage)
	f, _ := s.Open("file")
	fmt.Print(f)
}

func newMemoryStorage(tipo string) *Store { return &Store{} }
func newDiskStorage(tipo string) *Store   { return &Store{} }
func newTempStorage(tipo string) *Store   { return &Store{} }
