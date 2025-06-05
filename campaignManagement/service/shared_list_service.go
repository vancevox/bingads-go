package service

import (
	"fmt"

	"github.com/vancevox/bingads-go/config"

	"github.com/vancevox/bingads-go/campaignManagement/models"
)

// SharedListService 实现共享列表服务
type SharedListService struct {
	client *Client
}

// NewSharedListService 创建一个新的共享列表服务
func NewSharedListService(client *Client) *SharedListService {
	return &SharedListService{
		client: client,
	}
}

// GetListItemsBySharedList 获取共享列表中的项目
func (s *SharedListService) GetListItemsBySharedList(sharedList any, scope models.EntityScope) ([]models.SharedListItem, error) {
	var sharedListObj models.SharedList
	var sharedListType string

	// 根据不同的派生类型设置请求
	switch list := sharedList.(type) {
	case models.NegativeKeywordList:
		sharedListObj = list.SharedList
		sharedListType = "NegativeKeywordList"
	case models.PlacementExclusionList:
		sharedListObj = list.SharedList
		sharedListType = "PlacementExclusionList"
	case models.AccountNegativeKeywordList:
		sharedListObj = list.SharedList
		sharedListType = "AccountNegativeKeywordList"
	case models.BrandList:
		sharedListObj = list.SharedList
		sharedListType = "BrandList"
	case models.AccountPlacementExclusionList:
		sharedListObj = list.SharedList
		sharedListType = "AccountPlacementExclusionList"
	case models.AccountPlacementInclusionList:
		sharedListObj = list.SharedList
		sharedListType = "AccountPlacementInclusionList"
	case models.SharedList:
		sharedListObj = list
		sharedListType = list.ItemType
	default:
		return nil, fmt.Errorf("不支持的共享列表类型: %T", sharedList)
	}

	sharedListObj.ItemType = sharedListType
	// 创建请求
	request := models.GetListItemsBySharedListRequest{
		Namespace:         config.CampaignManagementNamespace,
		SharedList:        sharedListObj,
		SharedEntityScope: scope,
	}

	// 创建信封
	envelope := s.client.createEnvelope(models.SOAPActionGetListItemsBySharedList, "1")

	// 使用CampaignManagementBody包装请求
	envelope.Body = &models.CampaignManagementBody{
		GetListItemsBySharedListRequest: &request,
	}

	// 发送请求
	respBody, err := s.client.sendRequest(envelope, models.SOAPActionGetListItemsBySharedList)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var response models.CampaignManagementResponseEnvelope
	if err := s.client.processResponse(respBody, &response); err != nil {
		return nil, err
	}

	return response.Body.GetListItemsBySharedListResponse.ListItems, nil
}

// GetSharedEntities 获取共享实体
func (s *SharedListService) GetSharedEntities(entityType models.SharedEntityType, scope models.EntityScope) ([]models.SharedEntity, error) {
	// 创建请求
	request := models.GetSharedEntitiesRequest{
		Namespace:         config.CampaignManagementNamespace,
		SharedEntityType:  entityType,
		SharedEntityScope: scope,
	}

	// 创建信封
	envelope := s.client.createEnvelope(models.SOAPActionGetSharedEntities, "1")

	// 使用CampaignManagementBody包装请求
	envelope.Body = &models.CampaignManagementBody{
		GetSharedEntitiesRequest: &request,
	}

	// 发送请求
	respBody, err := s.client.sendRequest(envelope, models.SOAPActionGetSharedEntities)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var response models.CampaignManagementResponseEnvelope
	if err := s.client.processResponse(respBody, &response); err != nil {
		return nil, err
	}

	return response.Body.GetSharedEntitiesResponse.SharedEntities, nil
}

// GetSharedEntityAssociationsBySharedEntityIds 根据共享实体ID获取共享实体关联
func (s *SharedListService) GetSharedEntityAssociationsBySharedEntityIds(
	entityType models.EntityType,
	sharedEntityIds []int64,
	sharedEntityType models.SharedEntityType,
	scope models.EntityScope,
) ([]models.SharedEntityAssociation, []models.BatchError, error) {
	// 创建请求
	request := models.GetSharedEntityAssociationsBySharedEntityIdsRequest{
		Namespace:         config.CampaignManagementNamespace,
		EntityType:        entityType,
		SharedEntityIds:   sharedEntityIds,
		SharedEntityType:  sharedEntityType,
		SharedEntityScope: scope,
	}

	// 创建信封
	envelope := s.client.createEnvelope(models.SOAPActionGetSharedEntityAssociationsBySharedEntityIds, "1")

	// 使用CampaignManagementBody包装请求
	envelope.Body = &models.CampaignManagementBody{
		GetSharedEntityAssociationsBySharedEntityIdsRequest: &request,
	}

	// 发送请求
	respBody, err := s.client.sendRequest(envelope, models.SOAPActionGetSharedEntityAssociationsBySharedEntityIds)
	if err != nil {
		return nil, nil, err
	}

	// 解析响应
	var response models.CampaignManagementResponseEnvelope
	if err := s.client.processResponse(respBody, &response); err != nil {
		return nil, nil, err
	}

	resp := response.Body.GetSharedEntityAssociationsBySharedEntityIdsResponse
	return resp.Associations, resp.PartialErrors, nil
}

