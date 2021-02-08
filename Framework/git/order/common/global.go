package common

type Message struct {
	Code   int     `json:"code"`
	Remark string  `json:"remark"`
	Result *Result `json:"result,omitempty"`
}

type Result struct {
	Code   int    `json:"code"`
	Remark string `json:"remark,omitempty"`
}

type TheTrip struct {
	DepatureDate1 string `json:"depature_date_1"`
	DepatureDate2 string `json:"depature_date_2"`
	Provinsi      int64  `json:"provinsi"`
}

type TheTripResponse struct {
	Message       string          `json:"message"`
	Status        string          `json:"status"`
	TheTripDetail []TheTripDetail `json:"data"`
}

type TheTripDetail struct {
	AirlineName      string `json:"AirlineName,omitempty"`
	AirportName      string `json:"AirportName,omitempty"`
	CityName         string `json:"CityName,omitempty"`
	Currency         string `json:"Currency,omitempty"`
	DepartureDate    string `json:"DepartureDate,omitempty"`
	Description      string `json:"Description,omitempty"`
	Destination      string `json:"Destination,omitempty"`
	DetailTransit    string `json:"DetailTransit,omitempty"`
	DoubleType       string `json:"DoubleType,omitempty"`
	Duration         string `json:"Duration,omitempty"`
	Goods            string `json:"Goods,omitempty"`
	HotelName        string `json:"HotelName,omitempty"`
	HotelRating      string `json:"HotelRating,omitempty"`
	Lat              string `json:"Lat,omitempty"`
	LicenseNumber    string `json:"LicenseNumber,omitempty"`
	Logo             string `json:"Logo,omitempty"`
	Long             string `json:"Long,omitempty"`
	Origin           string `json:"Origin,omitempty"`
	OriginCity       string `json:"OriginCity,omitempty"`
	Price            string `json:"Price,omitempty"`
	PromoCode        string `json:"PromoCode,omitempty"`
	PromoDescription string `json:"PromoDescription,omitempty"`
	Provinsi         string `json:"Provinsi,omitempty"`
	QuadType         string `json:"QuadType,omitempty"`
	Rating           string `json:"Rating,omitempty"`
	ReturnDate       string `json:"ReturnDate,omitempty"`
	TermCondition    string `json:"TermCondition,omitempty"`
	Transit          string `json:"Transit,omitempty"`
	TravelID         string `json:"TravelID,omitempty"`
	TravelName       string `json:"TravelName,omitempty"`
	TripID           string `json:"TripID,omitempty"`
	TripleType       string `json:"TripleType,omitempty"`
}
