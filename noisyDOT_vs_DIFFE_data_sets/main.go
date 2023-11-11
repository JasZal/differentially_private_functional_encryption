package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
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
	//	test_randomData(secLevel, rounds, scaling, epsilon, debug, append)

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

/*

func test_randomData(secLevel, rounds int, scaling int64, epsilon float64, debug, append bool) {

	//the value by which the coordinates of encrypted vectors are bounded, for each dataset
	var bound int64 = 1000 //
	//aattributes and  records
	var attrA []int64 = []int64{5, 10, 20, 30}
	var recA []int64 = []int64{1, 10, 50, 100, 250, 500}

	fileNH := "resultRandomNH.txt"
	fileOT := "resultRandomOT_NH.txt"

	if !append {

		write(fileNH, fmt.Sprint("scaling=", scaling, ", avTime in Setup=Nanoseconds, Enc=Microseconds, KeyGen=Eval=Milliseconds \n"), append)
		write(fileOT, fmt.Sprint("scaling=", scaling, ", avTime in Setup=Nanoseconds, Enc=Microseconds, KeyGen=Eval=Milliseconds \n"), append)
		append = true
		write(fileNH, "records, Setup, Encryption, KeyGen, Evaluation;\n", append)
		write(fileOT, "records, Setup, Encryption, KeyGen, Evaluation;\n", append)

	}

	if debug {
		fmt.Print("scaling=", scaling, ", avTime in Setup=Nanoseconds, Enc=Microseconds, KeyGen=Eval=Milliseconds \n")
	}

	rand.Seed(14665)

	for _, attr := range attrA {
		write(fileNH, fmt.Sprint("attributes", attr, " = ["), append)
		write(fileOT, fmt.Sprint("attributes", attr, " = ["), append)

		for _, records := range recA {

			if debug {
				fmt.Println("attributes:  ", attr, " records: ", records)
			}

			write(fileOT, fmt.Sprint(records, ", "), append)
			write(fileNH, fmt.Sprint(records, ", "), append)
			//fill database
			//sample random values with at most scaling places after point, and bound as upper bound
			filename := "fileRandom.txt"
			write(filename, "", false)
			for j := 0; j < (int)(records); j++ {

				for i := 0; i < int(attr); i++ {
					if i == 0 {
						write(filename, fmt.Sprint(rand.Intn(2), ", "), true)
					} else if i == int(attr)-1 {
						write(filename, fmt.Sprint(rand.Intn(int((scaling/10)*bound)), "\n"), true)
					} else {
						write(filename, fmt.Sprint(rand.Intn(int((scaling/10)*bound)), ", "), true)
					}
				}
			}

			if debug {
				fmt.Print("file generated \n")
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
			boundX := big.NewInt(bound * scaling)
			boundY := big.NewInt(1)

			for i := 0; i < rounds; i++ {
				errNH, _, timeSetupNH, timeEncNH, timeKeyGenNH, timeEvalNH := testNH(filename, false, epsilon, scaling, secLevel, boundX, boundY)
				errOTNH, _, timeSetupOTNH, timeEncOTNH, timeKeyGenOTNH, timeEvalOTNH := testOTNH(filename, false, epsilon, scaling, secLevel, boundX, boundY)

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
			write(fileNH, fmt.Sprint(avTimeSetupNH/aNH, ",", avTimeEncNH/aNH, ",", avTimeKeyGenNH/aNH, ",", avTimeEvalNH/aNH, ";"), append)
			write(fileOT, fmt.Sprint(avTimeSetupOTNH/aOTNH, ",", avTimeEncOTNH/aOTNH, ",", avTimeKeyGenOTNH/aOTNH, ",", avTimeEvalOTNH/aOTNH, ";"), append)

		}

		write(fileNH, "]\n", append)
		write(fileOT, "]\n", append)

	}
}*/

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
