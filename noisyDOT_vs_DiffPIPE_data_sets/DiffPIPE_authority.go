package main

import (
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/JasZal/gofe/data"
	"github.com/JasZal/gofe/innerprod/noisy"
	"github.com/fentec-project/bn256"
	"github.com/google/differential-privacy/go/v2/noise"
)

type OTNHAuthority struct {
	secLevel     int
	vecLen       int
	numClient    int
	boundX       *big.Int
	boundY       *big.Int
	pubKey       *bn256.GT
	masterSecKey *noisy.OTNHMultiIPESecKey
	OTNHmulti    *noisy.OTNHMultiIPE
	epsilon      float64
	scaling      int64
}

func NewOTNHAuthority(secL int, vecL int, numC int, bX, bY *big.Int, e float64, scal int64) (*OTNHAuthority, time.Duration) {
	a := &OTNHAuthority{
		secLevel:  secL,
		vecLen:    vecL,
		numClient: numC,
		boundX:    bX,
		boundY:    bY,
		epsilon:   e,
		scaling:   scal,
	}

	start := time.Now()
	a.OTNHmulti = noisy.NewOTNHMultiIPE(a.secLevel, a.numClient, a.vecLen, a.boundX, a.boundY)
	timeSetup := time.Since(start)
	a.masterSecKey, a.pubKey, _ = a.OTNHmulti.GenerateKeys()

	return a, timeSetup
}

// TODO
func computeL0SensitivityOTNH(y data.Matrix) int64 {
	return 1.0
}

// TODO
func computeLInfSensitivityOTNH(y data.Matrix) float64 {
	return 1.0
}

func (a *OTNHAuthority) generateFunctionKeyOTNH(y data.Matrix, eps float64) (data.MatrixG2, error) {
	//check if key is permitted

	//check privacy budget

	//compute specification for noise (sensitivity)
	l0 := computeL0SensitivityOTNH(y)
	lInf := computeLInfSensitivityOTNH(y)

	//noise via gauss (or laplace)
	lap := noise.Laplace()
	noise, err := lap.AddNoiseFloat64(0, l0, lInf, eps, 0) //lap.AddNoiseFloat64(0.0, l0, lInf, eps, del)

	if err != nil {
		fmt.Println(err)
	}

	//scale noise
	noise *= float64(a.scaling)

	if noise >= 0 {
		noise = math.Ceil(noise)
	} else {
		noise = math.Floor(noise)
	}

	// derive a functional key for vector y
	key, err := a.OTNHmulti.DeriveKey(y, a.masterSecKey, int64(noise))
	if err != nil {
		fmt.Println("Error during key derivation:", err)
	}

	return key, nil
}

func (a OTNHAuthority) getOTNHSecretKey(pos int) data.Matrix {
	return a.masterSecKey.BHat[pos]
}

func (a OTNHAuthority) getOTNHPublicParams() *noisy.OTNHMultiIPE {
	return a.OTNHmulti
}
