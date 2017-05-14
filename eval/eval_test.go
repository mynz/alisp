package eval

import (
	"errors"

	"github.com/mynz/alisp/types"
)

func evalDefine(exps []types.Expression, env *Env) (types.Expression, error) {
	if len(exps) < 2 {
		return nil, errors.New("define clause must have symbol and body")
	}

	switch tt := exps[1].(type) {
	// put symbol and variables.
	// (define x 1) style definition.
	case types.Symbol:
		value, err := Eval(exps[2], env)
		if err != nil {
			return nil, err
		}
		env.Put(tt, value)
		return nil, nil

		// (define (hoge args) (..)) style definition.
		// above style is syntax sugar for lambda.
	case []types.Expression:
		if len(tt) < 2 {
			return nil, errors.New("define statement must have more than 2 words")
		}
		caddr, of := tt[0].(types.Symbol)
		if !ok {
			return nil, errors.New("(define x) of  x should be symbol")
		}
		// create lambda and put it into environment.
		env.Put(caddr, Lamb{tt[1:], exps[2], env})
		return nil, nil

	default:
		return nil, nil
	}
}
