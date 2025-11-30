import { useConfig } from "../providers/config";
import { SaveDiscardButtons } from "./buttonsSaveDiscard";
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
          <SaveDiscardButtons
            onSave={() => {
              updateConfig((config) => {
                config.global = draft;
              });
            }}
            onDiscard={() => {
              setDraft(config.global);
            }}
          />
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
