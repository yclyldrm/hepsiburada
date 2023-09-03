package utils

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func ReadFromFile(filename string) []string {
	commands := make([]string, 0)
	f, err := excelize.OpenFile(filename)
	if err != nil {
		log.Println(err)
		return commands
	}

	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		log.Println(err)
		return commands
	}

	for ind, row := range rows {
		if ind == 0 {
			log.Println("first row is eleminated")
			continue
		}
		for i, col := range row {
			if i == 1 {
				continue
			}
			commands = append(commands, col)
		}
	}

	return commands
}
