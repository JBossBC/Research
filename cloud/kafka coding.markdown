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

    this param indicates  the maximum of broker which can receive the news(default value is 1000012(B)),if producer send the news which length gether than  this maxValue,the producer client will thorw the recordTooLargeException error.If you will update this param,you should consider the 
