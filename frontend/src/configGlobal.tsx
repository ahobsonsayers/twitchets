import { CollapsibleCard } from "./components/collapsibleCard";
import { CommonSettings } from "./components/configCommonFields";
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
