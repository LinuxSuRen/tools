#!/bin/sh

curl -L https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-linux-amd64.tar.gz | tar xzv hd
mv hd /usr/bin/hd
hd completion
