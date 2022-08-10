# Deeply Realise Cache

## why we don't set a great deal of cache to balance the speed of memory and CPU? 

we should realise a question when we resolve this question. less cache will give us what disadvantage and benefit,and the answer lies that we will take less time to handle the question about the sync of caches that is to say more cache will take more time to synchronize the caches,but it’s undeniable that more cache will give us faster access speed,because it can better balance the IO rate between CPU and memory,when CPU need data which comes from memory.In summary,the number of cache need be setted in a suitable number so that computer can flexibly handle the speed of between synchronization and putting data.

## the policy of cache

when we handle the code,such x+y+z,we will obtain extremely slow  efficiency if the CPU fetches data one by one.When caching cache data,it is based on cache line that is to say memory will synchronize the data 'x' which includes the x and data around x,which can provide deeply efficiency hit data. 


## Cache consistency question

This question appear in multi core CPU.We can assume a situation,that We have thread A and thread B,which execute in parallel.thread A need data x,and thread B need data Y.the interesting thing is that data A and data B are located close to each other in memory,so that Cache L1 cache the other one when Cache L1 cache one(Cache the size of a cache is typically 64 bytes at each time).When thread A and thread B execute in parallel,how to ensure data consistency.This is somewhat similar to the transaction consistency in database.

## Cache Consistency Protocol

Different CPU core maybe use different Cache Consistency Protocol.Here we mainly introduce the MESI Procotol.In MESI Procotol,every cache line has four state,whic are represented by two bytes.As shown below

+ Invalid:the cache line is either already out of the cache or it's content is outdated.In order to achieve cache,the cache line which belongs to Invalid state will be ignored.In other words,the cache line which belongs to the state above is equal to no loading into Cache.
+ shared:the context of cache line  is consistent with  memory.Cache in this state only be read,but can't be written.MultiGroup cache can possess a cache line which comes from the same memory location  at the same time.
+ Exclusive:the cache line has the same effect with Shared cache line,that the cacheline's  context is consistent with  memory.The different place lies the others processor cant possess this cache line when processor has the cache line which belongs to the state above.if  processor has this cache line,the cache line in the others processor will become the invalid state.
+ Modified:the context of cache line is updated by processor.the cache line which belongs to the others processor become the invalid state,if processor update this cache line,the cache line in the others processor will become the invalid state.In additional,the context of modified state cache line should write to memory before  the modified state cache line is abandoned or become invalid.