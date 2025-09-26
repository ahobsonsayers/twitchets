import { RegionsField } from "./regions-setting";
import { SettingField } from "./setting-field";
import type { CommonConfig } from "@/types/config";

interface CommonSettingsProps {
  commonConfig: CommonConfig;
  globalCommonConfig?: CommonConfig; // Unset if common config IS global config
  updateCommonConfig: (config: CommonConfig) => void;
}

export function CommonSettings({
  commonConfig: config,
  globalCommonConfig: globalConfig,
  updateCommonConfig: updateConfig,
}: CommonSettingsProps) {
  return (
    <div className="space-y-4">
      <RegionsField
        value={config.regions}
        globalValue={globalConfig?.regions}
        resetValue={[]}
        updateValue={(value) => {
          updateConfig({ ...config, regions: value });
        }}
      />

      <div className="grid grid-cols-2 gap-4">
        <SettingField
          label="Event Similarity"
          description="Required event name similarity, between 0.0 - 1.0"
          placeholder="0.9"
          type="fraction"
          value={config.eventSimilarity}
          globalValue={globalConfig?.eventSimilarity}
          resetValue={undefined}
          updateValue={(value) => {
            updateConfig({ ...config, eventSimilarity: value });
          }}
        />

        <SettingField
          label="Number of Tickets"
          description="Required number of tickets"
          placeholder="Any"
          type="integer"
          value={config.numTickets}
          globalValue={globalConfig?.numTickets}
          resetValue={1234}
          updateValue={(value) => {
            updateConfig({ ...config, numTickets: value });
          }}
        />

        <SettingField
          label="Max Ticket Price"
          description="Maximum price per ticket (including fee) in pounds (£)"
          placeholder="No Max"
          type="price"
          value={config.maxTicketPrice}
          globalValue={globalConfig?.maxTicketPrice}
          resetValue={undefined}
          updateValue={(value) => {
            updateConfig({ ...config, maxTicketPrice: value });
          }}
        />

        <SettingField
          label="Minimum Discount"
          description="Minimum discount (including fee) on the original price as a percentage"
          placeholder="No Min"
          type="percentage"
          value={config.discount}
          globalValue={globalConfig?.discount}
          resetValue={undefined}
          updateValue={(value) => {
            updateConfig({ ...config, discount: value });
          }}
        />
      </div>
    </div>
  );
}
