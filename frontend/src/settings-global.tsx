import { CollapsibleCard } from "./components/collapsible-card";
import { Checkbox } from "./components/ui/checkbox";
import { Label } from "./components/ui/label";
import { REGIONS } from "./constants/regions";
import { useConfig } from "./providers/config";
import { ResetButton } from "./reset";
import { SettingField } from "@/components/setting-field";

export function GlobalSettings() {
  const { config, updateConfig } = useConfig();

  return (
    <CollapsibleCard
      title="Global Settings"
      description="Settings that apply to all tickets"
    >
      <div className="space-y-4">
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <Label>Regions</Label>
            <ResetButton
              resetType="default"
              onClick={() => {
                updateConfig((config) => {
                  config.global.regions = undefined;
                });
              }}
            />
          </div>
          <p className="text-muted-foreground text-sm">
            If no regions selected, all regions will be used
          </p>
          <div className="grid grid-cols-2">
            {REGIONS.map((region) => (
              <div className="flex items-center gap-2" key={region.code}>
                <Checkbox
                  checked={
                    config.global.regions?.includes(region.code) || false
                  }
                  onCheckedChange={(checked) => {
                    updateConfig((config) => {
                      var selectedRegions = config.global.regions || [];

                      if (checked && !selectedRegions.includes(region.code)) {
                        selectedRegions = [...selectedRegions, region.code];
                      } else if (!checked) {
                        selectedRegions = selectedRegions.filter(
                          (r) => r !== region.code,
                        );
                      }

                      config.global.regions = selectedRegions;
                    });
                  }}
                />
                <Label className="text-sm">{region.name}</Label>
              </div>
            ))}
          </div>
        </div>

        <div className="grid grid-cols-2 gap-4">
          <SettingField
            label="Event Similarity"
            description="Required event name similarity, between 0.0 - 1.0"
            placeholder="0.9"
            type="fraction"
            value={config.global.eventSimilarity}
            onChange={(value) =>
              updateConfig((config) => {
                config.global.eventSimilarity = value;
              })
            }
            onReset={() =>
              updateConfig((config) => {
                config.global.eventSimilarity = undefined;
              })
            }
          />
          <SettingField
            label="Number of Tickets"
            description="Required number of tickets"
            placeholder="Any"
            type="integer"
            value={config.global.numTickets}
            onChange={(value) =>
              updateConfig((config) => {
                config.global.numTickets = value;
              })
            }
            onReset={() =>
              updateConfig((config) => {
                config.global.numTickets = undefined;
              })
            }
          />
          <SettingField
            label="Max Ticket Price"
            description="Maximum price per ticket (including fee) in pounds (£)"
            placeholder="No Max"
            type="price"
            value={config.global.maxTicketPrice}
            onChange={(value) =>
              updateConfig((config) => {
                config.global.maxTicketPrice = value;
              })
            }
            onReset={() =>
              updateConfig((config) => {
                config.global.maxTicketPrice = undefined;
              })
            }
          />
          <SettingField
            label="Minimum Discount"
            description="Minimum discount (including fee) on the original price as a percentage"
            placeholder="No Min"
            type="percentage"
            value={config.global.discount}
            onChange={(value) =>
              updateConfig((config) => {
                config.global.discount = value;
              })
            }
            onReset={() =>
              updateConfig((config) => {
                config.global.discount = undefined;
              })
            }
          />
        </div>
      </div>
    </CollapsibleCard>
  );
}
