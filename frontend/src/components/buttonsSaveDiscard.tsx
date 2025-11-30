import { Button } from "./ui/button";

interface SaveDiscardButtonsProps {
  onSave: () => void;
  onDiscard: () => void;
}

export function SaveDiscardButtons({
  onSave,
  onDiscard,
}: SaveDiscardButtonsProps) {
  const handleSave = (e: React.MouseEvent) => {
    e.stopPropagation();
    onSave();
  };

  const handleDiscard = (e: React.MouseEvent) => {
    e.stopPropagation();
    onDiscard();
  };

  return (
    <div className="flex items-center gap-2">
      <Button type="button" variant="outline" size="sm" onClick={handleDiscard}>
        Discard
      </Button>

      <Button type="button" size="sm" onClick={handleSave}>
        Save
      </Button>
    </div>
  );
}
