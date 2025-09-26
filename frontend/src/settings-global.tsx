import { CollapsibleCard } from "./components/collapsible-card";
import { CommonSettings } from "./components/common-settings";
import { useConfig } from "./providers/config";

export function GlobalSettings() {
  const { config, setConfig } = useConfig();

  return (
    <CollapsibleCard
      title="Global Configuration"
      description="Configuration that apply to all tickets"
    >
      <CommonSettings
        commonConfig={config.global}
        updateCommonConfig={(commonConfig) => {
          setConfig((config) => ({
            ...config,
            global: { ...config.global, ...commonConfig },
          }));
        }}
      />
    </CollapsibleCard>
  );
}
