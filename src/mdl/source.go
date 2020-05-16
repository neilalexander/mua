package mdl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/matrix-org/gomatrix"
)

type SourceType byte
type SourceCode []byte

const (
	SourceTypeLua SourceType = iota
)

type Source struct {
	Type   SourceType `json:"type"`
	Source SourceCode `json:"source"`
}

func NewSourceFromJSON(j []byte) (*Source, error) {
	src := &Source{}
	if err := json.Unmarshal(j, src); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return src, nil
}

func NewSourceFromEvent(event gomatrix.Event) (*Source, error) {
	if event.Type != EVENT_TYPE {
		return nil, errors.New("event is incorrect type")
	}
	if event.Content == nil {
		return nil, errors.New("event has no content")
	}
	j, err := json.Marshal(event.Content)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}
	src := &Source{}
	if err = json.Unmarshal(j, src); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}
	return src, nil
}

func (b SourceCode) MarshalJSON() ([]byte, error) {
	lines := bytes.Split(b, []byte("\n"))
	for l := range lines {
		lines[l] = bytes.Trim(lines[l], "\r\n\t ")
	}
	byt := bytes.Join(lines, []byte("\n"))
	b64 := base64.RawStdEncoding.EncodeToString(byt)
	b64 = fmt.Sprintf("\"%s\"", b64)
	return []byte(b64), nil
}

func (b *SourceCode) UnmarshalJSON(in []byte) error {
	in = in[1 : len(in)-1]
	if byt, err := base64.RawStdEncoding.DecodeString(string(in)); err == nil {
		*b = SourceCode(byt)
		return nil
	} else {
		return err
	}
}
