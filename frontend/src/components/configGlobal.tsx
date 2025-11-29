import { useConfig } from "../providers/config";
import { CollapsibleCard } from "./cardCollapsible";
import { CommonFields } from "./configCommon";
import { Button } from "./ui/button";
import { isEqual } from "lodash";
import { useState, useEffect } from "react";

export function GlobalSettings() {
  const { config, updateConfig } = useConfig();

  const [draft, setDraft] = useState(config.global);

  useEffect(() => {
    setDraft(config.global);
  }, [config.global]);

  const hasChanges = !isEqual(draft, config.global);

  return (
    <CollapsibleCard
      title="Global Configuration"
      description="Configuration that apply to all tickets"
      action={
        hasChanges && (
          <div className="flex items-center gap-2">
            <Button
              variant="outline"
              size="sm"
              onClick={() => setDraft(config.global)}
            >
              Discard
            </Button>
            <Button
              size="sm"
              onClick={() => {
                updateConfig((config) => {
                  config.global = draft;
                });
              }}
            >
              Save
            </Button>
          </div>
        )
      }
    >
      <CommonFields
        commonConfig={draft}
        updateCommonConfig={(commonConfig) => {
          setDraft(commonConfig);
        }}
      />
    </CollapsibleCard>
  );
}
