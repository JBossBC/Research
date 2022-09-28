# Axios

axios是一个基于promise的http库,类似于jq的ajax,用于http请求，可以应用于浏览器端和node.js。


## 特性

+ 支持promise API
+ 拦截请求与响应,比如:在请求前添加授权和响应前做一些事情。
+ 转换请求数据和响应数据,比如:进行请求加密或者响应数据加密
+ 取消请求
+ **自动转换JSON数据**
+ 客户端支持防御XSRF??


**在axios中http中的context-type可以根据提交数据自动设置**


## axios常用的请求方法

方法列举:get、post、put、patch、delete

并发请求:axios.all()


## axios实例相关配置:

+ baseURL:请求的域名(基本地址)
+ timeout:请求的超时时长,超出后后端返回401(一般由后端定义,后端的接口需要处理时长较长的时候,如果请求的时长过长,后端处理不过来，就会阻塞，给服务器带来较大的压力。设置后，可以及时释放掉)
+ url:请求路径
+ method:请求方法
+ headers:请求头
+ params:将请求参数拼接到url上
+ data:将请求参数放置到请求体中
+ transformRequest:[function(data,headers){}],
+ transformResponse:[function(data){}],
+ timeout:number,
+ responseType:,

## axios响应
    
    {
      // `data` 由服务器提供的响应
      data: {},
    
      // `status` 来自服务器响应的 HTTP 状态码
      status: 200,
    
      // `statusText` 来自服务器响应的 HTTP 状态信息
      statusText: 'OK',
    
      // `headers` 服务器响应的头
      headers: {},
    
       // `config` 是为请求提供的配置信息
      config: {},
     // 'request'
      // `request` is the request that generated this response
      // It is the last ClientRequest instance in node.js (in redirects)
      // and an XMLHttpRequest instance the browser
      request: {}
    }


## 配置的优先顺序

配置会以一个优先顺序进行合并。这个顺序是:在lib/default.js找到的库的默认值,然后是实例的defaults属性,最后是请求的config参数,后者优先于前者

## 拦截器

在请求或响应被then或catch处理前拦截它们

> axios.interceptors.request
> axios.interceptors.response