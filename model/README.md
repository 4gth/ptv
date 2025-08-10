# `model` package

Typed request builders, parameter structs, and response payload types for PTV API endpoints.

## Core types

- `type Request struct { Path string; Parameters any; Payload any }`
- `type Requester interface { New() *Request }`
- `func NewRequest(funcType Requester) *Request`

Usage pattern:

```go
req := model.NewRequest(model.RoutesRequest{})
// Optionally set parameter values on req.Parameters (cast to the right type)
// Pass req.Path and req.Parameters to the HTTP client, then decode into req.Payload
```

## Endpoints

- `RoutesRequest` → `/v3/routes`
- `RoutesByRouteID` → `/v3/routes/{route_id}`
- `StopsByStopIDAndRouteType` → `/v3/stops/{stop_id}/route_type/{route_type}`
- `StopsByRouteIDAndRouteType` → `/v3/stops/route/{route_id}/route_type/{route_type}`
- `StopsByLatLng` → `/v3/stops/location/{latitude},{longitude}`
- `DeparturesByRouteTypeAndStopIDRequest` → `/v3/departures/route_type/{route_type}/stop/{stop_id}`
- `DeparturesByRouteTypeAndStopIDAndRouteID` → `/v3/departures/route_type/{route_type}/stop/{stop_id}/route/{route_id}`
- `DirectionsByRouteID` → `/v3/directions/route/{route_id}`
- `DirectionsByDirectionID` → `/v3/directions/{direction_id}`
- `DirectionsByRouteIDAndRouteType` → `/v3/directions/{direction_id}/route_type/{route_type}`
- `RunsByRouteID` → `/v3/runs/route/{route_id}`
- `RunsByRouteIDAndRouteType` → `/v3/runs/route/{route_id}/route_type/{route_type}`
- `RunsByRunRef` → `/v3/runs/{run_ref}`
- `RunsByRunRefAndRouteType` → `/v3/runs/{run_ref}/route_type/{route_type}`
- `PatternByRunRefAndRouteType` → `/v3/pattern/run/{run_ref}/route_type/{route_type}`
- `RouteType` → `/v3/route_types`
- `SearchTerm` → `/v3/search/{search_term}`
- `Outlet` → `/v3/outlets`
- `OutletsByLatLng` → `/v3/outlets/location/{latitude},{longitude}`
- `Disruption` → `/v3/disruptions`
- `DisruptionsByRouteID` → `/v3/disruptions/route/{route_id}`
- `DisruptionsByRouteIDAndStopID` → `/v3/disruptions/route/{route_id}/stop/{stop_id}`
- `DisruptionsByStopID` → `/v3/disruptions/stop/{stop_id}`
- `DisruptionByDisruptionID` → `/v3/disruptions/{disruption_id}`
- `DisruptionByModes` → `/v3/disruptions/modes`
- `FairEstimate` → `/v3/fare_estimate/min_zone/{min_zone}/max_zone/{max_zone}`

## Parameters and tags

- Path placeholders correspond to struct fields tagged with `path:"..."`.
- Query parameters come from struct fields tagged with `query:"..."`.
- See the source files for full field lists for each endpoint.