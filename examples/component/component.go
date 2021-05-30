package main

import (
	"errors"
	"fmt"
	"github.com/pjoc-team/healthz/api"
	"github.com/pjoc-team/healthz/registry"
	"math/rand"
	"time"
)

const (
	componentType = "demo"
)

var index = 0

type exampleComponent struct {
	status api.Status
}

// NewExampleComponent create example component
func NewExampleComponent() (*exampleComponent, error) {
	ex := &exampleComponent{}
	ex.status = api.Loading
	err := registry.RegisterHealth(componentType, fmt.Sprintf("demo-%d", index), ex)
	index++
	if err != nil {
		return nil, err
	}
	err = ex.Start()
	if err != nil {
		ex.status = api.Down
		return nil, err
	}
	ex.status = api.Up
	return ex, nil
}

func (e *exampleComponent) Do() {
	// component's logic
}

// Stat 供health registry统计状态
func (e *exampleComponent) Stat() ([]*api.Condition, error) {
	return []*api.Condition{
		{
			Type:    api.ConditionTypeHealth,
			Status:  e.status,
			Message: e.status.String(),
		},
	}, nil
}

func (e *exampleComponent) Start() error {
	time.Sleep(3 * time.Second)
	i := rand.Int()
	if i%2 == 0 { // 模拟失败，使得状态有可能为DOWN
		return errors.New("start failed")
	}
	return nil
}
