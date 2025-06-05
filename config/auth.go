package config

// AuthConfig 包含 Bing Ads API 的认证配置
type AuthConfig struct {
	// 开发者令牌，从 Bing Ads 开发者中心获取
	DeveloperToken string

	// 认证令牌，通过 OAuth 流程获取
	AuthenticationToken string

	// 客户 ID
	CustomerID string

	// 客户账户 ID
	CustomerAccountID string
}

// NewAuthConfig 创建一个新的认证配置
func NewAuthConfig(developerToken, authToken, customerID, accountID string) *AuthConfig {
	return &AuthConfig{
		DeveloperToken:      developerToken,
		AuthenticationToken: authToken,
		CustomerID:          customerID,
		CustomerAccountID:   accountID,
	}
}

// IsValid 检查认证配置是否有效
func (c *AuthConfig) IsValid() bool {
	// 开发者令牌和认证令牌是必需的
	if c.DeveloperToken == "" || c.AuthenticationToken == "" {
		return false
	}

	// 客户 ID 是必需的
	if c.CustomerID == "" {
		return false
	}

	return true
}
