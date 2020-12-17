参考地址：http://www.topgoer.com/%E5%85%B6%E4%BB%96/%E4%B8%AD%E6%96%87%E5%88%86%E8%AF%8D.html

包获取方法：go get github.com/yanyiwu/gojieba

我的想法

    分词如果跟前缀树一起使用是不是可以实现关键词典的模糊搜索：   
    前缀树代码参考我的另一个工程Go-Tool
    

问题
    因为这个库是通过C++实现的，所以需要下载gcc并设置环境变量，可参考：https://blog.csdn.net/benben_2015/article/details/80565676
    
   gcc环境下载地址：https://www.cnblogs.com/LUA123/p/11446185.html    
   
   注意这边的path需要加载系统变量上面，不是用户系统变量