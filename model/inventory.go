package model

import "gorm.io/gorm"

type EcomEnventorie struct {
	gorm.Model
	Id                     int     `json:"id"`
	ProductVariantId       int     `json:"product_variant_id"`
	OrgId                  int     `json:"org_id"`
	SkuOrgId               string  `json:"sku_org_id"`
	Cost                   float64 `json:"cost"`
	BasePrice              float64 `json:"base_price"`
	Unit                   string  `json:"unit"`
	ReceiptQuantity        int     `json:"receipt_quantity"`
	SellQuantity           int     `json:"sell_quantity"`
	SellTempQuantity       int     `json:"sell_temp_quantity"`
	Quantity               int     `json:"quantity"`
	Description            string  `json:"description"`
	Status                 string  `json:"status"`
	CreatedAt              string  `json:"created_at"`
	CreatedBy              string  `json:"created_by"`
	UpdatedAt              string  `json:"updated_at"`
	UpdatedBy              string  `json:"updated_by"`
	DeletedAt              string  `json:"deleted_at"`
	DeletedBy              string  `json:"deleted_by"`
	Resource               string  `json:"resource"`
	IsSync                 int     `json:"is_sync"`
	SyncStatus             string  `json:"sync_status"`
	LastSyncAt             string  `json:"last_sync_at"`
	IsAutoSync             int     `json:"is_auto_sync"`
	Countfail              int     `json:"count_fail"`
	IsApplyAllChildrenOrgs int     `json:"is_apply_all_children_orgs"`
	SpecialPrice           float64 `json:"special_price"`
	DisplayFromDate        string  `json:"display_from_date"`
	DisplayToDate          string  `json:"display_to_date"`
	MonthlyThreshold       int     `json:"monthly_threshold"`
	SellerSku              string  `json:"seller_sku"`
	InternalSell           int     `json:"internal_sell"`
	IsShip                 int     `json:"is_ship"`
	IsComeGet              int     `json:"is_come_get"`
}
