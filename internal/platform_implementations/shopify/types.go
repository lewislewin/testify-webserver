package shopify

type Order struct {
	ID                       int                   `json:"id,omitempty"`
	AdminGraphqlApiID        string                `json:"admin_graphql_api_id,omitempty"`
	AppID                    int                   `json:"app_id,omitempty"`
	BrowserIP                string                `json:"browser_ip,omitempty"`
	BuyerAcceptsMarketing    bool                  `json:"buyer_accepts_marketing,omitempty"`
	CancelReason             *string               `json:"cancel_reason,omitempty"`
	CancelledAt              *string               `json:"cancelled_at,omitempty"`
	CartToken                *string               `json:"cart_token,omitempty"`
	CheckoutID               int64                 `json:"checkout_id,omitempty"`
	CheckoutToken            string                `json:"checkout_token,omitempty"`
	ClientDetails            ClientDetails         `json:"client_details,omitempty"`
	ClosedAt                 *string               `json:"closed_at,omitempty"`
	ConfirmationNumber       string                `json:"confirmation_number,omitempty"`
	Confirmed                bool                  `json:"confirmed,omitempty"`
	ContactEmail             string                `json:"contact_email,omitempty"`
	CreatedAt                string                `json:"created_at,omitempty"`
	Currency                 string                `json:"currency,omitempty"`
	CurrentSubtotalPrice     string                `json:"current_subtotal_price,omitempty"`
	CurrentSubtotalPriceSet  PriceSet              `json:"current_subtotal_price_set,omitempty"`
	CurrentTotalDiscounts    string                `json:"current_total_discounts,omitempty"`
	CurrentTotalDiscountsSet PriceSet              `json:"current_total_discounts_set,omitempty"`
	CurrentTotalPrice        string                `json:"current_total_price,omitempty"`
	CurrentTotalPriceSet     PriceSet              `json:"current_total_price_set,omitempty"`
	CurrentTotalTax          string                `json:"current_total_tax,omitempty"`
	CurrentTotalTaxSet       PriceSet              `json:"current_total_tax_set,omitempty"`
	CustomerLocale           string                `json:"customer_locale,omitempty"`
	DiscountCodes            []DiscountCode        `json:"discount_codes,omitempty"`
	Email                    string                `json:"email,omitempty"`
	FinancialStatus          string                `json:"financial_status,omitempty"`
	FulfillmentStatus        *string               `json:"fulfillment_status,omitempty"`
	LandingSite              *string               `json:"landing_site,omitempty"`
	LocationID               *int64                `json:"location_id,omitempty"`
	Name                     string                `json:"name,omitempty"`
	Note                     *string               `json:"note,omitempty"`
	Number                   int                   `json:"number,omitempty"`
	OrderNumber              int                   `json:"order_number,omitempty"`
	OrderStatusURL           string                `json:"order_status_url,omitempty"`
	ProcessedAt              string                `json:"processed_at,omitempty"`
	Reference                string                `json:"reference,omitempty"`
	ReferringSite            *string               `json:"referring_site,omitempty"`
	SourceIdentifier         string                `json:"source_identifier,omitempty"`
	SourceName               string                `json:"source_name,omitempty"`
	SubtotalPrice            string                `json:"subtotal_price,omitempty"`
	SubtotalPriceSet         PriceSet              `json:"subtotal_price_set,omitempty"`
	Tags                     string                `json:"tags,omitempty"`
	TaxExempt                bool                  `json:"tax_exempt,omitempty"`
	TaxesIncluded            bool                  `json:"taxes_included,omitempty"`
	Test                     bool                  `json:"test,omitempty"`
	Token                    string                `json:"token,omitempty"`
	TotalDiscounts           string                `json:"total_discounts,omitempty"`
	TotalDiscountsSet        PriceSet              `json:"total_discounts_set,omitempty"`
	TotalLineItemsPrice      string                `json:"total_line_items_price,omitempty"`
	TotalLineItemsPriceSet   PriceSet              `json:"total_line_items_price_set,omitempty"`
	TotalOutstanding         string                `json:"total_outstanding,omitempty"`
	TotalPrice               string                `json:"total_price,omitempty"`
	TotalPriceSet            PriceSet              `json:"total_price_set,omitempty"`
	TotalShippingPriceSet    PriceSet              `json:"total_shipping_price_set,omitempty"`
	TotalTax                 string                `json:"total_tax,omitempty"`
	TotalTaxSet              PriceSet              `json:"total_tax_set,omitempty"`
	TotalTipReceived         string                `json:"total_tip_received,omitempty"`
	TotalWeight              int                   `json:"total_weight,omitempty"`
	Transactions             []Transaction         `json:transactions`
	UpdatedAt                string                `json:"updated_at,omitempty"`
	UserID                   int64                 `json:"user_id,omitempty"`
	BillingAddress           Address               `json:"billing_address,omitempty"`
	ShippingAddress          Address               `json:"shipping_address,omitempty"`
	Customer                 Customer              `json:"customer,omitempty"`
	DiscountApplications     []DiscountApplication `json:"discount_applications,omitempty"`
	Fulfillments             []Fulfillment         `json:"fulfillments,omitempty"`
	LineItems                []LineItem            `json:"line_items,omitempty"`
	PaymentTerms             *string               `json:"payment_terms,omitempty"`
	Refunds                  []Refund              `json:"refunds,omitempty"`
	ShippingLines            []ShippingLine        `json:"shipping_lines,omitempty"`
}

