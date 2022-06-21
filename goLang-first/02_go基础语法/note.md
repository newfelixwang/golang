### GO语言现状以及未来发展



## 特点

* 结合了一些现在有的语言缺点,进行整合使用
* 天然支持高并发等原生支持





## 下载Go, https://studygolang.com/dl

###　开发编辑器　　ＩＤＥＡ　ｖｓＣＯＤＥ　还有其他很多







### 完成第一个hello word



 ```
fmt.Println("hello word")
 ```





### 类型定义

```
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitaValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"

	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("hello word")
	variableZeroValue()
	variableInitaValue()
	variableTypeDeduction()
}
```

2/1########### go变量定义

