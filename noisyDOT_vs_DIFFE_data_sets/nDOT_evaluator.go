package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/fentec-project/bn256"
	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/fullysec"
)

type NHEvaluator struct {
	attr      int
	numClient int
	y         data.Matrix
	pubKey    *bn256.GT
	nhmulti   *fullysec.NHMultiIPE
}

func NewNHEvaluator(a int, numC int, pk *bn256.GT, fh *fullysec.NHMultiIPE) *NHEvaluator {
	e := &NHEvaluator{
		attr:      a,
		numClient: numC,
		pubKey:    pk,
		nhmulti:   fh,
	}

	return e
}

func (e NHEvaluator) generateY() data.Matrix {
	// generate inner product vectors and put them in a matrix
	y := make(data.Matrix, e.numClient)
	for i := 0; i < e.numClient; i++ {

		y[i] = data.NewConstantVector(e.attr, big.NewInt(0))
		y[i][0] = big.NewInt(1)

	}
	return y
}

func (e *NHEvaluator) evaluateNH(a *NHAuthority, cipher data.MatrixG1, eps float64) (*big.Int, time.Duration, time.Duration, error) {
	start := time.Now()
	funcKey, err := a.generateFunctionKeyNH(e.y, eps)
	timeKeyGen := time.Since(start)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}

	decryptor := fullysec.NewNHMultiIPEFromParams(e.nhmulti.Params)
	start = time.Now()
	xy, err := decryptor.Decrypt(cipher, funcKey, e.pubKey)
	timeEval := time.Since(start)
	if err != nil {
		fmt.Println(err)
	}

	return xy, timeKeyGen, timeEval, err
}
