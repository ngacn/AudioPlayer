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
ycs@linux-afno:~/go> warpten version
Warpten version: 0.0  

ycs@linux-afno:~/go> warpten playlists
Warpten playlists: {"Default":{}}  

ycs@linux-afno:~/go> warpten playlist -a NewList
Create playlists NewList: success  

ycs@linux-afno:~/go> warpten playlists
Warpten playlists: {"Default":{},"NewList":{}}  

ycs@linux-afno:~/go> warpten playlist -a NewList2
Create playlists NewList2: success  

ycs@linux-afno:~/go> warpten playlists
Warpten playlists: {"Default":{},"NewList":{},"NewList2":{}}  

ycs@linux-afno:~/go> warpten playlist -d NewList
Delete playlists NewList: success  

ycs@linux-afno:~/go> warpten playlists
Warpten playlists: {"Default":{},"NewList2":{}}
```
