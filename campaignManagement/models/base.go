package models

import (
	"encoding/xml"
	"fmt"

	"github.com/vancevox/bingads-go/base"
)

type AccountNegativeKeywordList struct {
	SharedList
}

type BrandList struct {
	SharedList
}

type AccountPlacementExclusionList struct {
	SharedList
}

type AccountPlacementInclusionList struct {
	SharedList
}

// KeyValuePairOfstringstring 键值对
type KeyValuePairOfstringstring struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// CampaignManagementBody 表示请求体
type CampaignManagementBody struct {
	XMLName                                             xml.Name                                             `xml:"s:Body"`
	GetListItemsBySharedListRequest                     *GetListItemsBySharedListRequest                     `xml:"GetListItemsBySharedListRequest,omitempty"`
	GetSharedEntitiesRequest                            *GetSharedEntitiesRequest                            `xml:"GetSharedEntitiesRequest,omitempty"`
	GetSharedEntityAssociationsBySharedEntityIdsRequest *GetSharedEntityAssociationsBySharedEntityIdsRequest `xml:"GetSharedEntityAssociationsBySharedEntityIdsRequest,omitempty"`
	AddListItemsToSharedListRequest                     *AddListItemsToSharedListRequest                     `xml:"AddListItemsToSharedListRequest,omitempty"`
	DeleteListItemsFromSharedListRequest                *DeleteListItemsFromSharedListRequest                `xml:"DeleteListItemsFromSharedListRequest,omitempty"`
}

// CampaignManagementResponseBody 表示响应体
type CampaignManagementResponseBody struct {
	XMLName                                              xml.Name                                              `xml:"Body"`
	Fault                                                *base.Fault                                           `xml:"Fault,omitempty"`
	GetListItemsBySharedListResponse                     *GetListItemsBySharedListResponse                     `xml:"GetListItemsBySharedListResponse,omitempty"`
	GetSharedEntitiesResponse                            *GetSharedEntitiesResponse                            `xml:"GetSharedEntitiesResponse,omitempty"`
	GetSharedEntityAssociationsBySharedEntityIdsResponse *GetSharedEntityAssociationsBySharedEntityIdsResponse `xml:"GetSharedEntityAssociationsBySharedEntityIdsResponse,omitempty"`
	AddListItemsToSharedListResponse                     *AddListItemsToSharedListResponse                     `xml:"AddListItemsToSharedListResponse,omitempty"`
	DeleteListItemsFromSharedListResponse                *DeleteListItemsFromSharedListResponse                `xml:"DeleteListItemsFromSharedListResponse,omitempty"`
}

// CampaignManagementEnvelope 表示完整的 SOAP 请求
type CampaignManagementEnvelope struct {
	XMLName xml.Name               `xml:"s:Envelope"`
	XmlnsI  string                 `xml:"xmlns:i,attr"`
	XmlnsS  string                 `xml:"xmlns:s,attr"`
	Header  base.RequestHeader     `xml:"s:Header"`
	Body    CampaignManagementBody `xml:"s:Body"`
}

// CampaignManagementResponseEnvelope 表示完整的 SOAP 响应
type CampaignManagementResponseEnvelope struct {
	XMLName xml.Name                       `xml:"Envelope"`
	XmlnsS  string                         `xml:"xmlns:s,attr,omitempty"`
	Header  base.ResponseHeader            `xml:"Header"`
	Body    CampaignManagementResponseBody `xml:"Body"`
}

