/////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-14 13:51:45
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

func Test_newFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFromFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeString(t *testing.T) {
	type args struct {
		newIni []string
		junk   string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeString(tt.args.newIni, tt.args.junk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findIndexByName(t *testing.T) {
	type args struct {
		names []string
		name  string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndexByName(tt.args.names, tt.args.name); got != tt.want {
				t.Errorf("findIndexByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseParamName(t *testing.T) {
	type args struct {
		paramname string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseParamName(tt.args.paramname); got != tt.want {
				t.Errorf("parseParamName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSections(t *testing.T) {
	type args struct {
		clearedIni []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSections(tt.args.clearedIni); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitParamString(t *testing.T) {
	type args struct {
		paramString string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitParamString(tt.args.paramString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitParamString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinParamStrings(t *testing.T) {
	type args struct {
		creds []string
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
			if got := joinParamStrings(tt.args.creds); got != tt.want {
				t.Errorf("joinParamStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeLastSpace(t *testing.T) {
	type args struct {
		paramname string
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
			if got := removeLastSpace(tt.args.paramname); got != tt.want {
				t.Errorf("removeLastSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkDubls(t *testing.T) {
	type args struct {
		paramnames []string
		paramname  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkDubls(tt.args.paramnames, tt.args.paramname); got != tt.want {
				t.Errorf("checkDubls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_renameDubls(t *testing.T) {
	type args struct {
		paramnames []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := renameDubls(tt.args.paramnames); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("renameDubls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkForEqual(t *testing.T) {
	type args struct {
		paramstring string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkForEqual(tt.args.paramstring); got != tt.want {
				t.Errorf("checkForEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkSpaceComma(t *testing.T) {
	type args struct {
		stringvalue string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkSpaceComma(tt.args.stringvalue); got != tt.want {
				t.Errorf("checkSpaceComma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_errorConstruct(t *testing.T) {
	type args struct {
		body []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorConstruct(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("errorConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_paramsConstruct(t *testing.T) {
	type args struct {
		body []string
	}
	tests := []struct {
		name string
		args args
		want *Params
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := paramsConstruct(tt.args.body); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("paramsConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sectionsConstruct(t *testing.T) {
	type args struct {
		sectionsMap map[string][]string
	}
	tests := []struct {
		name string
		args args
		want *Sections
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sectionsConstruct(tt.args.sectionsMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sectionsConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSectionsBodys(t *testing.T) {
	type args struct {
		clearedIni []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSectionsBodys(tt.args.clearedIni); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSectionsBodys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearSectionName(t *testing.T) {
	type args struct {
		sectionName string
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
			if got := clearSectionName(tt.args.sectionName); got != tt.want {
				t.Errorf("clearSectionName() = %v, want %v", got, tt.want)
			}
		})
	}
}
