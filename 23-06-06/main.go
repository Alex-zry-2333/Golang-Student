

/*
 * @Author       : Zry && 978524088@qq.com
 * @Date         : 2023-06-06 21:27:08
 * @LastEditors  : Zry && 978524088@qq.com
 * @LastEditTime : 2023-06-06 21:27:15
 * @FilePath     : \Golang\23-06-06\main.go
 * @Description  : 
 * 
 * Copyright (c) 2023 by 978524088@qq.com, All Rights Reserved. 
 */

package main  

import(  
	"fmt"  
)  

func main(){  
 // %d 表示整型数字，%s 表示字符串  
	var stockcode=123  
	var enddate="2023-6-6"  
	var url="Code=%d&endDate=%s"  
	var target_url=fmt.Sprintf(url,stockcode,enddate)  
	fmt.Println(target_url)  
} 