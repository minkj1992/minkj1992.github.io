# [일주일 만에 배우는 GO] CH.3 Real World Go


gin framework로 crud 서버를 만들며 배운 내용들을 기록해보겠습니다.

<!--more-->
<br />

## TIL

go에서는 모든 type을 나타내고자 할 때 비어진 interface를 사용합니다.(dynamic typing)

#### gin-gonic/context.go

```go
// MustBindWith binds the passed struct pointer using the specified binding engine.
// It will abort the request with HTTP 400 if any error occurs.
// See the binding package.
func (c *Context) MustBindWith(obj interface{}, b binding.Binding) error {
	if err := c.ShouldBindWith(obj, b); err != nil {
		c.AbortWithError(http.StatusBadRequest, err).SetType(ErrorTypeBind) // nolint: errcheck
		return err
	}
	return nil
}
```

## Update

crud를 만들면 update 부분을 깔끔하게 짜는것이 신경쓰입니다. 아래는 todo에서 짜본 update 코드입니다.

```go
func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.GetTodoById(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}

	c.BindJSON(&todo)
	err = models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
```

`c.BindJSON(&todo)`를

## conclustion

[2022-02-08T19:21:29+09:00] effective go로 공부하니까, 문서가 정말 좋긴한데, 예상보다 몇시간은 더 걸렸던 것 같습니다. 하지만 양질의 정보를 이렇게 빠르게 읽을 수 있어서 유익한 시간인 것 같네요.

<center>- 끝 -</center>

