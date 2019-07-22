# From https://github.com/therecipe/qt/tree/master/internal/docker/linux

FROM ubuntu:18.04

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true
ENV QT_API 5.9.0

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git pkg-config
RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev && apt-get -qq clean
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libqt*5-dev qt*5-dev qt*5-doc-html && apt-get -qq clean
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev \
	&& apt-get --no-install-recommends -qq -y install fontconfig libasound2 libegl1-mesa libnss3 libpci3 libxcomposite1 libxcursor1 libxi6 libxrandr2 libxtst6 && apt-get -qq clean && apt-get -qq update && apt-get --no-install-recommends -qq -y install fcitx-frontend-qt5 && apt-get -qq clean

RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check
RUN $GOPATH/bin/qtsetup generate

RUN mkdir $HOME/vorta
WORKDIR $HOME/vorta
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .
RUN cd $HOME/vorta && rm -rf deploy && $GOPATH/bin/qtdeploy -uic=false -quickcompiler -debug build
