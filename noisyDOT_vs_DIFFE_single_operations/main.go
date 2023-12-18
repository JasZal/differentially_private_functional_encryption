package main

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/JasZal/gofe/data"
	"github.com/JasZal/gofe/innerprod/noisy"
	"github.com/JasZal/gofe/sample"
)

var deb bool = true

func debug(s string, deb bool) {
	if deb {
		fmt.Printf(s)
	}
}

/* this function writes the data into files*/
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

func main() {
	//attributes to set for whished attribute range and client range
	attributesR := []int{5, 10, 20, 30}
	clientsR := []int{5, 10, 20, 30}

	//security Level k = 2
	secLevel := 2
	append := false

	//rounds over that the average should be computed
	rounds := 3

	file_noisyDOT := "results_noisy_DOT.txt"
	file_DIFFE := "results_DiffPIPE.txt"

	if !append {

		write(file_noisyDOT, "#clients, setup, encryption, function key generation, decryption -- time in millisec\n\n", append)
		write(file_DIFFE, "#clients, setup, encryption, function key generation, decryption -- time in millisec\n\n", append)

	}

	for _, attributes := range attributesR {

		write(file_noisyDOT, fmt.Sprintf("noisyDOT_attr%d = [\n", attributes), true)
		write(file_DIFFE, fmt.Sprintf("DiffPIPE_attr%d = [\n", attributes), true)

		for _, clients := range clientsR {
			failed_fh := 0
			failed_nh := 0
			//create file
			s := sample.NewUniform(big.NewInt(100)) // will sample uniformly from [0,100)
			x, _ := data.NewRandomMatrix(clients, attributes, s)

			debug(fmt.Sprintf("attributes: %d, clients: %d\n", attributes, clients), deb)

			//Attributes for time measurements
			var fh_time_Setup, fh_time_Enc, fh_time_KeyGen, fh_time_Dec, nh_time_Setup, nh_time_Enc, nh_time_KeyGen, nh_time_Dec int64

			for i := 1; i <= rounds; i++ {

				debug(fmt.Sprintf("round %d \n", i), deb)
				t_fh_time_Setup, t_fh_time_Enc, t_fh_time_KeyGen, t_fh_time_Dec, err_fh := testNHMIPE(clients, attributes, secLevel, x)
				if err_fh != nil {
					failed_fh += 1
				}
				t_nh_time_Setup, t_nh_time_Enc, t_nh_time_KeyGen, t_nh_time_Dec, err_nh := testOTNHMIPE(clients, attributes, secLevel, x)
				if err_nh != nil {
					failed_nh += 1
				}

				fh_time_Setup += int64(t_fh_time_Setup.Milliseconds())
				fh_time_Enc += int64(t_fh_time_Enc.Milliseconds())
				fh_time_KeyGen += int64(t_fh_time_KeyGen.Milliseconds())
				fh_time_Dec += int64(t_fh_time_Dec.Milliseconds())
				nh_time_Setup += int64(t_nh_time_Setup.Milliseconds())
				nh_time_Enc += int64(t_nh_time_Enc.Milliseconds())
				nh_time_KeyGen += int64(t_nh_time_KeyGen.Milliseconds())
				nh_time_Dec += int64(t_nh_time_Dec.Milliseconds())

			}

			write(file_noisyDOT, fmt.Sprint(clients, ", ", float64(fh_time_Setup)/float64(rounds-failed_fh), ",", float64(fh_time_Enc)/float64(rounds-failed_fh), ",", float64(fh_time_KeyGen)/float64(rounds-failed_fh), ",", float64(fh_time_Dec)/float64(rounds-failed_fh), "; \n"), true)

			write(file_DIFFE, fmt.Sprint(clients, ", ", float64(nh_time_Setup)/float64(rounds-failed_nh), ",", float64(nh_time_Enc)/float64(rounds-failed_nh), ",", float64(nh_time_KeyGen)/float64(rounds-failed_nh), ",", float64(nh_time_Dec)/float64(rounds-failed_nh), "; \n"), true)
		}

		write(file_noisyDOT, "];\n\n", true)
		write(file_DIFFE, "];\n\n", true)
	}

}

