# 单例模式

单例模式的实际使用案例是一个 [DataLoader](https://github.com/vektah/dataloaden)，按照 `dataloaden` 生成的 Loader 的代码来看，必须确保多次请求使用的 loader 是同一个 loader 实例，才能实现请求的收集和集中处理（去重）。

在 go 语言中，最快速实现单例的方式就是使用 `sync.Once`, `sync.Once` 能够保证所传入的方法只会被执行一次，这里用来给我们的单例 `userLoader` 赋值非常合适。

`DataLoader` 实际与 graphql 搭配使用更为合适。