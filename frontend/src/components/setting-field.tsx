import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ResetButton } from "@/reset";
import type { ChangeEvent } from "react";
import { NumericFormat, useNumericFormat } from "react-number-format";
import { withMask } from "use-mask-input";

interface SettingFieldProps<T extends string | number> {
  label: string;
  description: string;
  placeholder?: string;
  type: "text" | "number" | "integer" | "fraction" | "percentage" | "price";
  value?: T;
  onChange: (value?: T) => void;
  onReset?: () => void;
}

export function SettingField<T extends string | number>({
  label,
  description,
  placeholder,
  type,
  value,
  onChange,
  onReset,
}: SettingFieldProps<T>) {
  return (
    <div className="space-y-2">
      <div className="flex items-center justify-between">
        <Label>{label}</Label>
        {!!onReset && <ResetButton resetType="default" onClick={onReset} />}
      </div>
      <p className="text-muted-foreground text-sm">{description}</p>
      {/* TODO better to use use-mask-input - but this currently reverses numbers??? */}
      {(type === "text" && (
        <Input
          type="text"
          placeholder={placeholder}
          value={value ?? ""}
          onChange={(event) => onChange(event.target.value as T)}
        />
      )) || (
        <NumericFormat
          customInput={Input}
          value={value ?? ""}
          placeholder={placeholder}
          allowNegative={false}
          decimalScale={type === "integer" ? 0 : undefined}
          onValueChange={(values) => {
            onChange(values.floatValue as T);
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
