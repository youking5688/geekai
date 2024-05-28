package types

// * +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
// * Copyright 2023 The Geek-AI Authors. All rights reserved.
// * Use of this source code is governed by a Apache-2.0 license
// * that can be found in the LICENSE file.
// * @Author yangjian102621@163.com
// * +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

import (
	"fmt"
)

type AppConfig struct {
	Path           string `toml:"-"`
	Listen         string
	Session        Session
	AdminSession   Session
	ProxyURL       string
	MysqlDns       string                  // mysql 连接地址
	StaticDir      string                  // 静态资源目录
	StaticUrl      string                  // 静态资源 URL
	Redis          RedisConfig             // redis 连接信息
	ApiConfig      ApiConfig               // ChatPlus API authorization configs
	SMS            SMSConfig               // send mobile message config
	OSS            OSSConfig               // OSS config
	MjProxyConfigs []MjProxyConfig         // MJ proxy config
	MjPlusConfigs  []MjPlusConfig          // MJ plus config
	WeChatBot      bool                    // 是否启用微信机器人
	SdConfigs      []StableDiffusionConfig // sd AI draw service pool

	XXLConfig      XXLConfig
	AlipayConfig   AlipayConfig
	WxpayConfig    WxpayConfig
	HuPiPayConfig  HuPiPayConfig
	SmtpConfig     SmtpConfig // 邮件发送配置
	JPayConfig     JPayConfig // payjs 支付配置}
}	

type SmtpConfig struct {
	UseTls   bool // 是否使用 TLS 发送
	Host     string
	Port     int
	AppName  string // 应用名称
	From     string // 发件人邮箱地址
	Password string // 发件人邮箱密码
}

type ApiConfig struct {
	ApiURL string
	AppId  string
	Token  string
}

type MjProxyConfig struct {
	Enabled bool
	ApiURL  string // api 地址
	Mode    string // 绘画模式，可选值：fast/turbo/relax
	ApiKey  string
}

type StableDiffusionConfig struct {
	Enabled bool
	Model   string // 模型名称
	ApiURL  string
	ApiKey  string
}

type MjPlusConfig struct {
	Enabled bool   // 如果启用了 MidJourney Plus，将会自动禁用原生的MidJourney服务
	ApiURL  string // api 地址
	Mode    string // 绘画模式，可选值：fast/turbo/relax
	ApiKey  string
}

type AlipayConfig struct {
	Enabled         bool   // 是否启用该支付通道
	SandBox         bool   // 是否沙盒环境
	AppId           string // 应用 ID
	UserId          string // 支付宝用户 ID
	PrivateKey      string // 用户私钥文件路径
	PublicKey       string // 用户公钥文件路径
	AlipayPublicKey string // 支付宝公钥文件路径
	RootCert        string // Root 秘钥路径
	NotifyURL       string // 异步通知回调
	ReturnURL       string // 支付成功返回地址
}

type WxpayConfig struct {
	Enabled             bool   // 是否启用该支付通道
	SandBox             bool   // 是否沙盒环境
	AppId               string // 应用 ID
	WxAppSecret         string // 应用 Secret
	MchId               string // 商户 ID
	MchKey              string // 商户key
	CertificateSerialNo string // 商户key
	PrivateKey          string // 商户私密钥文件地址
	NotifyURL           string // 异步通知回调
	ReturnURL           string // 支付成功返回地址
}

type HuPiPayConfig struct { //虎皮椒第四方支付配置
	Enabled   bool   // 是否启用该支付通道
	Name      string // 支付名称，如：wechat/alipay
	AppId     string // App ID
	AppSecret string // app 密钥
	ApiURL    string // 支付网关
	NotifyURL string // 异步通知回调
	ReturnURL string // 支付成功返回地址
}

// JPayConfig PayJs 支付配置
type JPayConfig struct {
	Enabled    bool
	Name       string // 支付名称，默认 wechat
	AppId      string // 商户 ID
	PrivateKey string // 私钥
	ApiURL     string // API 网关
	NotifyURL  string // 异步回调地址
	ReturnURL  string // 支付成功返回地址
}

type XXLConfig struct { // XXL 任务调度配置
	Enabled      bool
	ServerAddr   string
	ExecutorIp   string
	ExecutorPort string
	AccessToken  string
	RegistryKey  string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

// LicenseKey 存储许可证书的 KEY
const LicenseKey = "Geek-AI-License"

type License struct {
	Key       string `json:"key"`        // 许可证书密钥
	MachineId string `json:"machine_id"` // 机器码
	UserNum   int    `json:"user_num"`   // 用户数量
	ExpiredAt int64  `json:"expired_at"` // 过期时间
	IsActive  bool   `json:"is_active"`  // 是否激活
}

func (c RedisConfig) Url() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type Platform string

const OpenAI = Platform("OpenAI")
const Azure = Platform("Azure")
const ChatGLM = Platform("ChatGLM")
const Baidu = Platform("Baidu")
const XunFei = Platform("XunFei")
const QWen = Platform("QWen")

type SystemConfig struct {
	Title         string `json:"title,omitempty"`
	AdminTitle    string `json:"admin_title,omitempty"`
	Logo          string `json:"logo,omitempty"`
	InitPower     int    `json:"init_power,omitempty"`      // 新用户注册赠送算力值
	DailyPower    int    `json:"daily_power,omitempty"`     // 每日赠送算力
	InvitePower   int    `json:"invite_power,omitempty"`    // 邀请新用户赠送算力值
	VipMonthPower int    `json:"vip_month_power,omitempty"` // VIP 会员每月赠送的算力值

	RegisterWays    []string `json:"register_ways,omitempty"`    // 注册方式：支持手机（mobile），邮箱注册（email），账号密码注册
	EnabledRegister bool     `json:"enabled_register,omitempty"` // 是否开放注册

	RewardImg     string  `json:"reward_img,omitempty"`     // 众筹收款二维码地址
	EnabledReward bool    `json:"enabled_reward,omitempty"` // 启用众筹功能
	PowerPrice    float64 `json:"power_price,omitempty"`    // 算力单价

	OrderPayTimeout int    `json:"order_pay_timeout,omitempty"` //订单支付超时时间
	VipInfoText     string `json:"vip_info_text,omitempty"`     // 会员页面充值说明
	DefaultModels   []int  `json:"default_models,omitempty"`    // 默认开通的 AI 模型

	MjPower       int `json:"mj_power,omitempty"`        // MJ 绘画消耗算力
	MjActionPower int `json:"mj_action_power,omitempty"` // MJ 操作（放大，变换）消耗算力
	SdPower       int `json:"sd_power,omitempty"`        // SD 绘画消耗算力
	DallPower     int `json:"dall_power,omitempty"`      // DALLE3 绘图消耗算力

	WechatCardURL string `json:"wechat_card_url,omitempty"` // 微信客服地址

	EnableContext bool `json:"enable_context,omitempty"`
	ContextDeep   int  `json:"context_deep,omitempty"`
}
