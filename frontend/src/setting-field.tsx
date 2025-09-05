import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ResetButton } from "@/reset";
import type { ChangeEvent } from "react";

interface SettingFieldProps<T extends string | number> {
  label: string;
  description: string;
  placeholder?: string;
  type: "text" | "number" | "integer";
  value: T | undefined;
  onChange: (value: T | undefined) => void;
  onReset: () => void;
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
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    const targetValue = event.target.value;
    if (targetValue === "") {
      onChange(undefined);
    } else if (type === "number") {
      onChange(event.target.valueAsNumber as T);
    } else if (type === "integer") {
      if (!Number.isNaN(event.target.valueAsNumber)) {
        onChange(Math.trunc(event.target.valueAsNumber) as T);
      }
    } else {
      onChange(targetValue as T);
    }
  };

  return (
    <div className="space-y-2">
      <div className="flex items-center justify-between">
        <Label>{label}</Label>
        <ResetButton resetType="default" onClick={onReset} />
      </div>
      <p className="text-muted-foreground text-sm">{description}</p>
      <Input
        type={type === "integer" ? "number" : type}
        placeholder={placeholder}
        value={value}
        onChange={handleChange}
      />
    </div>
  );
}
