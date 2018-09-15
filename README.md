# dockersearch
按照镜像名称与tag搜索dockerhub
由于本人要经常搜索arm版本的镜像，而网页搜索又不太方便，所以写了这个小工具  :dog: 

使用说明：
``` bash
dockersearch -h
dockersearch version: 0.0.1
Usage: dockersearch [-hrt]

Options:
  -h help
        this help
  -r queryRepository
        Input your queryRepository example: armv7/armhf-ubuntu or ubuntu
  -t queryTag
        Input your queryTag example: amd or arm and so on
```
参数说明：

-h	帮助信息

-r	要搜索的镜像名称	例如：armv7/armhf-ubuntu 、 ubuntu 等

-t	要搜索的镜像标签	例如：amd64、arm64 等

