This folder contains all source code, producing measuerments of Table 2.

The files authority, evaluator and user are the respective parties using each the noisy-DOT scheme and the DIFFE-scheme. The *.csv files contain the different data sets. 

In main.go there exist several variables, that can be adapted:
 - secLevel: contains security level, k = 2 currently
 - epsilon: epsilon associated with differntial privacy
 - scaling: scaling for fix-point arithmetic
 - rounds: rounds over which average is computed

The results after running the main.go file are saved in a .txt file. 
Each vector is named after the used dataset, e.g. lbw.csv, following bt the used scheme, e.g. noisyDOT. The attributes of the vector contain the average Time for setup (nanosec), encryption (microsec), keygen (millisec) and evaluation (millisec).

