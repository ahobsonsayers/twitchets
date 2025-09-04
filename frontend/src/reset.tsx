import { Button } from "./components/ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "./components/ui/tooltip";
import { Globe, RotateCcw } from "lucide-react";

interface ResetButtonProps {
  onClick: () => void;
  resetType: "default" | "global";
}

export function ResetButton({ onClick, resetType }: ResetButtonProps) {
  const Icon = resetType === "default" ? RotateCcw : Globe;

  const tooltip =
    resetType === "default" ? "Reset to default" : "Reset to global setting";

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Button
            className="text-muted-foreground hover:text-foreground h-8 w-8 p-0"
            type="button"
            variant="ghost"
            size="sm"
            onClick={onClick}
          >
            <Icon className="h-3 w-3" />
          </Button>
        </TooltipTrigger>
        <TooltipContent>
          <p>{tooltip}</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
