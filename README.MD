# Faithtop : GUI Framework for Go (Based on Gtk3)

---

# Installation

- [Windows](#windows)
- [Mac](#mac)
- [Linux](#linux)

## Windows

Download [msys2](https://www.msys2.org/)

Install gtk3 in msys2
```shell
pacman -S mingw-w64-x86_64-gtk3 git vim pkg-config mingw-w64-x86_64-gcc
```
[下载太慢？添加MSYS2中国源试一试](https://github.com/gofaith/faithtop.gtk3/wiki/MSYS2%E6%BA%90%E9%85%8D%E7%BD%AE)

Install Go inside msys2 directory, like `C:\msys2\opt\go`

Add those lines in file `C:\msys2\etc\profile`

```shell

export GOROOT=/opt/go
export GOPATH=/home/asd/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/mingw64/lib/pkgconfig
```

Open file `C:\msys2\mingw64\lib\pkgconfig\gdk-3.0.pc` , remove `-Wl,-luuid`,and add a line below:
```pc
Ldflags:-Wl,-luuid
```

Install Faithtop
```shell
go get -tags gtk_3_18 github.com/gofaith/faithtop
```

## Mac

```shell
brew install gtk+3
```

```shell
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/local/Cellar/libffi/3.2.1/lib/pkgconfig
```

```shell
go get -v -tags gtk_3_18 github.com/gofaith/faithtop
```

## Linux

# Get Started


# Deploy
