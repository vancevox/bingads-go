package common

import (
	"encoding/xml"
	"fmt"
)

// XMLHelper 提供XML序列化相关的辅助函数
type XMLHelper struct{}

// NewXMLHelper 创建一个新的XML辅助工具
func NewXMLHelper() *XMLHelper {
	return &XMLHelper{}
}

// MarshalIndent 将对象序列化为格式化的XML
func (h *XMLHelper) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	data, err := xml.MarshalIndent(v, prefix, indent)
	if err != nil {
		return nil, fmt.Errorf("XML序列化失败: %v", err)
	}
	return data, nil
}

// Marshal 将对象序列化为XML
func (h *XMLHelper) Marshal(v any) ([]byte, error) {
	data, err := xml.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("XML序列化失败: %v", err)
	}
	return data, nil
}

// Unmarshal 将XML反序列化为对象
func (h *XMLHelper) Unmarshal(data []byte, v any) error {
	if err := xml.Unmarshal(data, v); err != nil {
		return fmt.Errorf("XML反序列化失败: %v", err)
	}
	return nil
}

// FormatXML 格式化XML字符串
func (h *XMLHelper) FormatXML(data []byte) ([]byte, error) {
	var v interface{}
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, fmt.Errorf("XML解析失败: %v", err)
	}

	return h.MarshalIndent(v, "", "  ")
}
