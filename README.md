# Async-HTTPClient
Assignment for service computing : Asynchronous HTTPClient

## 练习要求
* 依据文档图6-1，用中文描述 Reactive 动机
* 使用 go HTTPClient 实现图 6-2 的 Naive Approach
* 为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3
* 对比两种实现，用数据说明 go 异步 REST 服务协作的优势
* *思考：* 是否存在一般性的解决方案？

### Reactive 动机
![](https://jersey.github.io/documentation/latest/images/rx-client-problem.png)

通过阅读文档[Motivation for Reactive Client](https://jersey.github.io/documentation/latest/rx-client.html#d0e5556)中6.1的内容，以及根据图6-1可总结出 Reactive 动机为：
1. 对于实现编排层，使用同步方法时，实现过程简单，易于阅读和直接调试。但这种方法速度较慢，浪费资源；而通过并行方法调用独立请求可以降低同步方法所需的时间，但实现起来比较困难，调试和维护也不方便。

![](https://jersey.github.io/documentation/latest/images/rx-client-sync-approach.png)

2. 使用异步方法获取数据时，当请求结束会调用一个回调方法，得到一个是否继续处理的信号，这种额外信息是同步方法无法得到的。但异步处理错误信息是很不方便的，且对于异步方法来说，与之前一样，实现起来比较困难，调试和维护不方便。

![](https://jersey.github.io/documentation/latest/images/rx-client-async-approach.png)

对于这两种情况，Reactive 式编程比异步方法更具可读性，速度也更快，而阅读和实现也很容易，也能更好处理错误信息。所以当要处理大量信息且希望同时具有异步方法和同步方法的优点时，使用 Reactive 式编程是很有效的。


### 使用 go HTTPClient 实现 Naive Approach

### 利用 Channel 搭建基于消息的异步机制

### go 异步 REST 服务协作的优势

### 思考：一般性的解决方案
