import { LinkedStatusTooltip } from "./statusLinked";
import { ResetButton } from "@/components/buttonReset";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { NumericFormat } from "react-number-format";

interface ConfigFieldProps<T extends string | number> {
  label: string;
  description: string;
  type: "text" | "number" | "integer" | "fraction" | "percentage" | "price";
  value?: T;
  showReset?: boolean; // Whether to show reset button
  resetValue?: T; // Value to set when reset is clicked
  defaultValuePlaceholder?: string; // Placeholder to use when field will use the default value (i.e. reset)
  showGlobalReset?: boolean; // Whether to show global reset button. Should only be true if field is not a global field
  globalValuePlaceholder?: string; // Placeholder to use when field will use the global value (i.e. global reset)
  updateValue: (newValue?: T) => void;
}

export function ConfigField<T extends string | number>({
  label,
  description,
  type,
  value,
  resetValue,
  globalValuePlaceholder,
  defaultValuePlaceholder,
  showReset = true,
  showGlobalReset = false,
  updateValue,
}: ConfigFieldProps<T>) {
  let fieldValue = value;

  // Determine whether value means field will use global/default
  // value, and there the placeholder to use.
  let isLinkedToGlobal = false;
  let placeholder = "";
  if (showGlobalReset && fieldValue === undefined) {
    placeholder = `${globalValuePlaceholder} (Global)`;
    isLinkedToGlobal = true;
  } else if (
    showReset &&
    (fieldValue === undefined ||
      (typeof fieldValue === "string" && fieldValue === "") ||
      (typeof fieldValue === "number" && fieldValue < 0))
  ) {
    placeholder = `${defaultValuePlaceholder} (Default)`;
    fieldValue = undefined; // Unset value so placeholder is shown
  }

  const renderInput = () => {
    if (type === "text") {
      return (
        <Input
          type="text"
          value={fieldValue ?? ""}
          placeholder={placeholder}
          onChange={(event) => updateValue(event.target.value as T)}
        />
      );
    }

    //  TODO better to use use-mask-input - but this currently reverses numbers???
    return (
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
          const val = values.floatValue;
          if (val === undefined) return true;
          if (type === "fraction") return val <= 1;
          if (type === "percentage") return val <= 100;
          return true;
        }}
        suffix={type === "percentage" ? "%" : ""}
        prefix={type === "price" ? "£" : ""}
      />
    );
  };

  return (
    <div className="space-y-2">
      <div className="flex">
        <div className="flex items-center space-x-2">
          <Label>{label}</Label>

          {showGlobalReset && (
            <LinkedStatusTooltip isLinked={isLinkedToGlobal} />
          )}
        </div>

        <div className="ml-auto flex items-center">
          {showGlobalReset && (
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

      {renderInput()}
    </div>
  );
}
