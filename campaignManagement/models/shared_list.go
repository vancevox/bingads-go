package models

import (
	"encoding/xml"
	"fmt"
)

// SharedEntity 表示共享实体的基础类型
type SharedEntity struct {
	AssociationCount        int                          `xml:"AssociationCount,omitempty"`
	ForwardCompatibilityMap []KeyValuePairOfstringstring `xml:"ForwardCompatibilityMap>KeyValuePairOfstringstring,omitempty"`
	Id                      int64                        `xml:"Id,omitempty"`
	Name                    string                       `xml:"Name,omitempty"`
	Type                    string                       `xml:"Type,omitempty"`
	ItemCount               int                          `xml:"ItemCount,omitempty"`
}

// SharedList 表示共享列表
type SharedList struct {
	SharedEntity
	ItemCount int    `xml:"ItemCount,omitempty"`
	ItemType  string `xml:"i:type,attr,omitempty"` // 用于指定具体类型
}

// NegativeKeywordList 表示负面关键词列表
type NegativeKeywordList struct {
	SharedList
}

// PlacementExclusionList 表示投放排除列表
type PlacementExclusionList struct {
	SharedList
}

// SharedListItem 表示共享列表项
type SharedListItemBase struct {
	XMLName                 xml.Name `xml:"SharedListItem"`
	Type                    string   `xml:"i:type,attr,omitempty"`
	ForwardCompatibilityMap string   `xml:"ForwardCompatibilityMap"`
	ListItemType            string   `xml:"ListItemType"`
}

// NegativeKeyword 表示负面关键词
type NegativeKeyword struct {
	SharedListItemBase
	Id        int64  `xml:"Id"`
	MatchType string `xml:"MatchType"`
	Text      string `xml:"Text"`
}

// NegativeSite 表示负面站点
type NegativeSite struct {
	SharedListItemBase
	Id  int64  `xml:"Id"`
	Url string `xml:"Url"`
}

// GetListItemsBySharedListRequest 请求结构体
type GetListItemsBySharedListRequest struct {
	XMLName           xml.Name    `xml:"GetListItemsBySharedListRequest"`
	Namespace         string      `xml:"xmlns,attr"`
	SharedList        SharedList  `xml:"SharedList"`
	SharedEntityScope EntityScope `xml:"SharedEntityScope"`
}

// SharedListItem 共享列表项基础结构
type SharedListItem struct {
	Type                    SharedListItemType           `xml:"Type"`
	ForwardCompatibilityMap []KeyValuePairOfstringstring `xml:"ForwardCompatibilityMap>KeyValuePairOfstringstring,omitempty"`
	ItemType                string                       `xml:"i:type,attr,omitempty"`

	// NegativeKeyword 类型的字段
	ID        int64  `xml:"Id,omitempty"`
	MatchType string `xml:"MatchType,omitempty"`
	Text      string `xml:"Text,omitempty"`

	// NegativeSite 或 Site 类型的字段
	Url string `xml:"Url,omitempty"`

	// BrandItem 类型的字段
	BrandId int64 `xml:"BrandId,omitempty"`
}

// GetListItemsBySharedListResponse 响应结构体
type GetListItemsBySharedListResponse struct {
	XMLName   xml.Name         `xml:"GetListItemsBySharedListResponse"`
	Namespace string           `xml:"xmlns,attr"`
	ListItems []SharedListItem `xml:"ListItems>SharedListItem,omitempty"`
}

