package base

import (
	"encoding/xml"
)

const (
	XSINamespace        = "http://www.w3.org/2001/XMLSchema-instance"
	DefaultEnvelopeAttr = "http://schemas.xmlsoap.org/soap/envelope/"
)

type Envelope struct {
	XMLName xml.Name      `xml:"s:Envelope"`
	XmlnsI  string        `xml:"xmlns:i,attr"`
	XmlnsS  string        `xml:"xmlns:s,attr"`
	Header  RequestHeader `xml:"s:Header"`
	Body    any           `xml:"s:Body"`
}

// MarshalXML 自定义 Envelope 的 XML 序列化
//func (e Envelope) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
//	// 创建 Envelope 开始标签，带命名空间属性
//	envelopeStart := xml.StartElement{
//		Name: xml.Name{Local: "s:Envelope"},
//		Attr: []xml.Attr{
//			{Name: xml.Name{Local: "xmlns:i"}, Value: e.XmlnsI},
//			{Name: xml.Name{Local: "xmlns:s"}, Value: e.XmlnsS},
//		},
//	}
//
//	if err := enc.EncodeToken(envelopeStart); err != nil {
//		return err
//	}
//
//	// 编码 Header
//	headerStart := xml.StartElement{Name: xml.Name{Local: "s:Header"}}
//
//	// 添加命名空间属性
//	if e.Header.Namespace != "" {
//		headerStart.Attr = append(headerStart.Attr, xml.Attr{
//			Name:  xml.Name{Local: "xmlns"},
//			Value: e.Header.Namespace,
//		})
//	}
//
//	if err := enc.EncodeToken(headerStart); err != nil {
//		return err
//	}
//
//	// Action 元素，带 mustUnderstand 属性
//	actionStart := xml.StartElement{
//		Name: xml.Name{Local: "Action"},
//		Attr: []xml.Attr{},
//	}
//
//	// mustUnderstand属性是无前缀的
//	if e.Header.MustUnderstand != "" {
//		actionStart.Attr = append(actionStart.Attr, xml.Attr{
//			Name:  xml.Name{Local: "mustUnderstand"},
//			Value: e.Header.MustUnderstand,
//		})
//	}
//
//	if err := enc.EncodeToken(actionStart); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(xml.CharData(e.Header.Action)); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(actionStart.End()); err != nil {
//		return err
//	}
//
//	// AuthenticationToken 元素
//	authStart := xml.StartElement{Name: xml.Name{Local: "AuthenticationToken"}}
//	if err := enc.EncodeToken(authStart); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(xml.CharData(e.Header.AuthenticationToken)); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(authStart.End()); err != nil {
//		return err
//	}
//
//	// CustomerAccountId 元素
//	acctStart := xml.StartElement{Name: xml.Name{Local: "CustomerAccountId"}}
//	if err := enc.EncodeToken(acctStart); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(xml.CharData(e.Header.CustomerAccountId)); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(acctStart.End()); err != nil {
//		return err
//	}
//
//	// CustomerId 元素
//	custStart := xml.StartElement{Name: xml.Name{Local: "CustomerId"}}
//	if err := enc.EncodeToken(custStart); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(xml.CharData(e.Header.CustomerId)); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(custStart.End()); err != nil {
//		return err
//	}
//
//	// DeveloperToken 元素
//	devStart := xml.StartElement{Name: xml.Name{Local: "DeveloperToken"}}
//	if err := enc.EncodeToken(devStart); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(xml.CharData(e.Header.DeveloperToken)); err != nil {
//		return err
//	}
//	if err := enc.EncodeToken(devStart.End()); err != nil {
//		return err
//	}
//
//	// 结束 Header
//	if err := enc.EncodeToken(headerStart.End()); err != nil {
//		return err
//	}
//
//	// 编码 Body
//	bodyStart := xml.StartElement{Name: xml.Name{Local: "s:Body"}}
//	if err := enc.EncodeElement(e.Body, bodyStart); err != nil {
//		return err
//	}
//
//	// 结束 Envelope
//	if err := enc.EncodeToken(envelopeStart.End()); err != nil {
//		return err
//	}
//
//	return nil
//}

// RequestHeader 表示 SOAP 请求头
type RequestHeader struct {
	XMLName             xml.Name `xml:"s:Header"`
	Namespace           string   `xml:"xmlns,attr"`
	Action              string   `xml:"Action"`
	MustUnderstand      string   `xml:"-"`
	AuthenticationToken string   `xml:"AuthenticationToken"`
	CustomerAccountId   string   `xml:"CustomerAccountId"`
	CustomerId          string   `xml:"CustomerId"`
	DeveloperToken      string   `xml:"DeveloperToken"`
}

// ResponseHeader 表示 SOAP 响应头
type ResponseHeader struct {
	XMLName    xml.Name `xml:"Header"`
	Namespace  string   `xml:"xmlns,attr,omitempty"`
	TrackingId string   `xml:"TrackingId"`
}

type Fault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      struct {
		XMLName          xml.Name
		AdApiFaultDetail struct {
			TrackingId string `xml:"TrackingId"`
			Errors     struct {
				AdApiError []struct {
					Code      int    `xml:"Code"`
					ErrorCode string `xml:"ErrorCode"`
					Message   string `xml:"Message"`
				} `xml:"AdApiError"`
			} `xml:"Errors"`
		} `xml:"AdApiFaultDetail"`
	} `xml:"detail"`
}
