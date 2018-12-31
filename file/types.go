package file

import "github.com/hbagdi/go-kong/kong"

type service struct {
	kong.Service `yaml:",inline"`
	Routes       []*route
}

type route struct {
	kong.Route `yaml:",inline"`
}

type upstream struct {
	kong.Upstream `yaml:",inline"`
	Targets       []*target
}

type target struct {
	kong.Target `yaml:",inline"`
}

type fileStructure struct {
	Services  []service
	Upstreams []upstream
}
