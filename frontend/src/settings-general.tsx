"use client";

import { CollapsibleCard } from "./components/collapsible-card";
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
import { Eye, EyeOff } from "lucide-react";
import { useState } from "react";

export function GeneralSettings() {
  const { config, updateConfig } = useConfig();

  const [showApiKey, setShowApiKey] = useState(false);

  const toggleShowApiKey = () => {
    setShowApiKey(!showApiKey);
  };

  return (
    <CollapsibleCard
      title="General Settings"
      description="General application configuration"
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
              value={config.apiKey}
              onChange={(event) => {
                updateConfig((config) => {
                  config.apiKey = event.target.value;
                });
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
            value={config.country}
            onValueChange={(value) => {
              updateConfig((c) => {
                c.country = value;
              });
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
