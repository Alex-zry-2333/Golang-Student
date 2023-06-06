

/*
 * @Author       : Zry && 978524088@qq.com
 * @Date         : 2023-06-01 22:41:33
 * @LastEditors  : Zry && 978524088@qq.com
 * @LastEditTime : 2023-06-06 20:22:30
 * @FilePath     : \Golang\test.go
 * @Description  : 
 * 
 * Copyright (c) 2023 by 978524088@qq.com, All Rights Reserved. 
 */

package main

import "fmt"

func main()  {
	fmt.Println("Hello ZRY's World!")
}

func init() {
	fmt.Print("Welcome to ") /*Print 不会换行*/
	fmt.Println("Init will be called") // init 函数将先于main 函数执行
}	