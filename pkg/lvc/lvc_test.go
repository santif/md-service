package lvc

import (
	"md-service/pkg/model"
	"testing"
)

func TestUpdateMarketData(t *testing.T) {
	lvc := NewLVC().(*lvc)

	md := &model.MarketData{
		Symbol: "AAPL",
	}

	lvc.UpdateMarketData(md)

	if lvc.marketData["AAPL"] != md {
		t.Errorf("Expected %v, got %v", md, lvc.marketData["AAPL"])
	}
}

func TestGetMarketData(t *testing.T) {
	lvc := NewLVC().(*lvc)

	md := &model.MarketData{
		Symbol: "AAPL",
	}

	lvc.marketData["AAPL"] = md

	if lvc.GetMarketData("AAPL") != md {
		t.Errorf("Expected %v, got %v", md, lvc.GetMarketData("AAPL"))
	}
}

func TestGetMarketDataNotFound(t *testing.T) {
	lvc := NewLVC().(*lvc)

	if lvc.GetMarketData("AAPL") != nil {
		t.Errorf("Expected nil, got %v", lvc.GetMarketData("AAPL"))
	}
}

func TestGetMarketDataAfterUpdate(t *testing.T) {
	lvc := NewLVC().(*lvc)

	md := &model.MarketData{
		Symbol: "AAPL",
	}

	lvc.UpdateMarketData(md)

	if lvc.GetMarketData("AAPL") != md {
		t.Errorf("Expected %v, got %v", md, lvc.GetMarketData("AAPL"))
	}
}

func TestConcurrentAccessToLVC(t *testing.T) {
	lvc := NewLVC().(*lvc)

	md := &model.MarketData{
		Symbol: "AAPL",
	}

	go func() {
		for i := 0; i < 1000000; i++ {
			lvc.UpdateMarketData(md)
		}
	}()

	go func() {
		for i := 0; i < 1000000; i++ {
			lvc.GetMarketData("AAPL")
		}
	}()

	for i := 0; i < 1000000; i++ {
		lvc.UpdateMarketData(md)
		lvc.GetMarketData("AAPL")
	}
}
