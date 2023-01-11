# HASH_BypassAV

 ## 免责声明

该工具仅用于安全研究，禁止使用工具发起非法攻击等违法行为，造成的后果使用者负

## 介绍

一个 go 语言实现的免杀框架，支持 msf / cs 生成的 c / bin 格式 shellcode。
支持 AES 加密、字符串混淆、反沙箱等规避方式。
可选免杀方式位于 `./core` 中。

## Example

```
go run main.go -m CreateFiber -strip -s shellcode.txt
```

可选参数：

| 参数  |                参数说明                 |  参数类型  |
|:---:|:-----------------------------------:|:------:|
| -m  |                使用模块                 | string |
| -s  |    shellcode文件（默认shellcode.txt）     | string |
| -c  |    c 格式 shellcode                      |  bool  |
| -enc |   加密方式                               | string |
| -key |   密钥                                  | string |
| -d  |    去除符号表        |  bool  |
| -hide  | 隐藏窗口 |  bool  |
| -sb  |           加入反沙箱           |  bool  |
| -strip  |        使用 go-strip 混淆生成的 exe        |  bool  |

## 参考

https://github.com/safe6Sec/GolangBypassAV

https://github.com/Nan3r/checkgo

