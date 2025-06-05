package unit

import (
	"testing"

	"github.com/vancevox/bingads-go/campaignManagement/models"
	"github.com/vancevox/bingads-go/campaignManagement/service"
	"github.com/vancevox/bingads-go/config"
)

func TestGetListItemsBySharedList(t *testing.T) {
	client := service.NewClient(&config.Config{
		Auth: &config.AuthConfig{
			DeveloperToken:      "DeveloperToken",
			AuthenticationToken: "AuthenticationToken",
			CustomerID:          "CustomerID",
			CustomerAccountID:   "CustomerAccountID",
		},
		API: config.DefaultConfig(),
	})
	client.Config.API.Debug = true
	sharedListService := client.SharedListService()
	list, err := sharedListService.GetListItemsBySharedList(models.PlacementExclusionList{
		SharedList: models.SharedList{
			SharedEntity: models.SharedEntity{
				Id: 123,
			},
		},
	}, models.EntityScopeCustomer)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("成功获取共享列表的项目: %+v", list)
	}
}

func TestGetSharedEntities(t *testing.T) {
	client := service.NewClient(&config.Config{
		Auth: &config.AuthConfig{
			DeveloperToken:      "DeveloperToken",
			AuthenticationToken: "AuthenticationToken",
			CustomerID:          "CustomerID",
			CustomerAccountID:   "CustomerAccountID",
		},
		API: config.DefaultConfig(),
	})
	client.Config.API.Debug = true
	sharedListService := client.SharedListService()
	list, err := sharedListService.GetSharedEntities(models.SharedEntityTypePlacementExclusionList, models.EntityScopeCustomer)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("成功获取共享实体列表: %+v", list)
	}
}

func TestGetSharedEntityAssociationsBySharedEntityIds(t *testing.T) {
	client := service.NewClient(&config.Config{
		Auth: &config.AuthConfig{
			DeveloperToken:      "DeveloperToken",
			AuthenticationToken: "AuthenticationToken",
			CustomerID:          "CustomerID",
			CustomerAccountID:   "CustomerAccountID",
		},
		API: config.DefaultConfig(),
	})
	client.Config.API.Debug = true
	sharedListService := client.SharedListService()
	list, _, err := sharedListService.GetSharedEntityAssociationsBySharedEntityIds(models.EntityTypeAccount, []int64{123}, models.SharedEntityTypePlacementExclusionList, models.EntityScopeCustomer)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("成功获取共享实体关联: %+v", list)
	}
}

func TestAddListItemsToSharedList(t *testing.T) {
	client := service.NewClient(&config.Config{
		Auth: &config.AuthConfig{
			DeveloperToken:      "DeveloperToken",
			AuthenticationToken: "AuthenticationToken",
			CustomerID:          "CustomerID",
			CustomerAccountID:   "CustomerAccountID",
		},
		API: config.DefaultConfig(),
	})
	client.Config.API.Debug = true
	sharedListService := client.SharedListService()
	list, partialErrors, err := sharedListService.AddListItemsToSharedList(models.PlacementExclusionList{
		SharedList: models.SharedList{
			SharedEntity: models.SharedEntity{
				Id: 123,
			},
		},
	}, []models.SharedListItem{
		{Url: "example1.com", Type: models.SharedListItemTypeNegativeSite},
		{Url: "example2.com", Type: models.SharedListItemTypeNegativeSite},
	}, models.EntityScopeCustomer)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("成功获取列表项ID: %+v", list)

		if len(partialErrors) > 0 {
			t.Logf("部分错误信息:")
			for i, e := range partialErrors {
				t.Logf("  错误 #%d: 代码=%d, 错误码=%s, 消息=%s, 索引=%d",
					i+1, e.Code, e.ErrorCode, e.Message, e.Index)
			}
		}
	}
}

func TestDeleteListItemsFromSharedList(t *testing.T) {
	client := service.NewClient(&config.Config{
		Auth: &config.AuthConfig{
			DeveloperToken:      "DeveloperToken",
			AuthenticationToken: "AuthenticationToken",
			CustomerID:          "CustomerID",
			CustomerAccountID:   "CustomerAccountID",
		},
		API: config.DefaultConfig(),
	})
	client.Config.API.Debug = true
	sharedListService := client.SharedListService()
	_, err := sharedListService.DeleteListItemsFromSharedList(models.PlacementExclusionList{
		SharedList: models.SharedList{
			SharedEntity: models.SharedEntity{
				Id: 123456789,
			},
		},
	}, []int64{123456, 123457}, models.EntityScopeCustomer)

	if err != nil {
		t.Error(err)
	} else {
		t.Logf("成功删除共享列表项")
	}
}
