package main

import (
	"time"

	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/fullysec"
)

type NHUser struct {
	vecLen    int
	x         data.Vector
	nhmulti   *fullysec.NHMultiIPE
	secretKey data.Matrix
}

func NewNHUser(vecL int, d data.Vector, sk data.Matrix, fh *fullysec.NHMultiIPE) *NHUser {
	u := &NHUser{
		vecLen:    vecL,
		x:         d,
		nhmulti:   fh,
		secretKey: sk,
	}

	return u
}

func (u NHUser) encryptNH(pos int, a *NHAuthority) (data.VectorG1, time.Duration) {

	client := fullysec.NewNHMultiIPEFromParams(u.nhmulti.Params)
	start := time.Now()
	c, _ := client.Encrypt(u.x, u.secretKey)
	timeEnc := time.Since(start)
	return c, timeEnc

}
