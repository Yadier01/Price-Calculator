package filemanger

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManger) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("failed to read line in file")
	}
	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println(err)
		file.Close()
		return nil, errors.New("failed scanning gangy")
	}

	file.Close()
	return lines, nil
}

func (fm FileManger) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to created file")
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to convert data to json")
	}
	file.Close()
	return nil
}

func New(input, output string) FileManger {
	return FileManger{
		InputFilePath:  input,
		OutputFilePath: output,
	}
}
