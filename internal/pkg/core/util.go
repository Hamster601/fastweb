package core

// NoCheckAndLogPath 不需要记录日志，和记录性能指标的接口
var WithoutCheckAndLogPath = map[string]struct{}{
	"/metrics":                  {}, // 上报监控数据
	"/debug/pprof/":             {}, // pprof相关
	"/debug/pprof/cmdline":      {},
	"/debug/pprof/profile":      {},
	"/debug/pprof/symbol":       {},
	"/debug/pprof/trace":        {},
	"/debug/pprof/allocs":       {},
	"/debug/pprof/block":        {},
	"/debug/pprof/goroutine":    {},
	"/debug/pprof/heap":         {},
	"/debug/pprof/mutex":        {},
	"/debug/pprof/threadcreate": {},
	"/favicon.ico":              {},
	"/system/health":            {}, // 健康检测
}
