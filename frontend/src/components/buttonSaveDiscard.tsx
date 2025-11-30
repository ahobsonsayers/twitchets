import { Button } from "./ui/button";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "./ui/tooltip";
import { Globe, RotateCcw } from "lucide-react";

interface SaveDiscardButtonsProps {
  onSave: () => void;
  onDiscard: () => void;
}

export function SaveDiscardButtons({
  onSave,
  onDiscard,
}: SaveDiscardButtonsProps) {
  return (
    <div className="flex items-center gap-2">
      <Button
        variant="outline"
        size="sm"
        onClick={(e) => {
          e.stopPropagation();
          onDiscard();
        }}
      >
        Discard
      </Button>
      <Button
        size="sm"
        onClick={(e) => {
          e.stopPropagation();
          onSave();
        }}
      >
        Save
      </Button>
    </div>
  );
}
