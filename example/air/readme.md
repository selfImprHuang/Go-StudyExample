获取命令： go get github.com/cosmtrek/air

参考：https://github.com/darjun/go-daily-lib

参考：https://segmentfault.com/a/1190000025186913

## 使用说明
air -c example.toml 是按照配置进行启动，可以查看github该工程的配置，但是我这边按照配置启动一直是失败的状态，没有去深究。

这边直接在【Terminal】上面输入 【air】命令也可以正常启动，他会找到最外层的main.go函数进行执行

当本工程中任意一个文件进行修改【保存】，都会触发重启操作，也就是说省去了重启的过程(监听程序文件的变化)