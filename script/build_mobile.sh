#!/usr/bin/env bash
# 构建移动端脚本

# gomobile bind [-target android|ios|iossimulator|macos|maccatalyst] [-bootclasspath <path>] [-classpath <path>] [-o output] [build flags] [package]
# gomobile bind ./kernel/
gomobile bind -target=android -o=./output/kernel.aar ./kernel