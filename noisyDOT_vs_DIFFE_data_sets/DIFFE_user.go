package main

import (
	"time"

	"github.com/fentec-project/gofe/data"
	"github.com/fentec-project/gofe/innerprod/fullysec"
)

type OTNHUser struct {
	vecLen    int
	x         data.Vector
	OTNHmulti *fullysec.OTNHMultiIPE
	secretKey data.Matrix
}

func NewOTNHUser(vecL int, d data.Vector, sk data.Matrix, fh *fullysec.OTNHMultiIPE) *OTNHUser {
	u := &OTNHUser{
		vecLen:    vecL,
		x:         d,
		OTNHmulti: fh,
		secretKey: sk,
	}

	return u
}

func (u OTNHUser) encryptOTNH(pos int, a *OTNHAuthority) (data.VectorG1, time.Duration) {

	client := fullysec.NewOTNHMultiIPEFromParams(u.OTNHmulti.Params)
	start := time.Now()
	c, _ := client.Encrypt(u.x, u.secretKey)
	timeEnc := time.Since(start)
	return c, timeEnc

}
