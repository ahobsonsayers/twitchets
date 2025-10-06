{{/* Based on the following original template, with custom enum type definition */}}
{{/* http://github.com/oapi-codegen/oapi-codegen/blob/0d94f2f2015e5fec2cd66273eeb8eff7d1f86bc1/pkg/codegen/schema.go */}}
{{ range .Types }}
{{ if .Schema.Description }} {{ toGoComment .Schema.Description .TypeName }} {{ else }}// {{ .TypeName }} defines model for {{ .JsonName }}. {{ end }}
{{ if eq (len .Schema.EnumValues) 0 -}}
type {{ .TypeName }} {{if .IsAlias }}={{ end }} {{ .Schema.TypeDecl }}
{{ else -}}
type {{ .TypeName }} enum.Member[{{ .Schema.TypeDecl }}]
{{ end }}
{{ end }}
