package models

type SOAPAction string

const (
	SOAPActionGetListItemsBySharedList                     SOAPAction = "GetListItemsBySharedList"
	SOAPActionGetSharedEntities                            SOAPAction = "GetSharedEntities"
	SOAPActionGetSharedEntityAssociationsBySharedEntityIds SOAPAction = "GetSharedEntityAssociationsBySharedEntityIds"
	SOAPActionAddListItemsToSharedList                     SOAPAction = "AddListItemsToSharedList"
	SOAPActionDeleteListItemsFromSharedList                SOAPAction = "DeleteListItemsFromSharedList"
)

type EntityScope string

const (
	EntityScopeAccount EntityScope = "Account"

	EntityScopeCustomer EntityScope = "Customer"
)

type SharedListItemType string

const (
	SharedListItemTypeNegativeSite    SharedListItemType = "NegativeSite"
	SharedListItemTypeNegativeKeyword SharedListItemType = "NegativeKeyword"
	SharedListItemTypeBrandItem       SharedListItemType = "BrandItem"
)
