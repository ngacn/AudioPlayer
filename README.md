这是测试用的repo, 会提交一些乱七八糟的东西，不要在意。。  

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
