"use client";

import { useConfig } from "../providers/config";
import type { Country } from "../types/config";
import { SaveDiscardButtons } from "./buttonsSaveDiscard";
import { CollapsibleCard } from "./cardCollapsible";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import {
  Select,
  SelectContent,
  SelectTrigger,
  SelectValue,
  SelectItem,
} from "./ui/select";
import { isEqual } from "lodash";
import { useState, useEffect } from "react";

export function GeneralSettings() {
  const { config, updateConfig } = useConfig();

  const [draft, setDraft] = useState({
    keysUrl: config.keysUrl,
    country: config.country,
  });

  // If the canonical config changes, reset the draft
  useEffect(() => {
    setDraft({
      keysUrl: config.keysUrl,
      country: config.country,
    });
  }, [config.keysUrl, config.country]);

  const hasChanges = !isEqual(draft, {
    keysUrl: config.keysUrl,
    country: config.country,
  });

  return (
    <CollapsibleCard
      title="General Configuration"
      description="General application configuration"
      action={
        hasChanges && (
          <SaveDiscardButtons
            onSave={() => {
              updateConfig((config) => {
                config.keysUrl = draft.keysUrl;
                config.country = draft.country;
              });
            }}
            onDiscard={() => {
              setDraft({
                keysUrl: config.keysUrl,
                country: config.country,
              });
            }}
          />
        )
      }
    >
      <div className="space-y-4">
        <div className="space-y-2">
          <Label>Keys URL</Label>
          <p className="text-muted-foreground text-sm">
            URL of Twickets keys.json (required) — see the README for details
          </p>
          <Input
            type="text"
            placeholder="Enter your keys URL"
            value={draft.keysUrl}
            onChange={(event) => {
              setDraft((prev) => ({ ...prev, keysUrl: event.target.value }));
            }}
          />
        </div>

        <div className="space-y-2">
          <Label>Country</Label>
          <p className="text-muted-foreground text-sm">
            Currently only GB is supported
          </p>
          <Select
            value={draft.country}
            onValueChange={(value) => {
              setDraft((prev) => ({ ...prev, country: value as Country }));
            }}
          >
            <SelectTrigger>
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="GB">GB (United Kingdom)</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </div>
    </CollapsibleCard>
  );
}