// MarshalXML 自定义 SharedList 的 XML 序列化
func (s SharedList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// 添加 i:type 属性（只有当确实需要时）
	if s.ItemType != "" {
		// 正确引用已经在根元素中定义的命名空间
		typeAttr := xml.Attr{
			Name: xml.Name{
				Space: "",       // 避免生成 xmlns:i 属性
				Local: "i:type", // 直接使用限定名称
			},
			Value: s.ItemType,
		}
		start.Attr = append(start.Attr, typeAttr)
	}

	// 开始 SharedList 元素
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 编码 ForwardCompatibilityMap 元素（如果有值）
	if len(s.ForwardCompatibilityMap) > 0 {
		fcMapStart := xml.StartElement{Name: xml.Name{Local: "ForwardCompatibilityMap"}}
		if err := e.EncodeToken(fcMapStart); err != nil {
			return err
		}

		for _, pair := range s.ForwardCompatibilityMap {
			// 编码 KeyValuePairOfstringstring 元素
			pairStart := xml.StartElement{Name: xml.Name{Local: "KeyValuePairOfstringstring"}}
			if err := e.EncodeToken(pairStart); err != nil {
				return err
			}

			// 编码 Key 元素
			keyStart := xml.StartElement{Name: xml.Name{Local: "key"}}
			if err := e.EncodeToken(keyStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.CharData(pair.Key)); err != nil {
				return err
			}
			if err := e.EncodeToken(keyStart.End()); err != nil {
				return err
			}

			// 编码 Value 元素
			valueStart := xml.StartElement{Name: xml.Name{Local: "value"}}
			if err := e.EncodeToken(valueStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.CharData(pair.Value)); err != nil {
				return err
			}
			if err := e.EncodeToken(valueStart.End()); err != nil {
				return err
			}

			// 结束 KeyValuePairOfstringstring 元素
			if err := e.EncodeToken(pairStart.End()); err != nil {
				return err
			}
		}

		// 结束 ForwardCompatibilityMap 元素
		if err := e.EncodeToken(fcMapStart.End()); err != nil {
			return err
		}
	}

	// 编码 Id 元素
	idStart := xml.StartElement{Name: xml.Name{Local: "Id"}}
	if err := e.EncodeToken(idStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", s.Id))); err != nil {
		return err
	}
	if err := e.EncodeToken(idStart.End()); err != nil {
		return err
	}

	// 编码 Name 元素 (如果存在)
	if s.Name != "" {
		nameStart := xml.StartElement{Name: xml.Name{Local: "Name"}}
		if err := e.EncodeToken(nameStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(s.Name)); err != nil {
			return err
		}
		if err := e.EncodeToken(nameStart.End()); err != nil {
			return err
		}
	}

	// 结束 SharedList 元素
	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// MarshalXML 自定义 SharedListItem 的 XML 序列化
func (item SharedListItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// 如果ItemType未设置，则使用Type字段的值
	if item.ItemType == "" && item.Type != "" {
		item.ItemType = string(item.Type)
	}

	// 添加 i:type 属性
	if item.ItemType != "" {
		typeAttr := xml.Attr{
			Name:  xml.Name{Local: "i:type"},
			Value: item.ItemType,
		}
		start.Attr = append(start.Attr, typeAttr)
	}

	// 开始 SharedListItem 元素
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 编码 ForwardCompatibilityMap 元素（如果有值）
	if len(item.ForwardCompatibilityMap) > 0 {
		fcMapStart := xml.StartElement{Name: xml.Name{Local: "ForwardCompatibilityMap"}}
		if err := e.EncodeToken(fcMapStart); err != nil {
			return err
		}

		for _, pair := range item.ForwardCompatibilityMap {
			// 编码 KeyValuePairOfstringstring 元素
			pairStart := xml.StartElement{Name: xml.Name{Local: "KeyValuePairOfstringstring"}}
			if err := e.EncodeToken(pairStart); err != nil {
				return err
			}

			// 编码 Key 元素
			keyStart := xml.StartElement{Name: xml.Name{Local: "key"}}
			if err := e.EncodeToken(keyStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.CharData(pair.Key)); err != nil {
				return err
			}
			if err := e.EncodeToken(keyStart.End()); err != nil {
				return err
			}

			// 编码 Value 元素
			valueStart := xml.StartElement{Name: xml.Name{Local: "value"}}
			if err := e.EncodeToken(valueStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.CharData(pair.Value)); err != nil {
				return err
			}
			if err := e.EncodeToken(valueStart.End()); err != nil {
				return err
			}

			// 结束 KeyValuePairOfstringstring 元素
			if err := e.EncodeToken(pairStart.End()); err != nil {
				return err
			}
		}

		// 结束 ForwardCompatibilityMap 元素
		if err := e.EncodeToken(fcMapStart.End()); err != nil {
			return err
		}
	}

	// 编码 Type 元素
	if item.Type != "" {
		typeStart := xml.StartElement{Name: xml.Name{Local: "Type"}}
		if err := e.EncodeToken(typeStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(string(item.Type))); err != nil {
			return err
		}
		if err := e.EncodeToken(typeStart.End()); err != nil {
			return err
		}
	}

	// 根据不同类型编码不同字段
	// 对于 NegativeKeyword 类型
	if item.ID != 0 {
		idStart := xml.StartElement{Name: xml.Name{Local: "Id"}}
		if err := e.EncodeToken(idStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", item.ID))); err != nil {
			return err
		}
		if err := e.EncodeToken(idStart.End()); err != nil {
			return err
		}
	}

	if item.MatchType != "" {
		matchTypeStart := xml.StartElement{Name: xml.Name{Local: "MatchType"}}
		if err := e.EncodeToken(matchTypeStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(item.MatchType)); err != nil {
			return err
		}
		if err := e.EncodeToken(matchTypeStart.End()); err != nil {
			return err
		}
	}

	if item.Text != "" {
		textStart := xml.StartElement{Name: xml.Name{Local: "Text"}}
		if err := e.EncodeToken(textStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(item.Text)); err != nil {
			return err
		}
		if err := e.EncodeToken(textStart.End()); err != nil {
			return err
		}
	}

	// 对于 NegativeSite 或 Site 类型
	if item.Url != "" {
		urlStart := xml.StartElement{Name: xml.Name{Local: "Url"}}
		if err := e.EncodeToken(urlStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(item.Url)); err != nil {
			return err
		}
		if err := e.EncodeToken(urlStart.End()); err != nil {
			return err
		}
	}

	// 对于 BrandItem 类型
	if item.BrandId != 0 {
		brandIdStart := xml.StartElement{Name: xml.Name{Local: "BrandId"}}
		if err := e.EncodeToken(brandIdStart); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", item.BrandId))); err != nil {
			return err
		}
		if err := e.EncodeToken(brandIdStart.End()); err != nil {
			return err
		}
	}

	// 结束 SharedListItem 元素
	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// MarshalXML 自定义 GetListItemsBySharedListRequest 的 XML 序列化
func (req GetListItemsBySharedListRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 添加 xmlns 属性
	namespaceStart := xml.StartElement{
		Name: start.Name,
		Attr: []xml.Attr{
			{Name: xml.Name{Local: "xmlns"}, Value: req.Namespace},
		},
	}
	if err := e.EncodeElement(struct{}{}, namespaceStart); err != nil {
		return err
	}

	// 编码 SharedList 元素
	sharedListStart := xml.StartElement{Name: xml.Name{Local: "SharedList"}}
	if err := e.EncodeElement(req.SharedList, sharedListStart); err != nil {
		return err
	}

	// 编码 SharedEntityScope 元素，不带任何属性
	scopeStart := xml.StartElement{
		Name: xml.Name{Local: "SharedEntityScope"},
	}
	if err := e.EncodeToken(scopeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityScope))); err != nil {
		return err
	}
	if err := e.EncodeToken(scopeStart.End()); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// SharedEntityType 表示共享实体类型
