"use client";

import { SaveDiscardButtons } from "./buttonSaveDiscard";
import { CollapsibleCard } from "./cardCollapsible";
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
  const [draft, setDraft] = useState<TicketConfig>(ticketConfig);

  useEffect(() => {
    setDraft(ticketConfig);
  }, [ticketConfig]);

  const hasChanges = !isEqual(ticketConfig, draft);

  return (
    <CollapsibleCard
      title={draft.event}
      action={
        <div className="flex items-center gap-2">
          {hasChanges && (
            <SaveDiscardButtons
              onSave={() => {
                setDraft(ticketConfig);
              }}
              onDiscard={() => {
                onUpdate(draft);
              }}
            />
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
          value={draft.event}
          showReset={false}
          updateValue={(value) => {
            if (typeof value === "string") {
              setDraft((prev) => ({ ...prev, event: value }));
            }
          }}
        />

        <CommonFields
          commonConfig={draft}
          globalCommonConfig={globalConfig}
          updateCommonConfig={(commonConfig) => {
            setDraft((prev) => ({ ...prev, ...commonConfig }));
          }}
        />
      </div>
    </CollapsibleCard>
  );
}
