#!/usr/bin/env bash
# 构建移动端脚本

CRTDIR=$(pwd)
 
# 判断是否有output文件夹
if [ ! -d "${CRTDIR}/output/android" ]; then
  mkdir ${CRTDIR}/output/android
fi


# gomobile bind [-target android|ios|iossimulator|macos|maccatalyst] [-bootclasspath <path>] [-classpath <path>] [-o output] [build flags] [package]
# gomobile bind ./kernel/
gomobile bind -target=android -o=./output/android/mobile.aar -ldflags '-s -w'  ./cmd/mobile