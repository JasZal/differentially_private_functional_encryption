FROM ubuntu:20.04

ARG  GO_VERSION=1.21.4
ARG ARCH=amd64

RUN apt update -y 
RUN apt install apt-transport-https curl gnupg -y

#Install bazel
RUN curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor >bazel-archive-keyring.gpg && \
   mv bazel-archive-keyring.gpg /etc/apt/trusted.gpg.d/ && \
   echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list && \
   apt  update -y && \
   apt install bazel=6.4.0 -y

#Install Go
RUN  curl -O -L "https://golang.org/dl/go${GO_VERSION}.linux-${ARCH}.tar.gz" && \
     tar -xf "go${GO_VERSION}.linux-${ARCH}.tar.gz" && \
     mv -v go /usr/local

ENV GOPATH=$HOME/go
ENV PATH /usr/local/go/bin:$PATH
#Build go

#COP . ~/differential-privacy

WORKDIR /root

RUN apt install git -y
RUN git clone https://github.com/google/differential-privacy.git && \
    cd differential-privacy && \
    bazel build

RUN git clone https://github.com/JasZal/gofe && \
   cd gofe &&  \
   go install github.com/JasZal/gofe

RUN git clone https://github.com/JasZal/differentially_private_functional_encryption