type LineItem struct {
	ID                         int64                `json:"id,omitempty"`
	AdminGraphqlApiID          string               `json:"admin_graphql_api_id,omitempty"`
	AttributedStaffs           []interface{}        `json:"attributed_staffs,omitempty"`
	FulfillableQuantity        int                  `json:"fulfillable_quantity,omitempty"`
	FulfillmentService         string               `json:"fulfillment_service,omitempty"`
	FulfillmentStatus          *string              `json:"fulfillment_status,omitempty"`
	GiftCard                   bool                 `json:"gift_card,omitempty"`
	Grams                      int                  `json:"grams,omitempty"`
	Name                       string               `json:"name,omitempty"`
	Price                      string               `json:"price,omitempty"`
	PriceSet                   PriceSet             `json:"price_set,omitempty"`
	ProductExists              bool                 `json:"product_exists,omitempty"`
	ProductID                  int64                `json:"product_id,omitempty"`
	Properties                 []interface{}        `json:"properties,omitempty"`
	Quantity                   int                  `json:"quantity,omitempty"`
	RequiresShipping           bool                 `json:"requires_shipping,omitempty"`
	SKU                        string               `json:"sku,omitempty"`
	Taxable                    bool                 `json:"taxable,omitempty"`
	Title                      string               `json:"title,omitempty"`
	TotalDiscount              string               `json:"total_discount,omitempty"`
	TotalDiscountSet           PriceSet             `json:"total_discount_set,omitempty"`
	VariantID                  int64                `json:"variant_id,omitempty"`
	VariantInventoryManagement *string              `json:"variant_inventory_management,omitempty"`
	VariantTitle               *string              `json:"variant_title,omitempty"`
	Vendor                     string               `json:"vendor,omitempty"`
	TaxLines                   []interface{}        `json:"tax_lines,omitempty"`
	Duties                     []interface{}        `json:"duties,omitempty"`
	DiscountAllocations        []DiscountAllocation `json:"discount_allocations,omitempty"`
}

type Address struct {
	ID           int64   `json:"id,omitempty"`
	CustomerID   int64   `json:"customer_id,omitempty"`
	FirstName    string  `json:"first_name,omitempty"`
	LastName     string  `json:"last_name,omitempty"`
	Company      *string `json:"company,omitempty"`
	Address1     string  `json:"address1,omitempty"`
	Address2     *string `json:"address2,omitempty"`
	City         string  `json:"city,omitempty"`
	Province     string  `json:"province,omitempty"`
	Country      string  `json:"country,omitempty"`
	Zip          string  `json:"zip,omitempty"`
	Phone        string  `json:"phone,omitempty"`
	Name         string  `json:"name,omitempty"`
	ProvinceCode string  `json:"province_code,omitempty"`
	CountryCode  string  `json:"country_code,omitempty"`
	CountryName  string  `json:"country_name,omitempty"`
	Default      bool    `json:"default,omitempty"`
}

type Transaction struct {
	ID                       int64                    `json:"id,omitempty"`
	AdminGraphqlApiID        string                   `json:"admin_graphql_api_id,omitempty"`
	Amount                   string                   `json:"amount,omitempty"`
	Authorization            string                   `json:"authorization,omitempty"`
	CreatedAt                string                   `json:"created_at,omitempty"`
	Currency                 string                   `json:"currency,omitempty"`
	DeviceID                 *string                  `json:"device_id,omitempty"`
	ErrorCode                *string                  `json:"error_code,omitempty"`
	Gateway                  string                   `json:"gateway,omitempty"`
	Kind                     string                   `json:"kind,omitempty"`
	LocationID               *string                  `json:"location_id,omitempty"`
	Message                  string                   `json:"message,omitempty"`
	OrderID                  int64                    `json:"order_id,omitempty"`
	ParentID                 int64                    `json:"parent_id,omitempty"`
	PaymentID                string                   `json:"payment_id,omitempty"`
	PaymentsRefundAttributes PaymentsRefundAttributes `json:"payments_refund_attributes,omitempty"`
	ProcessedAt              string                   `json:"processed_at,omitempty"`
	Receipt                  map[string]interface{}   `json:"receipt,omitempty"`
	SourceName               string                   `json:"source_name,omitempty"`
	Status                   string                   `json:"status,omitempty"`
	Test                     bool                     `json:"test,omitempty"`
	UserID                   int64                    `json:"user_id,omitempty"`
	PaymentDetails           PaymentDetails           `json:"payment_details,omitempty"`
}

