# JWT(Json web Token)


JWT相比于传统的session认证方式更节约服务器资源,并且对移动端和分布式更加友好(session会使用服务器内存资源),其优点如下:

+ 支持跨域访问:cookie是无法跨域的,而token由于没有用到cookie(前提是将token放到请求头中)，所以跨域后不会存在消息丢失问题
+ 无状态:token机制在服务端不需要存储session信息,因为token自身包含了所有登录用户的信息,所以可以减轻服务端压力
+ 更适用CDN:可以通过内容分发网络请求服务端的所有资料
+ 更适用于移动端:当客户端是非浏览器平台时,cookie是不被支持的，此时采用token认证方式会简单很多
+ 无需考虑CSRF:由于不再依赖cookie,所以采用token认证方式不会发生CSRF,所以也无需考虑CSRF防御

通俗来说,JWT本质上就是一个字符串,它是将用户信息保存到一个Json字符串中，然后进行编码后得到一个JWT token，并且这个JWT token带有签名信息，接收后可以校验是否被篡改，所以可以用于在各方之间安全地将信息作为Json对象传输。

## JWT认证过程

1. 首先，前端通过web表单将自己的用户名和密码发送到后端的接口，这个过程一般是一个POST请求。建议的方式是通过HTTPS，从而避免敏感信息被嗅探
2. 后端核对用户名和密码成功后,将包含用户信息的数据作为JWT的payload,将其与JWT Header分别进行Base64编码拼接后签名，形成一个JWT Token,形成的JWT Token就是一个如同111.ZZZ.XXX的字符串
3. 后端将JWT Token字符串作为等率成功的结果返回给前端。前端可以将返回的结果保存在浏览器中，退出登录时删除保存的JWT Token即可
4. 前端在每次请求时将JWT Token放入HTTP请求头中的Authorization属性中(解决XSS和XSRF问题)
5. 后端检查前端传过来的JWT Token,验证其有效性,比如检查签名是否正确、是否过期、token的接收方是否是自己等等
6. 验证通过后,后端解析出JWT Token中包含的用户信息，进行其他逻辑操作(一般是根据用户信息得到权限等),返回结果

![](https://img-blog.csdnimg.cn/img_convert/900b3e81f832b2f08c2e8aabb540536a.png)