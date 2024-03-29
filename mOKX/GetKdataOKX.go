package mOKX

import (
	"strconv"

	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mTime"
	jsoniter "github.com/json-iterator/go"
)

type OkxCandleDataType [9]string // Okx 原始数据
func GetKdataOKX(opt GetKdataOpt) (resData []TypeKd) {
	resData = []TypeKd{}

	if len(opt.InstID) < 2 {
		return
	}

	BarObj := GetBarOpt(opt.Bar)
	if BarObj.Interval < mTime.UnixTimeInt64.Minute {
		return
	}

	Size := 100

	// 时间设置
	now := mTime.GetUnixInt64()
	after := mTime.GetUnixInt64()
	// 时间必须大于6年前
	if opt.After > now-mTime.UnixTimeInt64.Day*2190 {
		after = opt.After
	}
	// 处理分页
	if opt.Page > 0 {
		pastTime := int64(opt.Page) * BarObj.Interval * int64(Size) // 一页数据 =  100 * 时间间隔
		after = after - pastTime                                    // 减去过去的时间节点
	}

	// 判断应该采取哪个接口获取数据  after 距离 now 有多少条数据?
	path := "/api/v5/market/candles"
	fromNowItem := (now - after) / BarObj.Interval
	if fromNowItem > 800 { // 大于 800 条就取历史数据
		path = "/api/v5/market/history-index-candles"
	}

	fetchData, err := FetchOKX(OptFetchOKX{
		Path: path,
		Data: map[string]any{
			"instId": opt.InstID,
			"bar":    BarObj.Okx,
			"after":  strconv.FormatInt(after, 10),
			"limit":  Size,
		},
		Method: "get",
	})
	if err != nil {
		return
	}
	var result TypeReq
	jsoniter.Unmarshal(fetchData, &result)
	if result.Code != "0" {
		return
	}

	jsonByte := mJson.ToJson(result.Data)

	var list []OkxCandleDataType
	err = jsoniter.Unmarshal(jsonByte, &list)
	if err != nil {
		return
	}
	if len(list) != Size {
		return
	}

	KdataList := []TypeKd{} // 声明存储
	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		kdata := TypeKd{
			InstID:   opt.InstID,
			TimeStr:  mTime.UnixFormat(item[0]),
			TimeUnix: mTime.ToUnixMsec(mTime.MsToTime(item[0], "0")),
			O:        item[1],
			H:        item[2],
			L:        item[3],
			C:        item[4],
			Vol:      item[5],
		}
		new_Kdata := NewKD(kdata, KdataList)
		KdataList = append(KdataList, new_Kdata)
	}

	if len(KdataList) == Size {
		resData = KdataList
	}

	return
}
