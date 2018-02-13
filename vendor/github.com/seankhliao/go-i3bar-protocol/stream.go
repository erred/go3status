package protocol

import (
	"encoding/json"
	"io"
)

// Encoder is a thin wrapper around json.Encoder
// Provides a BeginArray() to write a single [
// to start the infinite array
type Encoder struct {
	w   io.Writer
	enc *json.Encoder
}

// NewEncoder creates a new Encoder
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w, enc: json.NewEncoder(w)}
}

// BeginArray writes a single [
func (e *Encoder) BeginArray() error {
	_, err := e.w.Write([]byte("["))
	return err
}

// Encode passes through to json.Encode
func (e *Encoder) Encode(v interface{}) error {
	return e.enc.Encode(v)
}
