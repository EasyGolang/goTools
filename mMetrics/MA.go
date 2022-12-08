package mMetrics

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

func MA(opt EmaOpt) string {
	KDList := opt.CList
	n := opt.Cycle

	c_len := len(KDList) // K线总长
	c_n := n             // 长度
	if c_len < n {
		c_n = c_len
	}

	c_list := KDList[c_len-c_n:]
	ma_add := "0"

	var precision int32

	for _, KD := range c_list {
		// 数据源
		C := KD

		ma_add = mCount.Add(ma_add, C)
	}

	maRe := mCount.Div(ma_add, mStr.ToStr(c_n))

	maRe = mCount.CentRound(maRe, precision)

	return maRe
}
