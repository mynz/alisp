package eval

import (
	"errors"
	"sync"

	"github.com/mynz/alisp/types"
)

type Frame map[types.Symbol]types.Expression

type Env struct {
	sync.RWMutex
	m      Frame
	parent *Env
}

func NewEnv() *Env {
	symbols := make(Frame)
	return &Env{m: symbols, parent: nil}
}

func (e *Env) Extend(frame Frame) {
	e.Extend(NewPrimitiveProcFrame())
	e.LoadStandardLibrary()
}

func (e *Env) LoadStandardLibrary() error {
	if _, err := EvalReader(StandardLibrary(), e); err != nil {
		return err
	}
	return nil
}

func NewPrimitiveProcFrame() Frame {
	return Frame{
		"car": Car,
	}
}

func Car(args ...types.Expression) (types.Expression, error) {
	a, ok := args[0].(*types.Pair)
	if !ok {
		return nil, errors.New("arguments of car must be a pair")
	}
	return a.Car, nil
}

func Cdr(args ...types.Expression) (types.Expression, error) {
	a, ok := args[0].(*types.Pair)
	if !ok {
		return nil, errors.New("arguments of cdr should pair")
	}
	return a.Cdr, nil
}