type DiscountCode struct {
	Code   string `json:"code,omitempty"`
	Amount string `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
}

type DiscountApplication struct {
	TargetType       string  `json:"target_type,omitempty"`
	Type             string  `json:"type,omitempty"`
	Value            string  `json:"value,omitempty"`
	ValueType        string  `json:"value_type,omitempty"`
	AllocationMethod string  `json:"allocation_method,omitempty"`
	TargetSelection  string  `json:"target_selection,omitempty"`
	Title            *string `json:"title,omitempty"`
	Description      *string `json:"description,omitempty"`
	Code             *string `json:"code,omitempty"`
}

type DiscountAllocation struct {
	Amount                   string   `json:"amount,omitempty"`
	AmountSet                PriceSet `json:"amount_set,omitempty"`
	DiscountApplicationIndex int      `json:"discount_application_index,omitempty"`
}

type ShippingLine struct {
	ID                            int64                `json:"id,omitempty"`
	CarrierIdentifier             *string              `json:"carrier_identifier,omitempty"`
	Code                          string               `json:"code,omitempty"`
	DiscountedPrice               string               `json:"discounted_price,omitempty"`
	DiscountedPriceSet            PriceSet             `json:"discounted_price_set,omitempty"`
	Phone                         *string              `json:"phone,omitempty"`
	Price                         string               `json:"price,omitempty"`
	PriceSet                      PriceSet             `json:"price_set,omitempty"`
	RequestedFulfillmentServiceID *string              `json:"requested_fulfillment_service_id,omitempty"`
	Source                        string               `json:"source,omitempty"`
	Title                         string               `json:"title,omitempty"`
	TaxLines                      []interface{}        `json:"tax_lines,omitempty"`
	DiscountAllocations           []DiscountAllocation `json:"discount_allocations,omitempty"`
}

type ClientDetails struct {
	AcceptLanguage *string `json:"accept_language,omitempty"`
	BrowserHeight  *int    `json:"browser_height,omitempty"`
	BrowserIP      string  `json:"browser_ip,omitempty"`
	BrowserWidth   *int    `json:"browser_width,omitempty"`
	SessionHash    *string `json:"session_hash,omitempty"`
	UserAgent      string  `json:"user_agent,omitempty"`
}

type PriceSet struct {
	ShopMoney        Money `json:"shop_money,omitempty"`
	PresentmentMoney Money `json:"presentment_money,omitempty"`
}

type Money struct {
	Amount       string `json:"amount,omitempty"`
	CurrencyCode string `json:"currency_code,omitempty"`
}

type Fulfillment struct {
	ID                int64                  `json:"id,omitempty"`
	AdminGraphqlApiID string                 `json:"admin_graphql_api_id,omitempty"`
	CreatedAt         string                 `json:"created_at,omitempty"`
	LocationID        int64                  `json:"location_id,omitempty"`
	Name              string                 `json:"name,omitempty"`
	OrderID           int64                  `json:"order_id,omitempty"`
	OriginAddress     map[string]interface{} `json:"origin_address,omitempty"`
	Receipt           map[string]interface{} `json:"receipt,omitempty"`
	Service           string                 `json:"service,omitempty"`
	ShipmentStatus    *string                `json:"shipment_status,omitempty"`
	Status            string                 `json:"status,omitempty"`
	TrackingCompany   *string                `json:"tracking_company,omitempty"`
	TrackingNumber    *string                `json:"tracking_number,omitempty"`
	TrackingNumbers   []string               `json:"tracking_numbers,omitempty"`
	TrackingURL       *string                `json:"tracking_url,omitempty"`
	TrackingURLs      []string               `json:"tracking_urls,omitempty"`
	UpdatedAt         string                 `json:"updated_at,omitempty"`
	LineItems         []LineItem             `json:"line_items,omitempty"`
}

type Refund struct {
	ID                int64             `json:"id,omitempty"`
	AdminGraphqlApiID string            `json:"admin_graphql_api_id,omitempty"`
	CreatedAt         string            `json:"created_at,omitempty"`
	Note              *string           `json:"note,omitempty"`
	OrderID           int64             `json:"order_id,omitempty"`
	ProcessedAt       string            `json:"processed_at,omitempty"`
	Restock           bool              `json:"restock,omitempty"`
	TotalDutiesSet    PriceSet          `json:"total_duties_set,omitempty"`
	UserID            int64             `json:"user_id,omitempty"`
	OrderAdjustments  []OrderAdjustment `json:"order_adjustments,omitempty"`
	Transactions      []Transaction     `json:"transactions,omitempty"`
	RefundLineItems   []RefundLineItem  `json:"refund_line_items,omitempty"`
	Duties            []interface{}     `json:"duties,omitempty"`
}

type OrderAdjustment struct {
	ID           int64    `json:"id,omitempty"`
	Amount       string   `json:"amount,omitempty"`
	AmountSet    PriceSet `json:"amount_set,omitempty"`
	Kind         string   `json:"kind,omitempty"`
	OrderID      int64    `json:"order_id,omitempty"`
	Reason       string   `json:"reason,omitempty"`
	RefundID     int64    `json:"refund_id,omitempty"`
	TaxAmount    string   `json:"tax_amount,omitempty"`
	TaxAmountSet PriceSet `json:"tax_amount_set,omitempty"`
}

type RefundLineItem struct {
	ID          int64    `json:"id,omitempty"`
	LineItemID  int64    `json:"line_item_id,omitempty"`
	LocationID  int64    `json:"location_id,omitempty"`
	Quantity    int      `json:"quantity,omitempty"`
	RestockType string   `json:"restock_type,omitempty"`
	Subtotal    string   `json:"subtotal,omitempty"`
	SubtotalSet PriceSet `json:"subtotal_set,omitempty"`
	TotalTax    string   `json:"total_tax,omitempty"`
	TotalTaxSet PriceSet `json:"total_tax_set,omitempty"`
	LineItem    LineItem `json:"line_item,omitempty"`
}

type PaymentsRefundAttributes struct {
	Status                  string  `json:"status,omitempty"`
	AcquirerReferenceNumber *string `json:"acquirer_reference_number,omitempty"`
}

type PaymentDetails struct {
	CreditCardBin             string  `json:"credit_card_bin,omitempty"`
	AvsResultCode             string  `json:"avs_result_code,omitempty"`
	CvvResultCode             string  `json:"cvv_result_code,omitempty"`
	CreditCardNumber          string  `json:"credit_card_number,omitempty"`
	CreditCardCompany         string  `json:"credit_card_company,omitempty"`
	BuyerActionInfo           *string `json:"buyer_action_info,omitempty"`
	CreditCardName            string  `json:"credit_card_name,omitempty"`
	CreditCardWallet          *string `json:"credit_card_wallet,omitempty"`
	CreditCardExpirationMonth int     `json:"credit_card_expiration_month,omitempty"`
	CreditCardExpirationYear  int     `json:"credit_card_expiration_year,omitempty"`
}

type Product struct {
	ID                int64     `json:"id,omitempty"`
	Title             string    `json:"title,omitempty"`
	BodyHTML          string    `json:"body_html,omitempty"`
	Vendor            string    `json:"vendor,omitempty"`
	ProductType       string    `json:"product_type,omitempty"`
	CreatedAt         string    `json:"created_at,omitempty"`
	Handle            string    `json:"handle,omitempty"`
	UpdatedAt         string    `json:"updated_at,omitempty"`
	PublishedAt       string    `json:"published_at,omitempty"`
	TemplateSuffix    string    `json:"template_suffix,omitempty"`
	PublishedScope    string    `json:"published_scope,omitempty"`
	Tags              string    `json:"tags,omitempty"`
	Status            string    `json:"status,omitempty"`
	AdminGraphqlApiID string    `json:"admin_graphql_api_id,omitempty"`
	Variants          []Variant `json:"variants,omitempty"`
	Options           []Option  `json:"options,omitempty"`
	Images            []Image   `json:"images,omitempty"`
	Image             *Image    `json:"image,omitempty"`
}

type Variant struct {
	ID                   int64   `json:"id,omitempty"`
	ProductID            int64   `json:"product_id,omitempty"`
	Title                string  `json:"title,omitempty"`
	Price                string  `json:"price,omitempty"`
	SKU                  string  `json:"sku,omitempty"`
	Position             int     `json:"position,omitempty"`
	InventoryPolicy      string  `json:"inventory_policy,omitempty"`
	CompareAtPrice       *string `json:"compare_at_price,omitempty"`
	FulfillmentService   string  `json:"fulfillment_service,omitempty"`
	InventoryManagement  string  `json:"inventory_management,omitempty"`
	Option1              string  `json:"option1,omitempty"`
	Option2              *string `json:"option2,omitempty"`
	Option3              *string `json:"option3,omitempty"`
	CreatedAt            string  `json:"created_at,omitempty"`
	UpdatedAt            string  `json:"updated_at,omitempty"`
	Taxable              bool    `json:"taxable,omitempty"`
	Barcode              string  `json:"barcode,omitempty"`
	Grams                int     `json:"grams,omitempty"`
	Weight               float64 `json:"weight,omitempty"`
	WeightUnit           string  `json:"weight_unit,omitempty"`
	InventoryItemID      int64   `json:"inventory_item_id,omitempty"`
	InventoryQuantity    int     `json:"inventory_quantity,omitempty"`
	OldInventoryQuantity int     `json:"old_inventory_quantity,omitempty"`
	RequiresShipping     bool    `json:"requires_shipping,omitempty"`
	AdminGraphqlApiID    string  `json:"admin_graphql_api_id,omitempty"`
	ImageID              *int64  `json:"image_id,omitempty"`
}

type Option struct {
	ID        int64    `json:"id,omitempty"`
	ProductID int64    `json:"product_id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Position  int      `json:"position,omitempty"`
	Values    []string `json:"values,omitempty"`
}

