package alipay

/*
apiname	String		服务对应的名称，常量值为com.alipay.account.auth com.alipay.account.auth
method	String		接口名称，常量值为alipay.open.auth.sdk.code.get alipay.open.auth.sdk.code.get
app_id	String		支付宝分配给开发者的应用ID 2014123100022800
app_name	String	调用来源方的标识，常量值为mc mc
biz_type	String	调用业务的类型，常量值为openservice openservice
pid	String			签约的支付宝账号对应的支付宝唯一用户号，以2088开头的16位纯数字组成 2088123456789012
product_id	String	产品码，常量值为APP_FAST_LOGIN APP_FAST_LOGIN
scope	String		授权范围，常量值为kuaijie kuaijie
target_id	String	商户标识该次用户授权请求的ID，该值在商户端应保持唯一 kkkkk091125
auth_type	String	标识授权类型，取值范围：AUTHACCOUNT代表授权；LOGIN代表登录 AUTHACCOUNT
sign_type	String	商户生成签名字符串所使用的签名算法类型，目前支持RSA2和RSA，推荐使用RSA2 RSA2
sign	String		整个授权参数信息的签名
 */
type TokenPay struct {
	ApiName   string `json:"apiname,omitempty"`
	AppName   string `json:"app_name,omitempty"`
	AuthType  string `json:"auth_type,omitempty"`
	BizType   string `json:"biz_type,omitempty"`
	Method    string `json:"method,omitempty"`
	Pid       string `json:"pid,omitempty"`
	ProductId string `json:"product_id,omitempty"`
	Scope     string `json:"scope,omitempty"`
	SignType  string `json:"sign_type,omitempty"`
	TargetId  string `json:"target_id,omitempty"`
}

func (this TokenPay) APIName() string {
	return "alipay.trade.create"
}

func (this TokenPay) Params() map[string]string {
	var m = make(map[string]string)
	m["pid"] = this.Pid
	m["target_id"] = this.TargetId
	return m
}

func (this TokenPay) ExtJSONParamName() string {
	return ""
}

func (this TokenPay) ExtJSONParamValue() string {
	return marshal(this)
}
