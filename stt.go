package stt

import (
	"encoding/binary"
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"github.com/gordonklaus/portaudio"
)

// Transcriber transcribes an audio file
type Transcriber interface {
	Transcribe(path string) string
}

// RecordAndTranscribe records audio for Duration and uses the text to speech engine
// defined to format the speech audio into a string.
func RecordAndTranscribe(t Transcriber, duration time.Duration) string {
	timer := time.NewTimer(time.Second * duration)
	f, _ := ioutil.TempFile(os.TempDir(), "pa")
	defer f.Close()
	defer os.Remove(f.Name())
	defer timer.Stop()

	record(f, timer.C)

	exec.Command("afplay", f.Name()).Run()

	return t.Transcribe(f.Name())
}

// Slightly modified exact copy of the portaudio example
func record(f *os.File, timer <-chan time.Time) {
	// form chunk
	_, _ = f.WriteString("FORM")
	binary.Write(f, binary.BigEndian, int32(0))
	_, _ = f.WriteString("AIFF")

	// common chunk
	_, _ = f.WriteString("COMM")
	binary.Write(f, binary.BigEndian, int32(18))
	binary.Write(f, binary.BigEndian, int16(1))
	binary.Write(f, binary.BigEndian, int32(0))
	binary.Write(f, binary.BigEndian, int16(32))
	_, _ = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100

	// sound chunk
	_, _ = f.WriteString("SSND")
	binary.Write(f, binary.BigEndian, int32(0))
	binary.Write(f, binary.BigEndian, int32(0))
	binary.Write(f, binary.BigEndian, int32(0))
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, _ = f.Seek(4, 0)
		binary.Write(f, binary.BigEndian, int32(totalBytes))
		_, _ = f.Seek(22, 0)
		binary.Write(f, binary.BigEndian, int32(nSamples))
		_, _ = f.Seek(42, 0)
		binary.Write(f, binary.BigEndian, int32(4*nSamples+8))
		f.Close()
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	stream, _ := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	defer stream.Close()

	stream.Start()
	defer stream.Stop()

	for {
		stream.Read()
		binary.Write(f, binary.BigEndian, in)
		nSamples += len(in)
		select {
		case <-timer:
			return
		default:
		}
	}
}
