package wmsservice

type SKUHubValidationRequest struct {
	SKUCode string `json:"sku_code"`
	HubID   string `json:"hub_id"`
}
