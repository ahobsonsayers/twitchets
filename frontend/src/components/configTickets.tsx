"use client";

import { useConfig } from "../providers/config";
import { CollapsibleCard } from "./cardCollapsible";
import { Ticket } from "./configTicket";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./ui/card";
import { Button } from "@/components/ui/button";
import type { TicketConfig } from "@/types/config";
import { Plus } from "lucide-react";
import { useMemo, useState } from "react";

export function TicketsConfig() {
  const { config, updateConfig } = useConfig();
  const [isTicketsConfigOpen, setIsTicketsConfigOpen] = useState(false);

  const newEventName = "New Event";

  // Get tickets (ensuring they are ordered)
  const tickets = useMemo(() => {
    return sortTickets(config.tickets);
  }, [config.tickets]);

  const handleAddTicket = () => {
    const newTickets = [...tickets, { event: newEventName }];
    updateConfig((config) => {
      config.tickets = sortTickets(newTickets);
    });
    setIsTicketsConfigOpen(true);
  };

  const handleUpdateTicket = (updatedTicket: TicketConfig, index: number) => {
    const newTickets = [...tickets];
    newTickets[index] = updatedTicket;

    updateConfig((config) => {
      config.tickets = sortTickets(newTickets);
    });
  };

  const handleRemoveTicket = (index: number) => {
    updateConfig((config) => {
      const newTickets = [...config.tickets];
      newTickets.splice(index, 1);
      config.tickets = newTickets;
    });
  };

  return (
    <Card>
      <CardHeader className="flex w-full items-center">
        <div className="flex flex-col">
          <CardTitle>Ticket Configuration</CardTitle>
          <CardDescription>
            Individual ticket configurations (overrides global settings)
          </CardDescription>
        </div>
        <Button
          className="ml-auto"
          onClick={(e) => {
            e.stopPropagation();
            handleAddTicket();
          }}
        >
          <Plus className="size-5" />
          Add Ticket
        </Button>
      </CardHeader>
      <CardContent>
        <div className="flex flex-col gap-y-4">
          {config.tickets.length === 0 ? (
            <p className="text-muted-foreground text-center">
              No tickets configured. Click "Add Ticket" to get started.
            </p>
          ) : (
            tickets.map((ticket, ticketIndex) => {
              return (
                <Ticket
                  key={ticket.event}
                  ticketConfig={ticket}
                  globalConfig={config.global}
                  onUpdate={(updatedTicket) =>
                    handleUpdateTicket(updatedTicket, ticketIndex)
                  }
                  onRemove={() => handleRemoveTicket(ticketIndex)}
                />
              );
            })
          )}
        </div>
      </CardContent>
    </Card>
  );
}

function sortTickets(tickets: TicketConfig[]): TicketConfig[] {
  return [...tickets].sort((a, b) => {
    // "New Event" should always appear first
    if (a.event === "New Event") return -1;
    if (b.event === "New Event") return 1;

    // Otherwise, sort alphabetically
    return a.event.localeCompare(b.event);
  });
}
