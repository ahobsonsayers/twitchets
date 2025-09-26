import { LinkedStatusTooltip } from "./linkedStatusTooltip";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ResetButton } from "@/reset";
import { NumericFormat } from "react-number-format";

interface SettingFieldProps<T extends string | number> {
  label: string;
  description: string;
  placeholder?: string;
  type: "text" | "number" | "integer" | "fraction" | "percentage" | "price";
  value?: T;
  globalFallbackValue?: T;
  withGlobalFallback?: boolean;
  showReset?: boolean;
  updateValue: (newValue?: T) => void;
}

export function SettingField<T extends string | number>({
  label,
  description,
  placeholder,
  type,
  value,
  withGlobalFallback = false,
  globalFallbackValue = undefined,
  showReset = true,
  updateValue,
}: SettingFieldProps<T>) {
  // Determine the field value to display
  let fieldValue = value;
  let isLinkedToGlobal = false;
  if (fieldValue === undefined && withGlobalFallback) {
    // If no value is set, and we want to use global fallback
    fieldValue = globalFallbackValue;
    isLinkedToGlobal = true;
  } else if (typeof fieldValue === "number" && fieldValue < 0) {
    // If value is negative number, set undefined to show placeholder
    fieldValue = undefined;
  }

  // Determine the reset value based on type
  let resetValue: T;
  if (type === "text") {
    resetValue = "" as T;
  } else {
    resetValue = -1 as T;
  }

  return (
    <div className="space-y-2">
      <div className="flex">
        <div className="flex items-center space-x-2">
          <Label>{label}</Label>

          {withGlobalFallback && (
            <LinkedStatusTooltip isLinked={isLinkedToGlobal} />
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

          {showReset && (
            <ResetButton
              resetType="default"
              onClick={() => {
                updateValue(resetValue);
              }}
            />
          )}
        </div>
      </div>

      <p className="text-muted-foreground text-sm">{description}</p>

      {/* TODO better to use use-mask-input - but this currently reverses numbers??? */}
      {(type === "text" && (
        <Input
          type="text"
          value={fieldValue ?? ""}
          placeholder={placeholder}
          onChange={(event) => updateValue(event.target.value as T)}
        />
      )) || (
        <NumericFormat
          customInput={Input}
          value={fieldValue ?? ""}
          placeholder={placeholder}
          allowNegative={false}
          decimalScale={type === "integer" ? 0 : undefined}
          onValueChange={(values) => {
            // Only update if the value is not undefined
            // This prevents reverting from -1 (which is displayed as undefined) back to undefined
            if (values.floatValue !== undefined) {
              updateValue(values.floatValue as T);
            }
          }}
          isAllowed={(values) => {
            if (type === "fraction") {
              return values.floatValue === undefined || values.floatValue <= 1;
            }
            if (type === "percentage") {
              return (
                values.floatValue === undefined || values.floatValue <= 100
              );
            }
            return true;
          }}
          valueIsNumericString={true}
          suffix={type === "percentage" ? "%" : ""}
          prefix={type === "price" ? "£" : ""}
        />
      )}
    </div>
  );
}
