FROM opensuse/leap:latest as base

ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true

RUN zypper -q ref && zypper -n -q install --no-recommends curl git gzip tar pkg-config
RUN GO=go1.12.6.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

RUN zypper -q ref && zypper -n -q install --no-recommends -t pattern devel_basis && zypper clean -a
RUN zypper -q ref && zypper -n -q install --no-recommends libqt5-qt*-devel && zypper clean -a

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
