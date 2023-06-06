

/*
 * @Author       : Zry && 978524088@qq.com
 * @Date         : 2023-06-06 21:37:05
 * @LastEditors  : Zry && 978524088@qq.com
 * @LastEditTime : 2023-06-06 21:37:13
 * @FilePath     : \Golang\23-06-06\test_sprintf3.go
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
 
    // 声明数字变量
    const num1, num2, num3 = 5, 10, 15
 
    // 调用 Sprintf() 函数
    s := fmt.Sprintf("%d + %d = %d", num1, num2, num3)
 
    // 使用 WriteString() 函数将结果输出到终端 to write the
    //  "os.Stdout" 为字符串的内容
    io.WriteString(os.Stdout, s)
 
}