type SharedEntityType string

const (
	SharedEntityTypeNegativeKeywordList           SharedEntityType = "NegativeKeywordList"
	SharedEntityTypePlacementExclusionList        SharedEntityType = "PlacementExclusionList"
	SharedEntityTypeAccountNegativeKeywordList    SharedEntityType = "AccountNegativeKeywordList"
	SharedEntityTypeAccountPlacementExclusionList SharedEntityType = "AccountPlacementExclusionList"
	SharedEntityTypeAccountPlacementInclusionList SharedEntityType = "AccountPlacementInclusionList"
	SharedEntityTypeBrandList                     SharedEntityType = "BrandList"
)

// GetSharedEntitiesRequest 请求结构体
type GetSharedEntitiesRequest struct {
	XMLName           xml.Name         `xml:"GetSharedEntitiesRequest"`
	Namespace         string           `xml:"xmlns,attr"`
	SharedEntityType  SharedEntityType `xml:"SharedEntityType"`
	SharedEntityScope EntityScope      `xml:"SharedEntityScope"`
}

// GetSharedEntitiesResponse 响应结构体
type GetSharedEntitiesResponse struct {
	XMLName        xml.Name       `xml:"GetSharedEntitiesResponse"`
	Namespace      string         `xml:"xmlns,attr"`
	SharedEntities []SharedEntity `xml:"SharedEntities>SharedEntity,omitempty"`
}

