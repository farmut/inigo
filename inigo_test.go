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

func TestInifile_GetSectionParams(t *testing.T) {
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
			if got := i.GetSectionParams(tt.args.secname); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Inifile.GetSectionParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInifile_GetParamByName(t *testing.T) {
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
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Inifile{
				Sections: tt.fields.Sections,
			}
			if got := i.GetParamByName(tt.args.secname, tt.args.paramname); got != tt.want {
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
