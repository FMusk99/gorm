package model

import "gorm.io/gorm"

type EcomProduct struct {
	gorm.Model
	Id              int    `json:"id"`
	Code            string `json:"code"`
	RefCode         string `json:"ref_code"`
	Name            string `json:"name"`
	PackageContent  string `json:"package_content"`
	Description     string `json:"description"`
	TotalQuantity   int    `json:"total_quantity"`
	CategoryId      int    `json:"category_id"`
	CategoryTree    string `json:"category_tree"`
	CateTree        string `json:"cate_tree"`
	BrandId         int    `json:"brand_id"`
	SupplierId      int    `json:"supplier_id"`
	DistributorId   int    `json:"distributor_id"`
	ManufactureId   int    `json:"manufacture_id"`
	Tags            string `json:"tags"`
	WeightUnit      string `json:"weight_unit"`
	PackageUnit     string `json:"package_unit"`
	Weight          string `json:"weight"`
	Height          string `json:"height"`
	HeightUnit      string `json:"height_unit"`
	Width           string `json:"width"`
	WidthUnit       string `json:"width_unit"`
	Length          string `json:"length"`
	LengthUnit      string `json:"length_unit"`
	Image           string `json:"image"`
	OrgId           int    `json:"org_id"`
	Status          string `json:"status"`
	RefId           string `json:"ref_id"`
	CreatedAt       string `json:"created_at"`
	CreatedBy       string `json:"created_by"`
	UpdatedAt       string `json:"updated_at"`
	UpdatedBy       string `json:"updated_by"`
	DeletedAt       string `json:"deleted_at"`
	DeletedBy       string `json:"deleted_by"`
	Resource        string `json:"resource"`
	IsSync          int    `json:"is_sync"`
	SyncStatus      string `json:"sync_status"`
	LastSyncAt      string `json:"last_sync_at"`
	IsAutoSync      int    `json:"is_auto_sync"`
	CountFail       int    `json:"count_fail"`
	ManualCateId    string `json:"manual_cate_id"`
	ManualBrandName string `json:"manual_brand_name"`
	ManualTag       string `json:"manual_tag"`
	ManualId        string `json:"manual_id"`
	SupplierId1     int    `json:"supplier_id_1"`
	// Model             string `json:"model"`
}
