# EasyError-go

## 更方便的处理golang中的异常

### 示例1

``` golang
EasyError.HandError(err).PrintlnError().ThenGo(func() {
		fmt.Println("func ThenGo run!")
	}).Finaly(func() {
		fmt.Println("func Finaly run!")
	})
```



### 示例2

``` golang
EasyError.HandError(err).PrintlnError("string",1,1.1).ThenGo(func() {
		fmt.Println("func ThenGo run!")
	}).Finaly(func() {
		fmt.Println("func Finaly run!")
	})
```



### 函数

HandError(err error )   传入一个error类型的参数  返回一个HandContinue类 可以通过此类调用处理异常的方法

### type Hand func()

一个函数类型的Hand

### Class HandContinue 方法

(h  *HandContinue) ThenGo(hand Hand)  *HandContinue   

可以在ThenGo的参数函数中执行任意方法  ThenGo 只有在有异常时才会被执行 



(h *HandContinue) PrintlnError(log ...any) *HandContinue 

PrintlnError 里可以穿入多个任意类型的参数并拼接error进行打印并换行



(h *HandContinue) PrintError(log ...any) *HandContinue

PrintlnError 里可以穿入多个任意类型的参数并拼接error进行打印



(h *HandContinue) Panic() *HandContinue

Panic里被调用后自动panic  go程序中的error



(h *HandContinue) Finaly(head Hand)

Finaly方法在最终执行，无论有没有异常都会执行，不会返回任何参数
