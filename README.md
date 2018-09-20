# dockersearch
按照镜像名称与tag搜索dockerhub
由于本人要经常搜索arm版本的镜像，而网页搜索又不太方便，所以写了这个小工具  :dog: 

tag:0.0.1

使用说明：
``` bash
dockersearch -h
dockersearch version: 0.0.1
Usage: dockersearch [-hrt]

Options:
  -h help
        this help
  -r queryRepository
        要查询的镜像名称 例如: armv7/armhf-ubuntu or ubuntu
  -t queryTag
        要查询的标签 例如: amd or arm 等等
```
参数说明：

-h	帮助信息

-r	要搜索的镜像名称	例如：armv7/armhf-ubuntu 、 ubuntu 等

-t	要搜索的镜像标签	例如：amd64、arm64 等

tag:0.0.2

-r 支持多个参数，逗号分隔，且的关系

-t 不传此参数时，输出镜像详情页