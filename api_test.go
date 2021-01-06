package hmetro

import (
	"testing"
)

var cfg = &Config{
	ServiceUrl: "https://itapdev-api.ucitymetro.com",
	AppId:      "1609128932911",
	Secret:     "1343410230280720384",
	SecretAes:  "868971231616403394817a2360c4e8b2",
	Qrpage:     "https://itapdev.ucitymetro.com/appentry?path={path}&sign={sign}",
}

func TestAuthByMobile(t *testing.T) {
	userId, err := AuthByMobile(cfg, "13611703040")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("userId:", userId)
	t.Log("success...")
	// 手机号授权测试结果：OK 正常返回用户ID b89b4187202240b7a49007901305a17b
}

func TestProductInfo(t *testing.T) {
	prod, err := ProductInfo(cfg, "d001")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(prod)
	t.Log("success...")
	// 产品信息查询测试结果：返回产品信息比接口文档定义的多,且字段名与文档不一致，如：可乘车次数文档字段名 times,实际返回没有
}

func TestMonthlyTicketOpen(t *testing.T) {
	err := MonthlyTicketOpen(cfg, "b89b4187202240b7a49007901305a17b", "d001", Rand32())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("success...")
	// 开通月票接口测试结果： 无此商品代理授权.[3014]
}

// 月票核销推送 ,通知请求URL什么时候配置

//计次票二维码H5页面嵌入
func TestEntry(t *testing.T) {
	path := "/eleTicket/qrpage/1346708681668038656"
	url, err := Entry(cfg, path, "13611703040")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(url)
	// 测试结果：
}
