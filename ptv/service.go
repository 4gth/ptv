package ptv

import (
	"context"
	"fmt"

	"github.com/4gth/ptv/auth"
	"github.com/4gth/ptv/client"
	"github.com/4gth/ptv/model"
)

const (
	defaultHost   = "timetableapi.ptv.vic.gov.au"
	defaultScheme = "https"
)

type Option func(*Service)

type Service struct {
	c          *client.Client
	authWriter auth.AuthWriter
	host       string
	scheme     string
	basePath   string
}

func WithHost(host string) Option         { return func(s *Service) { s.host = host } }
func WithScheme(scheme string) Option     { return func(s *Service) { s.scheme = scheme } }
func WithBasePath(basePath string) Option { return func(s *Service) { s.basePath = basePath } }
func WithHTTPClient(hc *client.Client) Option {
	return func(s *Service) {
		if hc != nil {
			s.c = hc
		}
	}
}

func New(devid, apiKey string, opts ...Option) *Service {
	aw := auth.NewAuthWriter(&auth.Auth{DevID: devid, APIKey: apiKey})
	s := &Service{
		c:          client.NewClient(),
		authWriter: aw,
		host:       defaultHost,
		scheme:     defaultScheme,
	}
	for _, o := range opts {
		o(s)
	}
	s.c.SetDefaults(s.host, s.basePath, s.scheme, s.authWriter)
	return s
}

func NewFromEnv(opts ...Option) *Service {
	a := auth.NewAuthWriter(nil)
	if a == nil {
		panic("auth writer initialization failed")
	}
	s := &Service{
		c:          client.NewClient(),
		authWriter: a,
		host:       defaultHost,
		scheme:     defaultScheme,
	}
	for _, o := range opts {
		o(s)
	}
	s.c.SetDefaults(s.host, s.basePath, s.scheme, s.authWriter)
	return s
}

func do[T any, Q any](s *Service, ctx context.Context, req *model.Request[T, Q], set func(*Q)) (T, error) {
	var zero T
	if req == nil {
		return zero, fmt.Errorf("nil request")
	}
	if set != nil {
		set(&req.Parameters)
	}
	s.c.SetQuery(req.Path, req.Parameters)
	b, err := s.c.Get()
	if err != nil {
		return zero, err
	}
	if err := req.UnMarshalPayload(b); err != nil {
		return zero, err
	}
	return req.Payload, nil
}

// Departures

func (s *Service) DeparturesByRouteTypeAndStopID(ctx context.Context, p func(*model.DeparturesParameters)) (model.DeparturePayload, error) {
	return do(s, ctx, model.NewRequest(model.DeparturesByRouteTypeAndStopIDRequest{}), p)
}

func (s *Service) DeparturesByRouteTypeStopIDAndRouteID(ctx context.Context, p func(*model.DeparturesParametersByRoute)) (model.DeparturePayload, error) {
	return do(s, ctx, model.NewRequest(model.DeparturesByRouteTypeAndStopIDAndRouteID{}), p)
}

// Directions

func (s *Service) DirectionsByRouteID(ctx context.Context, p func(*model.DirectionsParametersByRoute)) (model.DirectionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DirectionsByRouteID{}), p)
}

func (s *Service) DirectionsByDirectionID(ctx context.Context, p func(*model.DirectionsParametersByDirection)) (model.DirectionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DirectionsByDirectionID{}), p)
}

func (s *Service) DirectionsByDirectionIDAndRouteType(ctx context.Context, p func(*model.DirectionsParametersByDirectionAndRouteType)) (model.DirectionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DirectionsByRouteIDAndRouteType{}), p)
}

// Disruptions

func (s *Service) Disruptions(ctx context.Context, p func(*model.DisruptionsParameters)) (model.DisruptionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.Disruption{}), p)
}

func (s *Service) DisruptionsByRouteID(ctx context.Context, p func(*model.DisruptionsParametersByRouteID)) (model.DisruptionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DisruptionsByRouteID{}), p)
}

func (s *Service) DisruptionsByRouteIDAndStopID(ctx context.Context, p func(*model.DisruptionsParametersByRouteIDAndStopID)) (model.DisruptionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DisruptionsByRouteIDAndStopID{}), p)
}

