package psphinx

import "github.com/kladd/pocketsphinx"

// Transcriber is the PocketSphinx implementation
type Transcriber struct{}

// Transcribe uses PocketSphinx to transcribe audio file f
func (t *Transcriber) Transcribe(path string) string {
	return pocketsphinx.TranscribeFile(path)
}
