# Changelog - inigo

### 0.1.0

__Changes__

- parsing of int64 in binary excluded now

  doc.go rewrited

  Comments added to all functions and types. Wrong comments fixed

  README.md template created

  Inifile.getSectionParams() is unexported now

  API methods are:

  func (i Inifile) GetParamString(secname, paramname string) string

  func (i Inifile) GetParamsEnabled(secname string) []string

  func (i Inifile) GetParamsDisabled(secname string) []string

  func (i Inifile) GetAllParams() []string

  func (i Inifile) GetParamValue(secname, paramname string) interface{}

  func (i Inifile) FindParam(paramname string) (string, interface{})

  func (i Inifile) GetErrors() []string

  Current version is v0.1.0alpha

  Tests needed


__Contributors__

- hIMEI

Released by hIMEI, Sun 28 Jan 2018 -
[see the diff](https://github.com/hIMEI29A/inigo/compare/a3fcca44d3d4c7b375acd9899ed615d653330a75...0.1.0#diff)
______________


