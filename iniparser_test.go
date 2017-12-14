////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-11 22:02:59
// @Copyright © 2017 hIMEI <himei@tuta.io>
// @license MIT
/////////////////////////////////////////////////////
//
//    ██╗███╗   ██╗██╗ ██████╗  ██████╗
//    ██║████╗  ██║██║██╔════╝ ██╔═══██╗
//    ██║██╔██╗ ██║██║██║  ███╗██║   ██║
//    ██║██║╚██╗██║██║██║   ██║██║   ██║
//    ██║██║ ╚████║██║╚██████╔╝╚██████╔╝
//    ╚═╝╚═╝  ╚═══╝╚═╝ ╚═════╝  ╚═════╝
//    ___.ini files parser___
/////////////////////////////////////////////////////

package inigo

import (
	"reflect"
	"testing"
)

func TestErrCheck(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrCheck(tt.args.value); got != tt.want {
				t.Errorf("ErrCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrFatal(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ErrFatal(tt.args.err)
		})
	}
}

func TestNewParser(t *testing.T) {
	tests := []struct {
		name string
		want *IniParser
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIniParser_getParserFunc(t *testing.T) {
	type fields struct {
		Checker   Checker
		Functions map[int]func(string) interface{}
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(value string) interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &IniParser{
				Checker:   tt.fields.Checker,
				Functions: tt.fields.Functions,
			}
			if got := parser.getParserFunc(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IniParser.getParserFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIniParser_Parse(t *testing.T) {
	type fields struct {
		Checker   Checker
		Functions map[int]func(string) interface{}
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &IniParser{
				Checker:   tt.fields.Checker,
				Functions: tt.fields.Functions,
			}
			if got := parser.Parse(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IniParser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkPreff(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkPreff(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkPreff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkDigs(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkDigs(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkDigs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkDot(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkDot(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkDot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkBin(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkBin(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkBin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkBool(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkBool(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkChar(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkChar(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkQuotes(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkQuotes(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkQuotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkColonsComma(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkColonsComma(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkColonsComma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecker_checkString(t *testing.T) {
	type fields struct {
		ENABLED Nbld
	}
	type args struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Checker{
				ENABLED: tt.fields.ENABLED,
			}
			if got := checker.checkString(tt.args.value); got != tt.want {
				t.Errorf("Checker.checkString() = %v, want %v", got, tt.want)
			}
		})
	}
}
