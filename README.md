# timerWood
Go 开发的记录 Linux 运行时间记录仪

## Quick Start
## Install
> git clone https://github.com/panguangyu/timerWood.git
> mv ./timeWood /usr/local/timerwood
> cd /usr/local/timerwood
> ln -s /usr/local/timerwood/main /usr/bin/timerwood
> ln -s /usr/local/timerwood/main /usr/bin/ti
> ./install

## usage
> ti "python test/test.py"                            # 计算脚本运行的时间 Calculate the running time of the script
> ti "python test/test.php"
> ti "ls"                                             # 计算Linux命令运行的时间 Calculate the running time of the Linux command
> ti "mysqldump ..."                                  # 计算mysql导出时间
> ti "wget http://..."                                # 计算wget下载的时间

## log
> cat /.Timerwood.dat
