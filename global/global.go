package global

import "flag"

var (
	RegularStartMoney  int
	OvertimeStartMoney int
)

func SetGlobal() {
	rMoney := flag.Int("rM", 800, "Regular start money (Default: 800)")
	otMoney := flag.Int("otM", 16000, "Overtime start money (Default: 16000)")
	flag.Parse()

	RegularStartMoney = *rMoney
	OvertimeStartMoney = *otMoney
}
