"use client";

import { CollapsibleCard } from "./components/collapsibleCard";
import { CommonSettings } from "./components/configCommonFields";
import { SettingField } from "./components/configField";
import { useConfig } from "./providers/config";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog";
import { Button } from "@/components/ui/button";
import type { TicketConfig } from "@/types/config";
import { Plus, Trash } from "lucide-react";
import type { MouseEvent } from "react";

export function TicketSettings() {
  const { config, setConfig } = useConfig();

  const handleAddTicket = (e: MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    setConfig((config) => {
      const newTickets = [...config.tickets];
      newTickets.push({
        event: `New Event ${config.tickets.length + 1}`,
      });
      return { ...config, tickets: newTickets };
    });
  };

  const handleTicketUpdate = (
    ticketUpdater: (ticket: TicketConfig) => TicketConfig,
    ticketIndex: number,
  ) => {
    setConfig((config) => {
      const newTickets = [...config.tickets];
      newTickets[ticketIndex] = ticketUpdater(newTickets[ticketIndex]);
      return { ...config, tickets: newTickets };
    });
  };

  const handleRemoveTicket =
    (index: number) => (e: MouseEvent<HTMLElement>) => {
      e.stopPropagation();
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
        <Button onClick={handleAddTicket}>
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
              <CollapsibleCard
                key={ticket.event}
                title={ticket.event || "New Event"}
                action={
                  <AlertDialog>
                    <AlertDialogTrigger asChild>
                      <Button variant="destructive" size="icon">
                        <Trash />
                      </Button>
                    </AlertDialogTrigger>
                    <AlertDialogContent>
                      <AlertDialogHeader>
                        <AlertDialogTitle>Delete Ticket</AlertDialogTitle>
                        <AlertDialogDescription>
                          Are you sure you want to delete "{ticket.event}"? This
                          action cannot be undone.
                        </AlertDialogDescription>
                      </AlertDialogHeader>
                      <AlertDialogFooter>
                        <AlertDialogCancel>Cancel</AlertDialogCancel>
                        <AlertDialogAction
                          className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
                          onClick={handleRemoveTicket(ticketIndex)}
                        >
                          Delete
                        </AlertDialogAction>
                      </AlertDialogFooter>
                    </AlertDialogContent>
                  </AlertDialog>
                }
              >
                <div className="space-y-4">
                  <SettingField
                    label="Event Name"
                    description="Name of the event to search for"
                    placeholder="Enter event name..."
                    type="text"
                    value={ticket.event}
                    resetValue={""}
                    updateValue={(value) => {
                      if (!value) return;
                      handleTicketUpdate(
                        (ticket) => ({
                          ...ticket,
                          event: value,
                        }),
                        ticketIndex,
                      );
                    }}
                  />
                  <CommonSettings
                    commonConfig={ticket}
                    globalCommonConfig={config.global}
                    updateCommonConfig={(commonConfig) => {
                      handleTicketUpdate(
                        (ticket) => ({ ...ticket, ...commonConfig }),
                        ticketIndex,
                      );
                    }}
                  />
                </div>
              </CollapsibleCard>
            );
          })
        )}
      </div>
    </CollapsibleCard>
  );
}
