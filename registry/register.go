package registry

import "github.com/pjoc-team/healthz/api"

// healthTypes 所有实现者
var healthTypes = make(map[string][]*componentAndName)

// componentAndName 组件和名字组合
type componentAndName struct {
	Name   string
	Health api.Health
}

// RegisterHealth 注册服务
func RegisterHealth(componentType string, name string, h api.Health) error {
	can := &componentAndName{
		Name:   name,
		Health: h,
	}
	components, ok := healthTypes[componentType]
	if ok {
		components = append(components, can)
		return nil
	}
	healthTypes[componentType] = []*componentAndName{can}
	return nil
}
