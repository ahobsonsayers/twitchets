"use client";

import { CollapsibleCard } from "./components/collapsible-card";
import { CommonSettings } from "./components/common-settings";
import { SettingField } from "./components/setting-field";
import { Label } from "./components/ui/label";
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
import { Plus, Trash } from "lucide-react";
import type { MouseEvent } from "react";

export function TicketSettings() {
  const { config, updateConfig } = useConfig();

  const handleAddTicket = (e: MouseEvent<HTMLElement>) => {
    e.stopPropagation();
    updateConfig((config) => {
      config.tickets.push({
        event: `New Event ${config.tickets.length + 1}`,
      });
    });
  };

  const handleRemoveTicket =
    (index: number) => (e: MouseEvent<HTMLElement>) => {
      e.stopPropagation();
      updateConfig((config) => {
        config.tickets.splice(index, 1);
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
          config.tickets.map((ticket, index) => {
            return (
              <CollapsibleCard
                key={index}
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
                          onClick={handleRemoveTicket(index)}
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
                    onChange={(value) =>
                      updateConfig((config) => {
                        config.tickets[index].event = value;
                      })
                    }
                    onReset={() =>
                      updateConfig((config) => {
                        config.tickets[index].event = undefined;
                      })
                    }
                  />
                  <CommonSettings
                    commonConfig={ticket}
                    globalCommonConfig={undefined}
                    updateCommonConfig={(commonConfig) => {
                      updateConfig((config) => {
                        config.tickets[index] = {
                          ...config.tickets[index],
                          ...commonConfig,
                        };
                      });
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
