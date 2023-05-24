package rbac

//APIModule API所属模块
//     |-> 2 --> { API权限
//     |
// 1 > |-> 3 --> { API权限集合
//     |
//     |-> 4 --> { API权限集合
type APIModule struct {
	ID         int           // 模块ID
	ModuleName string        // 模块名称
	ParentID   int           // 父模块ID
	APIs       []APIResource // 模块含有哪些API
}

//APIResource API资源定义
type APIResource struct {
	APIPath           string // API路径
	APIPathNameWithZh string // API中文名称
	HasState          bool   // 是否是有状态API
}

//APIWithAPIResource API和resource映射视图
var APIWithAPIResource map[string]APIResource
