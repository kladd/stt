# stt - Speech to text in Go

A speech to text library using portaudio to record and CMU PocketSphinx to transcribe.

### Example

The example below uses the CMU PocketSphinx driver provided by the included [psphinx](./engines/psphinx) package. See the [Engines](#engines) section for available speech to text engines.

```go
import (
	"github.com/kladd/stt"
	"github.com/kladd/stt/engines/psphinx"
)

func main() {
	transcriber := new(psphinx.Transcriber)

	// Records for 10 seconds before transcribing
	fmt.Println(stt.RecordAndTranscribe(transcriber, 10))
}
```

### Engines

* CMU PocketSphinx - [pshinx](./engines/pshinx)

### Caveats & Dependencies

Dependencies:

* portaudio
* cmu-sphinxbase
* cmu-pocketsphinx

Caveats:

* Not very accurate right now

### Installation

```bash
go get github.com/kladd/stt
```

### License

See [LICENSE](./LICENSE)

### Contributing

Appreciated. Fork, branch, commit, push, pull request.
