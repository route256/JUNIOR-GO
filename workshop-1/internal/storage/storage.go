package storage

import (
	"encoding/json"
	"io"

	"to-do-list/internal/tasks"
)

type File interface {
	io.ReadWriteCloser
	Truncate(size int64) error
	Seek(offset int64, whence int) (ret int64, err error)
}

type Storage struct {
	file File
}

func NewStorage(file File) *Storage {
	return &Storage{file: file}
}

func (s *Storage) Create(task tasks.Task) error {
	list, err := s.List()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(append(list, &task))
	if err != nil {
		return err
	}

	if err := s.file.Truncate(0); err != nil {
		return err
	}

	if _, err := s.file.Seek(0, 0); err != nil {
		return err
	}

	if _, err = s.file.Write(bytes); err != nil {
		return err
	}

	return nil
}

func (s *Storage) List() (tasks.TaskList, error) {
	readAll, err := io.ReadAll(s.file)
	if err != nil {
		return nil, err
	}

	var list tasks.TaskList

	if len(readAll) == 0 {
		return list, nil
	}

	if err := json.Unmarshal(readAll, &list); err != nil {
		return nil, err
	}

	return list, nil
}
