{{ if ne .Event "" -}}
*{{ .Event }}*
{{- end }}

{{ .Venue }}, {{ .Location }}
{{ .Date }} {{ .Time }}

{{ .NumTickets }} ticket(s) - {{ .TicketType }}
Ticket Price: {{ .TotalTicketPrice }}
Total Price: {{ .TotalPrice }}
{{ if eq .Discount "0.00%" }}
Discount: None
{{- else -}}
Discount: {{ .Discount }}
{{- end }}

Original Ticket Price: {{ .OriginalTicketPrice }}
Original Total Price: {{ .OriginalTotalPrice }}

{{ if ne .Link "" -}}
[Buy Link]({{ .Link }})
{{- end }}