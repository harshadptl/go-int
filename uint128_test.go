package goint

import (
	"reflect"
	"testing"
)

func TestNewInt128(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want *Uint128
	}{
		{"zero", args{n: 0}, &Uint128{0, 0}},
		{"1", args{n: 1}, &Uint128{0, 1}},
		{"10", args{n: 10}, &Uint128{0, 10}},
		{"100", args{n: 100}, &Uint128{0, 100}},
		{"1000", args{n: 1000}, &Uint128{0, 1000}},
		{"10000", args{n: 10000}, &Uint128{0, 10000}},
		{"100000", args{n: 100000}, &Uint128{0, 100000}},
		{"1000000", args{n: 1000000}, &Uint128{0, 1000000}},
		{"10000000", args{n: 10000000}, &Uint128{0, 10000000}},
		{"100000000", args{n: 100000000}, &Uint128{0, 100000000}},
		{"1000000000", args{n: 1000000000}, &Uint128{0, 1000000000}},
		{"10000000000", args{n: 10000000000}, &Uint128{0, 10000000000}},
		{"100000000000", args{n: 100000000000}, &Uint128{0, 100000000000}},
		{"1000000000000", args{n: 1000000000000}, &Uint128{0, 1000000000000}},
		{"10000000000000", args{n: 10000000000000}, &Uint128{0, 10000000000000}},
		{"100000000000000", args{n: 100000000000000}, &Uint128{0, 100000000000000}},
		{"1000000000000000", args{n: 1000000000000000}, &Uint128{0, 1000000000000000}},
		{"10000000000000000", args{n: 10000000000000000}, &Uint128{0, 10000000000000000}},
		{"100000000000000000", args{n: 100000000000000000}, &Uint128{0, 100000000000000000}},
		{"1000000000000000000", args{n: 1000000000000000000}, &Uint128{0, 1000000000000000000}},
		{"10000000000000000000", args{n: 10000000000000000000}, &Uint128{0, 10000000000000000000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInt128(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInt128() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func TestUint128_Add(t *testing.T) {
	type fields struct {
		High uint64
		Low  uint64
	}
	type args struct {
		a *Uint128
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Uint128{
				High: tt.fields.High,
				Low:  tt.fields.Low,
			}
			if err := n.Add(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUint128_AddUint64(t *testing.T) {
	type fields struct {
		High uint64
		Low  uint64
	}
	type args struct {
		a uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Uint128{
				High: tt.fields.High,
				Low:  tt.fields.Low,
			}
			if err := n.AddUint64(tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("AddUint64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}


 */