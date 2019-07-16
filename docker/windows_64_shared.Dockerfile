# From https://github.com/therecipe/qt/tree/master/internal/docker/

FROM ubuntu:16.04

ENV USER user
ENV HOME /home/$USER
ENV GOPATH $HOME/work
ENV PATH /usr/lib/mxe/usr/bin:/usr/local/go/bin:$PATH
ENV QT_DOCKER true
ENV QT_MXE true
ENV QT_MXE_ARCH amd64
ENV QT_MXE_STATIC false

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates git
RUN git clone -q https://github.com/mxe/mxe.git /usr/lib/mxe && cd /usr/lib/mxe && git checkout -f 884669f1faa0bda893c22356229ba98beace0f97
#TODO: update to 5.13
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install autoconf automake autopoint bash bison bzip2 flex g++ g++-multilib gettext git gperf intltool libc6-dev-i386 libgdk-pixbuf2.0-dev libltdl-dev libssl-dev libtool-bin libxml-parser-perl make openssl p7zip-full patch perl pkg-config python ruby scons sed unzip wget xz-utils lzip

RUN cd /usr/lib/mxe && make -j $(grep -c ^processor /proc/cpuinfo) MXE_TARGETS='x86_64-w64-mingw32.shared' qt5 && rm -rf /usr/lib/mxe/log && rm -rf /usr/lib/mxe/pkg

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install ca-certificates curl git
RUN GO=go1.12.6.linux-amd64.tar.gz && curl -sL --retry 10 --retry-delay 60 -O https://dl.google.com/go/$GO && tar -xzf $GO -C /usr/local
RUN /usr/local/go/bin/go get -tags=no_env github.com/therecipe/qt/cmd/...

RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install dbus fontconfig libx11-6 libx11-xcb1
RUN QT=qt-unified-linux-x64-online.run && curl -sL --retry 10 --retry-delay 60 -O https://download.qt.io/official_releases/online_installers/$QT && chmod +x $QT && QT_QPA_PLATFORM=minimal ./$QT --no-force-installations --script $GOPATH/src/github.com/therecipe/qt/internal/ci/iscript.qs LINUX=true
RUN find /opt/Qt/5.13.0 -type f -name "*.debug" -delete
RUN find /opt/Qt/Docs -type f ! -name "*.index" -delete
RUN apt-get -qq update && apt-get --no-install-recommends -qq -y install binutils
RUN find /opt/Qt/5.13.0/gcc_64 -type f ! -name "*.a" ! -name "*.la" ! -name "*.prl" -name "lib*" -exec strip -s {} \;


RUN $GOPATH/bin/qtsetup prep
RUN $GOPATH/bin/qtsetup check windows
RUN $GOPATH/bin/qtsetup generate windows
RUN $GOPATH/bin/qtsetup install windows

RUN mkdir $HOME/vorta
WORKDIR $HOME/vorta
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN cd $HOME/vorta && rm -rf deploy && $GOPATH/bin/qtdeploy -uic=false -quickcompiler -debug build