// MarshalXML 自定义 GetSharedEntitiesRequest 的 XML 序列化
func (req GetSharedEntitiesRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 添加 xmlns 属性
	if req.Namespace != "" {
		namespaceAttr := xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: req.Namespace,
		}
		start.Attr = append(start.Attr, namespaceAttr)
		if err := e.EncodeToken(start); err != nil {
			return err
		}
	}

	// 编码 SharedEntityType 元素
	typeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityType"}}
	if err := e.EncodeToken(typeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityType))); err != nil {
		return err
	}
	if err := e.EncodeToken(typeStart.End()); err != nil {
		return err
	}

	// 编码 SharedEntityScope 元素
	scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
	if err := e.EncodeToken(scopeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityScope))); err != nil {
		return err
	}
	if err := e.EncodeToken(scopeStart.End()); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// EntityType 表示实体类型
type EntityType string

const (
	EntityTypeCampaign EntityType = "Campaign"
	EntityTypeAccount  EntityType = "Account"
)

// SharedEntityAssociation 表示共享实体关联
type SharedEntityAssociation struct {
	EntityId               int64            `xml:"EntityId"`
	EntityType             EntityType       `xml:"EntityType"`
	SharedEntityCustomerId string           `xml:"SharedEntityCustomerId,omitempty"`
	SharedEntityId         int64            `xml:"SharedEntityId"`
	SharedEntityType       SharedEntityType `xml:"SharedEntityType"`
}

// BatchError 表示批处理错误
type BatchError struct {
	Code                    int                           `xml:"Code"`
	Details                 *string                       `xml:"Details,omitempty"`
	ErrorCode               string                        `xml:"ErrorCode,omitempty"`
	FieldPath               *string                       `xml:"FieldPath,omitempty"`
	ForwardCompatibilityMap *[]KeyValuePairOfstringstring `xml:"ForwardCompatibilityMap>KeyValuePairOfstringstring,omitempty"`
	Index                   int                           `xml:"Index"`
	Message                 string                        `xml:"Message,omitempty"`
	Type                    string                        `xml:"Type,omitempty"`
}

// GetSharedEntityAssociationsBySharedEntityIdsRequest 请求结构体
type GetSharedEntityAssociationsBySharedEntityIdsRequest struct {
	XMLName           xml.Name         `xml:"GetSharedEntityAssociationsBySharedEntityIdsRequest"`
	Namespace         string           `xml:"xmlns,attr"`
	EntityType        EntityType       `xml:"EntityType"`
	SharedEntityIds   []int64          `xml:"SharedEntityIds>a1:long"`
	SharedEntityType  SharedEntityType `xml:"SharedEntityType"`
	SharedEntityScope EntityScope      `xml:"SharedEntityScope"`
}

// GetSharedEntityAssociationsBySharedEntityIdsResponse 响应结构体
type GetSharedEntityAssociationsBySharedEntityIdsResponse struct {
	XMLName       xml.Name                  `xml:"GetSharedEntityAssociationsBySharedEntityIdsResponse"`
	Namespace     string                    `xml:"xmlns,attr"`
	Associations  []SharedEntityAssociation `xml:"Associations>SharedEntityAssociation,omitempty"`
	PartialErrors []BatchError              `xml:"PartialErrors>BatchError,omitempty"`
}

