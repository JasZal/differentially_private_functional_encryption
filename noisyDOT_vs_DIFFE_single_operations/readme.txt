This folder contains all source code for the comparison of noisy-DOT and DIFFE (Figures 5-7).

Start computation:

To start the computation please start the main.go file. 
In The main.go file, there are following attributes, that can be changed:
- attributesR: range for attribute size
- clientsR: range for number of clients 
- secLevel: security Level, currently k = 2
- rounds: rounds over that the average should be computed
-filenames: files where the data is stored
	- file_FH: file for data of noisy-DOT
	- file_NH file for data of DIFFE
	
	
Data:
The data is stored in the two files noisy-DOT and DIFFE. Measurements are in milliseconds. Each matrix corresponds to the measurements for attribute size stated in the variablename. The matrix contains the attributes:
numver of clients, time for setup,time for  encryption, time for function key generation,time for decryption 
