package msgpack

import (
	"github.com/pkg/errors"
	"github.com/poboisvert/hexagonal_golang_url/shortener"
	"github.com/vmihailenco/msgpack"
)

type Redirect struct{}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer err := Redirect.Decode")
	}
	return redirect, nil
}

func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer err := Redirect.Encode")
	}
	return rawMsg, nil
}