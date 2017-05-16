package eval

import (
	"errors"
	"fmt"

	"github.com/mynz/alisp/types"
)

// evalDefine evaluate (define ...) style expression.
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
		caddr, ok := tt[0].(types.Symbol)
		if !ok {
			return nil, errors.New("(define x) of  x should be symbol")
		}
		// create lambda and put it into environment.
		env.Put(caddr, Lambda{tt[1:], exps[2], env})
		return nil, nil

	default:
		return nil, nil
	}
}

func evalIf(predicate, consequent, alternative types.Expression, env *Env) (types.Expression, error) {
	bb, err := evalPredicate(predicate, env)
	if err != nil {
		return nil, err
	}
	if bb {
		return Eval(consequent, env)
	}
	return Eval(alternative, env)
}

func evalPredicate(exp types.Expression, env *Env) (types.Boolean, error) {
	b, err := Eval(exp, env)
	if err != nil {
		return false, err
	}
	bb, ok := b.(types.Boolean)
	if !ok {
		return false, fmt.Errorf("the expression should return types.Boolean, exps: %v", exp)
	}
	return bb, nil
}

func evalBegin(env *Env, exps ...types.Expression) (types.Expression, error) {
	var lastExp types.Expression
	for _, beginExp := range exps {
		l, err := Eval(beginExp, env)
		if err != nil {
			return nil, err
		}
		lastExp = l
	}
	return lastExp, nil
}
