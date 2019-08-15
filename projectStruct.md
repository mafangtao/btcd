addmgr 				- 地址管理
blockchain 			- 实现比特币区块处理和链选择规则
blockchain / fullblocktests 	- 提供一组用于测试共识验证规则的块测试
btcec 				- 实现对比特币脚本所需的椭圆曲线密码函数的支持
btcjson 			- 为底层的JSON-RPC命令和返回值提供一个扩展的API
chaincfg/chainhash 		- 提供通用的散列类型和相关函数，允许抽象特定的散列算法。
connmgr 			- 链接管理器
database 			- 为比特币区块链提供数据库接口,实现了ffldb
mempool 			- 未发掘的比特币交易池。
mining 				- pow挖矿
netsync 			- 同步管理器，同步区块链和交易
peer 				- 创建和管理比特币对等网络，及上行下行数据包的处理
rpcclient 			- 实现一个强大且易于使用的支持Websocket的比特币JSON-RPC客户端
txscript 			- 实现比特币交易脚本语言
wire 				- 实现比特币网络协议