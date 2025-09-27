import { CollapsibleCard } from "./components/collapsibleCard";
import { CommonFields } from "./components/configCommon";
import { useConfig } from "./providers/config";

export function GlobalSettings() {
  const { config, setConfig } = useConfig();

  return (
    <CollapsibleCard
      title="Global Configuration"
      description="Configuration that apply to all tickets"
    >
      <CommonFields
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
