package haki

type RouteHandler func(ctx Context) any
type HakiRoutes map[string]map[string][]RouteHandler

func (ctx *Haki) Get(path string, handlers ...RouteHandler) *Haki {
	return processMethod(ctx, path, "GET", handlers)
}

func (ctx *Haki) Post(path string, handlers ...RouteHandler) *Haki {
	return processMethod(ctx, path, "POST", handlers)
}

func (ctx *Haki) Put(path string, handlers ...RouteHandler) *Haki {
	return processMethod(ctx, path, "PUT", handlers)
}

func (ctx *Haki) Patch(path string, handlers ...RouteHandler) *Haki {
	return processMethod(ctx, path, "PATCH", handlers)
}

func (ctx *Haki) Delete(path string, handlers ...RouteHandler) *Haki {
	return processMethod(ctx, path, "DELETE", handlers)
}

func processMethod(h *Haki, route string, method string, handlers []RouteHandler) *Haki {
	if h.routes[route] == nil {
		h.routes[route] = make(map[string][]RouteHandler)
	}

	h.routes[route][method] = handlers

	if len(handlers) == 0 {
		println("No handlers implemented for route " + route)
	}

	return h
}
