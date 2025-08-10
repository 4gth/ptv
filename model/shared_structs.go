package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type RouteTypeEnum int

const (
	Train    RouteTypeEnum = iota
	Tram                   = 1
	Bus                    = 2
	Vline                  = 3
	Nightbus               = 4
)

func (r RouteTypeEnum) String() string {
	switch r {
	case Train:
		return "Train"
	case Tram:
		return "Tram"
	case Bus:
		return "Bus"
	case Vline:
		return "Vline"
	case Nightbus:
		return "NightBus"
	default:
		return "Unknown"
	}
}

type Parameters struct {
	DateUtc                      strfmt.DateTime `path:"date_utc"`
	DirectionID                  int32           `path:"direction_id"`
	DisruptionModes              []int32         `path:"disruption_modes"`
	DisruptionStatus             string          `path:"disruption_status"`
	Expand                       []int32         `path:"expand"`
	Gtfs                         bool            `path:"gtfs"`
	IncludeAddresses             bool            `path:"include_addresses"`
	IncludeAdvertisedInterchange bool            `path:"include_advertised_interchange"`
	IncludeCancelled             bool            `path:"include_cancelled"`
	IncludeGeopath               bool            `path:"include_geopath"`
	IncludeOutlets               bool            `path:"include_outlets"`
	IncludeSkippedStops          bool            `path:"include_skipped_stops"`
	IsJourneyInFreeTramZone      bool            `path:"is_journey_in_free_tram_zone"`
	JourneyTouchOffUtc           strfmt.DateTime `path:"journey_touch_off_utc"`
	JourneyTouchOnUtc            strfmt.DateTime `path:"journey_touch_on_utc"`
	Latitude                     float32         `path:"latitude"`
	Longitude                    float32         `path:"longitude"`
	LookBackwards                bool            `path:"look_backwards"`
	MatchRouteBySuburb           bool            `path:"match_route_by_suburb"`
	MatchStopByGtfsStopID        bool            `path:"match_stop_by_gtfs_stop_id"`
	MatchStopBySuburb            bool            `path:"match_stop_by_suburb"`
	MaxDistance                  float64         `path:"max_distance"`
	MaxResults                   int32           `path:"max_results"`
	MaxZone                      int32           `path:"max_zone"`
	MinZone                      int32           `path:"min_zone"`
	PlatformNumbers              []int32         `path:"platform_numbers"`
	RouteID                      int32           `path:"route_id"`
	RouteName                    string          `path:"route_name"`
	RouteNumber                  string          `path:"route_number"`
	RouteType                    int32           `path:"route_type"`
	RouteTypes                   []int32         `path:"route_types"`
	RunRef                       string          `path:"run_ref"`
	SearchTerm                   string          `path:"search_term"`
	StopAccessibility            bool            `path:"stop_accessibility"`
	StopAmenities                bool            `path:"stop_amenities"`
	StopContact                  bool            `path:"stop_contact"`
	StopDisruptions              bool            `path:"stop_disruptions"`
	StopID                       int32           `path:"stop_id"`
	StopIDs                      []int32         `path:"stop_ids"`
	StopLocation                 bool            `path:"stop_location"`
	StopStaffing                 bool            `path:"stop_staffing"`
	StopTicket                   bool            `path:"stop_ticket"`
	TravelledRouteTypes          []int32         `path:"travelled_route_types"`
}

type Departures struct {
	StopID                int       `json:"stop_id,omitempty"`
	RouteID               int       `json:"route_id,omitempty"`
	RunID                 int       `json:"run_id,omitempty"`
	RunRef                string    `json:"run_ref,omitempty"`
	DirectionID           int       `json:"direction_id,omitempty"`
	DisruptionIds         []int     `json:"disruption_ids,omitempty"`
	ScheduledDepartureUtc time.Time `json:"scheduled_departure_utc"`
	EstimatedDepartureUtc time.Time `json:"estimated_departure_utc"`
	AtPlatform            bool      `json:"at_platform,omitempty"`
	PlatformNumber        string    `json:"platform_number,omitempty"`
	Flags                 string    `json:"flags,omitempty"`
	DepartureSequence     int       `json:"departure_sequence,omitempty"`
	DepartureNote         string    `json:"departure_note,omitempty"`
}

