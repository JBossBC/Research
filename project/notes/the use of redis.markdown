# the use of redis

## Question
hash?

## redis数据类型

redis支持五种数据类型:string(字符串)，hash(哈希),list(列表),set(集合),zset(有序列表)

+ String(字符串)

string是redis最基本的类型，一个key对应一个value

string类型是二进制安全的，可以包含任何数据，比如jpg图片或者序列化的对象

string类型是redis最基本的数据类型，string类型的值最大能存储512MB

+ Hash(哈希)

redis hash是一个键值(key=>value)对集合

redis hash 是一个string类型的field和value1映射表，hash特别适合用于存储对象

实例中我们使用了redis HMSET、HGET命令，HMSET设置了两个field=>value对，HGET获取对应field对应的value
**每个hash可以存储2^32-1键值对**

+ List(列表)

redis列表是简单的字符串列表，按照插入顺序排序，你可以添加一个元素到列表的头部或者尾部

列表最多可以存储2^32-1元素

+ Set(集合)

Redis的set集合是string类型的无序集合

集合是通过哈希表实现的，所以添加，删除，查找的时间复杂度都是O(1)

+ zset(有序集合)

redis zset和set一样也是string类型元素的集合，且不允许重复的元素

不同的是每个元素都会关联一个double类型的分数，redis正是通过分数来为集合中的成员进行从小到大的排序

zset的成员是唯一的，但分数可以重复

redis -h host -p port -a  password

info 获取服务器的统计信息

select number 切换到指定的数据库
config get * 获取redis服务器所有的配置项,包括密码
config set configName configValue 设置config的值

keys * 获取当前数据库所有的key值

ping 查看服务是否运行

echo message打印字符串

auth password 验证密码是否正确