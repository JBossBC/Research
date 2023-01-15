# redis

### ping 检测redis服务是否启动

## redis key

+ del key 删除一个键
+ dump key 序列化给定得key,并返回被序列化的值
+ exists key 检查给定的key是否存在
+ expire key seconds 为给定的key设置过期时间
+ expireat key timestamp 与上面的类似，不同在于这个接收时间参数是unix时间戳
+ pexpire 以毫秒为单位
+ pexpireat key milliseconds-timestamp 设置key过期时间的时间戳以毫秒计
+ keys pattern 查找所有符合给定模式的key
+ move key db 将当前数据库的key移动到给定的数据库中
+ persist key:移除key的过期时间，key将持久保存
+ pttl key 以毫秒为单位返回key的剩余的过期时间
+ ttl key 以秒为单位，返回给定key的剩余生存时间
+ type key 返回key存储的值类型
+ rename key newkey 修改key的名称
+ renamenx key newkey 仅当newkey不存在时，将key改名为newkey

## redis string

+ set key value:设置指定key的值
+ get key: 获取指定key的值
+ getrange key value 返回key中字符串的子支付
+ getset key value 将给定的key设为value，并返回key的旧值
+ mget key1 key2 获取多个key的值
+ setnx key value 当key不存在时设置key的值
+ strlen key 返回key所存储的字符串的长度
+ mset key value key value 同时设置多个key-value对
+ msetnx key value key value 同时设置多个key-value对，仅当原key不存在时
+ append key value 如果key已经存在并且是个字符串，append将指定的value追加到key原来的值的末尾
+ incr key 将key中存储的数字值+1
+ decr key 将key中存储的数字值-1
+ PSETEX key milliseconds value 这个命令和 SETEX 命令相似，但它以毫秒为单位设置 key 的生存时间，而不是像 SETEX 命令那样，以秒为单位。
+ setrange key offset value 用value参数覆写给定key所存储的字符串，从偏移量offset开始

## redis hash

+ hdel key field1 field2 删除一个或多个哈希表字段
+ hexists key field 查看哈希表key中指定的字段是否存在
+ hget key field 获取在哈希表中指定字段的值
+ hgetall key 获取在哈希表中指定key的所有字段和值
+ hkeys key 获取哈希表key中所有的key
+ hlen key 获取哈希表key中字段的数量
+ hmget key field1 field2 获取所有给定字段的值
+ hmset key field1 value1 field2 value2 同时将多个 field-value对设置到哈希表key中
+ hsetnx key field value 只有当field不存在时，设置哈希表字段的值
+ hvals key 获取哈希表中的所有值

## redis list

Redis列表是简单的字符串列表，按照插入顺序排序。你可以添加一个元素到列表的头部（左边）或者尾部（右边）

一个列表最多可以包含 232 - 1 个元素 (4294967295, 每个列表超过40亿个元素)。

+ blpop key1 key2 timeout 移除并获取列表的第一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
+ brpop key1 key2 timeout 移除并获取列表的最后一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
+ brpopblpop source destination timeout 从列表中弹出一个值，将弹出的元素插入到另一个列表中并返回它；如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止
+ lindex key index 通过索引获取列表中的元素
+ linsert key before|after pivot value 在列表元素前或后插入元素
+ llen key 获取列表长度
+ lpop key 移除并获取列表中的第一个元素
+ lpush key value1 values 将一个值插入到列表头部
+ lpushx key value 将一个值插入到已存在的列表头部
+ lrange key start top 获取列表指定范围内的元素
+ lrem key count value 移除列表元素中与value相等的值(count>0从表头开始，移除数量为count;count<0从表尾开始，移除数量为count的绝对值)
+ lset key index value 通过索引设置列表元素的值
+ rpop key 移除列表的最后一个元素，返回值为移除的元素
+ rpoplpush source destination 移除列表的最后一个元素，并将该元素添加到另一个列表并返回
+ ltrim key start stop 对一个列表进行修剪，只保留指定区间内的与那苏，不在指定元素之间的元素都将被删除。
+ rpush key value1 value2 在列表中添加多个值到列表尾部
+ rpushx key value 为已经存在的列表添加值

## redis set

Redis 的 Set 是 String 类型的无序集合。集合成员是唯一的，这就意味着集合中不能出现重复的数据。

集合对象的编码可以是 intset 或者 hashtable。

Redis 中集合是通过哈希表实现的，所以添加，删除，查找的复杂度都是 O(1)。

集合中最大的成员数为 232 - 1 (4294967295, 每个集合可存储40多亿个成员)。

+ sadd key member1 member2 向一个集合中添加多个值
+ scard key 获取集合的成员数
+ sdiff key1 key2 返回第一个集合与其他集合之间的差异
+ sinter key1 key2 返回给定的所有集合的交集
+ smembers key 返回集合中的所有成员
+ smove source destination member 将member元素从source集合中移动到destination集合
+ spop key 移除并返回集合中的一个随即元素
+ srandmember key count 返回集合中一个或者多个随机数
+ srem key member1 member2 移除集合中一个或多个成员
+ sunion key1 key2 返回所有给定集合的并集


## redis sorted set

Redis 有序集合和集合一样也是 string 类型元素的集合,且不允许重复的成员。

不同的是每个元素都会关联一个 double 类型的分数。redis 正是通过分数来为集合中的成员进行从小到大的排序。

+ zadd key score1 member1 score2 member2: 向有序集合添加一个或多个成员，或者更新已经存在的成员的分数
+ zcard key 获取有序集合的成员数
+ zcount key min max 计算在有序集合中指定区间分数的成员数
+ zincrby key increment member:有序集合中对指定成员的分数加上increment
+ zrange key start stop withscores 通过索引区间返回有序集合指定区间内的成员
+ zrank key member 返回有序集合中指定成员的索引
+ zrem key member member2 移除有序集合中一个或者多个成员。
+ zscore key member 返回有序集中，成员的分数值
+ zrevrank key member 返回有序集合中指定成员的排名，有序集成员按分数值递减(倒序)
+ zunionstore destination numkeys key key2 计算一个或者多个有序集的并集，并存储到新的key中。
+ zscan key cursor match pattern count count 迭代有序集合中的元素(包括元素成员和分值),能增加搜索条件。
+ zremrangebyrank key start stop 移除有序集合中给定的排名区间内的所有成员
+ zremrangebyscore key start stop 移除有序集合中给定的分数区间内的所有成员