type Image struct {
	ID                int64   `json:"id,omitempty"`
	ProductID         int64   `json:"product_id,omitempty"`
	Position          int     `json:"position,omitempty"`
	CreatedAt         string  `json:"created_at,omitempty"`
	UpdatedAt         string  `json:"updated_at,omitempty"`
	Alt               *string `json:"alt,omitempty"`
	Width             int     `json:"width,omitempty"`
	Height            int     `json:"height,omitempty"`
	Src               string  `json:"src,omitempty"`
	VariantIDs        []int64 `json:"variant_ids,omitempty"`
	AdminGraphqlApiID string  `json:"admin_graphql_api_id,omitempty"`
}

type Customer struct {
	ID                        int64               `json:"id,omitempty"`
	Email                     string              `json:"email,omitempty"`
	CreatedAt                 string              `json:"created_at,omitempty"`
	UpdatedAt                 string              `json:"updated_at,omitempty"`
	FirstName                 string              `json:"first_name,omitempty"`
	LastName                  string              `json:"last_name,omitempty"`
	OrdersCount               int                 `json:"orders_count,omitempty"`
	State                     string              `json:"state,omitempty"`
	TotalSpent                string              `json:"total_spent,omitempty"`
	LastOrderID               int64               `json:"last_order_id,omitempty"`
	Note                      *string             `json:"note,omitempty"`
	VerifiedEmail             bool                `json:"verified_email,omitempty"`
	MultipassIdentifier       *string             `json:"multipass_identifier,omitempty"`
	TaxExempt                 bool                `json:"tax_exempt,omitempty"`
	Tags                      string              `json:"tags,omitempty"`
	LastOrderName             string              `json:"last_order_name,omitempty"`
	Currency                  string              `json:"currency,omitempty"`
	Phone                     string              `json:"phone,omitempty"`
	Addresses                 []Address           `json:"addresses,omitempty"`
	AcceptsMarketing          bool                `json:"accepts_marketing,omitempty"`
	AcceptsMarketingUpdatedAt *string             `json:"accepts_marketing_updated_at,omitempty"`
	MarketingOptInLevel       string              `json:"marketing_opt_in_level,omitempty"`
	TaxExemptions             []interface{}       `json:"tax_exemptions,omitempty"`
	EmailMarketingConsent     MarketingConsent    `json:"email_marketing_consent,omitempty"`
	SmsMarketingConsent       SmsMarketingConsent `json:"sms_marketing_consent,omitempty"`
	AdminGraphqlApiID         string              `json:"admin_graphql_api_id,omitempty"`
	DefaultAddress            Address             `json:"default_address,omitempty"`
}

type MarketingConsent struct {
	State            string  `json:"state,omitempty"`
	OptInLevel       string  `json:"opt_in_level,omitempty"`
	ConsentUpdatedAt *string `json:"consent_updated_at,omitempty"`
}

type SmsMarketingConsent struct {
	State                string  `json:"state,omitempty"`
	OptInLevel           string  `json:"opt_in_level,omitempty"`
	ConsentUpdatedAt     *string `json:"consent_updated_at,omitempty"`
	ConsentCollectedFrom string  `json:"consent_collected_from,omitempty"`
}
