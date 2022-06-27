# Run Program

## Question

<<<<<<< HEAD
eclipse是什么
=======
git -rebase的使用?

>>>>>>> 4090764 (0.1.0)
## 环境

集成环境:eclipse
jdk版本:1.8.0
maven版本:3.8.4

##注意:使用前必须进入专属VPN，否则获取不到相应资源##

+ 从SVN服务器上面导入项目 

代码地址:http://10.0.10.11:3344/svn/pt2022/trunk/


+ 将项目通过eclipse打开，开始注入maven依赖包。
此项目的部分依赖包是利用nexus搭建maven私服，所以我们不再一味从中央仓库中导入资源，必须配置maven的settings文件来标识私服的位置并进入所需要的权限配置。

   + 我们不再使用eclipse自定义的maven配置，通过在Window下面的Preferences设置maven的路径(Installations)以及全局配置文件(User Settings)，全局配置文件在http://10.0.10.11:3344/svn/pt2022/trunk/04_代码中的maven_settings.xml里面
   + 设置完自定义的maven配置后，我们使用maven build进行依赖导入就可以使用了