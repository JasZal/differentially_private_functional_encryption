package main

import (
	"time"

	"github.com/JasZal/gofe/data"
	"github.com/JasZal/gofe/innerprod/noisy"
)

type OTNHUser struct {
	vecLen    int
	x         data.Vector
	OTNHmulti *noisy.OTNHMultiIPE
	secretKey data.Matrix
}

func NewOTNHUser(vecL int, d data.Vector, sk data.Matrix, fh *noisy.OTNHMultiIPE) *OTNHUser {
	u := &OTNHUser{
		vecLen:    vecL,
		x:         d,
		OTNHmulti: fh,
		secretKey: sk,
	}

	return u
}

func (u OTNHUser) encryptOTNH(pos int, a *OTNHAuthority) (data.VectorG1, time.Duration) {

	client := noisy.NewOTNHMultiIPEFromParams(u.OTNHmulti.Params)
	start := time.Now()
	c, _ := client.Encrypt(u.x, u.secretKey)
	timeEnc := time.Since(start)
	return c, timeEnc

}
