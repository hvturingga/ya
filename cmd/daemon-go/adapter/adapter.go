package adapter

import (
	"context"
	"github.com/hvturingga/ya/ent"
)

const (
	SingBoxName = "sing-box"
	ClashName   = "mihomo"
)

type Adapter interface {
	Runner
	Killer
}

type Runner interface {
	Start()
}

type Killer interface {
	Stop()
}

type SingBox struct {
	ctx  context.Context
	user *ent.User
}

type Clash struct {
	ctx  context.Context
	user *ent.User
}

func NewAdapter(ctx context.Context, user *ent.User) Adapter {
	provider := user.QueryProvider().OnlyX(ctx)

	switch provider.Name {
	case SingBoxName:
		return &SingBox{
			ctx:  ctx,
			user: user,
		}
	case ClashName:
		return &Clash{
			ctx:  ctx,
			user: user,
		}
	default:
		panic("invalid provider name")
	}
}
