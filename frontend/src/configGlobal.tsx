import { CollapsibleCard } from "./components/collapsibleCard";
import { CommonFields } from "./components/configCommon";
import { useConfig } from "./providers/config";

export function GlobalSettings() {
  const { config, updateConfig } = useConfig();

  return (
    <CollapsibleCard
      title="Global Configuration"
      description="Configuration that apply to all tickets"
    >
      <CommonFields
        commonConfig={config.global}
        updateCommonConfig={(commonConfig) => {
          updateConfig((config) => {
            config.global = commonConfig;
          });
        }}
      />
    </CollapsibleCard>
  );
}
