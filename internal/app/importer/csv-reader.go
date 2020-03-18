package importer

import (
	"bufio"
	"github.com/dev-tim/message-board-api/internal/app/model"
	"github.com/gocarina/gocsv"
	"io"
	"os"
)

func ReadCSVFromFile(filename string) ([]*model.Message, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	csv, err := ReadCSV(reader)
	if err != nil {
		return nil, err
	}

	return csv, nil
}

func ReadCSV(in io.Reader) ([]*model.Message, error) {
	var messages = make([]*model.Message, 0)

	if err := gocsv.Unmarshal(in, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
