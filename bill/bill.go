package bill

import (
	"encoding/json"
	"os"
	"strconv"
)

type Item struct {
	Text  string
	Value int
	Date  string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	items := []Item{}

	if err := json.Unmarshal(file, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (i *Item) PrettyV() string {
	return "R$" + strconv.Itoa(i.Value)
}
