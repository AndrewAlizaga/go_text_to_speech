package tts

import (
	"errors"
	"testing"
)

func TestTts(t *testing.T) {

	var err error

	tts := TextToSpeech{
		Dir:    "audio",
		Voice:  0,
		Format: MP3,
		Engine: 0,
	}

	text := "Hello World"
	fileName := "test"

	t.Log("test 1 - basic tts")
	_, err = tts.TextToSpeech(text, fileName)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log("test 1 - basic tts - passed")

	t.Log("test 2 - nameles tts")
	dir, err := tts.TextToSpeech(text, "")
	if dir == "" {
		err = errors.New("default name not created")
	}
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log("test 2 - nameles tts - passed")

}
