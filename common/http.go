package common

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/vancevox/bingads-go/base"
	"github.com/vancevox/bingads-go/config"
)

// 使用相对路径的备用方案
// import (
//	"bytes"
//	"context"
//	"io"
//	"net/http"
//	"time"
//
//	"../base"
//	"../config"
// )

// HTTPClient 封装 HTTP 客户端功能
type HTTPClient struct {
	Client *resty.Client
	Config *config.Config
}

// NewHTTPClient 创建一个新的 HTTP 客户端
func NewHTTPClient(cfg *config.Config) *HTTPClient {
	client := resty.NewWithClient(&http.Client{
		Timeout: time.Duration(cfg.API.Timeout) * time.Second,
	})
	return &HTTPClient{
		Client: client,
		Config: cfg,
	}
}

// Post 发送 POST 请求
func (c *HTTPClient) Post(url string, action string, body []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.Config.API.Timeout)*time.Second)
	defer cancel()

	return c.PostWithContext(ctx, url, action, body)
}

// PostWithContext 使用指定的上下文发送 POST 请求
func (c *HTTPClient) PostWithContext(ctx context.Context, url string, action string, body []byte) ([]byte, error) {
	resp, err := c.Client.
		R().
		SetContext(ctx).
		SetBody(body).
		SetHeaders(map[string]string{
			"Content-Type": "text/xml; charset=utf-8",
			"SOAPAction":   action,
		}).
		Post(url)
	if err != nil {
		return nil, base.NewError(base.ErrNetworkFail, "发送 HTTP 请求失败", err)
	}

	// 检查状态码
	if resp.StatusCode() != http.StatusOK {
		return resp.Body(), base.NewError(base.ErrAPIError, fmt.Sprintf("API 返回非 200 状态码: %d", resp.StatusCode()), nil)
	}
	return resp.Body(), nil
}
