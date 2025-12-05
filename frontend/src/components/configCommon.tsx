import { ConfigField } from "./configField";
import { Regions } from "./configRegions";
import type { CommonConfig } from "@/types/config";

interface CommonFieldsProps {
  config: CommonConfig;
  globalConfig?: CommonConfig; // Unset if common config IS global config
  updateConfig: (config: CommonConfig) => void;
}

export function CommonFields({
  config,
  globalConfig,
  updateConfig,
}: CommonFieldsProps) {
  // If fields are for global config, globalConfig will not be set
  const isGlobal = !globalConfig;

  return (
    <div className="space-y-4">
      <Regions
        value={config.regions}
        withGlobalFallback={!isGlobal}
        globalFallbackValue={globalConfig?.regions}
        updateValue={(value) => {
          updateConfig({ ...config, regions: value });
        }}
      />

      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <ConfigField
          label="Event Similarity"
          description="Required event name similarity, between 0.0 - 1.0"
          type="fraction"
          value={config.eventSimilarity}
          showReset={true}
          resetValue={!isGlobal ? -1 : undefined} // Reset value for global is undefined
          defaultValuePlaceholder="0.9"
          showGlobalReset={!isGlobal}
          globalValuePlaceholder={
            globalConfig?.eventSimilarity?.toString() || "0.9"
          }
          updateValue={(value) => {
            updateConfig({ ...config, eventSimilarity: value });
          }}
        />

        <ConfigField
          label="Number of Tickets"
          description="Required number of tickets"
          type="integer"
          value={config.numTickets}
          showReset={true}
          resetValue={!isGlobal ? -1 : undefined} // Reset value for global is undefined
          defaultValuePlaceholder="Any"
          showGlobalReset={!isGlobal}
          globalValuePlaceholder={globalConfig?.numTickets?.toString() || "Any"}
          updateValue={(value) => {
            updateConfig({ ...config, numTickets: value });
          }}
        />

        <ConfigField
          label="Max Ticket Price"
          description="Maximum price per ticket (including fee) in pounds (£)"
          type="price"
          value={config.maxTicketPrice}
          showReset={true}
          resetValue={!isGlobal ? -1 : undefined} // Reset value for global is undefined
          defaultValuePlaceholder="No Max"
          showGlobalReset={!isGlobal}
          globalValuePlaceholder={
            globalConfig?.maxTicketPrice?.toString() || "No Max"
          }
          updateValue={(value) => {
            updateConfig({ ...config, maxTicketPrice: value });
          }}
        />

        <ConfigField
          label="Minimum Discount"
          description="Minimum discount (including fee) on the original price as a percentage"
          type="percentage"
          value={config.discount}
          showReset={true}
          resetValue={!isGlobal ? -1 : undefined} // Reset value for global is undefined
          defaultValuePlaceholder="No Min"
          showGlobalReset={!isGlobal}
          globalValuePlaceholder={
            globalConfig?.discount?.toString() || "No Min"
          }
          updateValue={(value) => {
            updateConfig({ ...config, discount: value });
          }}
        />
      </div>
    </div>
  );
}
