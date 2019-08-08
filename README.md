# 🐾 gTag

🐾 gTag is a `$(git tag)` workflow tool for semantic versioning (semver.org)

[![BSD License](https://img.shields.io/badge/license-BSD-blue.svg?style=flat)](LICENSE) 
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Flfaoro%2Fgtag.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Flfaoro%2Fgtag?ref=badge_shield)
[![Go Report Card](https://goreportcard.com/badge/github.com/lfaoro/gtag)](https://goreportcard.com/report/github.com/lfaoro/flares)

## Examples
```shell script
$ gtag zero
tag 0.0.0 created on commit c1788742

$ gtag minor
incremented 0.0.0 -> 0.1.0 on commit d846485e

$ gtag list
tag 0.0.0 [c1788742] by Leonardo on August 06 2019 at 16:12:46
tag 0.1.0 [d846485e] by Leonardo on August 08 2019 at 00:37:29

$ gtag major beta
tag 1.0.0-beta created on commit e894b051

$ gtag minor rc1
tag 1.1.0-rc1 created on commit 7706b9ad

$ gtag del
deleted 1.1.0-rc1 for commit 7706b9ad

$ gtag del --all
You're about to delete all tags, are you sure? y/n
```

## Quick Start

### [Video Tutorial](https://asciinema.org/a/261318)

### macOS
```bash
brew install lfaoro/tap/gtag
```

### Linux (soon)
```bash
curl dl.fireblaze.io/gtag.sh | bash
```

### Developers
> Go installer: https://golang.org/dl/
```bash
go get -d github.com/lfaoro/gtag
cd $GOPATH/src/github.com/lfaoro/gtag
make install
gtag -h
```

# Contributing

> Any help and suggestions are very welcome and appreciated. Start by opening an [issue](https://github.com/lfaoro/flares/issues/new).

- Fork the project
- Create your feature branch `git checkout -b my-new-feature`
- Commit your changes `git commit -am 'Add my feature'`
- Push to the branch `git push origin my-new-feature`
- Create a new pull request against the master branch

