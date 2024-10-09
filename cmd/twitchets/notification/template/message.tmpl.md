{{ if ne .Event "" -}}
**{{ .Event }}**
{{- end }}

{{ .Venue }}, {{ .Location }}
{{ .Date }} {{ .Time }}

{{ .NumTickets }} ticket(s) - {{ .TicketType }}
Ticket Price: {{ .TotalTicketPrice }}
Total Price: {{ .TotalPrice }}
{{ if lt .Discount 0.0 -}}
Discount: None
{{- else -}}
Discount: {{ printf "%.2f" .Discount }}%
{{- end }}

Original Ticket Price: {{ .OriginalTicketPrice }}
Original Total Price: {{ .OriginalTotalPrice }}

{{ if ne .Link "" -}}
[Buy Link]({{ .Link }})
{{- end }}
