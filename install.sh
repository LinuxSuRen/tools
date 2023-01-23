#!/usr/bin/env bash

ARCH=$(uname -m)
if [ ${ARCH} == 'x86_64' ]; then
  ARCH=amd64
elif [ ${ARCH} == 'armv6l' ]; then
  ARCH=arm6
elif [ ${ARCH} == 'armv7l' ]; then
  ARCH=arm7
elif [ ${ARCH} == 'arm64' ]; then
  ARCH=arm64
else
  echo "do not support ${ARCH}"
  echo "please feel free to help improve this at https://github.com/LinuxSuRen/tools/blob/master/install.sh"
  exit 1
fi

OS=$(uname -s)
if [ ${OS} == 'Darwin' ]; then
  OS=darwin
elif [ ${OS} == 'Linux' ]; then
  OS=linux
else
  echo "do not support ${OS}"
  echo "please feel free to help improve this at https://github.com/LinuxSuRen/tools/blob/master/install.sh"
  exit 2
fi

if [ -x "$(command -v curl)" ]; then
  curl -L https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-${OS}-${ARCH}.tar.gz | tar xzv hd
elif [ -x "$(command -v wget)" ]; then
  wget https://github.com/linuxsuren/http-downloader/releases/latest/download/hd-${OS}-${ARCH}.tar.gz -O - | tar xzv hd
else
  echo "curl or wget is required to download hd."
  exit 3
fi

if ($(./hd version> /dev/null 2> /dev/null)); then
  sudo install hd /usr/bin/hd
  hd fetch
  hd completion
  hd setup
else
  echo "hd was not downloaded successfully."
fi
