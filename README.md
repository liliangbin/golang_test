# golang_test
### test
 - 总结一下他特殊的地方
```go
    var a int , b = 1,3 
		func man x(nates int ,fsdfs int ) (int ,string ){
		}
```
切片的使用方法

#### go web 编程
##### go 在使用的时候，主要是使用net/http包，下面是在开发过程中遇到的问题。现在遇到的问题主要是对于ｈｔｔｐ包的使用不熟悉
- 参数。　关于我们使用http.request 这个包对应的函数，
- r.ParseForm()　　这个的目的是为了解析数据的参数。
- r.Form  这个里面存储着素有的数据，我们使用r.Form["some"][0]  来获取传傪为ｓｏｍｅ的第一个数据
- request 这我们可以把他看成一个ｍａｐ对象，他可以存储的多个数据，相同的参数的时候对应的一个数组。在传参数的时候，他的数不是放在formdata里面的，而是使用的x-www-form-urlencoded 这个东西来传送的，所以我们要仔细的使用他们。
- 这个东西是真的坑。。。。。。
##### we can create a restful api by json
- json 这个库是我们使用比较常用的库了，用来数据库的交互，配置文件的读取，ｒｅｑｕｅｓｔ请求的获取。我们都需要用到ｊｓｏｎ 函数
- 把数据加载成ｊｓｏｎ文件
```json
	if err:=json.NewEncoder(w).Encode(todo);err!=nil{
//这个地方我们使用json来把Ｓｔｒｕｔｓ文件呢给转成ｊｓｏｎ形式。　　ｗ　代表http.responseWriter ,todo 代表了一个struct 文件，
  ```
##### 中间件的使用，这个时候我们可以考虑使用negroni 
##### 路由我们可以考虑使用ｍｕｘ路由，这个还比较强大。
##### 数据库文件，gorm 这个orm 框架可以考虑一波，效率也还行
##### 具体的情况的话，我们可以考虑ｇｏ和一个内存数据库组合使用，增强他的并发能力。。有时间过看看别人的代码。同时他的面向对象编程，是我的一份短板。