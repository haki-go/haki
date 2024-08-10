package haki

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Haki struct {
	fiberApp *fiber.App
	routes   HakiRoutes
	config   Config
}

type Context struct {
	Request *fiber.Ctx
}

type Config struct {
	Prefix string
	Name   string
}

type IHaki interface {
	New(config Config) *Haki
	Cors(config ...cors.Config) *Haki
	Get(path string, handlers ...func(ctx Context) any) *Haki
	Post(path string, handlers ...func(ctx Context) any) *Haki
	Put(path string, handlers ...func(ctx Context) any) *Haki
	Patch(path string, handlers ...func(ctx Context) any) *Haki
	Delete(path string, handlers ...func(ctx Context) any) *Haki
	Use(handler *Haki) *Haki
	Listen(port string)
}

func (h *Haki) Cors(config ...cors.Config) *Haki {
	h.fiberApp.Use(cors.New(config...))

	return h
}

func New(config Config) *Haki {
	fiberApp := fiber.New()

	haki := Haki{}
	haki.fiberApp = fiberApp
	haki.config = config
	haki.routes = make(HakiRoutes)

	return &haki
}

func (h *Haki) Listen(port string) {
	h.applyRoutesTree()
	h.fiberApp.Listen(port)
}
