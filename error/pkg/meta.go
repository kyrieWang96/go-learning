package pkg

// 定义你的ErrorCode 可以有多种方式

const (
	ErrorNotExistData     ErrorType = iota // 数据不存在错误
	ErrorInvalidParameter                  // 非法参数
)
