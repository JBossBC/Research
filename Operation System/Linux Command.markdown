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


### 防火墙操作

+ 关闭防火墙:systemctl stop firewalld
+ 开机禁用防火墙 systemctl disable firewalld
+ 开机启动防火墙 systemctl enable firewalld
+ 开启防火墙 systemctl start filewalld
+ 检查防火墙状态 systemctl status firewalld

### 使用firewall-cmd配置端口

+ 查看防火墙状态:firewall-cmd --state
+ 重新加载配置:firewall-cmd --reload
+ 查看开放的端口: firewall-cmd --list-ports

### diff(比较文件内容差异)

diff[参数]文件1 文件2

常用参数

-a 	逐行比较文本文件

-b 	不检查空格字符的不同

-W	指定栏宽

-x	不比较选项中所指定的文件或目录

-X	将文件或目录类型存成文本文件

-y 	以并列的方式显示文件的异同之处

--brief	仅判断两个文件是否不同

--help 	查看帮助信息

--left-column 	若两个文件某一行内容相同，则仅在左侧的栏位显示该行内容

--suppress-common-lines 	在使用-y参数时，仅显示不同之处 


## less

-b 设置缓冲区的大小。

-i 忽略搜索时的大小写


常规操作

/string:向下搜索字符串的功能

?string:向上搜索字符串的功能

b:向上翻一页

d:向后翻一页

g:移动到第一行

G:移动到最后一行

&pattern:仅显示匹配模式的行，而不是整个文件


## free -m 查看内存和swap分区情况


## 使用dns查看域名ip (nslookup 域名)