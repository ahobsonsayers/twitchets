"use client";

import { CollapsibleCard } from "./components/collapsible-card";
import { Button } from "./components/ui/button";
import { Input } from "./components/ui/input";
import { Label } from "./components/ui/label";
import { Separator } from "./components/ui/separator";
import { useConfig } from "./providers/config";
import { Eye, EyeOff } from "lucide-react";
import { useState } from "react";

const defaultNtfyUrl = "https://ntfy.sh";

export function NotificationSettings() {
  const { config, updateConfig } = useConfig();

  const [showNtfyPassword, setShowNtfyPassword] = useState(false);

  const toggleShowNtfyPassword = () => {
    setShowNtfyPassword(!showNtfyPassword);
  };

  return (
    <CollapsibleCard title="Notification Settings">
      <div className="space-y-4">
        <div className="space-y-2">
          <div className="flex items-center gap-2">
            <h3>Ntfy</h3>
            {/* <Checkbox />
            <Label>Enabled</Label> */}
          </div>

          <div className="flex gap-2">
            <div className="flex-auto flex-col space-y-2">
              <Label>URL</Label>
              <Input
                type="text"
                placeholder={defaultNtfyUrl}
                value={
                  config.notification.ntfy?.url === defaultNtfyUrl
                    ? ""
                    : config.notification.ntfy?.url
                }
                onChange={(event) => {
                  updateConfig((config) => {
                    config.notification.ntfy.url =
                      event.target.value === ""
                        ? defaultNtfyUrl
                        : event.target.value;
                  });
                }}
              />
            </div>

            <div className="flex-auto flex-col space-y-2">
              <Label>Topic</Label>
              <Input
                type="text"
                value={config.notification.ntfy?.topic}
                onChange={(event) => {
                  updateConfig((config) => {
                    config.notification.ntfy.topic = event.target.value;
                  });
                }}
              />
            </div>
          </div>

          <div className="flex space-x-2">
            <div className="flex flex-1 flex-col space-y-2">
              <Label>Username</Label>
              <Input
                type="text"
                placeholder="Optional"
                value={config.notification.ntfy?.username}
                onChange={(event) => {
                  updateConfig((config) => {
                    config.notification.ntfy.username = event.target.value;
                  });
                }}
              />
            </div>

            <div className="flex flex-1 flex-col space-y-2">
              <Label>Password</Label>
              <div className="flex">
                <Input
                  type={showNtfyPassword ? "text" : "password"}
                  placeholder="Optional"
                  value={config.notification.ntfy?.password}
                  onChange={(event) => {
                    updateConfig((config) => {
                      config.notification.ntfy.password = event.target.value;
                    });
                  }}
                />
                <Button
                  type="button"
                  variant="ghost"
                  onClick={toggleShowNtfyPassword}
                >
                  {showNtfyPassword ? (
                    <EyeOff className="size-5" />
                  ) : (
                    <Eye className="size-5" />
                  )}
                </Button>
              </div>
            </div>
          </div>
        </div>
        <Separator />
        <h2 className="text-muted-foreground">
          Telegram and Gotify notification settings not yet supported in user
          interface
        </h2>
      </div>
    </CollapsibleCard>
  );
}
