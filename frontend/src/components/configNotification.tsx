"use client";

import { useConfig } from "../providers/config";
import { SaveDiscardButtons } from "./buttonSaveDiscard";
import { CollapsibleCard } from "./cardCollapsible";
import { Button } from "./ui/button";
import { Checkbox } from "./ui/checkbox";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Separator } from "./ui/separator";
import { isEqual, omit } from "lodash";
import { Eye, EyeOff } from "lucide-react";
import { useState, useEffect } from "react";

const defaultNtfyUrl = "https://ntfy.sh";

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
      title="Notification Settings"
      action={
        hasChanges && (
          <SaveDiscardButtons
            onSave={() => {
              updateConfig((config) => {
                config.notification = draft;
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
              className="ml-4"
              checked={!!draft.ntfy}
              onCheckedChange={(checked) => {
                setDraft((prev) => {
                  if (checked) {
                    return {
                      ...prev,
                      ntfy: {
                        url: defaultNtfyUrl,
                        topic: "",
                        username: "",
                        password: "",
                      },
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
            <>
              <div className="flex gap-2">
                <div className="flex-auto flex-col space-y-2">
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
                          ...prev.ntfy!,
                          url:
                            event.target.value === ""
                              ? defaultNtfyUrl
                              : event.target.value,
                        },
                      }));
                    }}
                  />
                </div>

                <div className="flex-auto flex-col space-y-2">
                  <Label>Topic</Label>
                  <Input
                    type="text"
                    value={draft.ntfy?.topic}
                    onChange={(event) => {
                      setDraft((prev) => ({
                        ...prev,
                        ntfy: {
                          ...prev.ntfy!,
                          topic: event.target.value,
                        },
                      }));
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
                    value={draft.ntfy?.username}
                    onChange={(event) => {
                      setDraft((prev) => ({
                        ...prev,
                        ntfy: {
                          ...prev.ntfy!,
                          username: event.target.value,
                        },
                      }));
                    }}
                  />
                </div>

                <div className="flex flex-1 flex-col space-y-2">
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
                            ...prev.ntfy!,
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
            </>
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
