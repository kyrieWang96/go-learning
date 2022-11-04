package pkg

// 定义你的ErrorCode

const (
	ErrorNotExistData     ErrorType = iota // 数据不存在错误
	ErrorInvalidParameter                  // 非法参数
)
