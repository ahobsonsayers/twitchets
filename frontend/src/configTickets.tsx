"use client";

import { CollapsibleCard } from "./components/collapsibleCard";
import { TicketItem } from "./components/configTicket";
import { useConfig } from "./providers/config";
import { Button } from "@/components/ui/button";
import type { TicketConfig } from "@/types/config";
import { Plus } from "lucide-react";
import { useMemo } from "react";

export function TicketSettings() {
  const { config, setConfig } = useConfig();

  // Get tickets (ensuring they are ordered)
  const tickets = useMemo(() => {
    return sortTickets(config.tickets);
  }, [config.tickets]);

  const handleAddTicket = () => {
    const newTickets = [...tickets, { event: "New Event" }];
    setConfig((config) => ({
      ...config,
      tickets: sortTickets(newTickets),
    }));
  };

  const handleUpdateTicket = (updatedTicket: TicketConfig, index: number) => {
    const newTickets = [...tickets, { event: "New Event" }];
    newTickets[index] = updatedTicket;
    setConfig((config) => ({
      ...config,
      tickets: sortTickets(newTickets),
    }));
  };

  const handleRemoveTicket = (index: number) => {
    setConfig((config) => {
      const newTickets = [...tickets];
      newTickets.splice(index, 1);
      return { ...config, tickets: newTickets };
    });
  };

  return (
    <CollapsibleCard
      title="Tickets Configuration"
      description="Individual ticket configurations (overrides global settings)"
      action={
        <Button
          onClick={(e) => {
            e.stopPropagation();
            handleAddTicket();
          }}
        >
          <Plus className="size-5" />
          Add Ticket
        </Button>
      }
    >
      <div className="flex flex-col gap-y-4">
        {config.tickets.length === 0 ? (
          <p className="text-muted-foreground text-center">
            No tickets configured. Click "Add Ticket" to get started.
          </p>
        ) : (
          tickets.map((ticket, ticketIndex) => {
            return (
              <TicketItem
                key={ticketIndex}
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
    </CollapsibleCard>
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
