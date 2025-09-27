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
import { type ReactNode, useEffect, useState } from "react";

interface CollapsibleCardProps {
  title: string;
  description?: string;
  isOpen?: boolean;
  setIsOpen?: (isOpen: boolean) => void;
  action?: ReactNode;
  children?: ReactNode;
}

export function CollapsibleCard({
  title,
  description,
  isOpen: isOpenProp,
  setIsOpen: setIsOpenProp,
  action,
  children,
}: CollapsibleCardProps) {
  const [isOpenInternal, setIsOpenInteral] = useState(false);

  const isOpen = isOpenProp ?? isOpenInternal;
  const setIsOpen = setIsOpenProp ?? setIsOpenInteral;

  const toggleCollapse = () => {
    setIsOpen(!isOpen);
  };

  return (
    <Card>
      <Collapsible open={isOpen} onOpenChange={setIsOpen}>
        <CollapsibleTrigger asChild onClick={toggleCollapse}>
          <CardHeader className="flex w-full items-center">
            <div className="flex flex-col">
              <CardTitle>{title}</CardTitle>
              {description && <CardDescription>{description}</CardDescription>}
            </div>
            <div className="ml-auto flex items-center gap-4">
              {action}
              {isOpen ? (
                <ChevronsDownUp className="size-5" />
              ) : (
                <ChevronsUpDown className="size-5" />
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
