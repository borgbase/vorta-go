# From https://github.com/therecipe/qt/tree/master/internal/docker/linux

FROM ubuntu:19.04

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true
ENV QT_DIR /usr/include/x86_64-linux-gnu/qt5/

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git
RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev  && apt-get -qq clean
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libqt*5-dev qt*5-dev qt*5-doc-html && apt-get -qq clean

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check
RUN $GOPATH/bin/qtsetup generate

RUN mkdir $HOME/vorta
WORKDIR $HOME/vorta
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN cd $HOME/vorta && $GOPATH/bin/qtdeploy -uic=false -quickcompiler -debug build
