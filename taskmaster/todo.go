package taskmaster

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type item struct {
	Task        string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(id int) {
	todo := t.GetItem(id)

	todo.Completed = true
	todo.CompletedAt = time.Now()

	(*t)[id-1] = todo
}

func (t *Todos) Delete(id int) {
	*t = append((*t)[:id-1], (*t)[id:]...)
}

func (t *Todos) GetItem(id int) item {
	if id < 1 || id > len(*t) {
		panic("Invalid id")
	}

	todo := (*t)[id-1]

	return todo
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Save(filename string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
