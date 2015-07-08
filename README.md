## GYAZO client for Linux

====

## Discription

Client tool for [GYAZO](http://gyazo.com) 

## Requirement

This tool use a [ImageMagick](http://www.imagemagick.org/script/index.php) and xclip.

## Install

```
$ go get github.com/jigyakkuma/gyagoyle
```

## Usage

If you run in a terminal:
```
$ gyagoyle
```

If you use a config file:
```
$ vi ~/.gyagoyle/config.toml
[[profile]]
name = "example.com"
endpoint = "http://example.com/gyazo/upload"
basicUser = "userName"
basicPassword = "password"
```

and
```
$ gyagoyle --profile example.com
```

It is convenient to register a shortcut on the menu.
