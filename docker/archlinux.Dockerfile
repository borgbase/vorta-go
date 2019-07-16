FROM archlinux/base

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/local/go/bin:$PATH
ENV QT_API 5.13.0
ENV QT_DOCKER true
ENV QT_PKG_CONFIG true

RUN pacman -Syyu --quiet || true
RUN pacman -S --noconfirm --needed --noprogressbar --quiet ca-certificates curl git tar pkg-config
RUN GO=go1.12.6.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

RUN pacman -Syyu --quiet || true
RUN pacman -S --noconfirm --needed --noprogressbar --quiet base-devel glibc pkg-config && pacman -Scc --noconfirm --noprogressbar --quiet
RUN pacman -S --noconfirm --needed --noprogressbar --quiet qt5 && pacman -Scc --noconfirm --noprogressbar --quiet

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
