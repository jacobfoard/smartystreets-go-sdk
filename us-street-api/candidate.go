package us_street

type (
	Candidate struct {
		InputID              string     `json:"input_id,omitempty"`
		InputIndex           int        `json:"input_index"`
		CandidateIndex       int        `json:"candidate_index"`
		Addressee            string     `json:"addressee,omitempty"`
		DeliveryLine1        string     `json:"delivery_line_1,omitempty"`
		DeliveryLine2        string     `json:"delivery_line_2,omitempty"`
		LastLine             string     `json:"last_line,omitempty"`
		DeliveryPointBarcode string     `json:"delivery_point_barcode,omitempty"`
		Components           Components `json:"components,omitempty"`
		Metadata             Metadata   `json:"metadata,omitempty"`
		Analysis             Analysis   `json:"analysis,omitempty"`
	}

	Components struct {
		PrimaryNumber            string `json:"primary_number,omitempty"`
		StreetPredirection       string `json:"street_predirection,omitempty"`
		StreetName               string `json:"street_name,omitempty"`
		StreetPostdirection      string `json:"street_postdirection,omitempty"`
		StreetSuffix             string `json:"street_suffix,omitempty"`
		SecondaryNumber          string `json:"secondary_number,omitempty"`
		SecondaryDesignator      string `json:"secondary_designator,omitempty"`
		ExtraSecondaryNumber     string `json:"extra_secondary_number,omitempty"`
		ExtraSecondaryDesignator string `json:"extra_secondary_designator,omitempty"`
		PMBNumber                string `json:"pmb_number,omitempty"`
		PMBDesignator            string `json:"pmb_designator,omitempty"`
		CityName                 string `json:"city_name,omitempty"`
		DefaultCityName          string `json:"default_city_name,omitempty"`
		StateAbbreviation        string `json:"state_abbreviation,omitempty"`
		ZIPCode                  string `json:"zipcode,omitempty"`
		Plus4Code                string `json:"plus4_code,omitempty"`
		DeliveryPoint            string `json:"delivery_point,omitempty"`
		DeliveryPointCheckDigit  string `json:"delivery_point_check_digit,omitempty"`
		Urbanization             string `json:"urbanization,omitempty"`
	}

	Metadata struct {
		RecordType               string  `json:"record_type,omitempty"`
		ZIPType                  string  `json:"zip_type,omitempty"`
		CountyFIPS               string  `json:"county_fips,omitempty"`
		CountyName               string  `json:"county_name,omitempty"`
		CarrierRoute             string  `json:"carrier_route,omitempty"`
		CongressionalDistrict    string  `json:"congressional_district,omitempty"`
		BuildingDefaultIndicator string  `json:"building_default_indicator,omitempty"`
		RDI                      string  `json:"rdi,omitempty"`
		ELOTSequence             string  `json:"elot_sequence,omitempty"`
		ELOTSort                 string  `json:"elot_sort,omitempty"`
		Latitude                 float64 `json:"latitude,omitempty"`
		Longitude                float64 `json:"longitude,omitempty"`
		Precision                string  `json:"precision,omitempty"`
		TimeZone                 string  `json:"time_zone,omitempty"`
		UTCOffset                float32 `json:"utc_offset,omitempty"`
		DST                      bool    `json:"dst,omitempty"`
	}

	Analysis struct {
		DPVMatchCode      string `json:"dpv_match_code,omitempty"`
		DPVFootnotes      string `json:"dpv_footnotes,omitempty"`
		DPVCMRACode       string `json:"dpv_cmra,omitempty"`
		DPVVacantCode     string `json:"dpv_vacant,omitempty"`
		Active            string `json:"active,omitempty"`
		Footnotes         string `json:"footnotes,omitempty"`
		LACSLinkCode      string `json:"lacslink_code,omitempty"`
		LACSLinkIndicator string `json:"lacslink_indicator,omitempty"`
		SuiteLinkMatch    bool   `json:"suitelink_match,omitempty"`
		EWSMatch          bool   `json:"ews_match,omitempty"`
	}
)
