这是测试用的repo, 会提交一些乱七八糟的东西，不要在意。。  

至今为止的代码结构  
------------------  
client===>实现一个restful api的客户端， gui构建在这个之上， 所有的界面操作都转化为restful请求，发送给服务器，  
用这个方法实现gui和播放器分离  
server===>实现restful api的服务器， 监听并接受请求，然后交由player处理请求  
player===>实现播放器， 播放器现在包含一个track列表和一个playlist管理器，之后声音引擎，解码插件都会在这个之下，  
server+player实现了完整的播放器功能， 但是除了restful api接口， 无法对其控制  
warpten===>程序的入口文件， 启动server, client,初始化player  
playlists===>播放列表相关的操作  
tracks===>每个track相关的操作  

如何编译
-------
 
1.官方下载go1.3.3.linux-amd64.tar.gz， 解压到/usr/local  
```
sudo tar -C /usr/local -xzf go1.3.3.linux-amd64.tar.gz  
```
Mac OS X下载go1.3.3.darwin-amd64-osx10.8.pkg， 直接双击安装  

2.创建一个workspace， 自己替换掉$HOME  
```
mkdir $HOME/go  
```

3.设置GOPATH  
``` 
export GOPATH=$HOME/go  
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin  
```

4.创建workspace的目录结构  
```  
mkdir $GOPATH/{src,pkg,bin}  
```

5.clone项目  
```  
cd $GOPATH/src  
git clone https://github.com/ngacn/AudioPlayer.git warpten  
```  
注意项目名需要是wrapten， 而不是AudioPlayer  

6.根据https://github.com/go-qml/qm配置go-qml环境，我的建议直接安装qt-creator
比如我现在是arch linux，就直接pacman -S qtcreator

7.安装go-qml
```
go get gopkg.in/qml.v1  
```
目录结构类似下面：  
```
.
├── bin
├── pkg
└── src
    ├── gopkg.in
    │   └── qml.v1
    └── warpten
        ├── client
        ├── player
        ├── playlists
        ├── README.md
        ├── server
        ├── tracks
        ├── warpten
        └── warpten-daemon

13 directories, 1 file

```  

8.编译  
```
go install warpten/warpten  
go install warpten/warpten-daemon
```

9.使用  
  
```
warpten  
```
发送测试命令给服务器并显示  
```
warpten-daemon version  
```

10.关闭
关闭gui, daemon进程也会被关闭

用curl测试warpten-daemon
![20141211184612](https://cloud.githubusercontent.com/assets/9798546/5392953/048362f6-8167-11e4-84ad-f05d187a2643.png)
