package libfrp

import (
	frpLog "frp_lib/pkg/util/log"
)

type FrpLogListener interface {
	Log(log string)
	Location() string
}

// type FrpLogListener struct {
// 	name string
// }

// func (l *FrpLogListener) Log(log string) {
// }

func SetFrpLogListener(l FrpLogListener) {
	frpLog.AppendListener(l)
}