type DisruptionModes struct {
	DisruptionModeName string `json:"disruption_mode_name"`
	DisruptionMode     int    `json:"disruption_mode"`
}

type Directions struct {
	RouteDirectionDescription string `json:"route_direction_description,omitempty"`
	DirectionID               int    `json:"direction_id,omitempty"`
	DirectionName             string `json:"direction_name,omitempty"`
	RouteID                   int    `json:"route_id,omitempty"`
	RouteType                 int    `json:"route_type,omitempty"`
}

type Direction struct {
	RouteDirectionID int    `json:"route_direction_id,omitempty"`
	DirectionID      int    `json:"direction_id,omitempty"`
	DirectionName    string `json:"direction_name,omitempty"`
	ServiceTime      string `json:"service_time,omitempty"`
}

type Routes struct {
	RouteType   int       `json:"route_type,omitempty"`
	RouteID     int       `json:"route_id,omitempty"`
	RouteName   string    `json:"route_name,omitempty"`
	RouteNumber string    `json:"route_number,omitempty"`
	RouteGtfsID string    `json:"route_gtfs_id,omitempty"`
	Direction   Direction `json:"direction"`
}

type Stops struct {
	StopID   int    `json:"stop_id,omitempty"`
	StopName string `json:"stop_name,omitempty"`
}

