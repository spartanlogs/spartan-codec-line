package codecs

import (
	"io"

	"github.com/spartanlogs/spartan/codecs"
	"github.com/spartanlogs/spartan/config"
	"github.com/spartanlogs/spartan/event"
	"github.com/spartanlogs/spartan/utils"
)

type lineConfig struct {
	delimiter string
}

var lineConfigSchema = []config.Setting{
	{
		Name:    "delimiter",
		Type:    config.String,
		Default: "\n",
	},
}

// The LineCodec reads plaintext with delimiter of \n
type LineCodec struct {
	config *lineConfig
}

func init() {
	codecs.Register("line", newLineCodec)
}

func newLineCodec(options utils.InterfaceMap) (codecs.Codec, error) {
	c := &LineCodec{
		config: &lineConfig{},
	}
	return c, c.setConfig(options)
}

func (c *LineCodec) setConfig(options utils.InterfaceMap) error {
	if err := config.VerifySettings(options, lineConfigSchema); err != nil {
		return err
	}

	c.config.delimiter = options.Get("delimiter").(string)

	return nil
}

// Encode event as a simple message.
func (c *LineCodec) Encode(e *event.Event) []byte {
	return []byte(e.String() + c.config.delimiter)
}

// EncodeWriter reads events from in and writes them to w
func (c *LineCodec) EncodeWriter(w io.Writer, in <-chan *event.Event) {}

// Decode creates a new event with message set to data.
func (c *LineCodec) Decode(data []byte) (*event.Event, error) {
	return event.New(string(data)), nil
}

// DecodeReader reads from r and creates an event sent to out
func (c *LineCodec) DecodeReader(r io.Reader, out chan<- *event.Event) {}
