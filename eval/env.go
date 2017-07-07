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

func Cons(args ...types.Expression) (types.Expression, error) {
	return &types.Pair{Car: args[0], Cdr: args[1]}, nil
}

func Print(args ...types.Expression) (types.Expression, error) {
	if len(args) == 1 {
		fmt.Println(args[0])
	} else {
		fmt.Println(args)
	}
	return nil, nil
}

func Add(args ...types.Expression) (types.Expression, error) {
	sum, ok := args[0].(types.Number)
	if !ok {
		return nil, fmt.Errorf("given args is not number: %#v", args[0])
	}
	for _, adder := range args[1:] {
		sum = sum + adder.(types.Number)
	}
	return sum, nil
}

func Subtract(args ...types.Expression) (types.Expression, error) {
	sub, ok := args[0].(types.Number)
	if !ok {
		return nil, fmt.Errorf("given args is not number: %#v", args[0])
	}
	for _, s := range args[1:] {
		sub = sub - s.(types.Number)
	}
	return sub, nil
}

func Multiply(args ...types.Expression) (types.Expression, error) {
	mul, ok := args[0].(types.Number)
	if !ok {
		return nil, fmt.Errorf("given args is not number: %#v", args[0])
	}
	for _, m := range args[1:] {
		mul = mul * m.(types.Number)
	}
	return mul, nil
}

func Divide(args ...types.Expression) (types.Expression, error) {
	div, ok := args[0].(types.Number)
	if !ok {
		return nil, fmt.Errorf("given args is not number: %#v", args[0])
	}
	for _, d := range args[1:] {
		div = div / d.(types.Number)
	}
	return div, nil
}
