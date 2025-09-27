"use client";

import { CollapsibleCard } from "./components/collapsibleCard";
import { Ticket } from "./components/configTicket";
import { useConfig } from "./providers/config";
import { Button } from "@/components/ui/button";
import type { TicketConfig } from "@/types/config";
import { Plus } from "lucide-react";
import { useMemo, useState } from "react";

export function TicketsConfig() {
  const { config, setConfig } = useConfig();
  const [isTicketsConfigOpen, setIsTicketsConfigOpen] = useState(false);

  const newEventName = "New Event";

  // Get tickets (ensuring they are ordered)
  const tickets = useMemo(() => {
    return sortTickets(config.tickets);
  }, [config.tickets]);

  const handleAddTicket = () => {
    const newTickets = [...tickets, { event: newEventName }];
    setConfig((config) => ({
      ...config,
      tickets: sortTickets(newTickets),
    }));
    setIsTicketsConfigOpen(true);
  };

  const handleUpdateTicket = (updatedTicket: TicketConfig, index: number) => {
    const newTickets = [...tickets];
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
      isOpen={isTicketsConfigOpen}
      setIsOpen={setIsTicketsConfigOpen}
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
              <Ticket
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
