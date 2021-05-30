package api

import "time"

const (
	// Up 健康
	Up Status = 1

	// Loading 加载中
	Loading Status = 5

	// Down 异常
	Down Status = 9
)

const (
	// ConditionTypeHealth 健康类型
	ConditionTypeHealth = "HEALTHY"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type Status
// Status 状态枚举值
type Status int

// ConditionType 类型
type ConditionType string

// Stat 组件状态
type Stat struct {
	Metadata   *Metadata    `json:"metadata,omitempty"`
	Conditions []*Condition `json:"conditions,omitempty"`
}

// Metadata 组件信息
type Metadata struct {
	Name              string    `json:"name,omitempty"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
}

// Condition 状态
type Condition struct {
	Type    ConditionType `json:"type,omitempty"`
	Status  Status        `json:"status,omitempty"`
	Message string        `json:"message,omitempty"`
}

// Level 返回级别，如果是多个condition一起取值，则取最大的一个作为终态值
func (i Status) Level() int {
	return int(i)
}

func (i Status) Marshal() ([]byte, error) {
	return []byte(i.String()), nil
}
