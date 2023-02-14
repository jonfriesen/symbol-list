package model

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		name          string
		newSecurities []*Security
		oldSecurities []*Security
		wantAdded     []*Security
		wantRemoved   []*Security
	}{
		{
			name:          "no securities",
			newSecurities: []*Security{},
			oldSecurities: []*Security{},
			wantAdded:     []*Security{},
			wantRemoved:   []*Security{},
		},
		{
			name: "new securities",
			newSecurities: []*Security{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			oldSecurities: []*Security{},
			wantAdded: []*Security{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			wantRemoved: []*Security{},
		},
		{
			name:          "removed securities",
			newSecurities: []*Security{},
			oldSecurities: []*Security{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			wantAdded: []*Security{},
			wantRemoved: []*Security{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
		},
		{
			name: "added and removed securities",
			newSecurities: []*Security{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			oldSecurities: []*Security{
				{Symbol: "GOOG"},
				{Symbol: "FB"},
			},
			wantAdded: []*Security{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			wantRemoved: []*Security{
				{Symbol: "GOOG"},
				{Symbol: "FB"},
			},
		},
		{
			name: "duplicate symbols with different exchanges",
			newSecurities: []*Security{
				{Symbol: "AAPL", Exchange: "EXCH1"},
				{Symbol: "AAPL", Exchange: "EXCH2"},
			},
			oldSecurities: []*Security{
				{Symbol: "AAPL", Exchange: "EXCH2"},
			},
			wantAdded: []*Security{
				{Symbol: "AAPL", Exchange: "EXCH1"},
			},
			wantRemoved: []*Security{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SecurityExportDiff{}
			if err := s.Diff(tt.newSecurities, tt.oldSecurities); err != nil {
				t.Errorf("Diff() error = %v", err)
				return
			}
			if !reflect.DeepEqual(s.Added, tt.wantAdded) {
				t.Errorf("Diff() gotAdded = %v, want %v", s.Added, tt.wantAdded)
			}
			if !reflect.DeepEqual(s.Removed, tt.wantRemoved) {
				t.Errorf("Diff() gotRemoved = %v, want %v", s.Removed, tt.wantRemoved)
			}
		})
	}
}

func TestCryptoDiff(t *testing.T) {
	tests := []struct {
		name        string
		new         []*Crypto
		old         []*Crypto
		wantAdded   []*Crypto
		wantRemoved []*Crypto
	}{
		{
			name:        "no crypto",
			new:         []*Crypto{},
			old:         []*Crypto{},
			wantAdded:   []*Crypto{},
			wantRemoved: []*Crypto{},
		},
		{
			name: "new crypto",
			new: []*Crypto{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			old: []*Crypto{},
			wantAdded: []*Crypto{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			wantRemoved: []*Crypto{},
		},
		{
			name: "removed crypto",
			new:  []*Crypto{},
			old: []*Crypto{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			wantAdded: []*Crypto{},
			wantRemoved: []*Crypto{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
		},
		{
			name: "added and removed crypto",
			new: []*Crypto{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			old: []*Crypto{
				{Symbol: "GOOG"},
				{Symbol: "FB"},
			},
			wantAdded: []*Crypto{
				{Symbol: "AAPL"},
				{Symbol: "MSFT"},
			},
			wantRemoved: []*Crypto{
				{Symbol: "GOOG"},
				{Symbol: "FB"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CryptoExportDiff{}
			if err := s.Diff(tt.new, tt.old); err != nil {
				t.Errorf("Diff() error = %v", err)
				return
			}
			if !reflect.DeepEqual(s.Added, tt.wantAdded) {
				t.Errorf("Diff() gotAdded = %v, want %v", s.Added, tt.wantAdded)
			}
			if !reflect.DeepEqual(s.Removed, tt.wantRemoved) {
				t.Errorf("Diff() gotRemoved = %v, want %v", s.Removed, tt.wantRemoved)
			}
		})
	}
}
