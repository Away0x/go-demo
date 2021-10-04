package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequiredServices []ServiceName // 定义服务依赖的其他服务
	ServiceUpdateURL string        // 接收服务注册中心的通知
	HeartbeatURL     string        // 心跳检测
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)

// 注册中心服务更新时, 发送给注册服务的通知 struct
type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
