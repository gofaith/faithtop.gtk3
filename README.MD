# Faithtop :Go语言的桌面GUI框架(基于Gtk3)

- [安装](#安装)
- [快速开始](#快速开始)
- [部署](#部署)

---

# 安装

- [Windows](#windows)
- [Mac](#mac)
- [Linux](#linux)

## Windows

Download [msys2](https://www.msys2.org/)

Install gtk3 in msys2
```shell
pacman -S mingw-w64-x86_64-gtk3
```
[下载太慢？添加MSYS2中国源试一试](https://github.com/gofaith/faithtop.gtk3/wiki/MSYS2%E6%BA%90%E9%85%8D%E7%BD%AE)



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

# 快速开始


# 部署
