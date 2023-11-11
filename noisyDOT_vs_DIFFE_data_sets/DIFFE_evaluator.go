package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/fentec-project/bn256"
	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/fullysec"
)

type OTNHEvaluator struct {
	attr      int
	numClient int
	y         data.Matrix
	pubKey    *bn256.GT
	OTNHmulti *fullysec.OTNHMultiIPE
}

func NewOTNHEvaluator(a int, numC int, pk *bn256.GT, fh *fullysec.OTNHMultiIPE) *OTNHEvaluator {
	e := &OTNHEvaluator{
		attr:      a,
		numClient: numC,
		pubKey:    pk,
		OTNHmulti: fh,
	}

	return e
}

func (e OTNHEvaluator) generateY() data.Matrix {
	// generate inner product vectors and put them in a matrix
	y := make(data.Matrix, e.numClient)
	for i := 0; i < e.numClient; i++ {

		y[i] = data.NewConstantVector(e.attr, big.NewInt(0))
		y[i][0] = big.NewInt(1)

	}
	return y
}

func (e *OTNHEvaluator) evaluateOTNH(a *OTNHAuthority, cipher data.MatrixG1, eps float64) (*big.Int, time.Duration, time.Duration, error) {
	start := time.Now()
	funcKey, err := a.generateFunctionKeyOTNH(e.y, eps)
	timeKeyGen := time.Since(start)
	if err != nil {
		fmt.Println(err)
		return nil, 0, 0, err
	}

	decryptor := fullysec.NewOTNHMultiIPEFromParams(e.OTNHmulti.Params)
	start = time.Now()
	xy, err := decryptor.Decrypt(cipher, funcKey, e.pubKey)
	timeEval := time.Since(start)
	if err != nil {
		fmt.Println(err)
	}

	return xy, timeKeyGen, timeEval, err
}
