import { Checkbox } from "./ui/checkbox";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "./ui/tooltip";
import { Label } from "@/components/ui/label";
import { type Region, REGIONS } from "@/constants/regions";
import { ResetButton } from "@/reset";
import { Link, Unlink } from "lucide-react";

interface RegionsFieldProps {
  label?: string;
  description?: string;
  value?: string[];
  withGlobalFallback?: boolean;
  globalFallbackValue?: string[];
  updateValue: (newValue?: string[]) => void;
}

export function RegionsField({
  label = "Regions",
  description = "If no regions selected, all regions will be used",
  value,
  withGlobalFallback = false,
  globalFallbackValue,
  updateValue,
}: RegionsFieldProps) {
  // Determine the field value to display
  let fieldValue = value;
  let isLinkedToGlobal = false;
  if (fieldValue === undefined && withGlobalFallback) {
    // If no value is set, and we want to use global fallback
    fieldValue = globalFallbackValue;
    isLinkedToGlobal = true;
  }

  const currentRegions = fieldValue || [];
  const resetValue: string[] = [];

  const handleOnCheckedChange = (regionCode: string, checked: boolean) => {
    const regionSet = new Set(currentRegions);

    if (checked) {
      regionSet.add(regionCode);
    } else {
      regionSet.delete(regionCode);
    }

    updateValue(Array.from(regionSet).sort());
  };

  return (
    <div className="space-y-2">
      <div className="flex">
        <div className="flex items-center space-x-2">
          <Label>{label}</Label>

          {withGlobalFallback && (
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  {isLinkedToGlobal ? (
                    <Link className="text-muted-foreground hover:text-foreground size-4" />
                  ) : (
                    <Unlink className="text-muted-foreground hover:text-foreground size-4" />
                  )}
                </TooltipTrigger>
                <TooltipContent>
                  <p>
                    {isLinkedToGlobal
                      ? "Linked to global value"
                      : "Unlinked from global value"}
                  </p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          )}
        </div>

        <div className="ml-auto flex items-center">
          {withGlobalFallback && (
            <ResetButton
              resetType="global"
              onClick={() => {
                // The global button sets the value to "undefined".
                // This causes the global value to be inherited.
                updateValue(undefined);
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
      </div>

      <p className="text-muted-foreground text-sm">{description}</p>

      <div className="grid grid-cols-2">
        {REGIONS.map((region) => (
          <div className="flex items-center gap-2" key={region.code}>
            <Checkbox
              checked={currentRegions.includes(region.code)}
              onCheckedChange={(checked: boolean) => {
                handleOnCheckedChange(region.code, checked);
              }}
            />
            <Label className="text-sm">{region.name}</Label>
          </div>
        ))}
      </div>
    </div>
  );
}