func (s *Service) DisruptionsByStopID(ctx context.Context, p func(*model.DisruptionsParametersByStopID)) (model.DisruptionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DisruptionsByStopID{}), p)
}

func (s *Service) DisruptionByDisruptionID(ctx context.Context, p func(*model.DisruptionsParametersByDisruptionID)) (model.DisruptionsPayload, error) {
	return do(s, ctx, model.NewRequest(model.DisruptionByDisruptionID{}), p)
}

func (s *Service) DisruptionModes(ctx context.Context, p func(*model.DisruptionsParameters)) (model.DisruptionModesPayload, error) {
	return do(s, ctx, model.NewRequest(model.DisruptionByModes{}), p)
}

// Fare estimate

func (s *Service) FareEstimate(ctx context.Context, p func(*model.FairEstimateParameters)) (model.FairEstimatePayload, error) {
	return do(s, ctx, model.NewRequest(model.FairEstimate{}), p)
}

// Outlets

func (s *Service) Outlets(ctx context.Context, p func(*model.OutletsParameters)) (model.OutletsPayload, error) {
	return do(s, ctx, model.NewRequest(model.Outlet{}), p)
}

func (s *Service) OutletsByLatLng(ctx context.Context, p func(*model.OutletsParametersByLatLng)) (model.OutletsPayload, error) {
	return do(s, ctx, model.NewRequest(model.OutletsByLatLng{}), p)
}

// Patterns

func (s *Service) PatternByRunRefAndRouteType(ctx context.Context, p func(*model.PatternsParameters)) (model.PatternsPayload, error) {
	return do(s, ctx, model.NewRequest(model.PatternByRunRefAndRouteType{}), p)
}

// Routes

func (s *Service) Routes(ctx context.Context, p func(*model.RoutesParameters)) (model.RoutePayload, error) {
	return do(s, ctx, model.NewRequest(model.RoutesRequest{}), p)
}

func (s *Service) RoutesByRouteID(ctx context.Context, p func(*model.RoutesParametersByRouteID)) (model.RouteIDPayload, error) {
	return do(s, ctx, model.NewRequest(model.RoutesByRouteID{}), p)
}

// Runs

func (s *Service) RunsByRouteID(ctx context.Context, p func(*model.RunsParametersByRouteID)) (model.RunsPayload, error) {
	return do(s, ctx, model.NewRequest(model.RunsByRouteID{}), p)
}

func (s *Service) RunsByRouteIDAndRouteType(ctx context.Context, p func(*model.RunsParametersByRouteIDAndRouteType)) (model.RunsPayload, error) {
	return do(s, ctx, model.NewRequest(model.RunsByRouteIDAndRouteType{}), p)
}

func (s *Service) RunsByRunRef(ctx context.Context, p func(*model.RunsParametersByRunRef)) (model.RunsPayload, error) {
	return do(s, ctx, model.NewRequest(model.RunsByRunRef{}), p)
}

func (s *Service) RunsByRunRefAndRouteType(ctx context.Context, p func(*model.RunsParametersByRunRefAndRouteType)) (model.RunsPayload, error) {
	return do(s, ctx, model.NewRequest(model.RunsByRunRefAndRouteType{}), p)
}

// Search

func (s *Service) Search(ctx context.Context, p func(*model.SearchParameters)) (model.SearchPayload, error) {
	return do(s, ctx, model.NewRequest(model.SearchTerm{}), p)
}

// Stops

func (s *Service) StopsByStopIDAndRouteType(ctx context.Context, p func(*model.StopParametersByStopIDAndRouteType)) (model.StopsPayload, error) {
	return do(s, ctx, model.NewRequest(model.StopsByStopIDAndRouteType{}), p)
}

func (s *Service) StopsByRouteIDAndRouteType(ctx context.Context, p func(*model.StopParametersByRouteIDAndRouteType)) (model.StopsPayload, error) {
	return do(s, ctx, model.NewRequest(model.StopsByRouteIDAndRouteType{}), p)
}

func (s *Service) StopsByLatLng(ctx context.Context, p func(*model.StopParametersByLatLng)) (model.StopsPayload, error) {
	return do(s, ctx, model.NewRequest(model.StopsByLatLng{}), p)
}
