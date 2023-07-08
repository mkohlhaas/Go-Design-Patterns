package creational

import (
	"errors"
	"fmt"
)

type Color int

const (
	White Color = iota
	Black
	Blue
)

type Shirt struct {
	Price float32
	SKU   string
	Color Color
}

////////////////
// Prototypes //
////////////////

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

type ItemInfoGetter interface {
	GetInfo() string
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

type ShirtCloner interface {
	GetClone(c Color) (ItemInfoGetter, error)
}

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(c Color) (ItemInfoGetter, error) {
	switch c {
	case White:
		newItem := *whitePrototype // Creates a new object, allocates memory (not just passing pointer)!!!
		return &newItem, nil       // This is the whole trick!!!
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}
