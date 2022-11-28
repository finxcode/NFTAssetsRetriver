package helper

import (
	"NftAssetsRetriever/global"
	"bufio"
	"errors"
	"os"
)

func ReadIn() ([]string, error) {
	var token_ids []string
	path := global.App.Config.File.Path
	filename := global.App.Config.File.Name
	file, err := os.Open(path + filename)
	if err != nil {
		return nil, errors.New("file open error")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		token_ids = append(token_ids, fileScanner.Text())
	}
	return token_ids, nil
}
