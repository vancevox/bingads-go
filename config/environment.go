package config

// 环境类型
type Environment string

const (
	// 生产环境
	Production Environment = "production"

	// 沙箱环境
	Sandbox Environment = "sandbox"
)

// API 端点
const (
	// 生产环境的 Campaign Management API 端点
	ProductionCampaignEndpoint = "https://campaign.api.bingads.microsoft.com/Api/Advertiser/CampaignManagement/v13/CampaignManagementService.svc"

	// 沙箱环境的 Campaign Management API 端点
	SandboxCampaignEndpoint = "https://campaign.api.sandbox.bingads.microsoft.com/Api/Advertiser/CampaignManagement/v13/CampaignManagementService.svc"
)

// 命名空间常量
const (
	// XML Schema Instance 命名空间
	XSINamespace = "http://www.w3.org/2001/XMLSchema-instance"

	// SOAP 信封命名空间
	SOAPEnvelopeNamespace = "http://schemas.xmlsoap.org/soap/envelope/"

	// Campaign Management API 命名空间
	CampaignManagementNamespace = "https://bingads.microsoft.com/CampaignManagement/v13"
)

// APIConfig 包含 Bing Ads API 的环境配置
type APIConfig struct {
	// 环境（生产或沙箱）
	Env Environment

	// 超时设置（秒）
	Timeout int

	// 重试次数
	MaxRetries int

	// 是否启用调试模式
	Debug bool
}

// DefaultConfig 返回默认的 API 配置
func DefaultConfig() *APIConfig {
	return &APIConfig{
		Env:        Production,
		Timeout:    30,
		MaxRetries: 3,
		Debug:      false,
	}
}

// GetCampaignEndpoint 根据环境获取 Campaign Management API 端点
func (c *APIConfig) GetCampaignEndpoint() string {
	if c.Env == Sandbox {
		return SandboxCampaignEndpoint
	}
	return ProductionCampaignEndpoint
}

// Config 包含所有 Bing Ads API 配置
type Config struct {
	// 认证配置
	Auth *AuthConfig

	// API 配置
	API *APIConfig
}

// NewConfig 创建一个新的配置
func NewConfig(auth *AuthConfig, api *APIConfig) *Config {
	if api == nil {
		api = DefaultConfig()
	}

	return &Config{
		Auth: auth,
		API:  api,
	}
}
