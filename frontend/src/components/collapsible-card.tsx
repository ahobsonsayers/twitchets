"use client";

import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from "@radix-ui/react-collapsible";
import { ChevronsDownUp, ChevronsUpDown } from "lucide-react";
import { type ReactNode, useState } from "react";

interface CollapsibleCardProps {
  title: string;
  description?: string;
  children?: ReactNode;
}

export function CollapsibleCard({
  title,
  description,
  children,
}: CollapsibleCardProps) {
  const [isCollapsed, setIsCollapsed] = useState(true);

  const toggleCollapse = () => {
    setIsCollapsed(!isCollapsed);
  };

  return (
    <Card>
      <Collapsible>
        <CollapsibleTrigger asChild onClick={toggleCollapse}>
          <CardHeader className="flex w-full items-center">
            <div className="flex flex-col">
              <CardTitle>{title}</CardTitle>
              {description && <CardDescription>{description}</CardDescription>}
            </div>
            <div className="ml-auto">
              {isCollapsed ? (
                <ChevronsUpDown className="h-5 w-5" />
              ) : (
                <ChevronsDownUp className="h-5 w-5" />
              )}
            </div>
          </CardHeader>
        </CollapsibleTrigger>
        <CollapsibleContent className="pt-4">
          <CardContent>{children}</CardContent>
        </CollapsibleContent>
      </Collapsible>
    </Card>
  );
}
