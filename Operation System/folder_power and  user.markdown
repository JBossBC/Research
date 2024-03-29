# 文件夹权限和用户


## etc/passwd 

/etc/passwd文件是系统用户配置文件，存储了系统中所有用户的基本信息，并且所有用户都可以对此文件执行r操作。


每行用户信息都以":"作为分隔符，划分为7个字段


|用户名|密码|UID|GID|描述信息|主目录|默认shell|
|--|--|--|--|--|--|--|
|root|x|0|0|root|/root|/bin/bash



**linux系统是通过UID来识别用户身份，分配用户权限./etc/passwd文件中定义了用户名和UID之间的对应关系。**



"x":表示此用户设有密码，但不是真正的密码,真正的密码保存在/etc/shadow文件.**linux把真正的加密密码串放置在/etc/shadow文件中，此文件只有root用户可以浏览和操作，最大限度地保证了密码地安全**


虽然x并不代表真正的密码，但如果删除了，那么系统会认为这个用户没有密码，从而导致只输入用户名而不输入密码就可以登录。

### UID

UID是一个0~65535之间的数(和端口号范围一致)，不同范围的数字表示不同的用户身份。


+ 0:超级用户
+ 1~499:系统用户
+ 500~65535:普通用户


### GID


表示用户初始组组ID号


1. 初始组

     指用户登录时就拥有这个用户组的相关权限。每个用户的初始组只能有1个，通过就是将和此用户的用户名相同的组名作为该用户的初始组

2. 附加组

     用户可以加入多个其他的用户组，并拥有这些组的权限。每个用户只能有一个初始组，除初始组外，用户再加入其他的用户组，这些用户组就是这个用户的附加组。附加组可以有多个，而且用户可以有这些附加组的权限。


### 描述信息

解释这个用户的意义


### 主目录

通常称为这个用户的家目录

### 默认shell

shell是linux的命令解释器，用户和linux内核之间沟通的桥梁。shell命令解释器的功能就是将用户输入的命令转换成系统可以识别的机器语言。通过情况下，linux系统默认使用的命令解释器就是/bin/bash



## 用户


linux系统是一个多用户多任务的分时操作系统，任何一个要使用系统资源的用户，都必须首先向系统管理员申请一个账号，然后以这个账号的身份进入系统。用户的账号一方面可以帮助系统管理员对使用系统的用户进行跟踪，并控制他们对系统资源的访问；另一方面也可以帮助用户组织文件，并为用户提供安全性保护。


实现用户账号的管理，要完成的工作主要有:

+ 用户账号的添加、删除和修改
+ 用户口令的管理
+ 用户组的管理



1. linux系统用户账号的管理

+ 用户账号的添加

    useradd 选项 用户(选项与/etc/passwd的参数设立有关，所属用户组、附加组、描述信息、主目录、shell、用户号)

+ 用户账号的删除
+ 

     userdel 选项(-r 将用户的主目录一起删除) 用户名(此命令删除用户sam在系统文件中（主要是/etc/passwd, /etc/shadow, /etc/group等）的记录，同时删除用户的主目录。) 

+ 修改账号

     usermod 选项 用户名  [修改用户账号就是根据实际情况更改用户的有关属性，如用户号、主目录、用户组、登录shell等(命令选项与useradd一样,可以为用户指定新的资源值)]

+ 用户口令的管理

     用户管理的一项用户口令的管理。用户账号刚创建时没有口令，但是被系统锁定，无法使用，必须为其指定口令后才可以使用，指定和修改用户口令的shell命令时passwd。超级用户可以为自己和其他用户指定口令。普通用户只能修改他自己的口令(passwd 选项 用户名 -i 锁定口令 -u 口令解锁 -d 使账号无口令 -f 强迫用户下次登录时修改口令)


## linux系统用户组的管理

每个用户都有一个用户组，系统可以对一个用户组中的所有用户进行集中管理。不同Linux 系统对用户组的规定有所不同，如Linux下的用户属于与它同名的用户组，这个用户组在创建用户时同时创建。

用户组的管理涉及用户组的添加、删除和修改。组的增加、删除和修改实际上就是对/etc/group文件的更新




1. 增加一个用户组:groupadd命令

     groupadd 选项 用户组(-g 组号) 

2. 