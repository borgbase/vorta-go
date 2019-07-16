DISTRO ?= ubuntu_19_04
.PHONY: darwin linux


linux:
	docker build -f docker/$(DISTRO).Dockerfile -t vorta/$(DISTRO) .
	docker cp $$(docker ps -alq):/home/user/vorta/deploy/linux/vorta deploy/linux
	upx deploy/linux/vorta
	mv deploy/linux/vorta deploy/linux/vorta-$(DISTRO)

darwin:
	qtdeploy -uic=false -quickcompiler build
	xattr -cr deploy/darwin/vorta-go.app
	codesign -f --deep --sign 'Developer ID Application: Manuel Riel (CNMSCAXT48)' deploy/darwin/vorta-go.app
	sleep 2; appdmg appdmg.json deploy/darwin/vorta-go.dmg