func testNHMIPE(clients, attributes int, secLevel int, x data.Matrix) (time.Duration, time.Duration, time.Duration, time.Duration, error) {

	boundX := big.NewInt(100)
	boundY := big.NewInt(1)

	//prepare database
	cipher := make(data.MatrixG1, clients)

	//generate Authority with Params and Keys
	start := time.Now()
	a_fhmulti := noisy.NewNHMultiIPE(secLevel, clients, attributes, boundX, boundY)
	a_masterSecKey, a_pubKey, _ := a_fhmulti.GenerateKeys()
	timeSetup := time.Since(start)

	//generate numClient Clients and fill cipher
	var timeEnc time.Duration
	for i := 0; i < clients; i++ {
		if i == 0 {

			client := noisy.NewNHMultiIPEFromParams(a_fhmulti.Params)
			start = time.Now()
			cipher[i], _ = client.Encrypt(x[i], a_masterSecKey.BHat[i])
			timeEnc = time.Since(start)
		} else {
			client := noisy.NewNHMultiIPEFromParams(a_fhmulti.Params)
			cipher[i], _ = client.Encrypt(x[i], a_masterSecKey.BHat[i])
		}

	}

	//generate evaluator, start decrypting
	// generate inner product vectors and put them in a matrix
	y := make(data.Matrix, clients)
	for i := 0; i < clients; i++ {
		y[i] = data.NewConstantVector(attributes, big.NewInt(0))
		y[i][0] = big.NewInt(1)
	}

	//time to generate key

	rand.Seed(time.Now().UnixNano())
	noise := rand.Int63n(21) - 10 //range -10 - 10

	start = time.Now()
	funcKey, err := a_fhmulti.DeriveKey(y, a_masterSecKey, noise)
	timeFK := time.Since(start)

	if err != nil {
		fmt.Println("Error during key derivation:", err)
		return 0, 0, 0, 0, err
	}

	decryptor := noisy.NewNHMultiIPEFromParams(a_fhmulti.Params)

	start = time.Now()
	xy, err := decryptor.Decrypt(cipher, funcKey, a_pubKey)
	timeD := time.Since(start)

	if err != nil {
		fmt.Println("Error Decrypting")
		fmt.Println(err)
		return 0, 0, 0, 0, err
	}

	debug(fmt.Sprintf("xy = %d,---------------finished FHMIPE---------------\n", xy), deb)

	return timeSetup, timeEnc, timeFK, timeD, nil
}

func testOTNHMIPE(clients, attributes int, secLevel int, x data.Matrix) (time.Duration, time.Duration, time.Duration, time.Duration, error) {

	boundX := big.NewInt(100)
	boundY := big.NewInt(1)

	//prepare database
	cipher := make(data.MatrixG1, clients)

	//generate Authority with Params and Keys
	start := time.Now()
	a_otnhmulti := noisy.NewOTNHMultiIPE(secLevel, clients, attributes, boundX, boundY)
	a_masterSecKey, a_pubKey, _ := a_otnhmulti.GenerateKeys()
	timeSetup := time.Since(start)

	var timeEnc time.Duration
	//generate numClient Clients and fill cipher
	for i := 0; i < clients; i++ {
		if i == 0 {
			client := noisy.NewOTNHMultiIPEFromParams(a_otnhmulti.Params)
			start = time.Now()
			cipher[i], _ = client.Encrypt(x[i], a_masterSecKey.BHat[i])
			timeEnc = time.Since(start)
		} else {
			client := noisy.NewOTNHMultiIPEFromParams(a_otnhmulti.Params)
			cipher[i], _ = client.Encrypt(x[i], a_masterSecKey.BHat[i])
		}

	}

	//generate evaluator, start decrypting
	// generate inner product vectors and put them in a matrix
	y := make(data.Matrix, clients)
	for i := 0; i < clients; i++ {
		y[i] = data.NewConstantVector(attributes, big.NewInt(0))
		y[i][0] = big.NewInt(1)
	}

	rand.Seed(time.Now().UnixNano())
	noise := rand.Int63n(21) - 10 //range -10 - 10
	debug(fmt.Sprintf("noise %d \n", noise), deb)

	start = time.Now()
	funcKey, err := a_otnhmulti.DeriveKey(y, a_masterSecKey, noise)
	timeFK := time.Since(start)
	if err != nil {
		fmt.Println("Error during key derivation:", err)
		return 0, 0, 0, 0, err
	}

	decryptor := noisy.NewOTNHMultiIPEFromParams(a_otnhmulti.Params)

	start = time.Now()
	xy, err := decryptor.Decrypt(cipher, funcKey, a_pubKey)
	timeD := time.Since(start)
	if err != nil {
		fmt.Println("Error Decrypting")
		fmt.Println(err)
		return 0, 0, 0, 0, err
	}

	debug(fmt.Sprintf("xy = %d,---------------finished OT-NHMIPE---------------\n", xy), deb)

	return timeSetup, timeEnc, timeFK, timeD, nil
}
