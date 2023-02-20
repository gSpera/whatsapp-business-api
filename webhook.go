package whatsapp

// An Entry object, as documentated in https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookEntryObject struct {
	Object string `json:"object"`
	Entry  []struct {
		ID      string                `json:"id"`
		Changes []WebhookChangeObject `json:"changes"`
	} `json:"entry"`
}

// An Change object, as documentated in https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components
type WebhookChangeObject struct {
	Field string `json:"field"`
	Value struct {
		MessagingProduct string `json:"messaging_product"`
		Metadata         struct {
			DisplayPhoneNumber string `json:"display_phone_number"`
			PhoneNumberID      int    `json:"phone_number_id"`
		} `json:"metadata"`

		Contacts struct {
			WaID int `json:"wa_id"`

			Profile struct {
				Name string `json:"name"`
			} `json:"profile"`
		} `json:"contacts"`

		Errors struct {
			Code    int    `json:"code"`
			Title   string `json:"title"`
			Message string `json:"message"`
			// ErrorData: Unkown Type
			Details string `json:"details"`
		} `json:"errors"`

		// Errors
		// Statuses
		Messages []WebhookMessage `json:"messages"`
	} `json:"value"`
}

type WebhookMessage struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	Timestamp int    `json:"timestamp,string"`
	Type      string `json:"type"`

	Identity struct {
		// Acknowledged: Unkown Type
		CreatedTimestamp string `json:"created_timestamp"`
		Hash             string `json:"hash"`
	} `json:"identity"`

	Context struct {
		ID                  string `json:"id"`
		Forwarded           bool   `json:"forwarded"`
		FrequentlyForwarded bool   `json:"frequently_forwarded"`
		For                 string `json:"for"`
		ReferredProduct     struct {
			CatalogID         string `json:"catalog_id"`
			ProductRetailerID string `json:"product_retailer_id"`
		} `json:"referred_product"`
	} `json:"context"`

	Image struct {
		ID       string `json:"id"`
		MimeType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
	} `json:"image"`

	Audio struct {
		ID       string `json:"id"`
		MimeType string `json:"mime_type"`
	}

	Document struct {
		ID       string `json:"id"`
		Caption  string `json:"caption"`
		Filename string `json:"filename"`
		MimeType string `json:"mime_type"`
		SHA256   string `json:"sha256"`
	}

	Button struct {
		Payload string `json:"payload"`
		Text    string `json:"text"`
	}
}

type WebhookStatus struct {
	Converstion struct {
		ID     string `json:"id"`
		Origin struct {
			Type string `json:"type"`
		} `json:"origin"`
		ExpirationTimeStamp string `json:"expiration_timestamp,string"`
	} `json:"conversation"`
}