// MarshalXML 自定义 GetSharedEntityAssociationsBySharedEntityIdsRequest 的 XML 序列化
func (req GetSharedEntityAssociationsBySharedEntityIdsRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 添加命名空间
	if req.Namespace != "" {
		namespaceAttr := xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: req.Namespace,
		}
		start.Attr = append(start.Attr, namespaceAttr)
		if err := e.EncodeToken(start); err != nil {
			return err
		}
	}

	// 编码 EntityType 元素
	entityTypeStart := xml.StartElement{Name: xml.Name{Local: "EntityType"}}
	if err := e.EncodeToken(entityTypeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.EntityType))); err != nil {
		return err
	}
	if err := e.EncodeToken(entityTypeStart.End()); err != nil {
		return err
	}

	// 编码 SharedEntityIds 元素
	if len(req.SharedEntityIds) > 0 {
		idsStart := xml.StartElement{
			Name: xml.Name{Local: "SharedEntityIds"},
			Attr: []xml.Attr{
				{
					Name:  xml.Name{Local: "xmlns:a1"},
					Value: "http://schemas.microsoft.com/2003/10/Serialization/Arrays",
				},
			},
		}
		if err := e.EncodeToken(idsStart); err != nil {
			return err
		}

		for _, id := range req.SharedEntityIds {
			idStart := xml.StartElement{Name: xml.Name{Local: "a1:long"}}
			if err := e.EncodeToken(idStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", id))); err != nil {
				return err
			}
			if err := e.EncodeToken(idStart.End()); err != nil {
				return err
			}
		}

		if err := e.EncodeToken(idsStart.End()); err != nil {
			return err
		}
	}

	// 编码 SharedEntityType 元素
	sharedEntityTypeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityType"}}
	if err := e.EncodeToken(sharedEntityTypeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityType))); err != nil {
		return err
	}
	if err := e.EncodeToken(sharedEntityTypeStart.End()); err != nil {
		return err
	}

	// 编码 SharedEntityScope 元素
	scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
	if err := e.EncodeToken(scopeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityScope))); err != nil {
		return err
	}
	if err := e.EncodeToken(scopeStart.End()); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// AddListItemsToSharedListRequest 请求结构体
type AddListItemsToSharedListRequest struct {
	XMLName           xml.Name         `xml:"AddListItemsToSharedListRequest"`
	Namespace         string           `xml:"xmlns,attr"`
	ListItems         []SharedListItem `xml:"ListItems>SharedListItem,omitempty"`
	SharedList        SharedList       `xml:"SharedList"`
	SharedEntityScope EntityScope      `xml:"SharedEntityScope"`
}

// AddListItemsToSharedListResponse 响应结构体
type AddListItemsToSharedListResponse struct {
	XMLName       xml.Name     `xml:"AddListItemsToSharedListResponse"`
	Namespace     string       `xml:"xmlns,attr"`
	ListItemIds   []int64      `xml:"ListItemIds>a:long,omitempty"`
	PartialErrors []BatchError `xml:"PartialErrors>BatchError,omitempty"`
}

