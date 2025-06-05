package service

import (
	"encoding/xml"
	"fmt"

	"github.com/vancevox/bingads/base"
	"github.com/vancevox/bingads/campaignManagement/models"
	"github.com/vancevox/bingads/common"
	"github.com/vancevox/bingads/config"
)

// Client 实现 CampaignManagementAPI 接口
type Client struct {
	Config     *config.Config
	HTTPClient *common.HTTPClient
	XMLHelper  *common.XMLHelper
}

// NewClient 创建一个新的 Campaign Management API 客户端
func NewClient(cfg *config.Config) *Client {
	return &Client{
		Config:     cfg,
		HTTPClient: common.NewHTTPClient(cfg),
		XMLHelper:  common.NewXMLHelper(),
	}
}

// SharedListService 返回共享列表服务
func (c *Client) SharedListService() models.SharedListService {
	return NewSharedListService(c)
}

// 创建 SOAP 请求头
func (c *Client) createRequestHeader(action models.SOAPAction, mustUnderstand string) base.RequestHeader {
	return base.RequestHeader{
		Namespace:           config.CampaignManagementNamespace,
		Action:              string(action),
		MustUnderstand:      mustUnderstand,
		AuthenticationToken: c.Config.Auth.AuthenticationToken,
		CustomerAccountId:   c.Config.Auth.CustomerAccountID,
		CustomerId:          c.Config.Auth.CustomerID,
		DeveloperToken:      c.Config.Auth.DeveloperToken,
	}
}

// 创建 SOAP 信封
func (c *Client) createEnvelope(action models.SOAPAction, mustUnderstand string) base.Envelope {
	return base.Envelope{
		XMLName: xml.Name{},
		XmlnsI:  config.XSINamespace,
		XmlnsS:  config.SOAPEnvelopeNamespace,
		Header:  c.createRequestHeader(action, mustUnderstand),
		Body:    nil,
	}
}

// 发送请求
func (c *Client) sendRequest(envelope base.Envelope, action models.SOAPAction) ([]byte, error) {
	// 序列化请求
	reqBody, err := c.XMLHelper.Marshal(envelope)
	if err != nil {
		return nil, base.NewError(base.ErrSerializationFail, "序列化请求失败", err)
	}

	// 添加 XML 声明
	reqBody = append([]byte(xml.Header), reqBody...)

	if c.Config.API.Debug {
		fmt.Println("请求体:", string(reqBody))
	}
	// 发送请求
	respBody, err := c.HTTPClient.Post(c.Config.API.GetCampaignEndpoint(), string(action), reqBody)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

// 处理响应
func (c *Client) processResponse(respBody []byte, respObj any) error {
	// 打印原始响应内容，用于调试
	if c.Config.API.Debug {
		fmt.Println("原始响应:", string(respBody))
	}

	var genericResp models.CampaignManagementResponseEnvelope
	// 先解析为通用结构，检查是否有错误
	//var genericResp GenericResponseEnvelope
	if err := c.XMLHelper.Unmarshal(respBody, &genericResp); err != nil {
		return base.NewError(base.ErrDeserializationFail, "反序列化响应失败", err)
	}

	// 检查是否有SOAP故障
	if genericResp.Body.Fault != nil {
		fault := genericResp.Body.Fault
		errorMsg := fmt.Sprintf("SOAP错误: %s - %s", fault.FaultCode, fault.FaultString)

		// 检查是否有详细错误信息
		if fault.Detail.AdApiFaultDetail.Errors.AdApiError != nil && len(fault.Detail.AdApiFaultDetail.Errors.AdApiError) > 0 {
			apiError := fault.Detail.AdApiFaultDetail.Errors.AdApiError[0]
			errorMsg = fmt.Sprintf("BingAds API错误 [%s]: %s", apiError.ErrorCode, apiError.Message)
		}

		// 记录跟踪ID（如果有）
		if genericResp.Header.TrackingId != "" {
			errorMsg += fmt.Sprintf(" (TrackingId: %s)", genericResp.Header.TrackingId)
		}

		return fmt.Errorf(errorMsg)
	}

	// 尝试将响应直接解析到目标对象
	if err := c.XMLHelper.Unmarshal(respBody, respObj); err != nil {
		return base.NewError(base.ErrDeserializationFail, fmt.Sprintf("反序列化响应对象失败: %v", err), nil)
	}

	return nil
}
