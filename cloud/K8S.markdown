# K8S

## Quesion 

swap分区的作用
ipvs??转发表,比iptables的转发效率更高
wget和yum的作用?
scp??


## 步骤

1. 准备主机
2. 主机名设置(hostname /hostnamectl set-hostname)
3. ip设置，主机dns设置(/etc/sysconf/network-srcipts/ens33)，主机名解析(/etc/host)，防火墙设置
4. 主机安全设置
5. 主机时钟同步
6. 关闭swap分区
7. 配置网桥过滤功能
8. 配置主机ipvs功能


### 防火墙设置

systemctl stop firewalld 关闭防火墙
systemctl disable firewalld 开机禁止防火墙
systemctl enable firewalld   开机启动防火墙
systemctl start firewalld  启动防火墙

systemctl status firewalld 查看防火墙状态


### 防火墙命令

firewall-cmd --参数



## 时钟同步

crontab -e 编辑定时任务
ntpdate 需要自己安装
yum -y install ntpdate
ntpdate time1.aliyun.com 同步时钟

### 定时任务日志

/var/log/cron.log


### 永久关闭swap分区

使用kubeadm部署必须关闭swap分区，修改配置文件后需要重启操作系统


### 安装ipset以及ipvsadm

yum -y install ipset ipvsadm


### 在所有节点执行如下脚本

**添加需要加载的模块**

cat > /etc/sysconfig/modules/ipvs.modules << EOF

\#!/bin/bash

modprobe -- ip_vs

modprobe -- ip_vs_rr

modprobe -- ip_vs_wrr

modprobe -- ip_vs_sh

modprobe -- nf_conntrack_ipv4

EOF


## 授权、运行、检查是否加载

chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack_ipv4

## 安装docker-ce

+ yum源获取

`wget -O /etc/yum.repos.d/docker-ce.repo https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/centos/docker-ce.repo`
    
### 清理yum的缓存 
1. yum clean all
2. yum makecache
3. yum update 


+ 查看docker-ce版本列表


**好像只有清华源有docker-ce**

    yum list docker-ce.x86_64 --showduplicates| sort -r

+ 安装指定版本docker-ce

    yum -y install --setopt=obsoletes=0 docker-ce-18.06.3.ce-3.el7

+ 启动docker服务器

systemctl enable docker
systemctl start docker

+ 修改docker-ce服务配置文件(为了后续使用/etc/docker/daemon.json来进行更多的配置)

     + 将/usr/lib/systemd/system/docker.service文件中`删除ExecStart=/usr/bin/dockerd后面有-H选项的所有内容(包含-H)`
     + 在/etc/docker/daemon.json中添加下面内容`{"exec-opts":["native.cgroupdriver=systemd"]}`

+ 重启docker服务

systemctl restart docker
systemctl status docker


### 安装kubeadm kubelet kubectl

注: kubeadm:初始化集群、管理集群等
kubelet:用于接收api-server指令，对pod声明周期进行管理
kubectl:集群命令行管理工具

[kubernetes]
name=Kubernates
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg           https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg


如果签名未通过，则把gpgcheck和gpgkey设置为0同时去掉gpgkey

**可以使用scp进行服务器与服务器之间的文件复制**
`scp /etc/yum.repo.d/k8s.repo work1:/etc/yum.repo.d/k8s.repo`

验证:yum list |grep kubeadm


安装上面三个软件:yum -y install kubeadm kubelet kubectl


### 软件设置

主要配置kubelet,如果不配置可能会导致k8s集群无法启动

> 为了实现docker使用的cgroupdrive与kubelet使用的cgroup的一致性，建议修改如下文件内容


`vim /etc/sysconfig/kubelet`
`KUBELET_EXTRA_ARGS="--cgroup-drivesystemd"`

> 设置为开机自动启动就可以了，因为没有生成配置文件，集群初始化后自动启动

`systemctl enable kubelet`


### k8s集群容器镜像准备(需要科学网上方式)

> 因为使用kubeadm部署集群,集群所有核心组件均以pod运行，需要为主机准备镜像，不同角色主机准备不同镜像。

+ Master主机的镜像

kubeadm config images list(查看集群使用的镜像)

将需要的镜像制作成脚本文件一键下载

`kubeadm config images list >> image.list`

`vi image.list `

    #! /bin/bash 
    img_list='image.list原先的内容'    
    for img in $image.list
    do
        docker pull $img
    done


`sh  image.list`(上面的k8s相关镜像需要科学上网获取)

