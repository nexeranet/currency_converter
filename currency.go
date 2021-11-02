package currency_converter

import (
	"github.com/nexeranet/currency_converter/pkg/whattomine"
)

type Converter struct {
	WhattomineApi *whattomine.WhatToMineApi
}

func NewConverter() *Converter {
	return &Converter{
		WhattomineApi: whattomine.NewWhatToMineApi(),
	}
}

func (c *Converter) Setup() error {
	err := c.WhattomineApi.Setup()
	if err != nil {
		return err
	}
	return nil
}

func (c *Converter) GetNetInfo(tag string) (whattomine.Coin, error) {
	return c.WhattomineApi.GetNetInfo(tag)
}
