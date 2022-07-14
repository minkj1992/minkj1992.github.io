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
프로젝트 일정이 다가와서, gin 공식문서는 대충보고 프로젝트를 1달동안 만들어보았습니다.

https://github.com/ExchangeDiary/exchange-diary.

만들면서 배우는게 가장 재밌게 배울 수 있는 거 같아서 만족스럽습니다.

<center>- 끝 -</center>

