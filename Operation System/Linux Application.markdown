# Linux Application

## Quesion 
free？
lsmod ?

systemctl?
sysctl?

net.bridge.bridge-nf-call-ip6tables =1 
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
vm.swappiness = 0

## SELINUX

SELinux(security enhanced linux)安全增强型linux系统，它是一个linux内核模块，也是linux的一个安全子系统

selinux的主要作用就是最大限度地减少系统中服务进程可访问地资源(最小权限原则)

SELinux有两个级别:

setenforce 0:表示警告(Permissive)
setenforce 1:表示强制(Enforcing)

命令: sestatus:查看SELinux状态
getenforce:查看当前SELinux级别


## swap分区

/etc/fstab
free -m



## 网桥过滤功能

添加网桥过滤以及地址转发
cat /etc/sysctl.d/k8s.conf

net.bridge.bridge-nf-call-ip6tables =1 
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
vm.swappiness = 0

加载br_netfilter模块 modprobe br_netfilter

查看是否被加载 lsmod | grep br_netfilter

加载网桥过滤配置文件 sysctl -p /etc/sysctl.d/k8s.conf


## 配置yum源为清华源

+ 修改/etc/yum.repo.d/CentOS.Base文件