
不要将 Context 存储在结构体中，应该将 Context 显示传递给每个函数，一般做为第一个参数。
一个 Context 被取消，它的子 Context 也会被取消。
一个 Context 可以被传递给多个 goroutine 同时使用。 

https://medium.com/codex/go-context-101-ebfaf655fa95

https://www.ardanlabs.com/blog/2019/09/context-package-semantics-in-go.html


