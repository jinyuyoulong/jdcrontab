#!/usr/bin/env bash
!/bin/bash

# 如果bin目录不存在则创建目录
if [ ! -d bin  ];then
  mkdir bin
else
  echo dir exist 存在
fi
cd bin
go build ../src/main/jdcrontab.go
./jdcrontab