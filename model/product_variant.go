package model

import "gorm.io/gorm"

type EcomProductVariant struct {
	gorm.Model
	Id              int     `json:"id"`
	ProductId       int     `json:"product_id"`
	Name            string  `json:"name"`
	Barcode         string  `json:"barcode"`
	IsDisplayed     int     `json:"is_displayed"`
	Sku             string  `json:"sku"`
	SellerSku       string  `json:"seller_sku"`
	WeightUnit      string  `json:"weight_unit"`
	PackageUnit     string  `json:"package_unit"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	HeightUnit      string  `json:"height_unit"`
	Width           float64 `json:"width"`
	WidthUnit       string  `json:"width_unit"`
	Length          int     `json:"length"`
	LengthUnit      string  `json:"length_unit"`
	Description     string  `json:"description"`
	OrgId           int     `json:"org_id"`
	Status          string  `json:"status"`
	Sorting         string  `json:"sorting"`
	RefId           string  `json:"ref_id"`
	CreatedAt       string  `json:"created_at"`
	CreatedBy       string  `json:"created_by"`
	UpdatedAt       string  `json:"updated_at"`
	UpdatedBy       string  `json:"updated_by"`
	DeletedAt       string  `json:"deleted_at"`
	DeletedBy       string  `json:"deleted_by"`
	Resource        string  `json:"resource"`
	IsSync          int     `json:"is_sync"`
	SyncStatus      string  `json:"sync_status"`
	LastSyncAt      string  `json:"last_sync_at"`
	IsAutoSync      int     `json:"is_auto_sync"`
	CountFail       int     `json:"count_fail"`
	SkuUnit         string  `json:"sku_unit"`
	PackageValue    int     `json:"package_value"`
	IsNotUsed       int     `json:"is_not_used"`
	DefaultImage    string  `json:"default_image"`
	IsManual        int     `json:"is_manual"`
	ManualId        string  `json:"manual_id"`
	GroupId         int     `json:"group_id"`
	GroupDetailId   int     `json:"group_detail_id"`
	MaterialGroupId int     `json:"material_group_id"`
	SegmentId       int     `json:"segment_id"`
	KidoProductCode string  `json:"kido_product_code"`
	SkuType         string  `json:"sku_type"`
}
