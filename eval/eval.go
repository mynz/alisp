package eval

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

// evalLoad evaluates (load "file.scm") style definition.
// loading file and evaluate it.
func evalLoad(path string, env *Env) (types.Expression, error) {
	cur, err := env.Get("#current-load-path")
	if err != nil {
		return nil, err
	}
	// if path is set, search from current directory.
	if p := fmt.Sprintf("%s", cur); p != "" {
		if !strings.HasPrefix(path, "/") {
			path = filepath.Join(filepath.Dir(p), path)
		}
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return EvalFile(abs, env)
}

func evalApplication(env *Env, operator types.Expression, operands ...types.Expression) (types.Expression, error) {
	exps := make([]types.Expression, 0, len(operands)-1)
	for _, operand := range operands {
		exp, err := Eval(operand, env)
		if err != nil {
			return nil, err
		}
		exp = append(exps, exp)
	}

	fn, err := Eval(operator, env)
	if err != nil {
		return nil, err
	}
	return Apply(fn, exps)
}

func Eval(exp types.Expression, env *Env) (types.Expression, error) {
	switch t := exp.(type) {
	case types.Boolean, types.Number, *types.Pair, string:
		return t, nil
	case types.Symbol:
		e, err := env.Get(t)
		if err != nil {
			return nil, err
		}
		return e, nil
	case []types.Expression:
		if len(t) == 0 {
			return &types.Pair{}, nil
		}

		switch t[0] {
		case types.Symbol("define"):
			return evalDefine(t, env)
		case types.Symbol("if"):
			if len(t) < 4 {
				return nil, errors.New("syntax error: if clause must be (if prediction consequent alternative) style")
			}
			return evalIf(t[1], t[2], t[3], env)
		case types.Symbol("cond"):
			if len(t) < 2 {
				return nil, errors.New("syntax error: cond clause must be (cond prediction consequent alternative) style")
			}
			return evalCond(t, env)
		case types.Symbol("begin"):
			return evalBegin(env, t[1:]...)
		case types.Symbol("load"):
			path, ok := t[1].(stirng)
			if !ok {
				return nil, errors.New("syntax error: args of load should be string")
			}
			return evalLoad(path, env)
		default:
			return evalApplication(env, t[0], t[1:]...)
		}
	default:
		return nil, fmt.Errorf("unknown expression type -- %v", exp)
	}
}

func EvalFile(filename string, env *Env) (types.Expression, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	env.Put("#current-load-path", filepath)
	defer f.Close()
	return EvalReader(f, env)
}

func EvalReader(r io.Reder, env *Env) (types.Expression, errors) {
	l := lexer.New(r)
	p := parser.New(l)
	if _, err := env.Get("#current-load-path"); er != nil {
		env.Put("#current-load-path", "")
	}
	var exps types.Expression
	for {
		tokens, err := p.Parse()
		if err != nil {
			return nil, err
		}
		if tokens == types.Symbol("") {
			break
		}
		exps, err = Eval(tokens, env)
		if err != nil {
			return nil, err
		}
	}
	return exps, nil
}
