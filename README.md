# STS2G
Struts2漏洞扫描Golang版.
![avatar](./pasted-95.png)  
#### 使用方法  
``` 
NAME:
   ST2SG - Struts2漏洞检测工具(Golang版)

USAGE:
   ST2SG.exe [global options] command [command options] [arguments...]

AUTHOR:
   x51 <x51enter@gmail.com>

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --mode value  specify working mode
   --vn value    Vulnerability number (default: 0)
   --url value   set target url
   --cmd value   exec command(only work on exploit mode)
   --data value  data for special vuln
   --help, -h    show help (default: false)
```
##### 默认扫描模式:  
```ST2SG --mode scan --url http://xxx.com/index.action```  
##### 指定漏洞扫描模式：  
```ST2SG --mode scan --url http://xxx.com/index.action --vn 16```  
##### 命令执行模式：  
```ST2SG --mode exec --url http://xxx.com/index.action --vn 16 --cmd "whoami"```  
##### 自定义参数模式：  
*该模式基于以上方法，分有两种情况，自定义GET参数名，和自定义POST数据包内容，POST方式需要在数据包中指定一下要测试的参数并用fuckit标记出来.*  
POST  
```ST2SG --mode scan --url http://xxx.com/index.action --vn 007 --data "name=fuckit&pass=qwer"```  
GET  
```ST2SG --mode scan --url http://xxx.com/index.action --vn 009 --data "name"```  
#### 待补充  
*上传Webshell功能*
#### 参考项目  
https://github.com/HatBoy/Struts2-Scan  
#### 测试环境  
https://github.com/vulhub/vulhub
