# Golang package management


**import导入包时，包名是从GOPATH开始计算的路径，使用/进行路径分隔;**
在使用IDE时，不用手动import包,例如使用fmt包，在敲入fmt.时，编译器会警告，此时敲回车，包会自动import

> import是用一个独一无二的字符串路径来指向包，而包的导入路径是基于工作目录的,是因为会在工作目录的src下查找包,即在GOROOT(或GOPATH)/src下找

vendor:随着Go 1.5 release版本的发布,vendor目录被添加到除了GOPATH和GOROOT之外的依赖目录查找的解决方案。在Go1.6之前,你需要手动的设置环境变量GO15VENDOREXPERIMENT-1才可以使GO找到vendor目录，然而在Go1.6之后，这个功能已经不需要配置环境变量就可以实现。即使使用vendor,也必须在GOPATH中,无论是通过IDE设置项目目录的GOPATH还好是通过环境变量设置GOPATH(这与go module 无关)

> + 在执行go build 或 go run,会按照一下规则顺序去查找包
> 
> + GO111MODULE=off时，如果一个包在vendor和$GOPATH下都存在，那么使用顺序为:
> 
>          优先使用vendor目录下的包          
>          如果未找到，则从上级目录的vendor路径下搜索，直到src的vendor路径下面搜索
>          在GOROOT/src目录下查找
>          在GOPATH/src下面查找依赖包
>          要么完整使用vendor下面的包,要么完整使用GOPATH/src下面的包,不会混合使用



## go module

golang提供了一个环境变量“GO111MODULE”，默认值为auto，如果当前目录里有 go.mod 文件，就使用 go modules，否则使用旧的 GOPATH 和 vendor 机制，因为在modules机制下go get只会下载go modules。

modules和传统GOPATH不同，不需要包含src，bin这样的子目录，一个源代码目录甚至是空目录都可以作为module，只要其中包含go.mod文件。

除了go.mod之外，go命令还维护一个名为go.sum的文件，其中包含特定模块版本内容的预期加密哈希，go命令使用go.sum文件确保这些模块的未来下载检索与第一次下载相同的位，以确保项目所依赖的模块不会出现意外更改，无论是出于恶意、意外还是其他原因。 go.mod和go.sum都应检入版本控制。

go.sum不需要手工维护，所以可以不用太关注。


    如果GO111MODULE=off，那么go命令行将不会使用新的module功能，相反的，它将会在vendor目录下和GOPATH目录中查找依赖包。也把这种模式叫GOPATH模式。
    
    如果GO111MODULE=on，那么go命令行就会使用modules功能，而不会访问GOPATH。也把这种模式称作module-aware模式，这种模式下，GOPATH不再在build时扮演导入的角色，但是尽管如此，它还是承担着存储下载依赖包的角色。它会将依赖包放在GOPATH/pkg/mod目录下。
    
    如果GO111MODULE=auto，这种模式是默认的模式，也就是说在你不设置的情况下，就是auto。这种情况下，go命令行会根据当前目录来决定是否启用module功能。只有当当前目录在GOPATH/src目录之外而且当前目录包含go.mod文件或者其子目录包含go.mod文件才会启用。


### go mod 命令介绍

+ 自动下载依赖包
+ 项目不必放在&GOPATH/src内了，modules和传统GOPATH不同，不需要包含src,bin这样的子目录,一个源代码目录甚至是空目录都可以作为module,只要其中包含go.mod文件
+ 项目内会生成一个go.mod文件,列出依赖包;所有来的第三方包会准确的指定版本号
+ 对于已经转移的包,可以使用replace申明替换，不需要修改代码