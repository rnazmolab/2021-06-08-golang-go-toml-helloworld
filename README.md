# 2021-06-08-golang-go-toml-v2-helloworld

## Link

[pelletier/go-toml: Go library for the TOML file format](https://github.com/pelletier/go-toml/tree/master)

## Purpose of this repository

Hello World to [pelletier/go-toml](https://github.com/pelletier/go-toml) **v2**.

## Results

```console
$ ls
go.mod  go.sum  main.go  README.md

$ go run ./main.go
=================
Version = 2
Name = 'go-toml'
Tags = ['go', 'toml']
foo_bar_version = '0.0.4'

=================
version: 2
name: go-toml
tags: [go toml]
foo_bar_version: 0.0.4
=================
```

## Env

```console
$ property
property v0.1.1 - A tiny Bash script to get OS and other
software version info. https://github.com/rnazmo/property
============================================================
OS NAME      : Kali GNU/Linux Rolling
OS VERSION   : 2021.2
Default Shell: Zsh
Bash VERSION : 5.1.4(1)-release (x86_64-pc-linux-gnu)
Zsh VERSION  : 5.8 (x86_64-debian-linux-gnu)
CPU ARCH     : x86_64
============================================================

$ go version
go version go1.16 linux/amd64
```

## Resources

[pelletier/go-toml at v2](https://github.com/pelletier/go-toml/tree/v2)

[toml · pkg.go.dev](https://pkg.go.dev/github.com/pelletier/go-toml/v2)

https://github.com/pelletier/go-toml/tree/618f0181ac76015c3efdaf8b8ab5a468293f6e82#automatic-field-name-guessing
