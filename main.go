package main

import (
	"fmt"
	"time"
)

type LOB interface {
	Product()
}

type LOBMultiproduct struct{}

func (c *LOBMultiproduct) Product() {
	fmt.Println("This is LOB Multiproduct")
}

type LOBDanaSyariah struct{}

func (c *LOBDanaSyariah) Product() {
	fmt.Println("This is LOB Dana Syariah")
}

type LOBNMC struct{}

func (c *LOBNMC) Product() {
	fmt.Println("This is LOB New Motorcycle")
}

type LOBUMC struct{}

func (c *LOBUMC) Product() {
	fmt.Println("This is LOB Use Motorcycle")
}

type LOBType string

const (
	Multiproduct  LOBType = "Every Saturday, Multiproduct Have Many Promo For Customer"
	DanaSyariah   LOBType = "Every Friday, Dana Syariah Have Many Promo For Customer"
	NewMotorcyle  LOBType = "Every Monday, New Motorcycle Have Many Promo For Customer"
	UseMotorcycle LOBType = "Every Tuesday, Use Motorcycle Have Many Promo For Customer"
)

type PromoLOB interface {
	AllProduct(time time.Time) LOB
	Type() LOBType
}

type TypePromoLOB struct{}

func (p *TypePromoLOB) AllProduct(time time.Time) LOB {
	if time.Weekday().String() == "Saturday" {
		return &LOBMultiproduct{}
	} else if time.Weekday().String() == "Friday" {
		return &LOBDanaSyariah{}
	} else if time.Weekday().String() == "Monday" {
		return &LOBNMC{}
	} else if time.Weekday().String() == "Tuesday" {
		return &LOBUMC{}
	} else {
		return nil
	}
}

func (p *TypePromoLOB) Type() LOBType {
	time := time.Now()
	if time.Weekday().String() == "Saturday" {
		return Multiproduct
	} else if time.Weekday().String() == "Friday" {
		return DanaSyariah
	} else if time.Weekday().String() == "Monday" {
		return NewMotorcyle
	} else if time.Weekday().String() == "Tuesday" {
		return UseMotorcycle
	} else {
		return "No Promo For Today"
	}
}

func main() {
	var promoLOB PromoLOB
	promoLOB = &TypePromoLOB{}

	promo := promoLOB.AllProduct(time.Now())
	promo.Product()
	fmt.Println(promoLOB.Type())

}
