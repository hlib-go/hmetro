package hmetro

/*

通联：
appId 商户编码：1609128932911
秘钥：1343410230280720384
接口服务地址：https://itapdev-api.ucitymetro.com

产品编号： c003

*/

type Config struct {
	ServiceUrl string `json:"serviceUrl"`
	AppId      string `json:"appId"`
	Secret     string `json:"secret"`
	SecretAes  string `json:"secretAes"`
	Qrpage     string `json:"qrpage"` // 计次票二维码H5页面，参数 {path} 二维码地址  {sign} 手机号加密字符串  示例：https://itapdev.ucitymetro.com/appentry?path=/eleTicket/qrpage/1325725767652478976&sign=手机号加密字符串
	Path       string `json:"path"`   // /ticket/qrcode-nbhy/11111111111111111
}
