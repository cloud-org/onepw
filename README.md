## onepw [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/cloud-org/onepw/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/cloud-org/onepw)](https://goreportcard.com/report/github.com/cloud-org/onepw)

### Install

check [Releases](https://github.com/cloud-org/onepw/releases)

### What's this

onepw is a command line tool for managing passwords. You **MUST** remember the `master password`, and don't tell anyone!

### Principles

1. Generate Key by master password

```
o--------o             o-----o
| Master | KDF: scrypt |     |
| Pass   |============>| Key |
| Word   |             |     |
o--------o             o-----o
```

2. Encrypt account and password

```
o-----------o
|           |
| Random IV |==o
|           |  |                o------------o
o-----------o  | CFB Encrypter  |            |
               |===============>| CipherText |
o-----------o  | AES Cipher     |            |
|           |  | with Key       o------------o
| PlainText |==o
|           |
o-----------o
```

### Get Started

[Get Started](./quickstart.md)

### CHANGELOG

[ChangeLog](./CHANGELOG.md)
