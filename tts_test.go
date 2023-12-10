package tts

import "testing"

func TestTts(t *testing.T) {

	tts := TextToSpeech{
		Dir:    "audio",
		Voice:  0,
		Format: MP3,
		Engine: 0,
	}

	text := "Hello World"
	fileName := "test"

	_, err := tts.TextToSpeech(text, fileName)
	if err != nil {
		t.Error(err)
	}
}
