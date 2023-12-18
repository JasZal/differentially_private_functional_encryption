package main

import (
	"time"

	"github.com/JasZal/gofe/data"
	"github.com/JasZal/gofe/innerprod/noisy"
)

type NHUser struct {
	vecLen    int
	x         data.Vector
	nhmulti   *noisy.NHMultiIPE
	secretKey data.Matrix
}

func NewNHUser(vecL int, d data.Vector, sk data.Matrix, fh *noisy.NHMultiIPE) *NHUser {
	u := &NHUser{
		vecLen:    vecL,
		x:         d,
		nhmulti:   fh,
		secretKey: sk,
	}

	return u
}

func (u NHUser) encryptNH(pos int, a *NHAuthority) (data.VectorG1, time.Duration) {

	client := noisy.NewNHMultiIPEFromParams(u.nhmulti.Params)
	start := time.Now()
	c, _ := client.Encrypt(u.x, u.secretKey)
	timeEnc := time.Since(start)
	return c, timeEnc

}
