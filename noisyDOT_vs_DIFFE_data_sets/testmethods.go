package main

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/fentec-project/gofe/data"
)

func debug(s string, deb bool) {
	if deb {
		fmt.Printf(s)
	}
}

func testNH(filename string, deb bool, epsilon float64, scaling int64, secLevel int, boundX *big.Int, boundY *big.Int) (error, int, time.Duration, time.Duration, time.Duration, time.Duration) {

	count, numClient := computeTrueCount(filename)
	//read file
	x, _ := readFile(filename, scaling)
	attributes := len(x[0])

	debug(fmt.Sprintf("count: %d, numClient: %d, attributes: %d\n", count, numClient, attributes), deb)
	debug(fmt.Sprintf("true xy: %d, true result: %.6f \n", count, float64(count)*100/float64(numClient)), deb)

	//prepare database
	cipher := make(data.MatrixG1, numClient)

	//generate Authority with Params and Keys
	a, timeSetup := NewNHAuthority(secLevel, attributes, numClient, boundX, boundY, epsilon, scaling)
	debug("generatet NH authority\n", deb)

	//generate numClient Clients
	clients := make([]*NHUser, numClient)
	for i := 0; i < numClient; i++ {
		clients[i] = NewNHUser(attributes, x[i], a.getNHSecretKey(i), a.getNHPublicParams())
	}

	var timeEnc time.Duration
	//fill cipher
	for i := 0; i < numClient; i++ {
		if i == 0 {
			cipher[i], timeEnc = clients[i].encryptNH(i, a)
		}
		cipher[i], _ = clients[i].encryptNH(i, a)
	}
	debug("filled encrypted database\n", deb)

	//generate evaluator, start decrypting
	evaluator := NewNHEvaluator(attributes, numClient, a.pubKey, a.getNHPublicParams())
	evaluator.y = evaluator.generateY()

	debug("start evaluation\n", deb)

	xy, timeKeyGen, timeEval, err := evaluator.evaluateNH(a, cipher, epsilon)
	if err != nil {
		return err, numClient, 0, 0, 0, 0
	}

	var r float64 = float64(xy.Int64())
	r = (r / float64(scaling))

	debug(fmt.Sprintf("time in sec: keygen %f, eval %f\n", timeKeyGen.Seconds(), timeEval.Seconds()), deb)
	debug(fmt.Sprintf("scaling: %d, xy: %d; result: %.6f ; percent: %.6f\n", scaling, xy, r, r*100/float64(numClient)), deb)
	debug("---------------finished---------------\n", deb)

	return nil, numClient, timeSetup, timeEnc, timeKeyGen, timeEval
}

func testOTNH(filename string, deb bool, epsilon float64, scaling int64, secLevel int, boundX *big.Int, boundY *big.Int) (error, int, time.Duration, time.Duration, time.Duration, time.Duration) {

	count, numClient := computeTrueCount(filename)
	//read file
	x, _ := readFile(filename, scaling)
	attributes := len(x[0])

	debug(fmt.Sprintf("count: %d, numClient: %d, attributes: %d\n", count, numClient, attributes), deb)
	debug(fmt.Sprintf("true xy: %d, true result: %.6f \n", count, float64(count)*100/float64(numClient)), deb)

	//prepare database
	cipher := make(data.MatrixG1, numClient)

	//generate Authority with Params and Keys
	a, timeSetup := NewOTNHAuthority(secLevel, attributes, numClient, boundX, boundY, epsilon, scaling)
	debug("generatet OTNH authority\n", deb)

	//generate numClient Clients
	clients := make([]*OTNHUser, numClient)
	for i := 0; i < numClient; i++ {
		clients[i] = NewOTNHUser(attributes, x[i], a.getOTNHSecretKey(i), a.getOTNHPublicParams())
	}

	var timeEnc time.Duration
	//fill cipher
	for i := 0; i < numClient; i++ {
		if i == 0 {
			cipher[i], timeEnc = clients[i].encryptOTNH(i, a)
		}
		cipher[i], _ = clients[i].encryptOTNH(i, a)
	}
	debug("filled encrypted database\n", deb)

	//generate evaluator, start decrypting
	evaluator := NewOTNHEvaluator(attributes, numClient, a.pubKey, a.getOTNHPublicParams())
	evaluator.y = evaluator.generateY()

	debug("start evaluation\n", deb)

	xy, timeKeyGen, timeEval, err := evaluator.evaluateOTNH(a, cipher, epsilon)
	if err != nil {
		return err, numClient, 0, 0, 0, 0
	}

	var r float64 = float64(xy.Int64())
	r = (r / float64(scaling))

	debug(fmt.Sprintf("time in sec: keygen %f, eval %f\n", timeKeyGen.Seconds(), timeEval.Seconds()), deb)
	debug(fmt.Sprintf("scaling: %d, xy: %d; result: %.6f ; percent: %.6f\n", scaling, xy, r, r*100/float64(numClient)), deb)
	debug("---------------finished---------------\n", deb)

	return nil, numClient, timeSetup, timeEnc, timeKeyGen, timeEval
}

func computeTrueCount(filename string) (int64, int) {
	count := 0

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for i := 0; i < len(records); i++ {
		h, _ := strconv.Atoi(records[i][0])
		count += h
	}

	return int64(count), len(records)

}

func readFile(filename string, scaling int64) ([]data.Vector, error) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	//big.NewInt(int64(1))
	var m = make([]data.Vector, len(records))

	var tmp []*big.Int

	for i := 0; i < len(records); i++ {
		r := records[i]
		tmp = make([]*big.Int, len(r))

		for j := 0; j < len(tmp); j++ {
			t, _ := strconv.Atoi(r[j])
			tmp[j] = big.NewInt(int64(t) * scaling)
		}
		m[i] = data.NewVector(tmp)
	}

	return m, err
}