// MarshalXML 自定义 CampaignManagementBody 的 XML 序列化
func (b CampaignManagementBody) MarshalXML(enc *xml.Encoder, start xml.StartElement) error {
	// 编码 Body 开始标签
	if err := enc.EncodeToken(start); err != nil {
		return err
	}

	// 编码 GetListItemsBySharedListRequest
	if b.GetListItemsBySharedListRequest != nil {
		reqStart := xml.StartElement{Name: xml.Name{Local: "GetListItemsBySharedListRequest"}}

		// 设置命名空间属性
		if b.GetListItemsBySharedListRequest.Namespace != "" {
			reqStart.Attr = append(reqStart.Attr, xml.Attr{
				Name:  xml.Name{Local: "xmlns"},
				Value: b.GetListItemsBySharedListRequest.Namespace,
			})
		}

		// 编码带命名空间的 GetListItemsBySharedListRequest 开始标签
		if err := enc.EncodeToken(reqStart); err != nil {
			return err
		}

		// 编码 SharedList 元素
		slStart := xml.StartElement{Name: xml.Name{Local: "SharedList"}}
		if err := enc.EncodeElement(b.GetListItemsBySharedListRequest.SharedList, slStart); err != nil {
			return err
		}

		// 编码 SharedEntityScope 元素
		scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
		if err := enc.EncodeToken(scopeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.GetListItemsBySharedListRequest.SharedEntityScope))); err != nil {
			return err
		}
		if err := enc.EncodeToken(scopeStart.End()); err != nil {
			return err
		}

		// 结束 GetListItemsBySharedListRequest
		if err := enc.EncodeToken(reqStart.End()); err != nil {
			return err
		}
	}

	// 编码 GetSharedEntitiesRequest
	if b.GetSharedEntitiesRequest != nil {
		reqStart := xml.StartElement{Name: xml.Name{Local: "GetSharedEntitiesRequest"}}

		// 设置命名空间属性
		if b.GetSharedEntitiesRequest.Namespace != "" {
			reqStart.Attr = append(reqStart.Attr, xml.Attr{
				Name:  xml.Name{Local: "xmlns"},
				Value: b.GetSharedEntitiesRequest.Namespace,
			})
		}

		// 编码带命名空间的 GetSharedEntitiesRequest 开始标签
		if err := enc.EncodeToken(reqStart); err != nil {
			return err
		}

		// 编码 SharedEntityType 元素
		typeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityType"}}
		if err := enc.EncodeToken(typeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.GetSharedEntitiesRequest.SharedEntityType))); err != nil {
			return err
		}
		if err := enc.EncodeToken(typeStart.End()); err != nil {
			return err
		}

		// 编码 SharedEntityScope 元素
		scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
		if err := enc.EncodeToken(scopeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.GetSharedEntitiesRequest.SharedEntityScope))); err != nil {
			return err
		}
		if err := enc.EncodeToken(scopeStart.End()); err != nil {
			return err
		}

		// 结束 GetSharedEntitiesRequest
		if err := enc.EncodeToken(reqStart.End()); err != nil {
			return err
		}
	}

	// 编码 GetSharedEntityAssociationsBySharedEntityIdsRequest
	if b.GetSharedEntityAssociationsBySharedEntityIdsRequest != nil {
		reqStart := xml.StartElement{Name: xml.Name{Local: "GetSharedEntityAssociationsBySharedEntityIdsRequest"}}

		// 设置命名空间属性
		if b.GetSharedEntityAssociationsBySharedEntityIdsRequest.Namespace != "" {
			reqStart.Attr = append(reqStart.Attr, xml.Attr{
				Name:  xml.Name{Local: "xmlns"},
				Value: b.GetSharedEntityAssociationsBySharedEntityIdsRequest.Namespace,
			})
		}

		// 编码带命名空间的 GetSharedEntityAssociationsBySharedEntityIdsRequest 开始标签
		if err := enc.EncodeToken(reqStart); err != nil {
			return err
		}

		// 编码 EntityType 元素
		entityTypeStart := xml.StartElement{Name: xml.Name{Local: "EntityType"}}
		if err := enc.EncodeToken(entityTypeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.GetSharedEntityAssociationsBySharedEntityIdsRequest.EntityType))); err != nil {
			return err
		}
		if err := enc.EncodeToken(entityTypeStart.End()); err != nil {
			return err
		}

		// 编码 SharedEntityIds 元素
		if len(b.GetSharedEntityAssociationsBySharedEntityIdsRequest.SharedEntityIds) > 0 {
			idsStart := xml.StartElement{
				Name: xml.Name{Local: "SharedEntityIds"},
				Attr: []xml.Attr{
					{
						Name:  xml.Name{Local: "xmlns:a1"},
						Value: "http://schemas.microsoft.com/2003/10/Serialization/Arrays",
					},
				},
			}
			if err := enc.EncodeToken(idsStart); err != nil {
				return err
			}

			for _, id := range b.GetSharedEntityAssociationsBySharedEntityIdsRequest.SharedEntityIds {
				idStart := xml.StartElement{Name: xml.Name{Local: "a1:long"}}
				if err := enc.EncodeToken(idStart); err != nil {
					return err
				}
				if err := enc.EncodeToken(xml.CharData(fmt.Sprintf("%d", id))); err != nil {
					return err
				}
				if err := enc.EncodeToken(idStart.End()); err != nil {
					return err
				}
			}

			if err := enc.EncodeToken(idsStart.End()); err != nil {
				return err
			}
		}

		// 编码 SharedEntityType 元素
		typeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityType"}}
		if err := enc.EncodeToken(typeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.GetSharedEntityAssociationsBySharedEntityIdsRequest.SharedEntityType))); err != nil {
			return err
		}
		if err := enc.EncodeToken(typeStart.End()); err != nil {
			return err
		}

		// 编码 SharedEntityScope 元素
		scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
		if err := enc.EncodeToken(scopeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.GetSharedEntityAssociationsBySharedEntityIdsRequest.SharedEntityScope))); err != nil {
			return err
		}
		if err := enc.EncodeToken(scopeStart.End()); err != nil {
			return err
		}

		// 结束 GetSharedEntityAssociationsBySharedEntityIdsRequest
		if err := enc.EncodeToken(reqStart.End()); err != nil {
			return err
		}
	}

	// 编码 AddListItemsToSharedListRequest
	if b.AddListItemsToSharedListRequest != nil {
		reqStart := xml.StartElement{Name: xml.Name{Local: "AddListItemsToSharedListRequest"}}

		// 设置命名空间属性
		if b.AddListItemsToSharedListRequest.Namespace != "" {
			reqStart.Attr = append(reqStart.Attr, xml.Attr{
				Name:  xml.Name{Local: "xmlns"},
				Value: b.AddListItemsToSharedListRequest.Namespace,
			})
		}

		// 编码带命名空间的 AddListItemsToSharedListRequest 开始标签
		if err := enc.EncodeToken(reqStart); err != nil {
			return err
		}

		// 编码 ListItems 元素
		if len(b.AddListItemsToSharedListRequest.ListItems) > 0 {
			listItemsStart := xml.StartElement{Name: xml.Name{Local: "ListItems"}}
			if err := enc.EncodeToken(listItemsStart); err != nil {
				return err
			}

			for _, item := range b.AddListItemsToSharedListRequest.ListItems {
				// 使用共享列表项自己的MarshalXML方法来处理序列化
				itemStart := xml.StartElement{Name: xml.Name{Local: "SharedListItem"}}
				if err := enc.EncodeElement(item, itemStart); err != nil {
					return err
				}
			}

			// 结束 ListItems 元素
			if err := enc.EncodeToken(listItemsStart.End()); err != nil {
				return err
			}
		}

		// 编码 SharedList 元素
		slStart := xml.StartElement{Name: xml.Name{Local: "SharedList"}}
		if err := enc.EncodeElement(b.AddListItemsToSharedListRequest.SharedList, slStart); err != nil {
			return err
		}

		// 编码 SharedEntityScope 元素
		scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
		if err := enc.EncodeToken(scopeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.AddListItemsToSharedListRequest.SharedEntityScope))); err != nil {
			return err
		}
		if err := enc.EncodeToken(scopeStart.End()); err != nil {
			return err
		}

		// 结束 AddListItemsToSharedListRequest
		if err := enc.EncodeToken(reqStart.End()); err != nil {
			return err
		}
	}

	// 编码 DeleteListItemsFromSharedListRequest
	if b.DeleteListItemsFromSharedListRequest != nil {
		reqStart := xml.StartElement{Name: xml.Name{Local: "DeleteListItemsFromSharedListRequest"}}

		// 设置命名空间属性
		if b.DeleteListItemsFromSharedListRequest.Namespace != "" {
			reqStart.Attr = append(reqStart.Attr, xml.Attr{
				Name:  xml.Name{Local: "xmlns"},
				Value: b.DeleteListItemsFromSharedListRequest.Namespace,
			})
		}

		// 编码带命名空间的 DeleteListItemsFromSharedListRequest 开始标签
		if err := enc.EncodeToken(reqStart); err != nil {
			return err
		}

		// 编码 ListItemIds 元素
		if len(b.DeleteListItemsFromSharedListRequest.ListItemIds) > 0 {
			listItemIdsStart := xml.StartElement{
				Name: xml.Name{Local: "ListItemIds"},
				Attr: []xml.Attr{
					{
						Name:  xml.Name{Local: "xmlns:a1"},
						Value: "http://schemas.microsoft.com/2003/10/Serialization/Arrays",
					},
				},
			}
			if err := enc.EncodeToken(listItemIdsStart); err != nil {
				return err
			}

			for _, id := range b.DeleteListItemsFromSharedListRequest.ListItemIds {
				idStart := xml.StartElement{Name: xml.Name{Local: "a1:long"}}
				if err := enc.EncodeToken(idStart); err != nil {
					return err
				}
				if err := enc.EncodeToken(xml.CharData(fmt.Sprintf("%d", id))); err != nil {
					return err
				}
				if err := enc.EncodeToken(idStart.End()); err != nil {
					return err
				}
			}

			// 结束 ListItemIds 元素
			if err := enc.EncodeToken(listItemIdsStart.End()); err != nil {
				return err
			}
		}

		// 编码 SharedList 元素
		slStart := xml.StartElement{Name: xml.Name{Local: "SharedList"}}
		if err := enc.EncodeElement(b.DeleteListItemsFromSharedListRequest.SharedList, slStart); err != nil {
			return err
		}

		// 编码 SharedEntityScope 元素
		scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
		if err := enc.EncodeToken(scopeStart); err != nil {
			return err
		}
		if err := enc.EncodeToken(xml.CharData(string(b.DeleteListItemsFromSharedListRequest.SharedEntityScope))); err != nil {
			return err
		}
		if err := enc.EncodeToken(scopeStart.End()); err != nil {
			return err
		}

		// 结束 DeleteListItemsFromSharedListRequest
		if err := enc.EncodeToken(reqStart.End()); err != nil {
			return err
		}
	}

	// 结束 Body
	if err := enc.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}
