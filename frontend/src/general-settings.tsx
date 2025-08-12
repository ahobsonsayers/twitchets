"use client";

import { CollapsibleCard } from "./components/collapsible-card";
import { Button } from "@/components/ui/button";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@/components/ui/collapsible";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Eye, EyeOff } from "lucide-react";
import { useState } from "react";

interface GeneralSettingsProps {}

export function GeneralSettings({}: GeneralSettingsProps) {
  const [showApiKey, setShowApiKey] = useState(false);

  return (
    <CollapsibleCard title="hello" description="hello">
      hello
    </CollapsibleCard>
  );

  //   return (
  //     <CollapsibleSection
  //       id="general"
  //       title="General Settings"
  //       description="General application configuration"
  //     >
  //       <div className="space-y-4">
  //         <div className="space-y-2">
  //           <Label htmlFor="apiKey">API Key</Label>
  //           <p className="text-sm text-muted-foreground">
  //             Twickets API key (required)
  //           </p>
  //           <div className="relative">
  //             <Input
  //               id="apiKey"
  //               type={showApiKey ? "text" : "password"}
  //               value={config.apiKey}
  //               onChange={(e) => onUpdateConfig(["apiKey"], e.target.value)}
  //               placeholder="Enter your API key"
  //             />
  //             <Button
  //               type="button"
  //               variant="ghost"
  //               size="sm"
  //               className="absolute right-0 top-0 h-full px-3"
  //               onClick={onShowApiKeyToggle}
  //             >
  //               {showApiKey ? (
  //                 <EyeOff className="h-4 w-4" />
  //               ) : (
  //                 <Eye className="h-4 w-4" />
  //               )}
  //             </Button>
  //           </div>
  //         </div>
  //         <div className="space-y-2">
  //           <Label htmlFor="country">Country</Label>
  //           <p className="text-sm text-muted-foreground">
  //             Currently only GB is supported
  //           </p>
  //           <Select
  //             value={config.country}
  //             onValueChange={(value) => onUpdateConfig(["country"], value)}
  //           >
  //             <SelectTrigger>
  //               <SelectValue />
  //             </SelectTrigger>
  //             <SelectContent>
  //               <SelectItem value="GB">GB (United Kingdom)</SelectItem>
  //             </SelectContent>
  //           </Select>
  //         </div>
  //       </div>
  //     </CollapsibleSection>
  //   );
}
