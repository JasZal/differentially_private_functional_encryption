# Differentially Private Functional Encryption

This repository is a proof of concept and demonstrates how a secure noisy functional encryption scheme can be used. It is implemented under go version 1.21.4. For help installing go, please check the official [website](https://go.dev/doc/install).


The example uses a forked and modified implementation of the [gofe library](https://github.com/JasZal/gofe) and the [differential privacy library](https://github.com/google/differential-privacy). 
The addition to the gofe library are two schemes, one a message-and-noise-hiding noisy FE scheme (called noisyDOT) and the other a single-message-and-noise-hiding noisy FE scheme (called DIFFE). 
Our example shows a comparison between the two schemes for single operations as well as for statistic analysis on real data sets.

# TLTR

To run the example first install the forked [gofe library](https://github.com/JasZal/gofe) and the [differential privacy library](https://github.com/google/differential-privacy). Make sure you have installed all dependencies, e.g. bazel.
For comparison of the two schemes concerning the single operations navigate into the 'noisyDOT_vs_DiffPIPE_single_operations' folder and run main.go.
To compare the two over real datasets navigate to 'noisyDOT_vs_DiffPIPE_data_sets' folder and run main.go

To generate figures out of the computed results for the 'noisyDOT_vs_DiffPIPE_single_operations' experiment, you can use the file 'buildfigures_single_operations.m' in the main folder. This file contains octave code, that can be compiled either locally or with an online octave compiler. Please copy your results in the respective place. 

More explanations can be found in the file "artifact_template.md". Or below in the long version. 
We additionally provide a dockerfile "Dockerfile" which installs all above statet dependencies. 




## Description
This artifact is the source code that was used to measure the data linked to the Figures 8 - 10 and Table 2 in the paper [Differentially Private Functional Encryption](). 

## Basic Requirements

### Hardware Requirements
at least 8 GB RAM

### Software Requirements
- OS: Ubuntu (at least version 20.04)
- Software: go (at least version 1.21.4), bazel (at least version 6.4.0)

### Estimated Time and Storage Consumption
(measured on VM with 8GB RAM - at capacity)
Exp. 1 - Single Operations: about 3min 
Exp. 2 - Data Sets: about 35min 


## Set up the environment
Either use the provided dockerfile or follow the above instruction:

(Assuming Ubuntu 20.04)
- install go on your system (https://go.dev/doc/install)
```bash
curl -O -L "https://golang.org/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz" 
tar -xf "go${GO_VERSION}.linux-${ARCH}.tar.gz" && mv -v go /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin' >>$HOME/.profile
echo 'export PATH=$PATH:$HOME/go/bin' >>$HOME/.profile
```
  
- install bazel on your system, at least version 6.4.0 (https://bazel.build/install)
```bash
curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor >bazel-archive-keyring.gpg
mv bazel-archive-keyring.gpg /etc/apt/trusted.gpg.d/
echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list 
sudo apt  update -y 
sudo apt install bazel=6.4.0
```

- clone and install the differntial privacy library from google (https://github.com/google/differential-privacy)
```bash 
git clone https://github.com/google/differential-privacy.git
cd differential-privacy
cd go
bazel build ...
cd ../..
```

- clone and install the forked gofe library (https://github.com/JasZal/gofe)
```bash
git clone https://github.com/JasZal/gofe
cd gofe
go install github.com/JasZal/gofe/...
cd ..
```

- clone the artifact (https://github.com/JasZal/differentially_private_functional_encryption/tree/V0)
```bash
git clone https://github.com/JasZal/differentially_private_functional_encryption
cd differentially_private_functional_encryption
```

now you can run the source code of the experiment by typing ```go run .``` in one of the two folders from the artifact (see below in subsection Experiments).
Expected results are described in the readme files.


## Artifact Evaluation

### Main Results and Claims
Main results regarding the implementation are first, a proof of concept, meaning, the scheme is implementable and runs in reasonable time. Second we show, that DiffPIPE is faster than noisyDOT. (See Figures 8-10 (Exp.1) and Table 2 (Exp. 2) in the Paper)


### Experiments
Expected results are described in the readme files.

#### Experiment 1: Single Operations
This experiment measures the single operations of noisyDot and DiffPIPE 
Navigate to the folder ``differentially_private_functional_encryption/noisyDOT_vs_DiffPIPE_single_operations``
run ```go run .```
results: time in millisec for setup, encryption, key generation and evaluation for different numbers of attributs and clients printed on the shell and also saved in the file "results_DiffPIPE.txt" and "results_noisy_DOT.txt".

To generate similar figures (8-10) like the ones in the paper, please use the octave code in file "buildfigures_single_operations.m" in the main folder. 
If you do not have octave or matlab installed, you can use an online compiler, e.g. https://www.mycompiler.io/new/octave
There are two variables, that can be adapted:
-- xaxis, that determines, the range of the xaxis. Depending on your results, this should be reduced, to get a meaningful figure. The paper uses 500.
-- fig: the online compilers that have a good visualization only allow to plot one figure at a time. If you use for example "https://www.mycompiler.io/new/octave" you can choose wich figure you want to plot, adapting the variable "fig". 


#### Experiment 2: Data Sets
This experiment measures the single operations of noisyDot and DiffPIPE
Navigate to the folder ``differentially_private_functional_encryption/noisyDOT_vs_DiffPIPE_data_sets``
run ```go run .```
results: time for setup [nanosec], encryption [microsec], key generation [millisec]  and evaluation [millisec] for different data sets printed on the shell and also saved in the file "results..."



