package main

import (
	"context"
	routeguide "google.golang.org/grpc/examples/route_guide"
	"google.golang.org/protobuf/proto"
	"math"
)

type routeGuideServer struct {
	savedFeatures []*routeguide.Feature
}

func (r *routeGuideServer) GetFeature(ctx context.Context, point *routeguide.Point) (*routeguide.Feature, error) {
	for _, feature := range r.savedFeatures {
		if proto.Equal(feature.Location, point) {
			return feature, nil
		}
	}
	return &routeguide.Feature{Location: point}, nil
}

func (r *routeGuideServer) ListFeatures(rectangle *routeguide.Rectangle, server routeguide.RouteGuide_ListFeaturesServer) error {
	for _, feature := range r.savedFeatures {
		if inRange(feature.Location, rectangle) {
			if err := server.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *routeGuideServer) RecordRoute(server routeguide.RouteGuide_RecordRouteServer) error {
	panic("implement me")
}

func (r *routeGuideServer) RouteChat(server routeguide.RouteGuide_RouteChatServer) error {
	panic("implement me")
}

func (r *routeGuideServer) mustEmbedUnimplementedRouteGuideServer() {
	panic("implement me")
}

func inRange(point *routeguide.Point, rect *routeguide.Rectangle) bool {
	left := math.Min(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	right := math.Max(float64(rect.Lo.Longitude), float64(rect.Hi.Longitude))
	top := math.Max(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))
	bottom := math.Min(float64(rect.Lo.Latitude), float64(rect.Hi.Latitude))

	if float64(point.Longitude) >= left &&
		float64(point.Longitude) <= right &&
		float64(point.Latitude) >= bottom &&
		float64(point.Latitude) <= top {
		return true
	}
	return false
}
