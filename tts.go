package tts

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	uuid "github.com/satori/go.uuid"
)

type AudioFormat int32

const (
	MP3 AudioFormat = 0
	WAV AudioFormat = 1
	MP4 AudioFormat = 3
)

type FileOverWriteConfig struct{}

type TextToSpeech struct {
	Dir   string
	Voice int32
	//Language            string
	Format              AudioFormat
	Engine              int32
	FileOverwrite       bool
	FileOverWriteConfig FileOverWriteConfig
}

func (t *TextToSpeech) TextToSpeech(text, fileName string) (fileDir string, err error) {

	if fileName == "" {
		fileName = uuid.NewV4().String()
	}

	if t.Dir == "" {
		t.Dir = "audio"
	}

	fileDir = t.Dir + "/" + fileName + "." + getAudioFormat(t.Format)

	if err = t.createFolder(); err == nil {

		err = t.createFile(text, fileDir)

	}
	return
}

func (t *TextToSpeech) createFolder() (err error) {

	directory, err := os.Open(t.Dir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(t.Dir, 0755)
	} else {
		directory.Close()
	}

	return
}

func (t *TextToSpeech) createFile(text, fileName string) (err error) {

	var writeFile = true

	file, err := os.Open(fileName)

	if err == nil {

		file.Close()
		if !t.FileOverwrite {
			writeFile = false
			err = fmt.Errorf("File already exists - to overwrite set FileOverwrite = true")
		}
	}

	err = nil

	if writeFile {

		switch t.Engine {

		//google translate enginge
		case 0:
			downloadURL := fmt.Sprintf("https://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=tw-ob", url.QueryEscape(text), "en")
			response, err := t.downloadFile(downloadURL)

			if err == nil {
				defer response.Body.Close()
				//write to file
				output, err := os.Create(fileName)

				if err == nil {
					io.Copy(output, response.Body)
				}
			}

		}

	}

	return
}

// future proxy // configuration functionalities
func (t *TextToSpeech) downloadFile(url string) (resp *http.Response, err error) {

	resp, err = http.Get(url)

	return
}

func getAudioFormat(f AudioFormat) (format string) {

	switch f {

	case MP3:
		format = "mp3"
	case WAV:
		format = "wav"
	case MP4:
		format = "mp4"
	default:
		format = "mp3"
	}
	return
}
