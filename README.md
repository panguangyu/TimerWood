[![License](https://img.shields.io/badge/license-Apache%202-green.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Build Status](https://travis-ci.org/xialonghua/kotmvp.svg?branch=master)](https://travis-ci.org/xialonghua/kotmvp) 

# timerWood
Go 开发的记录 Linux 命令或程序执行的时间工具

## Quick Start
## Install
```go
git clone https://github.com/panguangyu/timerWood.git
mv ./timerWood /usr/local/timerwood
cd /usr/local/timerwood
ln -s /usr/local/timerwood/main /usr/bin/ti
./install
```

## Usage
```go
> ti "python test/test.py"                            # 计算脚本运行的时间
> ti "python test/test.php"
> ti "ls"                                             # 计算 Linux 命令运行的时间
> ti "mysqldump ..."                                  # 计算 mysql 导出时间
> ti "wget http://..."                                # 计算 wget 下载的时间
```

## Log
```
cat /.Timerwood.dat
```