// MarshalXML 自定义 AddListItemsToSharedListRequest 的 XML 序列化
func (req AddListItemsToSharedListRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 添加 xmlns 属性
	if req.Namespace != "" {
		namespaceAttr := xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: req.Namespace,
		}
		start.Attr = append(start.Attr, namespaceAttr)
		if err := e.EncodeToken(start); err != nil {
			return err
		}
	}

	// 编码 ListItems 元素
	if len(req.ListItems) > 0 {
		listItemsStart := xml.StartElement{Name: xml.Name{Local: "ListItems"}}
		if err := e.EncodeToken(listItemsStart); err != nil {
			return err
		}

		for _, item := range req.ListItems {
			itemStart := xml.StartElement{
				Name: xml.Name{Local: "SharedListItem"},
				Attr: []xml.Attr{},
			}

			// 添加类型属性
			if item.ItemType != "" {
				typeAttr := xml.Attr{
					Name:  xml.Name{Local: "i:type"},
					Value: item.ItemType,
				}
				itemStart.Attr = append(itemStart.Attr, typeAttr)
			}

			if err := e.EncodeToken(itemStart); err != nil {
				return err
			}

			// 编码 ForwardCompatibilityMap 元素（如果有）
			if len(item.ForwardCompatibilityMap) > 0 {
				fcMapStart := xml.StartElement{Name: xml.Name{Local: "ForwardCompatibilityMap"}}
				if err := e.EncodeToken(fcMapStart); err != nil {
					return err
				}

				for _, pair := range item.ForwardCompatibilityMap {
					pairStart := xml.StartElement{Name: xml.Name{Local: "KeyValuePairOfstringstring"}}
					if err := e.EncodeToken(pairStart); err != nil {
						return err
					}

					// 编码 Key 元素
					keyStart := xml.StartElement{Name: xml.Name{Local: "key"}}
					if err := e.EncodeToken(keyStart); err != nil {
						return err
					}
					if err := e.EncodeToken(xml.CharData(pair.Key)); err != nil {
						return err
					}
					if err := e.EncodeToken(keyStart.End()); err != nil {
						return err
					}

					// 编码 Value 元素
					valueStart := xml.StartElement{Name: xml.Name{Local: "value"}}
					if err := e.EncodeToken(valueStart); err != nil {
						return err
					}
					if err := e.EncodeToken(xml.CharData(pair.Value)); err != nil {
						return err
					}
					if err := e.EncodeToken(valueStart.End()); err != nil {
						return err
					}

					// 结束 KeyValuePairOfstringstring 元素
					if err := e.EncodeToken(pairStart.End()); err != nil {
						return err
					}
				}

				// 结束 ForwardCompatibilityMap 元素
				if err := e.EncodeToken(fcMapStart.End()); err != nil {
					return err
				}
			}

			// 编码 Type 元素
			if item.Type != "" {
				typeStart := xml.StartElement{Name: xml.Name{Local: "Type"}}
				if err := e.EncodeToken(typeStart); err != nil {
					return err
				}
				if err := e.EncodeToken(xml.CharData(item.Type)); err != nil {
					return err
				}
				if err := e.EncodeToken(typeStart.End()); err != nil {
					return err
				}
			}

			// 根据不同类型编码不同字段
			// 对于 NegativeKeyword 类型
			if item.ID != 0 {
				idStart := xml.StartElement{Name: xml.Name{Local: "Id"}}
				if err := e.EncodeToken(idStart); err != nil {
					return err
				}
				if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", item.ID))); err != nil {
					return err
				}
				if err := e.EncodeToken(idStart.End()); err != nil {
					return err
				}
			}

			if item.MatchType != "" {
				matchTypeStart := xml.StartElement{Name: xml.Name{Local: "MatchType"}}
				if err := e.EncodeToken(matchTypeStart); err != nil {
					return err
				}
				if err := e.EncodeToken(xml.CharData(item.MatchType)); err != nil {
					return err
				}
				if err := e.EncodeToken(matchTypeStart.End()); err != nil {
					return err
				}
			}

			if item.Text != "" {
				textStart := xml.StartElement{Name: xml.Name{Local: "Text"}}
				if err := e.EncodeToken(textStart); err != nil {
					return err
				}
				if err := e.EncodeToken(xml.CharData(item.Text)); err != nil {
					return err
				}
				if err := e.EncodeToken(textStart.End()); err != nil {
					return err
				}
			}

			// 对于 NegativeSite 或 Site 类型
			if item.Url != "" {
				urlStart := xml.StartElement{Name: xml.Name{Local: "Url"}}
				if err := e.EncodeToken(urlStart); err != nil {
					return err
				}
				if err := e.EncodeToken(xml.CharData(item.Url)); err != nil {
					return err
				}
				if err := e.EncodeToken(urlStart.End()); err != nil {
					return err
				}
			}

			// 对于 BrandItem 类型
			if item.BrandId != 0 {
				brandIdStart := xml.StartElement{Name: xml.Name{Local: "BrandId"}}
				if err := e.EncodeToken(brandIdStart); err != nil {
					return err
				}
				if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", item.BrandId))); err != nil {
					return err
				}
				if err := e.EncodeToken(brandIdStart.End()); err != nil {
					return err
				}
			}

			// 结束 SharedListItem 元素
			if err := e.EncodeToken(itemStart.End()); err != nil {
				return err
			}
		}

		// 结束 ListItems 元素
		if err := e.EncodeToken(listItemsStart.End()); err != nil {
			return err
		}
	}

	// 编码 SharedList 元素
	sharedListStart := xml.StartElement{Name: xml.Name{Local: "SharedList"}}
	if err := e.EncodeElement(req.SharedList, sharedListStart); err != nil {
		return err
	}

	// 编码 SharedEntityScope 元素
	scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
	if err := e.EncodeToken(scopeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityScope))); err != nil {
		return err
	}
	if err := e.EncodeToken(scopeStart.End()); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}

