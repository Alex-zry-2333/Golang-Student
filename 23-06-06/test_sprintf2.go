

/*
 * @Author       : Zry && 978524088@qq.com
 * @Date         : 2023-06-06 21:31:29
 * @LastEditors  : Zry && 978524088@qq.com
 * @LastEditTime : 2023-06-06 21:32:07
 * @FilePath     : \Golang\23-06-06\test_sprintf.go
 * @Description  : 
 * 
 * Copyright (c) 2023 by 978524088@qq.com, All Rights Reserved. 
 */

package main


import (
    "fmt"
    "io"
    "os"
)


func main() {
    // go 中格式化字符串并赋值给新串，使用 fmt.Sprintf
    // %s 表示字符串
    var stockcode="000987"
    var enddate="2023-06-06"
    var url="Code=%s&endDate=%s"
    var target_url=fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)

    // 另外一个实例，%d 表示整型
    const name, age = "Zry", 26
    s := fmt.Sprintf("%s is %d years old.\n", name, age)
    io.WriteString(os.Stdout, s) // 简单起见，忽略一些错误
}