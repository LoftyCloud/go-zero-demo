# 基于数据库表生成model架构

#!/usr/bin/env bash

# 使用方法：
# ./genModel.sh usercenter user
# ./genModel.sh usercenter user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

#生成的表名
tables=$2  # 获取第二个参数（表名）
#表生成的genmodel目录
modeldir=./genModel

# 数据库配置
host=127.0.0.1
# host=localhost
port=3306
dbname=$1   # 获取第一个参数（数据库名）
username=go-zero-root 
# passwd=PXDN93VRKUm8TeE7
passwd=root

echo "开始创建库：$dbname 的表：$tables" 
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero

# goctl model mysql datasource -url="go-zero-root:root@tcp(127.0.0.1:3306)/zero-demo" -table="user,userdata" -dir=./genModel -cache=true -style=goZero