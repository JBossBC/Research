# Docker

**刚开始使用Docker时，建议以宽容(permissive)模式运行SELinux,这样SElinux将只把错误写进日志，而非强制执行。如果以强制(enforcing)模式运行SELinux，那么很有可能会遇到各种莫名奇妙的permission denied错误**(这是我经常遇到的错误)


> 查看SELinux处于什么模式，可以通过执行sestatus命令的结果得知
> 要将SELinux设为宽容模式，只需执行sudo setenforce 0


因为Docker运行时需要特殊权限，所以默认执行命令时都必须在前面加上sudo。但这样做确实使人厌烦，一个可行的解决方法是把用户放进docker用户组里。在ubuntu下你可以输入`sudo usermod -aG docker $USER`
如果docker用户组不存在，这个命令会创建它，并且把当前用户添加到组里。然后，你需要先注销并重新登入系统。