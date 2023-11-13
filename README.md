# Differentially Private Functional Encryption

This repository is a proof of concept and demonstrates how a secure noisy functional encryption scheme can be used. It is implemented under go version 1.21.4. For help installing go, please check the official [website](https://go.dev/doc/install).\n

The example uses a forked and modified implementation of the [gofe library](https://github.com/JasZal/gofe) and the [differential privacy library](https://github.com/google/differential-privacy). 
The addition to the gofe library are two schemes, one a message-and-noise-hiding noisy FE scheme (called noisyDOT) and the other a single-message-and-noise-hiding noisy FE scheme (called DIFFE). 
Our example shows a comparison between the two schemes for single operations as well as for statistic analysis on real data sets.

# How to run the example

To run the example first install the forked [gofe library](https://github.com/JasZal/gofe) and the [differential privacy library](https://github.com/google/differential-privacy). Make sure you have installed all dependencies, e.g. bazel in version 6.2.1.
For comparison of the two schemes concerning the single operations navigate into the 'noisyDOT_vs_DIFFE_single_operations' folder and run main.go.
To compare the two over real datasets navigate to 'noisyDOT_vs_DIFFE_data_sets' folder and run main.go
