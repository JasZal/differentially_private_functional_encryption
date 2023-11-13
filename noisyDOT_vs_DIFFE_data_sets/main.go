package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/JasZal/gofe/data"
)

func main() {
	//params
	secLevel := 2
	epsilon := 1.0
	debug := true
	append := false
	//variables
	var scaling int64 = 100000
	rounds := 1

	test_realData(secLevel, rounds, scaling, epsilon, debug, append)

}

func test_realData(secLevel, rounds int, scaling int64, epsilon float64, debug, append bool) {
	dbfilename := []string{"datasets/lbw.csv", "datasets/pcs.csv", "datasets/uis.csv", "datasets/nhanes3.csv"}
	//the value by which the coordinates of encrypted vectors are bounded, for each dataset
	bounds := []int64{5000, 150, 60, 750} //
	//nahnes, pcs minscaling = 10,
	file := "resultDatasets_noisyDOT_vs_DIFFE.txt"

	if !append {
		write(file, fmt.Sprint("scaling=", scaling, ", avTime in Setup=Nanoseconds, Enc=Microseconds, KeyGen=Eval=Milliseconds \n"), append)
		append = true
		write(file, "Setup, Encryption, KeyGen, Evaluation\n", append)
	}
	if debug {
		fmt.Println("Setup, Encryption, KeyGen, Evaluation")
	}

	for f, filename := range dbfilename {

		if debug {
			fmt.Println(filename)
		}

		//params
		failedAttemptsNH := 0
		failedAttemptsOTNH := 0
		avTimeKeyGenNH := 0.0
		avTimeKeyGenOTNH := 0.0
		avTimeEvalNH := 0.0
		avTimeEvalOTNH := 0.0
		avTimeSetupNH := 0.0
		avTimeSetupOTNH := 0.0
		avTimeEncNH := 0.0
		avTimeEncOTNH := 0.0
		boundX := big.NewInt(bounds[f] * scaling)
		boundY := big.NewInt(1)

		for i := 0; i < rounds; i++ {
			errNH, _, timeSetupNH, timeEncNH, timeKeyGenNH, timeEvalNH := testNH(dbfilename[f], false, epsilon, scaling, secLevel, boundX, boundY)
			errOTNH, _, timeSetupOTNH, timeEncOTNH, timeKeyGenOTNH, timeEvalOTNH := testOTNH(dbfilename[f], false, epsilon, scaling, secLevel, boundX, boundY)

			if debug {
				fmt.Println(" NH - round ", i+1, ": ", timeSetupNH.Nanoseconds(), ", ", timeEncNH.Microseconds(), ", ", timeKeyGenNH.Milliseconds(), ", ", timeEvalNH.Milliseconds())
				fmt.Println(" OT-NH - round ", i+1, ": ", timeSetupOTNH.Nanoseconds(), ", ", timeEncOTNH.Microseconds(), ", ", timeKeyGenOTNH.Milliseconds(), ", ", timeEvalOTNH.Milliseconds())
			}

			if errNH == nil {
				avTimeSetupNH += float64(timeSetupNH.Nanoseconds())
				avTimeEncNH += float64(timeEncNH.Microseconds())
				avTimeKeyGenNH += float64(timeKeyGenNH.Milliseconds())
				avTimeEvalNH += float64(timeEvalNH.Milliseconds())
			} else {
				failedAttemptsNH += 1
			}

			if errOTNH == nil {
				avTimeSetupOTNH += float64(timeSetupOTNH.Nanoseconds())
				avTimeEncOTNH += float64(timeEncOTNH.Microseconds())
				avTimeKeyGenOTNH += float64(timeKeyGenOTNH.Milliseconds())
				avTimeEvalOTNH += float64(timeEvalOTNH.Milliseconds())
			} else {
				failedAttemptsOTNH += 1
			}

		}
		aNH := float64(rounds - failedAttemptsNH)
		aOTNH := float64(rounds - failedAttemptsOTNH)
		write(file, fmt.Sprint(filename, "- noisyDOT = [", avTimeSetupNH/aNH, ",", avTimeEncNH/aNH, ",", avTimeKeyGenNH/aNH, ",", avTimeEvalNH/aNH, "]\n"), append)
		write(file, fmt.Sprint(filename, "- DIFFE = [", avTimeSetupOTNH/aOTNH, ",", avTimeEncOTNH/aOTNH, ",", avTimeKeyGenOTNH/aOTNH, ",", avTimeEvalOTNH/aOTNH, "]\n"), append)

	}
}

func write(filename string, message string, append bool) {

	var file *os.File
	var err error

	if append {
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	} else {
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		log.Fatal(err)
	}

}

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
