package api

// Health 健康检查接口
type Health interface {
	// Stat 返回健康状态
	Stat() ([]*Condition, error)
}
