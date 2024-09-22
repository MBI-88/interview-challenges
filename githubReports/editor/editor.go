package editor

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"os/exec"
	"bytes"

)

// Editor  function to open  an editor
func Editor(fields map[string]string) error{
	editorName := "nano"
	temp, er := os.CreateTemp("./temp","")
	if er != nil {return er}
	tempFileName := temp.Name()
	defer os.Remove(tempFileName)
	encoder := json.NewEncoder(temp)
	encoder.SetIndent("","  ")
	if err  := encoder.Encode(fields); err != nil {
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}

	cmd := exec.Command(editorName,tempFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	edited,err := ioutil.ReadFile(tempFileName)
	if err != nil {return err}
	if err := json.Unmarshal(removeUTF8BOM(edited),&fields); err != nil {
		return err
	}
	return nil
}

func removeUTF8BOM(s []byte) []byte {
	utf8Bom := []byte{239, 187, 191}
	return bytes.TrimPrefix(s, utf8Bom)
}