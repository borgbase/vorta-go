FROM ubuntu:18.04 as base

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git
RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

FROM ubuntu:18.04
LABEL maintainer therecipe

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true

COPY --from=base /usr/local/go /usr/local/go
COPY --from=base $GOPATH/bin $GOPATH/bin
COPY --from=base $GOPATH/src/github.com/therecipe/qt $GOPATH/src/github.com/therecipe/qt

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install build-essential libglib2.0-dev libglu1-mesa-dev libpulse-dev  && apt-get -qq clean
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install libqt*5-dev qt*5-dev qt*5-doc-html && apt-get -qq clean

RUN $GOPATH/bin/qtsetup

COPY . $HOME/src
RUN cd $HOME/src && $GOPATH/bin/qtdeploy -uic=false -quickcompiler -debug build
