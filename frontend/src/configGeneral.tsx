"use client";

import { CollapsibleCard } from "./components/collapsibleCard";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { Label } from "./components/ui/label";
import {
  Select,
  SelectContent,
  SelectTrigger,
  SelectValue,
  SelectItem,
} from "./components/ui/select";
import { useConfig } from "./providers/config";
import type { Country } from "./types/config";
import { isEqual } from "lodash";
import { Eye, EyeOff } from "lucide-react";
import { useState, useEffect } from "react";

export function GeneralSettings() {
  const { config, updateConfig } = useConfig();

  const [draft, setDraft] = useState({
    apiKey: config.apiKey,
    country: config.country,
  });
  const [showApiKey, setShowApiKey] = useState(false);

  // If the canonical config changes, reset the draft
  useEffect(() => {
    setDraft({ apiKey: config.apiKey, country: config.country });
  }, [config.apiKey, config.country]);

  const hasChanges = !isEqual(draft, {
    apiKey: config.apiKey,
    country: config.country,
  });

  const toggleShowApiKey = () => {
    setShowApiKey(!showApiKey);
  };

  return (
    <CollapsibleCard
      title="General Settings"
      description="General application configuration"
      action={
        hasChanges && (
          <div className="flex items-center gap-2">
            <Button
              variant="outline"
              size="sm"
              onClick={() =>
                setDraft({ apiKey: config.apiKey, country: config.country })
              }
            >
              Discard
            </Button>
            <Button
              size="sm"
              onClick={() => {
                updateConfig((config) => {
                  config.apiKey = draft.apiKey;
                  config.country = draft.country;
                });
              }}
            >
              Save
            </Button>
          </div>
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
      </div>
    </CollapsibleCard>
  );
}
