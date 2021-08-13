# 欢迎使用 durl 短链服务

## 体验地址 ->>> [durl.fun](https://durl.fun)

## durl介绍:
durl 是一个分布式的高性能短链服务,逻辑简单,部署方便.

## 背景
发现在github中已有的短链服务中,非分布式服务做到快速扩容,并且有些项目是用的redis作为数据缓存,性能上不够优秀.

## 短链服务介绍
短链接，通俗来说，就是将长的URL网址，通过程序计算等方式，转换为简短的网址字符串。

## 使用场景
微博和Twitter都有140字数的限制，如果分享一个长网址，很容易就超出限制。

营销短信，字数的限制,当字数过长: 1.不美观 2.超出字符额外收费。

生成二维码的原始链接,当原始链接过长时,生成的二维码过于复杂,导致一些像素较低的手机无法扫描.

### durl的三个模块:

portal: 首页可以通过页面进行短链生成.公司内部或者公司外部可以通过页面生成短链接.

openApi: 对内开放api,增删改查.

jump: 只服务短链跳转.作为专门的跳转服务.

这样分为三个模块的原因,是因为根据需要进行部署,需要那个就部署那个.

因为这个项目的结构原因,整个项目三个模块之间没有耦合,可以随意增加pod数量,来提高系统性能.
一般来说openApi可以部署为只内网访问,jump作为专门的跳转服务,如果有需要页面服务就部署portal.

## 特征:
1. [beego](https://github.com/beego/beego) 为项目web框架.
2. 使用了 [xorm](https://github.com/xormplus/xorm) 来实现持久数据存储, 项目已测试 mysql 与 mongo.
3. 使用了 [mcache](https://github.com/songangweb/mcache) 来实现内存缓存.
4. 因使用内存缓存作为缓存池,实际使用中,项目本身的性能瓶颈更多体现在数据库自身.
5. 项目内存消耗大多为缓存内存所用容量,可通过配置文件进行内存大小限制.


## 如何使用

### 本地调试:
进入各个项目, 修改完数据库, 直接run就可以

### 系统部署:
在 durl/build 目录下提供有dockerfile demo. 可以根据需要进行修改后部署.


## 项目流程解析   详细了解durl项目时,此模块内容必看

[项目详解](https://github.com/songangweb/durl/wiki/Explain)

[comment]: <> ([项目目录结构]&#40;https://github.com/songangweb/durl/wiki/Directory&#41;)

[comment]: <> ([配置文件详解]&#40;https://github.com/songangweb/durl/wiki/Explain&#41;)

## openApi

[接口文档](https://github.com/songangweb/durl/wiki/OpenApi)


## JetBrains操作系统许可证

durl 是根据JetBrains sro授予的免费JetBrains开源许可证与GoLand一起开发的，因此在此我要表示感谢。

[免费申请 jetbrains 全家桶](https://zhuanlan.zhihu.com/p/264139984?utm_source=wechat_session)


## 交流
#### 如果文档中未能覆盖的任何疑问,欢迎您发送邮件到<songangweb@foxmail.com>,我会尽快答复。
#### 您可以在提出使用中需要改进的地方,我会考虑合理性并尽快修改。
#### 如果您发现 bug 请及时提 issue,我会尽快确认并修改。
#### 有劳点一下 star，一个小小的 star 是作者回答问题的动力 🤝
