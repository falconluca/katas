# Gin解析HTTP的5种参数类型

## 1. 路径参数(path)

举例: `http://api.xxx.com/v1/user/luca` 

这里的 luca 就是路径参数

## 2. 查询字符串参数(query)

举例: `/weclome?firstname=Luca&lastName=Yeung`

firstname 和 lastname 就是查询字符串参数

## 3. 表单参数(form)

举例: `curl -XPOST -F'username=Luca' -F'password=123' http://xxx.com/login`

username 和 password 就是表单参数

## 4. 消息体参数(body)

举例：`curl -XPOST -H'Content-Type: application/json' -d'{"username": "luca", "password": "123"}' http://xxx.com/login`

username 和 password 就是消息体参数

## 5. HTTP头参数(header)

举例：`curl -XPOST -H'name: luca' http://xxx.com/hello`

name 就是HTTP头参数

## Gin解析方式

第1⃣️种

```go
gin.Default.GET("/user/:name/:id", nil)
```

第2⃣️种

```go
name := c.Params("name")
action := c.Params("action")
```

第3⃣️种

```go
// 将路径参数绑定到结构体中
type Person struct {
	ID string `url:"id" binding:"required,uuid"`
	Name string `url:"name" binding:"required"`
}

if err := c.ShouldBindUri(&person); err != nil {
	return 
}
```

两个函数: 

1. `ShoudBindWith(obj interface{}, b binding.Binding) error`
2. `MustBindWith(obj interface{}, b binding.Binding) error`