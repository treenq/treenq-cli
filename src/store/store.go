package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/treenq/treenq-cli/src/models"
)

func NewStore(w io.ReadWriteCloser) (*Store, func() error, error) {
	var data []models.Context
	if err := json.NewDecoder(w).Decode(&data); err != nil && !errors.Is(err, io.EOF) {
		return nil, nil, fmt.Errorf("failed to decode config: %w", err)
	}

	s := &Store{data: data, w: w}

	close := func() error {
		list := s.GetContexts()
		defer w.Close()
		if err := json.NewEncoder(w).Encode(list); err != nil {
			return err
		}

		return nil
	}

	return s, close, nil
}

type Store struct {
	data []models.Context
	w    io.ReadWriter
}

func (s *Store) SetActiveContext(name string) error {
	// declare found flag,
	// if the given context not found - returns error
	var found bool

	// make all the existing contexts inactive except of the "name"
	for i := range s.data {
		active := s.data[i].Name == name
		s.data[i].Active = active
		if active {
			found = true
		}
	}

	if !found {
		return models.ErrContextNotFound
	}
	return nil
}

func (s *Store) GetContexts() []models.Context {
	return s.data[:]
}

func (s *Store) NewContext(ctx models.Context) error {
	for i := range s.data {
		// make all the existing contexts inactive
		s.data[i].Active = false

		// validate unique name
		if s.data[i].Name == ctx.Name {
			return models.ErrContextAlreadyExists
		}
	}

	// insert new context
	ctx.Active = true
	s.data = append(s.data, ctx)
	return nil
}
