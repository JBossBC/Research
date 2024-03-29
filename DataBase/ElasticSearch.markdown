# ElasticSearch

Elasticsearch是一个开源的搜索引擎，建立在一个全文搜索引擎库Apache Lucene基础之上。Lucene可以说是当下最先进、高性能、全功能的搜索引擎库。

Elasticsearch内部使用Lucene做索引与搜索，但是他的目的是使全文检索变得简单，通过隐藏Lucene的复杂性，取而代之的提供一套简单一致的RESTful API


然后，Elasticsearch不仅仅是lucene，并且也不仅仅是一个全文搜索引擎。它可以被下面这样准确的形容:

+ 一个分布式的实时文档存储，每个字段可以被索引与搜索
+ 一个分布式实时分析搜索引擎
+ 能胜任上百个服务节点的扩展，并支持PB级别的结构化或者非结构化数据。


## 面向文档

在应用程序中对象很少只是一个简单的key、value的列表。通常，他们拥有更复杂的数据结构，可能包括日期、地理信息、其他对象或者数组等。

使用关系型数据库的行和列存储，这相当于是把一个表现力丰富的对象塞到一个非常大的电子表格中:为了适应表结构，你必须设法将这个对象扁平化-通常一个字段对应一列而且每次查询时又需要将其重新构造为对象。


Elasticsearch是面向文档的，意味着它存储整个对象或文档。elasticsearch不仅存储文档，而且索引每个文档的内容，使之可以被检索。在Elasticsearch中，我们对文档进行索引、检索、排序和过滤-而不是对行列数据。这是一种完全不同的思考数据的方式，也是elasticsearch能支持复杂全文检索的原因。


### Json

elasticsearch 使用JavaScript object notation 作为文档的序列化格式。json序列化为大多数编程语言所支持，并且已经成为nosql领域的标准格式。它简单、简介、易于阅读。