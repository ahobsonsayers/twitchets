"use client";

import { CollapsibleCard } from "./components/collapsibleCard";
import { TicketItem } from "./components/configTicket";
import { useConfig } from "./providers/config";
import { Button } from "@/components/ui/button";
import type { TicketConfig } from "@/types/config";
import { Plus } from "lucide-react";

export function TicketSettings() {
  const { config, setConfig } = useConfig();

  const handleAddTicket = () => {
    setConfig((config) => {
      const newTickets = [...config.tickets];
      newTickets.push({
        event: "New Event",
      });
      return { ...config, tickets: newTickets };
    });
  };

  const handleUpdateTicket = (updatedTicket: TicketConfig, index: number) => {
    setConfig((config) => {
      const newTickets = [...config.tickets];
      newTickets[index] = updatedTicket;
      return { ...config, tickets: newTickets };
    });
  };

  const handleRemoveTicket = (index: number) => {
    setConfig((config) => {
      const newTickets = [...config.tickets];
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
          config.tickets.map((ticket, ticketIndex) => {
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
