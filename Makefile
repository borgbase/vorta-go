OUT := vorta
PKG := gitlab.com/group/project
VERSION := $(shell git describe --always --long --dirty)
DISTRO ?= ubuntu
DISTRO_VERSION ?= 19.04

.PHONY: darwin linux

all: test

linux:
	docker build -f build/docker/$(DISTRO).Dockerfile -t vorta/$(DISTRO) .
	ID=$$(docker create vorta/$(DISTRO)) && docker cp $$ID:/home/user/vorta/deploy/linux/vorta deploy/linux/vorta && docker rm -v $$ID
	upx deploy/linux/vorta
	nfpm -f build/package/nfpm.yaml pkg --target deploy/linux/vorta_v0.0.1-$(DISTRO)_amd64.deb
	mv deploy/linux/vorta deploy/linux/vorta_v0.0.1-$(DISTRO)_amd64.bin

darwin:
	QT_HOMEBREW=true qtdeploy -uic=false -quickcompiler -ldflags '-X vorta/ui.version=${VERSION}' build
	xattr -cr deploy/darwin/vorta-go.app
	codesign -f --deep --sign 'Developer ID Application: Manuel Riel (CNMSCAXT48)' deploy/darwin/vorta-go.app
	sleep 2; appdmg appdmg.json deploy/darwin/vorta-${VERSION}.dmg

test:
	QT_HOMEBREW=true qtdeploy -uic=false -quickcompiler -ldflags '-X vorta/ui.version=${VERSION}' test
