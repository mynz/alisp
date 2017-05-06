package types

import (
	"fmt"
	"strings"
)

type Expression interface{}

type Number float64
type Symbol string
type Boolean bool

func (b Boolean) String() string {
	if b {
		return "#t"
	}
	return "#f"
}

type Pair struct {
	Car Expression
	Cdr Expression
}

func (p *Pair) String() string {
	if p.IsNull() {
		return "()"
	}
	if p.IsList() {
		tokens := []string{}
		pp := p
		for {
			if pp.IsNull() {
				break
			}
			tokens = append(tokens, fmt.Sprintf("%v", pp.Car))
			// TODO
			switch cdr := pp.Cdr.(type) {
			case *Pair:
				pp = cdr
			default:
				break
			}
		}
		return fmt.Sprintf("(%s)", strings.Join(tokens, " "))
	}
	return fmt.Sprintf("(%v . %v)", p.Car, p.Cdr)
}

func (p *Pair) IsNull() bool {
	return p.Car == nil && p.Cdr == nil
}

func (p *Pair) IsList() bool {
	pp := p
	for {
		if pp.IsNull() {
			return true
		}
		switch cdr := pp.Cdr.(type) {
		case *Pair:
			pp = cdr
		default:
			return false
		}
	}
}

func (p *Pair) Append(exp Expression) *Pair {
	pp := p
	for {
		if pp.IsNull() {
			break
		}
		pp = pp.Cdr.(*Pair)
	}
	pp.Car = exp
	pp.Cdr = &Pair{}
	return pp
}

func NewList(args ...Expression) *Pair {
	// In normal, p is prefer to be defined by var statement because of no allocation.
	// But in this case, for empty list should be return empty pair.
	p := &Pair{}
	for i := len(args) - 1; i >= 0; i-- {
		p = &Pair{args[i], p}
	}
	return p
}
