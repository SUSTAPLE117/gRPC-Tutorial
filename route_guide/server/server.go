package main

import (
	"context"
	routeguide "google.golang.org/grpc/examples/route_guide"
)

type routeGuideServer struct {
}

func (r *routeGuideServer) GetFeature(ctx context.Context, point *routeguide.Point) (*routeguide.Feature, error) {
	panic("implement me")
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
