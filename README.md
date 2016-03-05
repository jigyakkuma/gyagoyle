## GYAZO client for Linux

[![wercker status](https://app.wercker.com/status/f5615d48864cf5abc373172a277c8c40/s "wercker status")](https://app.wercker.com/project/bykey/f5615d48864cf5abc373172a277c8c40)

## Discription

Client tool for [GYAZO](http://gyazo.com) 

## Requirement

This tool use a [ImageMagick](http://www.imagemagick.org/script/index.php) or Gnome Screenshot, and xclip.

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
