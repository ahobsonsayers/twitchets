import { CollapsibleCard } from "./components/collapsible-card";
import { Checkbox } from "./components/ui/checkbox";
import { Input } from "./components/ui/input";
import { REGIONS } from "./constants/regions";
import { useConfig } from "./providers/config";
import { ResetButton } from "./reset";
import { Label } from "@radix-ui/react-label";

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
          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label>Event Similarity</Label>
              <ResetButton
                resetType="default"
                onClick={() => {
                  updateConfig((config) => {
                    config.global.eventSimilarity = undefined;  
                  });
                }}
              />
            </div>
            <p className="text-muted-foreground text-sm">
              Required event name similarity, between 0.0 - 1.0
            </p>
            <Input
              type="string"
              placeholder="0.9"
              value={config.global.eventSimilarity}
              onChange={(event) => {
                updateConfig((config) => {
                  config.global.eventSimilarity = event.target.value;
                });
              }}
            />
          </div>

          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label>Number of Tickets</Label>
              <ResetButton
                resetType="default"
                onClick={() => {
                  updateConfig((config) => {
                    config.global.numTickets = undefined;
                  });
                }}
              />
            </div>
            <p className="text-muted-foreground text-sm">
              Required number of tickets
            </p>
            <Input
              type="string"
              placeholder="Any"
              value={config.global.numTickets}
              onChange={(event) => {
                updateConfig((config) => {
                  config.global.numTickets = event.target.value;
                });
              }}
            />
          </div>

          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label>Max Ticket Price</Label>
              <ResetButton
                resetType="default"
                onClick={() => {
                  updateConfig((config) => {
                    config.global.maxTicketPrice = undefined;
                  });
                }}
              />
            </div>
            <p className="text-muted-foreground text-sm">
              Maximum price per ticket (including fee) in pounds (£)
            </p>
            <Input
              type="string"
              placeholder="No Max"
              value={config.global.maxTicketPrice}
              onChange={(event) => {
                updateConfig((config) => {
                  config.global.maxTicketPrice = event.target.value;
                });
              }}
            />
          </div>

          <div className="space-y-2">
            <div className="flex items-center justify-between">
              <Label>Minimum Discount</Label>
              <ResetButton
                resetType="default"
                onClick={() => {
                  updateConfig((config) => {
                    config.global.discount = undefined;
                  });
                }}
              />
            </div>
            <p className="text-muted-foreground text-sm">
              Minimum discount (including fee) on the original price as a
              percentage
            </p>
            <Input
              type="string"
              placeholder="No Min"
              value={config.global.discount}
              onChange={(event) => {
                updateConfig((config) => {
                  config.global.discount = event.target.value;
                });
              }}
            />
          </div>
        </div>
      </div>
    </CollapsibleCard>
  );
}
