# 将数据表转换成结构体，并将内容复制到剪切板

#### 本代码只限公司项目使用

- 将内容复制到剪切板：

```go
github.com/atotto/clipboard
```

- 复制到项目中后不要忘记检查tinyint的字段（默认为int类型），根据实际情况，修改成bool

  #### 运行命令

  ```go
  go run main -t="数据表名称"
  ```

  