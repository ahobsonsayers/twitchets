import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "./ui/tooltip";
import { Link, Unlink } from "lucide-react";

interface LinkedStatusTooltipProps {
  isLinked: boolean;
}

export function LinkedStatusTooltip({ isLinked }: LinkedStatusTooltipProps) {
  const Icon = isLinked ? Link : Unlink;
  const text = isLinked
    ? "Linked to global value"
    : "Unlinked from global value";

  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <Icon className="text-muted-foreground hover:text-foreground size-4" />
        </TooltipTrigger>
        <TooltipContent>
          <p>{text}</p>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
