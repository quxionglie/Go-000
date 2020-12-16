学习笔记

第二周：异常处理

#1.Week02 作业题目：

问题：1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
为什么，应该怎么做请写出代码？

回答：为了屏蔽dao层不同库(mysql/redis等)的差异，需要吃掉sql.ErrNoRows错误，抛出包装好的固定错误。

sql.ErrNoRows说明
```go
// ErrNoRows is returned by Scan when QueryRow doesn't return a
// row. In such a case, QueryRow returns a placeholder *Row value that
// defers this error until a Scan.
var ErrNoRows = errors.New("sql: no rows in result set")
```

标准答案：

```
以下为毛剑老师给的week2作业答案，仅供大家参考～

dao:

return errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s error: %v", sql, err))


biz:

if errors.Is(err, code.NotFound} {

}

助教：在dao层wrap有比较多的优点：1. dao接口被多个server调用，不必所有人都wrap一遍。dao可以直接将一些上下文信息封进去，例如部分不敏感的查询参数；2. dao接口可以考虑自定义error，用于屏蔽不同sql底层的错误；
而实际上中，业务方不关心具体的数据库错误，大部分只关心两个问题：有没有找到数据？没有找到数据是本身就没有数据，还是你DB有问题，所以在dao wrap的时候堆栈信息只是作为一个重要考虑点，但不是唯一考虑的因素
```

#2.其它

##2.1为什么查询不到记录会返回  sql.ErrNoRows 错误,而不是nil?