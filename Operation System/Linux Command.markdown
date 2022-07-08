# Linux Command

### 添加环境变量

`export 需要添加的路径:$PATH` 


### Nano的使用

Nano是一个简单的没有花梢及华丽效果的文本编辑器。Nano在做简单文本文件编辑时相当不错,,可以满足一些基本操作.

直接在命令行模式下输入nano 文件名，就可以对文件进行简单的编辑。

文件编辑中常用快捷键：ctrl+X 离开nano软件，若有修改过的文件会提示是否保存；

ctrl+O 保存文件；   ctrl+W 查询字符串；

ctrl +C 说明目前光标所在处的行数和列数等信息；

ctrl+ _ 可以直接输入行号，让光标快速移到该行；

### netstat

netstat -tunlp 用于限制tcp，udp的端口和进程等相关情况

`netstat -tunlp | grep 端口号`

+ -t(tcp)仅显示tcp相关选项
+ -u(udp)仅显示udp相关选项
+ -n 拒绝显示别名，能显示数字的全部转化为数字
+ -l 仅列出在listen(监听)的服务状态
+ -p 显示建立相关链接的程序名

### kill 

kill ??

### 统计当前目录下的文件数量(包括子文件)

`ls -lR|grep "^-"| wc -l`

