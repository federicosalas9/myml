package myml

type User struct {
	ID               int64    `json:"id"`
	Nickname         string `json:"nickname"`
	RegistrationDate string `json:"registration_date"`
	CountryID        string `json:"country_id"`
	Address          struct {
		City  string `json:"city"`
		State string `json:"state"`
	} `json:"address"`
	UserType         string      `json:"user_type"`
	Tags             []string    `json:"tags"`
	Logo             interface{} `json:"logo"`
	Points           int         `json:"points"`
	SiteID           string      `json:"site_id"`
	Permalink        string      `json:"permalink"`
	SellerReputation struct {
		LevelID           interface{} `json:"level_id"`
		PowerSellerStatus interface{} `json:"power_seller_status"`
		Transactions      struct {
			Canceled  int    `json:"canceled"`
			Completed int    `json:"completed"`
			Period    string `json:"period"`
			Ratings   struct {
			} `json:"ratings"`
			Total int `json:"total"`
		} `json:"transactions"`
	} `json:"seller_reputation"`
	BuyerReputation struct {
		Tags []interface{} `json:"tags"`
	} `json:"buyer_reputation"`
	Status struct {
		SiteStatus string `json:"site_status"`
	} `json:"status"`
}

type Site struct {
	ID                 string        `json:"id"`
	Name               string        `json:"name"`
	CountryID          string        `json:"country_id"`
	SaleFeesMode       string        `json:"sale_fees_mode"`
	MercadopagoVersion int           `json:"mercadopago_version"`
	DefaultCurrencyID  string        `json:"default_currency_id"`
	ImmediatePayment   string        `json:"immediate_payment"`
	PaymentMethodIds   []interface{} `json:"payment_method_ids"`
	Settings           struct {
	} `json:"settings"`
	Currencies []interface{} `json:"currencies"`
	Categories []interface{} `json:"categories"`
}

type Categories []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	User     User       `json:"user"`
	Site     Site       `json:"site"`
	Categories Categories `json:"categories"`
}