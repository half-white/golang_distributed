package registry

type Registration struct {
	ServiceName      ServiceName   //服务名称
	ServiceURL       string        //服务URL
	RequiredServices []ServiceName //该服务依赖的服务名称
	ServiceUpdateURL string        //与服务注册中心进行通信
	HeartbeatURL     string        //用来做心跳检查的URL
}

type ServiceName string

// 所有的服务列表
const (
	LogService    = ServiceName("LogService")
	GradeService  = ServiceName("GradeService")
	PortalService = ServiceName("PortalService")
)

// 用于服务状态更新的结构体
type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
