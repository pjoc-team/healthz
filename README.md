### 架构图

```plantuml
!includeurl https://raw.githubusercontent.com/blademainer/plantuml-style-c4/master/c4_component.puml

Container(kubelet, "Kubelet", "Runtime in node")

Boundary(app, "App"){
  Component(appc, "AppCode", "App Code")
  
  Component(http, "HTTP API", "提供http服务给kubernetes探测服务是否readiness")
  Component(registry, "HealthRegistry", "提供API注册/Health状态查询")
  Component(api, "HealthAPI", "提供注册Health的函数: Register(type, APIInterface)")
  Component(componenta, "ComponentA", "组件A，比如MySQL")
  Component(componentb, "ComponentB", "组件B，比如RocketMQ")
  Component(componentx, "Component...", "组件xxx")

  componenta -U-> api: implements
  componentb -U-> api: implements
  componentx -U-> api: implements

  componenta --> registry: Register(type, componenta)
  componentb --> registry: Register(type, componentb)
  componentx --> registry: Register(type, componentx)
  
  http --> registry
  appc --> componenta
  appc --> componentb
  appc --> componentx
}

kubelet -> http: readiness
```