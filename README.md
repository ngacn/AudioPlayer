这是测试用的repo, 会提交一些乱七八糟的东西，不要在意。。  
更新了win下环境搭建，往下拉到最下面

至今为止的代码结构  
------------------  

player  
实现播放器， 播放器现在包含一个track列表和一个playlist管理器，之后声音引擎，解码插件都会在这个之下  
server+player实现了完整的播放器功能， 但是除了restful api接口， 无法对其控制  

playlists  
播放列表相关的操作  

README.md  
本文件  

server  
实现restful api的服务器， 监听并接受请求，然后交由player处理请求  

tracks  
每个track相关的操作  

utils  
工具， 比如生成uuid  

warpten  
qt和c++写的gui客户端， 请求warpten-daemon获取显示内容， 不运行任何实际播放功能  

warpten-daemon  
daemon程序的入口文件，启动server， 初始化player  

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
        ├── player
        ├── playlists
        ├── README.md
        ├── server
        ├── tracks
        ├── utils
        ├── warpten
        └── warpten-daemon

11 directories, 1 file

```  

6.编译warpten-daemon和warpten  
```
go install warpten/warpten-daemon
```  
qtcreator打开src/warpten/warpten.pro, 不用多说

7.使用  
运行qtcreator编译出的二进制可执行程序， 注意go编译出的warpten-daemon需要在系统path里，  
就是可以在命令行里不使用路径运行

8.关闭  
关闭gui, daemon进程也会被关闭

用curl测试warpten-daemon
![20141211184612](https://cloud.githubusercontent.com/assets/9798546/5392953/048362f6-8167-11e4-84ad-f05d187a2643.png)

qt客户端目前。。
![20141212192440](https://cloud.githubusercontent.com/assets/9798546/5411363/6140e530-823a-11e4-91a2-d9584423bb50.png)


###　Window下编译环境配置
1.下载golang：https://golang.org/dl/，选go1.4.windows-amd64.msi下载并安装

2.配置环境变量
```
path加入 C:\Go\bin
GOPATH=D:\MyProjects\GIT\AudioPlayer（代码目录要根据具体路径修改）
GOROOT=C:\Go（go安装目录要根据具体路径修改）
```

3.打开cmd，输入go，如果出现提示语，证明已经配好go环境

4.在AudioPlayer目录下建立目录src/warpten，并把所有文件与目录拷贝一份到src/warpten

5.编译warpten-daemon
```
go install warpten/warpten-daemon
```  

6.下载qt环境http://www.qt.io/download-open-source/#，选择Qt 5.4.0 for Windows 32-bit (MinGW 4.9.1, 852 MB)下载并安装

7.qtcreator打开src/warpten/warpten.pro, 不用多说

8.剩下的步骤都一样
