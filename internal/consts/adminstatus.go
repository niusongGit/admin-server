package consts

type AdminStatus int64

const (
	AdminStatusNo  AdminStatus = iota // 管理员账户开启
	AdminStatusOff                    // 管理员账户禁用
)

var AdminStatusMsg = map[AdminStatus]string{
	AdminStatusNo:  "管理员账户开启",
	AdminStatusOff: "管理员账户禁用",
}

func (c AdminStatus) Int() int64 {
	return int64(c)
}

func (c AdminStatus) Equals(i int64) bool {
	return c.Int() == i
}

func (c AdminStatus) GetCarCategoryMsg() string {
	return AdminStatusMsg[c]
}
