package haki

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Haki struct {
	fiberApp *fiber.App
	routes   HakiRoutes
	config   HakiConfig
}

type HakiRequestContext struct {
	Request *fiber.Ctx
}

type HakiConfig struct {
	Prefix string
	Name   string
}

type IHaki interface {
	New(config HakiConfig) *Haki
	Cors(config ...cors.Config) *Haki
	Get(path string, handlers ...func(ctx HakiRequestContext) any) *Haki
	Post(path string, handlers ...func(ctx HakiRequestContext) any) *Haki
	Put(path string, handlers ...func(ctx HakiRequestContext) any) *Haki
	Patch(path string, handlers ...func(ctx HakiRequestContext) any) *Haki
	Delete(path string, handlers ...func(ctx HakiRequestContext) any) *Haki
	Use(handler *Haki) *Haki
	Listen(port string)
}

func (h *Haki) Cors(config ...cors.Config) *Haki {
	h.fiberApp.Use(cors.New(config...))

	return h
}

func New(config HakiConfig) *Haki {
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
