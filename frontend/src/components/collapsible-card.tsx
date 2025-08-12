"use client";

import { Button } from "@/components/ui/button";
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
import { ChevronDown, ChevronRight } from "lucide-react";
import type React from "react";
import { useState, useEffect, ReactNode } from "react";

interface CollapsibleCardProps {
  title: string;
  description: string;
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
        <CollapsibleTrigger onClick={toggleCollapse}>
          <CardHeader className="flex items-center">
            <CardTitle>{title}</CardTitle>
            {description && <CardDescription>{description}</CardDescription>}
            <div className="alignRight">
              {isCollapsed ? (
                <ChevronRight className="h-4 w-4" />
              ) : (
                <ChevronDown className="h-4 w-4" />
              )}
            </div>
          </CardHeader>
        </CollapsibleTrigger>
        <CollapsibleContent>
          <CardContent>{children}</CardContent>
        </CollapsibleContent>
      </Collapsible>
    </Card>
  );
}
