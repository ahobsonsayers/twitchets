"use client";

import { useConfig } from "../providers/config";
import type { Country } from "../types/config";
import { SaveDiscardButtons } from "./buttonsSaveDiscard";
import { CollapsibleCard } from "./cardCollapsible";
import { ConfigField } from "./configField";
import { Button } from "./ui/button";
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
import { Eye, EyeOff } from "lucide-react";
import { useState, useEffect } from "react";

export function GeneralSettings() {
  const { config, updateConfig } = useConfig();

  const [draft, setDraft] = useState({
    apiKey: config.apiKey,
    country: config.country,
    flaresolverrUrl: config.flaresolverrUrl,
  });
  const [showApiKey, setShowApiKey] = useState(false);

  // If the canonical config changes, reset the draft
  useEffect(() => {
    setDraft({
      apiKey: config.apiKey,
      country: config.country,
      flaresolverrUrl: config.flaresolverrUrl,
    });
  }, [config.apiKey, config.country, config.flaresolverrUrl]);

  const hasChanges = !isEqual(draft, {
    apiKey: config.apiKey,
    country: config.country,
    flaresolverrUrl: config.flaresolverrUrl,
  });

  const toggleShowApiKey = () => {
    setShowApiKey(!showApiKey);
  };

  return (
    <CollapsibleCard
      title="General Configuration"
      description="General application configuration"
      action={
        hasChanges && (
          <SaveDiscardButtons
            onSave={() => {
              updateConfig((config) => {
                config.apiKey = draft.apiKey;
                config.flaresolverrUrl = draft.flaresolverrUrl;
                config.country = draft.country;
              });
            }}
            onDiscard={() => {
              setDraft({
                apiKey: config.apiKey,
                flaresolverrUrl: config.flaresolverrUrl,
                country: config.country,
              });
            }}
          />
        )
      }
    >
      <div className="space-y-4">
        <div className="space-y-2">
          <Label>API Key</Label>
          <p className="text-muted-foreground text-sm">
            Twickets API key (required)
          </p>
          <div className="flex">
            <Input
              type={showApiKey ? "text" : "password"}
              placeholder="Enter your API key"
              value={draft.apiKey}
              onChange={(event) => {
                setDraft((prev) => ({ ...prev, apiKey: event.target.value }));
              }}
            />
            <Button
              type="button"
              variant="ghost"
              className="ml-auto"
              onClick={toggleShowApiKey}
            >
              {showApiKey ? (
                <EyeOff className="size-5" />
              ) : (
                <Eye className="size-5" />
              )}
            </Button>
          </div>
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

        <ConfigField
          label="FlareSolverr URL"
          description="Optional: URL of FlareSolverr proxy server for bypassing Cloudflare (requires restart)"
          type="text"
          value={draft.flaresolverrUrl}
          showReset={false}
          updateValue={(value) => {
            setDraft((prev) => ({ ...prev, flaresolverrUrl: value }));
          }}
        />
      </div>
    </CollapsibleCard>
  );
}
