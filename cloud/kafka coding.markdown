# kafka coding

## the important configure for kafka


1. zookeeper.connect

     the param indicates the cluster of zookeeper(including the port). it havnt default value,at the same time,it's required.When  it has multiple nodes in the zookeeper cluster,you can separate with commas.

2. listeners

     the param indicates the broker monitor the address of connecting the client,configuration format is protocol://hostname:port,...。The kafka support the type of protocol including PLAINTEXT、SSL、SASL_SSL and so on.If cant open the security certification,it can use the simple protocol: PLAINTEXT. you can use the commas to support multiple address.

     advertised.listeners connected with listeners,default value is null.the difference place is the advertised.listeners are useful to IaaS environment.In this situation,you can set advertised.listeners to bind the public network for the use of external client.For this situation,the listeners params be set for binding the private IP which be used by broker 

3. broker.id

     the param indicates the unique identification in the cluster of kafka,default value is -1.this params are related to meta.properties and broker.id.generation.enable,reserved.broker.max.id.

4. log.dir and log.dirs

     the all of news save on disk in kafka, the params above be set to configure the root of log file.log .dir  can configure the single catalogue，and the logs.dir can configure the multiple catalogue(separating with commas).the default value is /tmp/kafka-logs


5. message.max.bytes

    this param indicates  the maximum of broker which can receive the news(default value is 1000012(B)),if producer send the news which length gether than  this maxValue,the producer client will thorw the recordTooLargeException error.If you will update this param,you should consider the  max.request.size(Client param) 、max.message.bytes(topic param).In order to avoid a series of influence which  are caused by updating this param,i suggest that you should consider 
    the feasibility of splitting messages.



## kafka producer init required param

+ **bootstrap.servers**:indicate the address which producer client need to connect  in kafka cluster.you should set greater than two broker to keep the highly available.
+ **key.serializer and value.serializer**:the news which is accepted by broker must keep the form of byte array.



## kafka producer builder news(create producerRecord object)




## kafka producer send news

+ sync
+ async
+ fire-and-forget(send messages to kafka without caring whether the message arrive correctly)

+ send function returns a Future object.
    
     Future represents a life cycle of task,and provides the corresponding method to judge the task whether finish or cancel.

      kafka producer will thorw two types of exceptions: retryable exception and non-retryable exception.For retryable exception,you can set retries param to avoid this situation.As long as the exception recover in the specified  times,the system cant throw exception





 



