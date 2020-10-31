package main

import (
	"context"
	routeguide "google.golang.org/grpc/examples/route_guide"
	"google.golang.org/protobuf/proto"
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
	panic("implement me")
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
