package pre

import (
	"bytes"
	"encoding/gob"
	"labda/eval"
	"labda/std"
)

func Encode(expr eval.Expr) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(&expr)
	return buffer.Bytes(), err
}

func Decode(data []byte) (eval.Expr, error) {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	var expr eval.Expr
	err := decoder.Decode(&expr)
	return expr, err
}

func init() {
	gob.Register(&eval.Abstraction{})
	gob.Register(&eval.PathedAbstraction{})
	gob.Register(&eval.Application{})
	gob.Register(&eval.Variable{})
	gob.Register(&eval.StringLit{})
	gob.Register(&eval.NumberLit{})
	gob.Register(std.BuiltinExpr{})
}
