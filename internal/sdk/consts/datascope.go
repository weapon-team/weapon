package consts

// 数据权限
const (
	DataScope_All      = 1 // 全部数据
	DataScope_DeptBlow = 2 // 部门及以下数据
	DataScope_DeptOnly = 3 // 本部门数据
	DataScope_Personal = 4 // 个人数据
	DataScope_Custom   = 5 // 自定义数据
)
