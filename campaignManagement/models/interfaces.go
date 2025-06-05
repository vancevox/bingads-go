package models

// CampaignManagementAPI 定义Campaign Management API的操作
type CampaignManagementAPI interface {
	// SharedListService 返回共享列表服务
	SharedListService() SharedListService
}

// SharedListService 定义共享列表相关的操作
type SharedListService interface {
	// GetListItemsBySharedList 获取共享列表中的项目
	GetListItemsBySharedList(sharedList any, scope EntityScope) ([]SharedListItem, error)

	// GetSharedEntities 获取共享实体
	GetSharedEntities(entityType SharedEntityType, scope EntityScope) ([]SharedEntity, error)

	// GetSharedEntityAssociationsBySharedEntityIds 根据共享实体ID获取共享实体关联
	GetSharedEntityAssociationsBySharedEntityIds(entityType EntityType, sharedEntityIds []int64, sharedEntityType SharedEntityType, scope EntityScope) ([]SharedEntityAssociation, []BatchError, error)

	// AddListItemsToSharedList 向共享列表添加项目
	AddListItemsToSharedList(sharedList any, listItems []SharedListItem, scope EntityScope) ([]int64, []BatchError, error)

	// DeleteListItemsFromSharedList 从共享列表删除项目
	DeleteListItemsFromSharedList(sharedList any, listItemIds []int64, scope EntityScope) ([]BatchError, error)
}

// CampaignService 定义广告系列相关的操作
type CampaignService interface {
	// 此处添加广告系列相关的方法
}

// AdGroupService 定义广告组相关的操作
type AdGroupService interface {
	// 此处添加广告组相关的方法
}

// AdService 定义广告相关的操作
type AdService interface {
	// 此处添加广告相关的方法
}

// KeywordService 定义关键词相关的操作
type KeywordService interface {
	// 此处添加关键词相关的方法
}

// TargetingService 定义定位相关的操作
type TargetingService interface {
	// 此处添加定位相关的方法
}
