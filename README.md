这是测试用的repo, 会提交一些乱七八糟的东西，不要在意。。   
破菊好厉害  
po菊好厉害 --dawn  
 
如何编译 
------- 
 
**Linux** 
 
1.官方下载go1.3.3.linux-amd64.tar.gz， 解压到/usr/local 
```
sudo tar -C /usr/local -xzf go1.3.3.linux-amd64.tar.gz 
```

2.创建一个workspace， 自己替换掉$HOME  
```
mkdir $HOME/go  
```

3.设置GOPATH， 如果只是临时用就不要写到.profile
``` 
echo "export GOPATH=\$HOME/go" >> ~/.profile  
echo "export PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> ~/.profile  
source ~/.profile  
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
 
目录结构如下：  
```  
├── bin  
├── pkg  
└── src  
     └── warpten  
          ├── README.md  
          ├── server  
          │    └── server.go  
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
warpten -d -t &  
```
发送给服务器并显示  
```  
warpten -t -s="PLAY"  
```
 
**Max OS X**  

1.官方下载go1.3.3.darwin-amd64-osx10.8.pkg， 双击安装  

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
 
目录结构如下：  
```  
├── bin  
├── pkg  
└── src  
     └── warpten  
          ├── README.md  
          ├── server  
          │    └── server.go  
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
warpten -d -t &  
```
发送给服务器并显示  
```  
warpten -t -s="PLAY"  
```
**Win**  