// DeleteListItemsFromSharedListRequest 请求结构体
type DeleteListItemsFromSharedListRequest struct {
	XMLName           xml.Name    `xml:"DeleteListItemsFromSharedListRequest"`
	Namespace         string      `xml:"xmlns,attr"`
	ListItemIds       []int64     `xml:"ListItemIds>a1:long,omitempty"`
	SharedList        SharedList  `xml:"SharedList"`
	SharedEntityScope EntityScope `xml:"SharedEntityScope"`
}

// DeleteListItemsFromSharedListResponse 响应结构体
type DeleteListItemsFromSharedListResponse struct {
	XMLName       xml.Name     `xml:"DeleteListItemsFromSharedListResponse"`
	Namespace     string       `xml:"xmlns,attr"`
	PartialErrors []BatchError `xml:"PartialErrors>BatchError,omitempty"`
}

// MarshalXML 自定义 DeleteListItemsFromSharedListRequest 的 XML 序列化
func (req DeleteListItemsFromSharedListRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(start); err != nil {
		return err
	}

	// 添加 xmlns 属性
	if req.Namespace != "" {
		namespaceAttr := xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: req.Namespace,
		}
		start.Attr = append(start.Attr, namespaceAttr)
		if err := e.EncodeToken(start); err != nil {
			return err
		}
	}

	// 编码 ListItemIds 元素
	if len(req.ListItemIds) > 0 {
		listItemIdsStart := xml.StartElement{
			Name: xml.Name{Local: "ListItemIds"},
			Attr: []xml.Attr{
				{
					Name:  xml.Name{Local: "xmlns:a1"},
					Value: "http://schemas.microsoft.com/2003/10/Serialization/Arrays",
				},
			},
		}
		if err := e.EncodeToken(listItemIdsStart); err != nil {
			return err
		}

		for _, id := range req.ListItemIds {
			idStart := xml.StartElement{Name: xml.Name{Local: "a1:long"}}
			if err := e.EncodeToken(idStart); err != nil {
				return err
			}
			if err := e.EncodeToken(xml.CharData(fmt.Sprintf("%d", id))); err != nil {
				return err
			}
			if err := e.EncodeToken(idStart.End()); err != nil {
				return err
			}
		}

		// 结束 ListItemIds 元素
		if err := e.EncodeToken(listItemIdsStart.End()); err != nil {
			return err
		}
	}

	// 编码 SharedList 元素
	sharedListStart := xml.StartElement{Name: xml.Name{Local: "SharedList"}}
	if err := e.EncodeElement(req.SharedList, sharedListStart); err != nil {
		return err
	}

	// 编码 SharedEntityScope 元素
	scopeStart := xml.StartElement{Name: xml.Name{Local: "SharedEntityScope"}}
	if err := e.EncodeToken(scopeStart); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.CharData(string(req.SharedEntityScope))); err != nil {
		return err
	}
	if err := e.EncodeToken(scopeStart.End()); err != nil {
		return err
	}

	if err := e.EncodeToken(start.End()); err != nil {
		return err
	}

	return nil
}
