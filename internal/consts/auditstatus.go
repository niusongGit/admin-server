package consts

type AuditStatus int64

const (
	AuditWaiting AuditStatus = iota //0待审核
	AuditPass                       //1审核成功
	AuditDenied                     //2审核未通过
)

var AuditStatusMsg = map[AuditStatus]string{
	AuditWaiting: "待审核",
	AuditPass:    "审核成功",
	AuditDenied:  "审核未通过",
}

func (c AuditStatus) Int() int64 {
	return int64(c)
}

func (c AuditStatus) Equals(i int64) bool {
	return c.Int() == i
}

func (c AuditStatus) GetAuditStatusMsg() string {
	return AuditStatusMsg[c]
}
