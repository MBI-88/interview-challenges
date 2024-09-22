package poster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

func (om OmdAPI) getPath() (string, string) {
	return fmt.Sprintf("./%s", om.Title), fmt.Sprintf("%s.json", om.Title)
}

func (om OmdAPI) createDir() error {
	path, fileName := om.getPath()
	if _, er := os.Stat(path); os.IsNotExist(er) {
		if err := os.Mkdir(path, 0700); err != nil {
			return err
		}
	}
	// Creating json file
	if _, err := os.Stat(path + "/" + fileName); os.IsNotExist(err) {
		if _, err := os.Create(path + "/" + fileName); err != nil {
			return err
		}
	}
	return nil
}

// Serializer method. Save data in json file
func (om OmdAPI) serializer() error {
	if err := om.createDir(); err != nil {
		return err
	}
	path, fileName := om.getPath()
	file, err := os.OpenFile(path+"/"+fileName, os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encode := json.NewEncoder(file)
	encode.SetIndent("", "  ")
	if err := encode.Encode(om); err != nil {
		return err
	}

	return nil
}

// Unserializer method. Return data from json file
func (om OmdAPI) unserializer() error {
	path, fileName := om.getPath()
	file, err := os.OpenFile(path+"/"+fileName, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	if err := json.NewDecoder(file).Decode(&om); err != nil {
		return err
	}
	return nil
}

// getPoster method save movie poster
func (om OmdAPI) getPoster() error {
	if om.Poster == "" {
		return fmt.Errorf("error poster url is empty")
	}
	resp, err := http.Get(om.Poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error in response %d", resp.StatusCode)
	}
	path, _ := om.getPath()
	fileName := fmt.Sprintf("%s.jpg", om.Title)
	if _, err := os.Stat(path + "/" + fileName); os.IsNotExist(err) {
		file, err := os.Create(path + "/" + fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		img, _ := ioutil.ReadAll(resp.Body)
		if _, err := file.Write(img); err != nil {
			return err
		}
	}
	return nil
}

// Search method. Search a movie in omdapi
func (om *OmdAPI) Search(movie string, poster bool) error {
	url := fmt.Sprintf("%s?apikey=%s&t=%s", api, key, movie)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error bad request %d", resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(om); err != nil {
		return fmt.Errorf("error in json decoder %s", err)
	}
	if om.Title != "" {
		if err := om.serializer(); err != nil {
			return fmt.Errorf("error in serilizer json file %s", err)
		}
		if poster {
			if err := om.getPoster(); err != nil {
				return fmt.Errorf("error downloading poster %s", err)
			}
		}
		if err := om.Show(); err != nil {
			return err
		}
	}

	return nil
}

// Show data from json file
func (om OmdAPI) Show() error {
	title := fmt.Sprintf("Data from move %s", om.Title)
	temp := template.Must(template.New(title).Parse(`
		-----------------------------------
		| Title: -> {{.Title}}
		-----------------------------------
		| Year:  -> {{.Year}}
		-----------------------------------       
		| Rated: -> {{.Rated}}
		-----------------------------------     
		| Released: ->  {{.Released}}
		-----------------------------------   
		| Runtime:  ->  {{.Runtime}}
		-----------------------------------   
		| Genre:  ->  {{.Genre}}
		----------------------------------- 
		| Director:  ->  {{.Director}}
		----------------------------------- 
		| Writer:  ->  {{.Writer}}
		----------------------------------- 
		| Actors:  ->  {{.Actors}}
		----------------------------------- 
		| Plot:  ->  {{.Plot}}
		----------------------------------- 
		| Language:  ->  {{.Language}}
		----------------------------------- 
		| Country:  ->  {{.Country}}
		----------------------------------- 
		| Awards:  ->  {{.Awards}}
		----------------------------------- 
		| Ratings:cc
		  {{range .Ratings}}
		| Source:  ->  {{.Source}}
		| Value:  ->  {{.Value}}
		{{end}}
		----------------------------------- 
		| Metascore:  ->  {{.Metascore}}
		----------------------------------- 
		| ImdbRating:  ->  {{.ImdbRating}}
		----------------------------------- 
		| ImdbVotes:  ->  {{.ImdbVotes}}
		----------------------------------- 
		| ImdbID:  ->  {{.ImdbID}}
		----------------------------------- 
		| Type:  ->  {{.Type}}
		----------------------------------- 
		| DVD:  ->  {{.DVD}}
		----------------------------------- 
		| BoxOffice:  ->  {{.BoxOffice}}
		----------------------------------- 
		| Production:  ->  {{.Production}}
		----------------------------------- 
		| Website:  ->  {{.Website}}
		----------------------------------- 
		| Response:  ->  {{.Response}}
		----------------------------------- 
	`))
	if err := temp.Execute(os.Stdout, om); err != nil {
		return err
	}

	return nil
}