// AddListItemsToSharedList 向共享列表添加项目
func (s *SharedListService) AddListItemsToSharedList(sharedList any, listItems []models.SharedListItem, scope models.EntityScope) ([]int64, []models.BatchError, error) {
	var sharedListObj models.SharedList
	var sharedListType string

	// 根据不同的派生类型设置请求
	switch list := sharedList.(type) {
	case models.NegativeKeywordList:
		sharedListObj = list.SharedList
		sharedListType = "NegativeKeywordList"
	case models.PlacementExclusionList:
		sharedListObj = list.SharedList
		sharedListType = "PlacementExclusionList"
	case models.AccountNegativeKeywordList:
		sharedListObj = list.SharedList
		sharedListType = "AccountNegativeKeywordList"
	case models.BrandList:
		sharedListObj = list.SharedList
		sharedListType = "BrandList"
	case models.AccountPlacementExclusionList:
		sharedListObj = list.SharedList
		sharedListType = "AccountPlacementExclusionList"
	case models.AccountPlacementInclusionList:
		sharedListObj = list.SharedList
		sharedListType = "AccountPlacementInclusionList"
	case models.SharedList:
		sharedListObj = list
		sharedListType = list.ItemType
	default:
		return nil, nil, fmt.Errorf("不支持的共享列表类型: %T", sharedList)
	}

	sharedListObj.ItemType = sharedListType

	// 确保每个列表项都有正确的ItemType设置
	for i := range listItems {
		if listItems[i].ItemType == "" {
			listItems[i].ItemType = string(listItems[i].Type)
		}
	}

	// 创建请求
	request := models.AddListItemsToSharedListRequest{
		Namespace:         config.CampaignManagementNamespace,
		ListItems:         listItems,
		SharedList:        sharedListObj,
		SharedEntityScope: scope,
	}

	// 创建信封
	envelope := s.client.createEnvelope(models.SOAPActionAddListItemsToSharedList, "1")

	// 使用CampaignManagementBody包装请求
	envelope.Body = &models.CampaignManagementBody{
		AddListItemsToSharedListRequest: &request,
	}

	// 发送请求
	respBody, err := s.client.sendRequest(envelope, models.SOAPActionAddListItemsToSharedList)
	if err != nil {
		return nil, nil, err
	}

	// 解析响应
	var response models.CampaignManagementResponseEnvelope
	if err := s.client.processResponse(respBody, &response); err != nil {
		return nil, nil, err
	}

	resp := response.Body.AddListItemsToSharedListResponse
	return resp.ListItemIds, resp.PartialErrors, nil
}

// DeleteListItemsFromSharedList 从共享列表中删除项目
func (s *SharedListService) DeleteListItemsFromSharedList(sharedList any, listItemIds []int64, scope models.EntityScope) ([]models.BatchError, error) {
	var sharedListObj models.SharedList
	var sharedListType string

	// 根据不同的派生类型设置请求
	switch list := sharedList.(type) {
	case models.NegativeKeywordList:
		sharedListObj = list.SharedList
		sharedListType = "NegativeKeywordList"
	case models.PlacementExclusionList:
		sharedListObj = list.SharedList
		sharedListType = "PlacementExclusionList"
	case models.AccountNegativeKeywordList:
		sharedListObj = list.SharedList
		sharedListType = "AccountNegativeKeywordList"
	case models.BrandList:
		sharedListObj = list.SharedList
		sharedListType = "BrandList"
	case models.AccountPlacementExclusionList:
		sharedListObj = list.SharedList
		sharedListType = "AccountPlacementExclusionList"
	case models.AccountPlacementInclusionList:
		sharedListObj = list.SharedList
		sharedListType = "AccountPlacementInclusionList"
	case models.SharedList:
		sharedListObj = list
		sharedListType = list.ItemType
	default:
		return nil, fmt.Errorf("不支持的共享列表类型: %T", sharedList)
	}

	sharedListObj.ItemType = sharedListType

	// 创建请求
	request := models.DeleteListItemsFromSharedListRequest{
		Namespace:         config.CampaignManagementNamespace,
		ListItemIds:       listItemIds,
		SharedList:        sharedListObj,
		SharedEntityScope: scope,
	}

	// 创建信封
	envelope := s.client.createEnvelope(models.SOAPActionDeleteListItemsFromSharedList, "1")

	// 使用CampaignManagementBody包装请求
	envelope.Body = &models.CampaignManagementBody{
		DeleteListItemsFromSharedListRequest: &request,
	}

	// 发送请求
	respBody, err := s.client.sendRequest(envelope, models.SOAPActionDeleteListItemsFromSharedList)
	if err != nil {
		return nil, err
	}

	// 解析响应
	var response models.CampaignManagementResponseEnvelope
	if err := s.client.processResponse(respBody, &response); err != nil {
		return nil, err
	}

	resp := response.Body.DeleteListItemsFromSharedListResponse
	return resp.PartialErrors, nil
}