type General struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type MetroTrain struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type MetroTram struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type MetroBus struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type RegionalTrain struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type RegionalCoach struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type RegionalBus struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type SchoolBus struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type Telebus struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type NightBus struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type Ferry struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type InterstateTrain struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type Skybus struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type Taxi struct {
	DisruptionID     int       `json:"disruption_id,omitempty"`
	Title            string    `json:"title,omitempty"`
	URL              string    `json:"url,omitempty"`
	Description      string    `json:"description,omitempty"`
	DisruptionStatus string    `json:"disruption_status,omitempty"`
	DisruptionType   string    `json:"disruption_type,omitempty"`
	PublishedOn      time.Time `json:"published_on"`
	LastUpdated      time.Time `json:"last_updated"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	Routes           []Routes  `json:"routes,omitempty"`
	Stops            []Stops   `json:"stops,omitempty"`
	Colour           string    `json:"colour,omitempty"`
	DisplayOnBoard   bool      `json:"display_on_board,omitempty"`
	DisplayStatus    bool      `json:"display_status,omitempty"`
}

type Disruptions struct {
	General         []General         `json:"general,omitempty"`
	MetroTrain      []MetroTrain      `json:"metro_train,omitempty"`
	MetroTram       []MetroTram       `json:"metro_tram,omitempty"`
	MetroBus        []MetroBus        `json:"metro_bus,omitempty"`
	RegionalTrain   []RegionalTrain   `json:"regional_train,omitempty"`
	RegionalCoach   []RegionalCoach   `json:"regional_coach,omitempty"`
	RegionalBus     []RegionalBus     `json:"regional_bus,omitempty"`
	SchoolBus       []SchoolBus       `json:"school_bus,omitempty"`
	Telebus         []Telebus         `json:"telebus,omitempty"`
	NightBus        []NightBus        `json:"night_bus,omitempty"`
	Ferry           []Ferry           `json:"ferry,omitempty"`
	InterstateTrain []InterstateTrain `json:"interstate_train,omitempty"`
	Skybus          []Skybus          `json:"skybus,omitempty"`
	Taxi            []Taxi            `json:"taxi,omitempty"`
}

type FareEstimateResultStatus struct {
	StatusCode int    `json:"StatusCode,omitempty"`
	Message    string `json:"Message,omitempty"`
}

type ZoneInfo struct {
	MinZone     int   `json:"MinZone,omitempty"`
	MaxZone     int   `json:"MaxZone,omitempty"`
	UniqueZones []int `json:"UniqueZones,omitempty"`
}

type PassengerFares struct {
	PassengerType       string  `json:"PassengerType,omitempty"`
	Fare2HourPeak       float32 `json:"Fare2HourPeak,omitempty"`
	Fare2HourOffPeak    float32 `json:"Fare2HourOffPeak,omitempty"`
	FareDailyPeak       float32 `json:"FareDailyPeak,omitempty"`
	FareDailyOffPeak    float32 `json:"FareDailyOffPeak,omitempty"`
	Pass7Days           float32 `json:"Pass7Days,omitempty"`
	Pass28To69DayPerDay float32 `json:"Pass28To69DayPerDay,omitempty"`
	Pass70PlusDayPerDay float32 `json:"Pass70PlusDayPerDay,omitempty"`
	WeekendCap          float32 `json:"WeekendCap,omitempty"`
	HolidayCap          float32 `json:"HolidayCap,omitempty"`
}

type FareEstimateResult struct {
	IsEarlyBird             bool             `json:"IsEarlyBird,omitempty"`
	IsJourneyInFreeTramZone bool             `json:"IsJourneyInFreeTramZone,omitempty"`
	IsThisWeekendJourney    bool             `json:"IsThisWeekendJourney,omitempty"`
	ZoneInfo                ZoneInfo         `json:"ZoneInfo"`
	PassengerFares          []PassengerFares `json:"PassengerFares,omitempty"`
}

type Outlets struct {
	OutletSlidSpid         string `json:"outlet_slid_spid,omitempty"`
	OutletName             string `json:"outlet_name,omitempty"`
	OutletBusiness         string `json:"outlet_business,omitempty"`
	OutletLatitude         int    `json:"outlet_latitude,omitempty"`
	OutletLongitude        int    `json:"outlet_longitude,omitempty"`
	OutletSuburb           string `json:"outlet_suburb,omitempty"`
	OutletPostcode         int    `json:"outlet_postcode,omitempty"`
	OutletBusinessHourMon  string `json:"outlet_business_hour_mon,omitempty"`
	OutletBusinessHourTue  string `json:"outlet_business_hour_tue,omitempty"`
	OutletBusinessHourWed  string `json:"outlet_business_hour_wed,omitempty"`
	OutletBusinessHourThur string `json:"outlet_business_hour_thur,omitempty"`
	OutletBusinessHourFri  string `json:"outlet_business_hour_fri,omitempty"`
	OutletBusinessHourSat  string `json:"outlet_business_hour_sat,omitempty"`
	OutletBusinessHourSun  string `json:"outlet_business_hour_sun,omitempty"`
	OutletNotes            string `json:"outlet_notes,omitempty"`
}

type SkippedStops struct {
	StopDistance  int    `json:"stop_distance,omitempty"`
	StopSuburb    string `json:"stop_suburb,omitempty"`
	StopName      string `json:"stop_name,omitempty"`
	StopID        int    `json:"stop_id,omitempty"`
	RouteType     int    `json:"route_type,omitempty"`
	StopLatitude  int    `json:"stop_latitude,omitempty"`
	StopLongitude int    `json:"stop_longitude,omitempty"`
	StopLandmark  string `json:"stop_landmark,omitempty"`
	StopSequence  int    `json:"stop_sequence,omitempty"`
}

type RouteServiceStatus struct {
	Description string    `json:"description,omitempty"`
	Timestamp   time.Time `json:"timestamp"`
}

type Geopath struct{}

type Route struct {
	RouteServiceStatus RouteServiceStatus `json:"route_service_status"`
	RouteType          int                `json:"route_type,omitempty"`
	RouteID            int                `json:"route_id,omitempty"`
	RouteName          string             `json:"route_name,omitempty"`
	RouteNumber        string             `json:"route_number,omitempty"`
	RouteGtfsID        string             `json:"route_gtfs_id,omitempty"`
	Geopath            []Geopath          `json:"geopath,omitempty"`
}

type RouteTypes struct {
	RouteTypeName string `json:"route_type_name,omitempty"`
	RouteType     int    `json:"route_type,omitempty"`
}

type VehiclePosition struct {
	Latitude    int       `json:"latitude,omitempty"`
	Longitude   int       `json:"longitude,omitempty"`
	Easting     int       `json:"easting,omitempty"`
	Northing    int       `json:"northing,omitempty"`
	Direction   string    `json:"direction,omitempty"`
	Bearing     int       `json:"bearing,omitempty"`
	Supplier    string    `json:"supplier,omitempty"`
	DatetimeUtc time.Time `json:"datetime_utc"`
	ExpiryTime  time.Time `json:"expiry_time"`
}

type VehicleDescriptor struct {
	Operator       string `json:"operator,omitempty"`
	ID             string `json:"id,omitempty"`
	LowFloor       bool   `json:"low_floor,omitempty"`
	AirConditioned bool   `json:"air_conditioned,omitempty"`
	Description    string `json:"description,omitempty"`
	Supplier       string `json:"supplier,omitempty"`
	Length         string `json:"length,omitempty"`
}

type Feeder struct {
	RunRef          string `json:"run_ref,omitempty"`
	RouteID         int    `json:"route_id,omitempty"`
	StopID          int    `json:"stop_id,omitempty"`
	Advertised      bool   `json:"advertised,omitempty"`
	DirectionID     int    `json:"direction_id,omitempty"`
	DestinationName string `json:"destination_name,omitempty"`
}

type Distributor struct {
	RunRef          string `json:"run_ref,omitempty"`
	RouteID         int    `json:"route_id,omitempty"`
	StopID          int    `json:"stop_id,omitempty"`
	Advertised      bool   `json:"advertised,omitempty"`
	DirectionID     int    `json:"direction_id,omitempty"`
	DestinationName string `json:"destination_name,omitempty"`
}

type Interchange struct {
	Feeder      Feeder      `json:"feeder"`
	Distributor Distributor `json:"distributor"`
}

type Runs struct {
	RunID             int               `json:"run_id,omitempty"`
	RunRef            string            `json:"run_ref,omitempty"`
	RouteID           int               `json:"route_id,omitempty"`
	RouteType         int               `json:"route_type,omitempty"`
	FinalStopID       int               `json:"final_stop_id,omitempty"`
	DestinationName   string            `json:"destination_name,omitempty"`
	Status            string            `json:"status,omitempty"`
	DirectionID       int               `json:"direction_id,omitempty"`
	RunSequence       int               `json:"run_sequence,omitempty"`
	ExpressStopCount  int               `json:"express_stop_count,omitempty"`
	VehiclePosition   VehiclePosition   `json:"vehicle_position"`
	VehicleDescriptor VehicleDescriptor `json:"vehicle_descriptor"`
	Geopath           []Geopath         `json:"geopath,omitempty"`
	Interchange       Interchange       `json:"interchange"`
	RunNote           string            `json:"run_note,omitempty"`
	ExternalService   int               `json:"externalService,omitempty"`
}

type Gps struct {
	Latitude  int `json:"latitude,omitempty"`
	Longitude int `json:"longitude,omitempty"`
}

type StopLocation struct {
	Gps Gps `json:"gps"`
}

type StopAmenities struct {
	Toilet     bool   `json:"toilet,omitempty"`
	TaxiRank   bool   `json:"taxi_rank,omitempty"`
	CarParking string `json:"car_parking,omitempty"`
	Cctv       bool   `json:"cctv,omitempty"`
}

type Wheelchair struct {
	AccessibleRamp         bool `json:"accessible_ramp,omitempty"`
	Parking                bool `json:"parking,omitempty"`
	Telephone              bool `json:"telephone,omitempty"`
	Toilet                 bool `json:"toilet,omitempty"`
	LowTicketCounter       bool `json:"low_ticket_counter,omitempty"`
	Manouvering            bool `json:"manouvering,omitempty"`
	RaisedPlatform         bool `json:"raised_platform,omitempty"`
	Ramp                   bool `json:"ramp,omitempty"`
	SecondaryPath          bool `json:"secondary_path,omitempty"`
	RaisedPlatformShelther bool `json:"raised_platform_shelther,omitempty"`
	SteepRamp              bool `json:"steep_ramp,omitempty"`
}

type StopAccessibility struct {
	Lighting                      bool       `json:"lighting,omitempty"`
	PlatformNumber                int        `json:"platform_number,omitempty"`
	AudioCustomerInformation      bool       `json:"audio_customer_information,omitempty"`
	Escalator                     bool       `json:"escalator,omitempty"`
	HearingLoop                   bool       `json:"hearing_loop,omitempty"`
	Lift                          bool       `json:"lift,omitempty"`
	Stairs                        bool       `json:"stairs,omitempty"`
	StopAccessible                bool       `json:"stop_accessible,omitempty"`
	TactileGroundSurfaceIndicator bool       `json:"tactile_ground_surface_indicator,omitempty"`
	WaitingRoom                   bool       `json:"waiting_room,omitempty"`
	Wheelchair                    Wheelchair `json:"wheelchair"`
}

type StopStaffing struct {
	FriAmFrom        string `json:"fri_am_from,omitempty"`
	FriAmTo          string `json:"fri_am_to,omitempty"`
	FriPmFrom        string `json:"fri_pm_from,omitempty"`
	FriPmTo          string `json:"fri_pm_to,omitempty"`
	MonAmFrom        string `json:"mon_am_from,omitempty"`
	MonAmTo          string `json:"mon_am_to,omitempty"`
	MonPmFrom        string `json:"mon_pm_from,omitempty"`
	MonPmTo          string `json:"mon_pm_to,omitempty"`
	PhAdditionalText string `json:"ph_additional_text,omitempty"`
	PhFrom           string `json:"ph_from,omitempty"`
	PhTo             string `json:"ph_to,omitempty"`
	SatAmFrom        string `json:"sat_am_from,omitempty"`
	SatAmTo          string `json:"sat_am_to,omitempty"`
	SatPmFrom        string `json:"sat_pm_from,omitempty"`
	SatPmTo          string `json:"sat_pm_to,omitempty"`
	SunAmFrom        string `json:"sun_am_from,omitempty"`
	SunAmTo          string `json:"sun_am_to,omitempty"`
	SunPmFrom        string `json:"sun_pm_from,omitempty"`
	SunPmTo          string `json:"sun_pm_to,omitempty"`
	ThuAmFrom        string `json:"thu_am_from,omitempty"`
	ThuAmTo          string `json:"thu_am_to,omitempty"`
	ThuPmFrom        string `json:"thu_pm_from,omitempty"`
	ThuPmTo          string `json:"thu_pm_to,omitempty"`
	TueAmFrom        string `json:"tue_am_from,omitempty"`
	TueAmTo          string `json:"tue_am_to,omitempty"`
	TuePmFrom        string `json:"tue_pm_from,omitempty"`
	TuePmTo          string `json:"tue_pm_to,omitempty"`
	WedAmFrom        string `json:"wed_am_from,omitempty"`
	WedAmTo          string `json:"wed_am_to,omitempty"`
	WedPmFrom        string `json:"wed_pm_from,omitempty"`
	WedPmTo          string `json:"wed_pm_To,omitempty"`
}

type Stop struct {
	DisruptionIds      []int             `json:"disruption_ids,omitempty"`
	StationType        string            `json:"station_type,omitempty"`
	StationDescription string            `json:"station_description,omitempty"`
	RouteType          int               `json:"route_type,omitempty"`
	StopLocation       StopLocation      `json:"stop_location"`
	StopAmenities      StopAmenities     `json:"stop_amenities"`
	StopAccessibility  StopAccessibility `json:"stop_accessibility"`
	StopStaffing       StopStaffing      `json:"stop_staffing"`
	Routes             []Routes          `json:"routes,omitempty"`
	StopID             int               `json:"stop_id,omitempty"`
	StopName           string            `json:"stop_name,omitempty"`
	StopLandmark       string            `json:"stop_landmark,omitempty"`
}

type Status struct {
	Version string `json:"version,omitempty"`
	Health  int    `json:"health,omitempty"`
}
