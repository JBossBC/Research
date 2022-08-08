# K8S

## Quesion 

swap分区的作用
ipvs??转发表,比iptables的转发效率更高



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


授权、运行、检查是否加载

chmod 755 /etc/sysconfig/modules/ipvs.modules && bash /etc/sysconfig/modules/ipvs.modules && lsmod | grep -e ip_vs -e nf_conntrack_ipv4

