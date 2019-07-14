# Vorta-Go

An implementation of [Vorta](https://github.com/borgbase/vorta) in Golang to improve deployment and packaging.

Work in progress and NOT functional. Use the [Python version](https://github.com/borgbase/vorta) if you need something that works.

## Development

1. Follow the [official steps](https://github.com/therecipe/qt/wiki/Installation) to set up a Go project in **Module Mode**.
2. Test app using `$ qtdeploy -debug -uic=false -quickcompiler test`
3. Package for deployment `$ qtdeploy -uic=false -quickcompiler build`

Important folders:

- `/ui` has `.ui` files provided by Qt Designer
- `/qml` has icons and other assets
