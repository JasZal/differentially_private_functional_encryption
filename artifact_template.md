# Artifact Appendix

Paper title: Differentially Private Functional Encryption

Artifacts HotCRP Id: #47 

Requested Badge: Reproducible

## Description
This artifact is the source code that was used to measure the data linked to the Figures 8 - 10 and Table 2 in the Paper. 

### Security/Privacy Issues and Ethical Concerns
- none

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


## Environment
Please access the artifact through the provided github link. All related data and software components as well as the description how to build and run it, are in the README file.

### Accessibility
https://github.com/JasZal/differentially_private_functional_encryption 
tag V0


### Set up the environment
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


### Testing the Environment
Try to start the experiments (see above). If it starts to print debug messages, e.g. name of database or number of attributes/clients, it works.


## Artifact Evaluation

### Main Results and Claims
Main results regarding the implementation are first, a proof of concept, meaning, the scheme is implementable and runs in reasonable time. Second we show, that DiffPIPE is faster than noisyDOT. (See Figures 8-10 (Exp.1) and Table 2 (Exp. 2) in the Paper)


### Experiments

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
results: time in millisec for setup, encryption, key generation and evaluation for different data sets printed on the shell and also saved in the file "results..."

## Limitations
no limitations

## Notes on Reusability
The adaptet library can be used in any setting regarding privacy preserving analysis that can be done with linear functions + differential privacy.
