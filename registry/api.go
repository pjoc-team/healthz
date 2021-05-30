package registry

import (
	"encoding/json"
	api "github.com/pjoc-team/healthz/api"
	"github.com/pjoc-team/tracing/logger"
	"time"
)

const (
	healthzName = "healthz"
)

// HealthInfo 服务状态
type HealthInfo struct {
	Status   string
	Items    []*api.Stat
	Metadata *api.Metadata
}

func (i *HealthInfo) String() string {
	marshal, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(marshal)
}

// Registry registry
type Registry struct {
}

// HealthStatus 健康状态
func HealthStatus() (*HealthInfo, error) {
	log := logger.Log()
	info := &HealthInfo{
		Metadata: &api.Metadata{
			Name:              healthzName,
			CreationTimestamp: time.Now(),
		},
	}

	status := api.Up

	for tp, healths := range healthTypes {
		for _, health := range healths {
			md := &api.Metadata{
				Name:              health.Name,
				CreationTimestamp: time.Now(),
			}

			conditions, err := health.Health.Stat()
			if err != nil {
				log.Errorf(
					"failed to check stat of component, type: %v, name: %v, err: %v", tp,
					health.Name, err.Error(),
				)
				return nil, err
			}
			// 如果condition内的status高于当前status，则取condition内的status
			for _, condition := range conditions {
				if condition.Status.Level() > status.Level() {
					status = condition.Status
				}
			}

			stat := &api.Stat{
				Metadata:   md,
				Conditions: conditions,
			}

			info.Items = append(info.Items, stat)
		}
	}
	info.Status = status.String()
	return info, nil
}
