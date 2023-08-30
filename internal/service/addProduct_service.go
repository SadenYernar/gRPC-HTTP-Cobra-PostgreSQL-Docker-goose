package service

import "context"

type IAddProduct interface {
	AddProduct(ctx context.Context)
}
