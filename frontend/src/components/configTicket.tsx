"use client";

import { CollapsibleCard } from "./collapsibleCard";
import { CommonFields } from "./configCommon";
import { ConfigField } from "./configField";
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
import type { CommonConfig, TicketConfig } from "@/types/config";
import { isEqual } from "lodash";
import { Save, Trash, X } from "lucide-react";
import { useEffect, useState } from "react";

interface TicketProps {
  ticketConfig: TicketConfig;
  globalConfig: CommonConfig;
  isOpen?: boolean;
  onUpdate: (updatedTicket: TicketConfig) => void;
  onRemove: () => void;
}

export function Ticket({
  ticketConfig,
  globalConfig,
  onUpdate,
  onRemove,
}: TicketProps) {
  const [localTicket, setLocalTicket] = useState<TicketConfig>(ticketConfig);

  // If the canonical ticket changes, reset the local state
  useEffect(() => {
    setLocalTicket(ticketConfig);
  }, [ticketConfig]);

  const hasChanges = !isEqual(ticketConfig, localTicket);

  return (
    <CollapsibleCard
      title={localTicket.event}
      action={
        <div className="flex items-center gap-2">
          {hasChanges && (
            <>
              <Button
                variant="outline"
                size="sm"
                onClick={(e) => {
                  e.stopPropagation();
                  setLocalTicket(ticketConfig);
                }}
              >
                <X className="size-4" />
                Discard
              </Button>

              <Button
                size="sm"
                onClick={(e) => {
                  e.stopPropagation();
                  onUpdate(localTicket);
                }}
              >
                <Save className="size-4" />
                Save
              </Button>
            </>
          )}

          <AlertDialog>
            <AlertDialogTrigger asChild>
              <Button
                variant="destructive"
                size="icon"
                onClick={(e) => e.stopPropagation()}
              >
                <Trash />
              </Button>
            </AlertDialogTrigger>
            <AlertDialogContent>
              <AlertDialogHeader>
                <AlertDialogTitle>Delete Ticket</AlertDialogTitle>
                <AlertDialogDescription>
                  Are you sure you want to delete "{ticketConfig.event}"? This
                  action cannot be undone.
                </AlertDialogDescription>
              </AlertDialogHeader>
              <AlertDialogFooter>
                <AlertDialogCancel>Cancel</AlertDialogCancel>
                <AlertDialogAction
                  className="bg-destructive text-destructive-foreground hover:bg-destructive/90"
                  onClick={onRemove}
                >
                  Delete
                </AlertDialogAction>
              </AlertDialogFooter>
            </AlertDialogContent>
          </AlertDialog>
        </div>
      }
    >
      <div className="space-y-4">
        <ConfigField
          label="Event Name"
          description="Name of the event to search for"
          placeholder="Enter event name..."
          type="text"
          value={localTicket.event}
          showReset={false}
          updateValue={(value) => {
            if (typeof value === "string") {
              setLocalTicket((current) => ({ ...current, event: value }));
            }
          }}
        />

        <CommonFields
          commonConfig={localTicket}
          globalCommonConfig={globalConfig}
          updateCommonConfig={(commonConfig) => {
            setLocalTicket((current) => ({ ...current, ...commonConfig }));
          }}
        />
      </div>
    </CollapsibleCard>
  );
}
