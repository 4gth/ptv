package model

import "github.com/go-openapi/strfmt"

type FairEstimatePayload struct {
	FareEstimateResultStatus FareEstimateResultStatus `json:"FareEstimateResultStatus"`
	FareEstimateResult       FareEstimateResult       `json:"FareEstimateResult"`
}

type FairEstimateParameters struct {
	MinZone                 int32           `path:"min_zone"`
	MaxZone                 int32           `path:"max_zone"`
	JourneyTouchOnUtc       strfmt.DateTime `query:"journey_touch_on_utc"`
	JourneyTouchOffUtc      strfmt.DateTime `query:"journey_touch_off_utc"`
	IsJourneyInFreeTramZone bool            `query:"is_journey_in_free_tram_zone"`
	TravelledRouteTypes     []int32         `query:"travelled_route_types"`
}

type FairEstimate struct{}

func (FairEstimate) New() *Request {
	return &Request{
		Path:       "/v3/fare_estimate/min_zone/{min_zone}/max_zone/{max_zone}",
		Parameters: FairEstimateParameters{},
		Payload:    FairEstimatePayload{},
	}
}
