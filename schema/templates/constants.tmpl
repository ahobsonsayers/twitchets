{{/* Based on the following original template, with custom enum values (as vars instead of consts) */}}
{{/* http://github.com/oapi-codegen/oapi-codegen/blob/0d94f2f2015e5fec2cd66273eeb8eff7d1f86bc1/pkg/codegen/schema.go */}}
{{- if gt (len .SecuritySchemeProviderNames) 0 }}
const (
{{ range $ProviderName := .SecuritySchemeProviderNames }}
	{{- $ProviderName | sanitizeGoIdentity | ucFirst }}Scopes = "{{ $ProviderName }}.Scopes"
{{ end }}
)
{{ end }}
{{ range $Enum := .EnumDefinitions }}
// Defines values for {{ $Enum.TypeName }}.
var (
	{{ $Enum.TypeName | lcFirst }}Builder = enum.NewBuilder[{{$Enum.Schema.TypeDecl}}, {{ $Enum.TypeName }}]()
	
	{{ range $name, $value := $Enum.GetValues -}}
		{{ $Enum.TypeName }}{{ $name }} = {{ $Enum.TypeName | lcFirst }}Builder.Add({{ $Enum.TypeName }}{ {{ $Enum.ValueWrapper }}{{ $value }}{{ $Enum.ValueWrapper }} })
	{{ end }}

	{{ $Enum.TypeName }}s = {{ $Enum.TypeName | lcFirst }}Builder.Enum()
)
{{ end }}
