# TCP

TCP 协议，全称：Transmission Control Protocol，传输控制协议。位于IP层之上，应用层之下。是一个基于字节流（而非消息）的，面向连接的，可靠的传输层通信协议。

## 常量

### 数据包大小

todo  [TCP 协议简介](http://www.ruanyifeng.com/blog/2017/06/tcp-protocol.html)

### 传输速度

北京到上海，1000公里需要多少ms

## 连接建立与断开

关于TCP最常见的一个知识点，即是：三次握手和四次挥手。

三次握手，为什么需要三次，很多文章解释得不对误人子弟。这篇文章讲得比较清楚：[“三次握手，四次挥手”你真的懂吗？](https://zhuanlan.zhihu.com/p/53374516)

> 换个易于理解的视角来看为什么要3次握手。
> 
> 客户端和服务端通信前要进行连接，“3次握手”的作用就是双方都能明确自己和对方的收、发能力是正常的。
> 
> 第一次握手：客户端发送网络包，服务端收到了。这样服务端就能得出结论：客户端的发送能力、服务端的接收能力是正常的。
> 
> 第二次握手：服务端发包，客户端收到了。这样客户端就能得出结论：服务端的接收、发送能力，客户端的接收、发送能力是正常的。 从客户端的视角来看，我接到了服务端发送过来的响应数据包，说明服务端接收到了我在第一次握手时发送的网络包，并且成功发送了响应数据包，这就说明，服务端的接收、发送能力正常。而另一方面，我收到了服务端的响应数据包，说明我第一次发送的网络包成功到达服务端，这样，我自己的发送和接收能力也是正常的。
> 
> 第三次握手：客户端发包，服务端收到了。这样服务端就能得出结论：客户端的接收、发送能力，服务端的发送、接收能力是正常的。 第一、二次握手后，服务端并不知道客户端的接收能力以及自己的发送能力是否正常。而在第三次握手时，服务端收到了客户端对第二次握手作的回应。从服务端的角度，我在第二次握手时的响应数据发送出去了，客户端接收到了。所以，我的发送能力是正常的。而客户端的接收能力也是正常的。
> 
> 经历了上面的三次握手过程，客户端和服务端都确认了自己的接收、发送能力是正常的。之后就可以正常通信了。
> 
> 每次都是接收到数据包的一方可以得到一些结论，发送的一方其实没有任何头绪。我虽然有发包的动作，但是我怎么知道我有没有发出去，而对方有没有接收到呢？
> 
> 而从上面的过程可以看到，最少是需要三次握手过程的。两次达不到让双方都得出自己、对方的接收、发送能力都正常的结论。其实每次收到网络包的一方至少是可以得到：对方的发送、我方的接收是正常的。而每一步都是有关联的，下一次的“响应”是由于第一次的“请求”触发，因此每次握手其实是可以得到额外的结论的。比如第三次握手时，服务端收到数据包，表明看服务端只能得到客户端的发送能力、服务端的接收能力是正常的，但是结合第二次，说明服务端在第二次发送的响应包，客户端接收到了，并且作出了响应，从而得到额外的结论：客户端的接收、服务端的发送是正常的。

从工程角度上讲，为了确认发送方和接收方的发送能力和接收能力，需要各自进行至少一次发送和接收，因而需要三次握手。

而四次挥手的核心在于关闭发送能力和接收能力。

## TCP 状态转移 

todo

## 可靠性机制设计

todo

## 性能缺陷

todo