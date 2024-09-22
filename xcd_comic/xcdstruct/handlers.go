package xcdstruct

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	dirName  = "./index"
	fileName = "index.json"
	path     = dirName + "/" + fileName
)

func urlComic(id int) string {
	return fmt.Sprintf("https://xkcd.com/%d/info.0.json", id)
}

// SearchComic search a comic
func SearchComic(comicID int) (*XcdComic, error) {
	resp, err := http.Get(urlComic(comicID))
	if err != nil {
		return nil, err
	}
	
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK { 
		return nil,fmt.Errorf("bad request %d",resp.StatusCode)
	}
	var data XcdComic
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil

}

// Serializer function save in dirs
func Serializer(data *Comics) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		if err := os.Mkdir(dirName, 0700); err != nil {
			return err
		}
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}
	file, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		return err
	}

	return nil
}

// Unserializer function unserialze form json data
func Unserializer(data *Comics) error {
	if _, err := os.Stat(path); os.IsExist(err) {
		return err
	}
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return err
	}
	return nil
}
