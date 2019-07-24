DISTRO ?= 0ubuntu0.19.04.1
.PHONY: darwin linux


linux:
	docker build -f build/docker/$(DISTRO).Dockerfile -t vorta/$(DISTRO) .
	ID=$$(docker create vorta/$(DISTRO)) && docker cp $$ID:/home/user/vorta/deploy/linux/vorta deploy/linux/vorta && docker rm -v $$ID
	upx deploy/linux/vorta
	nfpm -f build/package/nfpm.yaml pkg --target deploy/linux/vorta_v0.0.1-$(DISTRO)_amd64.deb
	mv deploy/linux/vorta deploy/linux/vorta_v0.0.1-$(DISTRO)_amd64.bin


darwin:
	qtdeploy -uic=false -quickcompiler build
	xattr -cr deploy/darwin/vorta-go.app
	codesign -f --deep --sign 'Developer ID Application: Manuel Riel (CNMSCAXT48)' deploy/darwin/vorta-go.app
	sleep 2; appdmg appdmg.json deploy/darwin/vorta-go.dmg
