import { ConfigField } from "./configField";
import { Regions } from "./configRegions";
import type { CommonConfig } from "@/types/config";

interface CommonFieldsProps {
  commonConfig: CommonConfig;
  globalCommonConfig?: CommonConfig; // Unset if common config IS global config
  updateCommonConfig: (config: CommonConfig) => void;
}

export function CommonFields({
  commonConfig: config,
  globalCommonConfig: globalConfig,
  updateCommonConfig: updateConfig,
}: CommonFieldsProps) {
  return (
    <div className="space-y-4">
      <Regions
        value={config.regions}
        withGlobalFallback={!!globalConfig}
        globalFallbackValue={globalConfig?.regions}
        updateValue={(value) => {
          updateConfig({ ...config, regions: value });
        }}
      />

      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <ConfigField
          label="Event Similarity"
          description="Required event name similarity, between 0.0 - 1.0"
          placeholder="0.9"
          type="fraction"
          value={config.eventSimilarity}
          withGlobalFallback={!!globalConfig}
          globalFallbackValue={globalConfig?.eventSimilarity}
          showReset={true}
          updateValue={(value) => {
            updateConfig({ ...config, eventSimilarity: value });
          }}
        />

        <ConfigField
          label="Number of Tickets"
          description="Required number of tickets"
          placeholder="Any"
          type="integer"
          value={config.numTickets}
          globalFallbackValue={globalConfig?.numTickets}
          withGlobalFallback={!!globalConfig}
          showReset={true}
          updateValue={(value) => {
            updateConfig({ ...config, numTickets: value });
          }}
        />

        <ConfigField
          label="Max Ticket Price"
          description="Maximum price per ticket (including fee) in pounds (£)"
          placeholder="No Max"
          type="price"
          value={config.maxTicketPrice}
          globalFallbackValue={globalConfig?.maxTicketPrice}
          withGlobalFallback={!!globalConfig}
          showReset={true}
          updateValue={(value) => {
            updateConfig({ ...config, maxTicketPrice: value });
          }}
        />

        <ConfigField
          label="Minimum Discount"
          description="Minimum discount (including fee) on the original price as a percentage"
          placeholder="No Min"
          type="percentage"
          value={config.discount}
          globalFallbackValue={globalConfig?.discount}
          withGlobalFallback={!!globalConfig}
          showReset={true}
          updateValue={(value) => {
            updateConfig({ ...config, discount: value });
          }}
        />
      </div>
    </div>
  );
}
