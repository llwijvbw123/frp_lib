package frpc

import (
	"embed"

	"github.com/Bellegar/frp_lib/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
