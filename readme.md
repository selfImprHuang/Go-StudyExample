# 更新日志

2020/4/27 ：创建项目，添加使用示例如下：
map与struct之间转换的处理工具使用示例，参考代码地址：https://github.com/mitchellh/mapstructure
json与struct之间转换处理工具使用示例


2020/5/15 : 添加CloneExample,测试go语言的对象内存和深度克隆的实现关于对象克隆可参考文章[知乎-](https://zhuanlan.zhihu.com/p/59125443)、[知乎-](https://zhuanlan.zhihu.com/p/58065429)


2020/7/8 : 添加两种二维码生成使用示例，添加两种图片压缩方式、添加水印示例，添加文字到图片


2020/7/13 : 新增 网络及端口测试、传值和传址的比较、切片测试、Map测试、通道测试、测试init调用、测试error、defer关键字、recover测试、测试携程的异常捕获、矩阵旋转

2020/12/9 ：新增Gin框架使用示例

2020/12/15: 新增通过gopsutil包读取服务器信息的方法

2020/12/15: 新增通过pinyin包获取中文拼音的使用示例

2020/12/16: 新增雪花算法实现

2020/12/16: 新增gjson的使用示例

2020/12/17：新增yanyiwu中文分词工具使用示例

2020/12/17: 新增cronTab与go cron的定时任务使用示例

2020/12/17：新增sdk包Logger的使用示例

# mod vendor模式加载包
通过go mod的方式加载的github上面的包会有报红的问题，但是包本身是可以运行的，这样就是会有一个问题，如果你想要点击去看方法的内容，没办法做到

查阅了网上提供的处理方式，通过设置GOPATH 和Go Modules(vgo)可以解决相应的问题，但是我在这样处理之后依然会出现报红的问题

后面通过使用vendor包的方式来处理，就避免了报红的问题

处理步骤如下 ：
- 通过 go get github.com/kardianos/govendor 命令下载govendor命令
- 通过 go mod vendor 切换到vendor管理
