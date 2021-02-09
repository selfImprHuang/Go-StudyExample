参考地址：[https://segmentfault.com/a/1190000021997004](https://segmentfault.com/a/1190000021997004)

工程地址：[https://github.com/imdario/mergo](https://github.com/imdario/mergo)

获取方式：go get github.com/imdario/mergo

注意事项
   
    mergo不会赋值非导出字段；
    map中对应的键名首字母会转为小写；
    mergo可嵌套赋值，我们演示的只有一层结构。