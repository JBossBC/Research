# Docker

+ Docker daemon(Docker守护进程)

Docker daemon是一个运行在宿主机(DOCKER_HOST)的后台进程。我们可通过Docker客户端与之通信

+ Client(Docker客户端)

Docker客户端是Docker的用户界面，它可以接收用户命令和配置标识，并与Docker daemon通信。Docker build等都是Docker的相关命令

+ Images(Docker 镜像)

Docker镜像是一个只读模板，它包含创建Docker容器的说明。它和系统安装光盘有点类似--我们使用系统安装光盘安装系统。同理，我们可以使用Docker镜像运行Docker镜像中的程序。

+ Container

容器是镜像的可运行实例。镜像和容器的关系有点类似于面向对象中，类和对象的关系。我们可以通过Docker API或者CLI命令来启停、移动、删除容器


+ Registry

Docker Registry是一个集中存储与分发镜像的服务。我们构建完Docker镜像后，就可在当前宿主机上运行，但如果想要在其他机器上运行这个镜像，我们就需要手动拷贝。此时，我们可借助Docker Registry来避免镜像的手动拷贝

一个Docker Registry包含多个Docker仓库；每个仓库可包含多个镜像标签;每个标签对应一个Docker镜像。这与maven的仓库有点类似。如果把Docker Registry比作Maven仓库的话，那么Docker仓库就可以理解为某jar包的路径，而镜像标签则可理解为jar包的版本号。

Docker Registry可分为公有和私有两种，最常见的Docker Registry莫过于官方的Docker hub。这也是默认的Docker Registry

----

**刚开始使用Docker时，建议以宽容(permissive)模式运行SELinux,这样SElinux将只把错误写进日志，而非强制执行。如果以强制(enforcing)模式运行SELinux，那么很有可能会遇到各种莫名奇妙的permission denied错误**(这是我经常遇到的错误)


> 查看SELinux处于什么模式，可以通过执行sestatus命令的结果得知
> 要将SELinux设为宽容模式，只需执行sudo setenforce 0


因为Docker运行时需要特殊权限，所以默认执行命令时都必须在前面加上sudo。但这样做确实使人厌烦，一个可行的解决方法是把用户放进docker用户组里。在ubuntu下你可以输入`sudo usermod -aG docker $USER`
如果docker用户组不存在，这个命令会创建它，并且把当前用户添加到组里。然后，你需要先注销并重新登入系统。