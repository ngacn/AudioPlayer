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
 
目录结构类似下面：  
```
.
├── bin
├── pkg
└── src
    └── warpten
        ├── README.md
        ├── client
        │   ├── cli.go
        │   ├── commands.go
        │   └── utils.go
        ├── server
        │   └── server.go
        └── warpten
            └── warpten.go
```  

6.编译  
```
go install warpten/warpten  
```

7.使用  
  
开启服务器  
```
warpten -d &  
或  
warpten -d -t & 使用Tcp socket  
```
发送给服务器并显示  
```
warpten version  
或  
warpten -t version 使用Tcp socket  
```

```
查看播放器版本
ycs@linux-afno:~/go/src/warpten> warpten version
Warpten version: 0.0 
查看所有播放列表
ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":[]}
添加一个名字NewList的播放列表
ycs@linux-afno:~/go/src/warpten> warpten playlist -a NewList
Create playlists NewList: success
查看所有播放列表
ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":[],"NewList":[]}
查看播放列表NewList的内容
ycs@linux-afno:~/go/src/warpten> warpten playlist NewList
Get playlist NewList: []
添加播放列表NewList2
ycs@linux-afno:~/go/src/warpten> warpten playlist -a NewList2
Create playlists NewList2: success
查看所有播放列表
ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":[],"NewList":[],"NewList2":[]}
删除播放列表NewList
ycs@linux-afno:~/go/src/warpten> warpten playlist -d NewList
Delete playlists NewList: success
查看所有播放列表
ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":[],"NewList2":[]}
删除播放列表NewList3
ycs@linux-afno:~/go/src/warpten> warpten playlist -d NewList3
Delete playlists NewList3: NewList3 not exists
添加播放列表NewList2
ycs@linux-afno:~/go/src/warpten> warpten playlist -a NewList2
Create playlists NewList2: NewList2 exists
查看所有播放列表
ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":[],"NewList2":[]}
```

更新一些track相关的例子
```
ycs@linux-afno:~/go/src/warpten> warpten playlist -a NewList
Create playlist NewList: success
ycs@linux-afno:~/go/src/warpten> warpten tracks
debug=> map[]
Warpten tracks: {}
ycs@linux-afno:~/go/src/warpten> warpten track -a -pl="NewList" /tmp/1.mp3
Create track /tmp/1.mp3: success
ycs@linux-afno:~/go/src/warpten> warpten track -a -pl="NewList" /tmp/3.mp3                                                                                                          
Create track /tmp/3.mp3: success
ycs@linux-afno:~/go/src/warpten> warpten track -a /tmp/4.mp3                                                                                                                        
Create track /tmp/4.mp3: success

ycs@linux-afno:~/go/src/warpten> warpten tracks
debug=> map[d0a819e1-ff36-4eac-8d81-4071e6c12012:0xc208036d60 27253a71-b8b2-49f2-bb93-0d13a54d2bac:0xc2080368d0 c71f96b6-5693-4296-becb-a6f02f43b54c:0xc208036be0]
Warpten tracks: {"27253a71-b8b2-49f2-bb93-0d13a54d2bac":{},"c71f96b6-5693-4296-becb-a6f02f43b54c":{},"d0a819e1-ff36-4eac-8d81-4071e6c12012":{}}

ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":["c71f96b6-5693-4296-becb-a6f02f43b54c"],"NewList":["d0a819e1-ff36-4eac-8d81-4071e6c12012","27253a71-b8b2-49f2-bb93-0d13a54d2bac"]}
ycs@linux-afno:~/go/src/warpten> warpten track -d -pl="NewList" d0a819e1-ff36-4eac-8d81-4071e6c12012                                                                                
Delete track d0a819e1-ff36-4eac-8d81-4071e6c12012: success

ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":["c71f96b6-5693-4296-becb-a6f02f43b54c"],"NewList":["27253a71-b8b2-49f2-bb93-0d13a54d2bac"]}
ycs@linux-afno:~/go/src/warpten> warpten track 27253a71-b8b2-49f2-bb93-0d13a54d2bac
debug=> &{/tmp/3.mp3}
Get track 27253a71-b8b2-49f2-bb93-0d13a54d2bac: {}
ycs@linux-afno:~/go/src/warpten> warpten playlist -d NewList
Delete playlist NewList: success
ycs@linux-afno:~/go/src/warpten> warpten playlists
Warpten playlists: {"Default":["c71f96b6-5693-4296-becb-a6f02f43b54c"]}
ycs@linux-afno:~/go/src/warpten> warpten tracks
debug=> map[c71f96b6-5693-4296-becb-a6f02f43b54c:0xc208036be0]
Warpten tracks: {"c71f96b6-5693-4296-becb-a6f02f43b54c":{}}
```
