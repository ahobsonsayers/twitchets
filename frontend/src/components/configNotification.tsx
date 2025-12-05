"use client";

import { useConfig } from "../providers/config";
import { SaveDiscardButtons } from "./buttonsSaveDiscard";
import { CollapsibleCard } from "./cardCollapsible";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Separator } from "./ui/separator";
import type { NtfyConfig } from "@/types/config";
import { isEqual, omit } from "lodash";
import { Eye, EyeOff } from "lucide-react";
import { useState, useEffect } from "react";

const defaultNtfyUrl = "https://ntfy.sh";

function newNtfyConfig(): NtfyConfig {
  return {
    url: defaultNtfyUrl,
    topic: "",
    username: "",
    password: "",
  };
}

export function NotificationSettings() {
  const { config, updateConfig } = useConfig();

  const [draft, setDraft] = useState(config.notification);
  const [showNtfyPassword, setShowNtfyPassword] = useState(false);

  useEffect(() => {
    setDraft(config.notification);
  }, [config.notification]);

  const hasChanges = !isEqual(draft, config.notification);

  const toggleShowNtfyPassword = () => {
    setShowNtfyPassword(!showNtfyPassword);
  };

  return (
    <CollapsibleCard
      title="Notification Configuration"
      description="Notification services configuration"
      action={
        hasChanges && (
          <SaveDiscardButtons
            onSave={() => {
              updateConfig((config) => {
                config.notification = draft;

                // Add/remove ntfy from global config
                const hasNtfy = !!draft.ntfy;
                if (hasNtfy) {
                  if (!config.global.notification?.includes("ntfy")) {
                    config.global.notification = [
                      ...(config.global.notification || []),
                      "ntfy",
                    ];
                  }
                } else {
                  config.global.notification =
                    config.global.notification?.filter((n) => n !== "ntfy");
                }
              });
            }}
            onDiscard={() => {
              setDraft(config.notification);
            }}
          />
        )
      }
    >
      <div className="space-y-4">
        <div className="space-y-2">
          <div className="flex items-center gap-2">
            <h3>Ntfy</h3>
            <Checkbox
              className="ml-2"
              checked={!!draft.ntfy}
              onCheckedChange={(checked) => {
                setDraft((prev) => {
                  if (checked) {
                    return {
                      ...prev,
                      ntfy: newNtfyConfig(),
                    };
                  } else {
                    return omit(prev, "ntfy");
                  }
                });
              }}
            />
            <Label>Enabled</Label>
          </div>

          {draft.ntfy && (
            <div className="grid grid-cols-1 gap-2 md:grid-cols-2">
              <div className="space-y-2">
                <Label>URL</Label>
                <Input
                  type="text"
                  placeholder={defaultNtfyUrl}
                  value={
                    draft.ntfy?.url === defaultNtfyUrl ? "" : draft.ntfy?.url
                  }
                  onChange={(event) => {
                    setDraft((prev) => ({
                      ...prev,
                      ntfy: {
                        ...(prev.ntfy || newNtfyConfig()),
                        url:
                          event.target.value === ""
                            ? defaultNtfyUrl
                            : event.target.value,
                      },
                    }));
                  }}
                />
              </div>

              <div className="space-y-2">
                <Label>Topic</Label>
                <Input
                  type="text"
                  value={draft.ntfy?.topic}
                  onChange={(event) => {
                    setDraft((prev) => ({
                      ...prev,
                      ntfy: {
                        ...(prev.ntfy || newNtfyConfig()),
                        topic: event.target.value,
                      },
                    }));
                  }}
                />
              </div>

              <div className="space-y-2">
                <Label>Username</Label>
                <Input
                  type="text"
                  placeholder="Optional"
                  value={draft.ntfy?.username}
                  onChange={(event) => {
                    setDraft((prev) => ({
                      ...prev,
                      ntfy: {
                        ...(prev.ntfy || newNtfyConfig()),
                        username: event.target.value,
                      },
                    }));
                  }}
                />
              </div>

              <div className="space-y-2">
                <Label>Password</Label>
                <div className="flex">
                  <Input
                    type={showNtfyPassword ? "text" : "password"}
                    placeholder="Optional"
                    value={draft.ntfy?.password}
                    onChange={(event) => {
                      setDraft((prev) => ({
                        ...prev,
                        ntfy: {
                          ...(prev.ntfy || newNtfyConfig()),
                          password: event.target.value,
                        },
                      }));
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
          )}
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
