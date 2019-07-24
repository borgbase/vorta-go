# From https://github.com/therecipe/qt/tree/master/internal/docker/linux

FROM fedora:30

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true
ENV QT_DIR /usr/include/qt5/
ENV QT_VERSION 5.12.1

RUN dnf --refresh install -y qt5-qtbase qt5-devel qt5-qtbase-doc glib2-devel mesa-libGLU-devel pulseaudio-libs-devel git curl

RUN GO=go1.12.4.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

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
