import { Checkbox } from "./ui/checkbox";
import { Label } from "@/components/ui/label";
import { type Region, REGIONS } from "@/constants/regions";
import { ResetButton } from "@/reset";

interface RegionsFieldProps {
  label?: string;
  description?: string;
  value?: string[];
  globalValue?: string[];
  resetValue?: string[];
  updateValue: (newValue?: string[]) => void;
}

export function RegionsField({
  label = "Regions",
  description = "If no regions selected, all regions will be used",
  value,
  globalValue,
  resetValue,
  updateValue,
}: RegionsFieldProps) {
  const currentRegions = value || [];

  const handleOnCheckedChange = (region: Region, checked: boolean) => {
    var newRegions = currentRegions;

    if (!checked) {
      newRegions = newRegions.filter((r) => r !== region.code);
    } else if (!newRegions.includes(region.code)) {
      newRegions = [...newRegions, region.code];
    }

    updateValue(newRegions);
  };

  return (
    <div className="space-y-2">
      <div className="flex items-center justify-between">
        <Label>{label}</Label>
        {globalValue && (
          <ResetButton
            resetType="global"
            onClick={() => {
              updateValue(globalValue);
            }}
          />
        )}
        <ResetButton
          resetType="default"
          onClick={() => {
            updateValue(resetValue);
          }}
        />
      </div>

      <p className="text-muted-foreground text-sm">{description}</p>

      <div className="grid grid-cols-2">
        {REGIONS.map((region) => (
          <div className="flex items-center gap-2" key={region.code}>
            <Checkbox
              checked={currentRegions.includes(region.code)}
              onCheckedChange={(checked: boolean) => {
                handleOnCheckedChange(region, checked);
              }}
            />
            <Label className="text-sm">{region.name}</Label>
          </div>
        ))}
      </div>
    </div>
  );
}
