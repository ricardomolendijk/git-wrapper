package main

import (
	"fmt"
	"strings"
)

// formatTicket1 formats the ticket list as [#ticket] for each ticket.
func formatTicket1(ticket string) string {
	if ticket == "" {
		return ""
	}

	tickets := strings.Split(ticket, ",")
	var formattedTickets []string
	for _, t := range tickets {
		formattedTickets = append(formattedTickets, fmt.Sprintf("[%s]", strings.ToUpper(strings.TrimSpace(t))))
	}

	return strings.Join(formattedTickets, " ")
}

// formatTicket2 formats the ticket list as - #ticket for each ticket.
func formatTicket2(ticket string) string {
	if ticket == "" {
		return ""
	}

	tickets := strings.Split(ticket, ",")
	var formattedTickets []string
	for _, t := range tickets {
		formattedTickets = append(formattedTickets, fmt.Sprintf("- %s", strings.ToUpper(strings.TrimSpace(t))))
	}

	return strings.Join(formattedTickets, "\n")
}
