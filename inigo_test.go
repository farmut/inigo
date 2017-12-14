/////////////////////////////////////////////////////
// @Author: hIMEI
// @Date:   2017-12-06 07:15:25
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

func TestNewIniFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want *Inifile
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIniFile(tt.args.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIniFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_IniToJson(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			got, err := i.IniToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("Inifile.IniToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.IniToJson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_GetSectionsNames(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetSectionsNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetSectionsNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_PrintSectionsNames(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			i.PrintSectionsNames()
		})
	}
}

func TestInifile_GetSectionByName(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		secname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Params
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetSectionByName(tt.args.secname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetSectionByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_GetParamByName(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		params    *Params
		paramname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetParamByName(tt.args.params, tt.args.paramname); got != tt.want {
				t.Errorf("Inifile.GetParamByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_GetParamsEnabled(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		secname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetParamsEnabled(tt.args.secname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetParamsEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_PrintParamsEnabled(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		secname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			i.PrintParamsEnabled(tt.args.secname)
		})
	}
}

func TestInifile_GetParamsDisabled(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		secname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetParamsDisabled(tt.args.secname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetParamsDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_PrintParamsDisabled(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		secname string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			i.PrintParamsDisabled(tt.args.secname)
		})
	}
}

func TestInifile_GetAllParams(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetAllParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetAllParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_GetValue(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	type args struct {
		secname   string
		paramname string
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
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetValue(tt.args.secname, tt.args.paramname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_PrintAllParams(t *testing.T) {
	type fields struct {
		Sections *Sections
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			i.PrintAllParams()
		})
	}
}

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
