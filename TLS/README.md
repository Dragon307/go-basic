## 使用Go实现TLS服务端和客户端

传输层安全协议（Transport Layer Security，缩写TLS），及其前身安全套接层（Secure Sockets Layer， SSL）是一种安全协议，目的是为互联网
通信提供安全及数据完整性保障。

SSL 包含记录层Record Layer）和传输层，记录层协议确定了传输层数据的封装格式。传输层安全协议使用X.509认证，之后利用非对称加密演算来对通信方做身份认证，
之后交换对称密钥作为会谈密钥（Session key）。这个会谈密钥是用来将通信两方交换的数据做加密，保证两个应用间通信的保密性和可靠性，使客户与服务器应用之
间的通信不被攻击者窃听。

## 证书生成

生成服务端的私钥

>   openssl genrsa -out server.key 2048

生成服务端证书

>  openssl req -new -x509 -key server.key -out server.pem -days 3650

客户端证书生成

生成客户端私钥

> openssl genrsa -out client.key 2048

生成客户端证书

> openssl req -new -x509 -key client.key -out client.pem -days 3650