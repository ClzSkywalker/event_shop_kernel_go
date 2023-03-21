#!/usr/bin/env bash
# 构建移动端脚本

CRTDIR=$(pwd)
 
# 判断是否有output文件夹
if [ ! -d "${CRTDIR}/output/windows" ]; then
  mkdir ${CRTDIR}/output/windows
fi

go env -w GOOS=windows
go build -ldflags '-s -w' -o ./output/windows/kernel.exe ./main.go