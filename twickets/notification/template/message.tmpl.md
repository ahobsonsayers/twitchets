{{ .Venue }}, {{ .Location }}
{{ .Date }} {{ .Time }}

Num Tickets: {{ .NumTickets }} ticket(s)
Ticket Price: {{ .TotalTicketPrice }}
Total Price: {{ .TotalPrice }}
{{ if lt .Discount 0.0 }}
Discount: None
{{ else }}
Discount: {{ .Discount }}%
{{ end }}

Original Ticket Price: {{ .OriginalTicketPrice }}
Original Total Price: {{ .OriginalTotalPrice }}
{{ if .Link != "" }}
Buy: {{ .Link }}
{{ end }}