#!/bin/bash

# 下载并解压缩可执行文件
wget https://github.com/sly17zql/cli88/releases/download/v0.0.1/cli88.tar.gz

if [ $? -ne 0 ]; then
  echo "Download failed. Exiting."
  exit 1
fi

tar zxvf cli88.tar.gz cli88

# 移动可执行文件到系统的bin目录
sudo mv cli88 /usr/local/bin/

# 清理
rm cli88.tar.gz

echo "cli88 已成功安装"