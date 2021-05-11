## caches
[caching cn](https://aws.amazon.com/cn/caching/)

### caching-challenges-and-strategies
[local caches](https://d1.awsstatic.com/builderslibrary/pdfs/caching-challenges-and-strategies.pdf?did=ba_card-body&trk=ba_card-body)
1. when to use cache
    1. help to decrease the latency
        *. but because of low usage, be scaled down
        *. and then a traffic happened
    2. latency(efficiency) vs tolerant/eventual consistency
2. local cache
    1. quick and easy to implement
    2. inconsistent from server to server
    3. 下游服务不足以支撑上游服务而无感知
    3. cold start, 为填充cache而造成峰值流量
3. external cache
    1. 解决了很多上面的问题
    2. 因为是全局的，所以有好的cache coherence
    3. 总体负载打到下游流量的情况好转，isn’t proportional to fleet size
    4. 问题是，总体系统的复杂度提升和运转负载

## mysql vs redis
常用的数据库是MySQL、Redis
二者该如何选择？
1. MySQL Redis支持的qps是多少
    1. 如何评估qps
2. 业务本身的qps是多少
3. 峰值流量是平峰的几倍
4. MySQL中的分库分表带来什么样的优点，它对支持的qps有什么样的影响