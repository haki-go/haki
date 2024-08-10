package haki

func (h *Haki) Use(handler *Haki) *Haki {
	for route, methods := range handler.routes {
		for method := range methods {
			targetRoute := handler.config.Prefix + route

			if h.routes[targetRoute] == nil {
				h.routes[targetRoute] = make(map[string][]RouteHandler)
			}

			h.routes[targetRoute][method] = handler.routes[route][method]
		}
	}

	return h
}
