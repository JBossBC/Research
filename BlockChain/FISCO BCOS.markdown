# FISCO BCOS

## FISCO BCOS 环境搭建

支持的系统环境
Ubuntu 16.04+


安装依赖openssl、curl

>** ubuntu**
> 
> sudo apt install -y openssl curl
> 

在fisco-bcos操作目录下执行下载脚本

    curl -#LO https://github.com/FISCO-BCOS/FISCO-BCOS/releases/download/v2.7.1/build_chain.sh
将下载的脚本设置权限

    chmod u+x build_chain.sh

执行脚本,生成4节点的fisco链,最后输出All completed，表示节点生成成功

    bash build_chain.sh -l 127.0.0.1:4 30300,20200,8545

启动所有节点
    
    bash nodes/127.0.0.1/start_all.sh


## 配置以及使用控制台

+ 安装Java环境
+ 获取控制台下载脚本，并执行下载控制台
   ` curl -LO https://github.com/FISCO-BCOS/console/releases/download/v2.7.1/download_console.sh && bash download_console.sh`
+ 拷贝控制台配置文件
`cp -n console/conf/config-example.toml console/conf/config.toml
`
+ 配置控制台证书
`cp -r nodes/127.0.0.1/sdk/* console/conf/`
+ 到console目录下面启动并使用控制台
	`bash start.sh	`
+ 进入控制台后
     
       获取客户端版本`getNodeVersion`
       获取节点信息`getPeers`


## 部署智能合约以及调用

### 部署智能合约
`deploy 合约名` 

example合约在console/contracts/solidity里面

**部署HeeloWorld.sol合约**
`deploy HelloWorld`

### 调用合约

+ **查看当前块高** `getBlockNumber`

+ **调用get接口获取name变量地址 此处的合约地址是deploy指令返回的地址** `call HelloWorld 0xf192536949b27bf4c78e051ba1244deac2814f4a get`

+ **调用set接口改变name** `call HelloWorld 0xf192536949b27bf4c78e051ba1244deac2814f4a set "Hello,FISCO BCOS"`

+ **再次调用get接口查看是否生效** `call HelloWorld 0xf192536949b27bf4c78e051ba1244deac2814f4a get`

+ **查看BlockNumber的变化** `getBlockNumber`

+ **退出控制台** `quit`

