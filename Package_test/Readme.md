

<!--
 * @Author       : Zry && 978524088@qq.com
 * @Date         : 2023-06-06 20:34:11
 * @LastEditors  : Zry && 978524088@qq.com
 * @LastEditTime : 2023-06-06 20:35:51
 * @FilePath     : \Golang\Package_test\Readme.md
 * @Description  : 
 * 
 * Copyright (c) 2023 by 978524088@qq.com, All Rights Reserved. 
-->

关于包，根据本地测试得出以下几点：

- 文件名与包名没有直接关系，不一定要将文件名与包名定成同一个。
- 文件夹名与包名没有直接关系，并非需要一致。
- 同一个文件夹下的文件只能有一个包名，否则编译报错。

文件结构
```shell
--main.go
--myMath
----myMath1.go
----myMath2.go
```
