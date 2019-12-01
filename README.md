# Vorta-Go

An implementation of [Vorta](https://github.com/borgbase/vorta) in Golang to improve deployment and packaging.

Functional, but still missing some features. Use the [Python version](https://github.com/borgbase/vorta) if you need something that works.

## Progress

Borg Commands:
- [x] `init`
- [x] `create`
- [x] `info`
- [x] `check`
- [ ] [`extract`](https://github.com/therecipe/examples/blob/master/advanced/widgets/treeview/main.go)
- [x] `mount`
- [ ] `delete`
- [ ] `diff`
- [ ] `list-archive`
- [x] `list-repo`
- [x] `prune`
- [ ] `umount` and mount status
- [x] `version`

Backend Functionality
- [x] Keychain/SecretService
- [x] Background scheduler
- [x] Single App
- [x] Read and parse existing SSH keys
- [x] Create new SSH key
- [x] Check Borg version for available features
- [ ] [Tests](https://github.com/therecipe/examples/tree/master/test/widgets)
- [x] Read list of WiFis
- [ ] Check Wifi before scheduled backup
- [ ] ~~Password fallback to database~~
- [ ] Translations
- [x] Notifications from scheduler
- [ ] Settings an related functionality (icon color, etc)
- [ ] [Sparkle updates for macOS](https://github.com/therecipe/qt/issues/743#issuecomment-444689169)

Packaging (via Docker)
- [x] macOS
- [x] Ubuntu 19.04
- [ ] Debian 10
- [ ] Fedora 30
- [x] Archlinux
- [ ] Windows?

Other issues:
- [x] Exclusions
- [x] Cancel button
- [x] Backup status (icon, menu)
- [ ] Bug: exclusions are lost

## Development

1. Follow the [official steps](https://github.com/therecipe/qt/wiki/Installation) to set up a Go project in **Module Mode**. E.g. for macOS

```
export GO111MODULE=on
go install -v -tags=no_env github.com/therecipe/qt/cmd/...
go mod vendor
git clone https://github.com/therecipe/env_darwin_amd64_513.git vendor/github.com/therecipe/env_darwin_amd64_513
```

2. Test app using `$ qtdeploy -debug -uic=false -quickcompiler test`
3. Package for deployment `$ qtdeploy -uic=false -quickcompiler build` or `make darwin`

Important folders:

- `/ui` has `.ui` files provided by Qt Designer
- `/qml` has icons and other assets


## Deployment

See the `Makefile` for different deployment options. Needs Docker installed. E.g.

- `$ make darwin`
- `$ make linux DISTRO=archlinux`

For Linux, Qt5 is linked dynamically to match your distro's look and feel. Install required Qt5 packages like this:

- Ubuntu/Debian: `$ apt install qt5-default libqt5qml5`
- Archlinux: `$ pacman -S qt5`
- Fedora: `$ yum install qt5-qtbase`

## Translations (work in progress)
- extract strings: `lupdate -extensions ui ui/*.ui -ts qml/i18n/ui_en.ts`
- merge .ts files: `lconvert -i primary.ts secondary.ts -o complete.ts` 
- compile .ts to .qm: `lrelease qml/i18n/ui_de.ts -qm qml/i18n/ui_de.qm`

## License and Credits
- Thank you to all the people who already contributed to Vorta: [code](https://github.com/borgbase/vorta/graphs/contributors), [translations](https://github.com/borgbase/vorta/issues/159)
- Licensed under GPLv3. See [LICENSE.txt](LICENSE.txt) for details.
- Icons by [FontAwesome](https://fontawesome.com)
