# qiniutil

> 七牛云对象存储的命令行工具,当前仅支持linux,上传成功至七牛云后,会将七牛云的外链拷贝至剪切板,让markdown文档写起来更舒服

## 特性

1. 支持直接上传本地文件
2. 支持上传网络中的文件

## 下载

> go get github.com/siskinc/qiniutil

## 使用

### 查看帮助

``` shell
qiniutil -h
```

``` text
Usage of qiniutil:
  -ak string
        Qiniu Access Key
  -config string
        config file path (default "/home/siskinc/.qiniutilrc.json")
  -f string
        The parameter is upload file path
  -hf string
        The parameter is http url of upload file
  -key string
        file key
  -sk string
        Qiniu Secret Key
```

### 参数意义

1. ak: 七牛云的Access Key

2. sk: 七牛云的Secret Key

3. config: 指定qiniutil的配置文件

4. f: 指定本地文件

5. hf: 指定网络中的文件，例如：![http://www.baidu.com/a.jpg](http://www.baidu.com/a.jpg)

6. key: 指定对象存储中的值，可不填，默认取文件的md5

### 配置文件

> 默认会读取~/.qiniutilrc.json文件  

配置如下：

``` json
{
    "access_key": "",
    "secret_key": "",
    "bucket": "",
    "qiniu_dns": "",
    "zone": ""
}
```

access_key和secret_key就在个人面板->密钥管理里面查看  
bucket就是你的对象存储名  
qiniu_dns就是你外链的前缀，在空间概览里面，例如：
123abc.bkt.clouddn.com  
zone就是你对象存储所在区域：

1. Huadong: 华东

2. Huabei: 华北

3. Huanan: 华南

4. Beimei: 北美

5. Xinjiapo: 新加坡

### 使用案例

#### 上传本地文件

``` shell
qinutil -f /home/siskinc/Pictures/新建文件夹/a.jpg
```

#### 上传网络中的文件

``` shell
qinutil -hf http://www.baidu.com/a.jpg
```