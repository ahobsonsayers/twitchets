"use client";

import { Button } from "./components/ui/button";
import { GeneralSettings } from "./general-settings";
import { ConfigProvider } from "./providers/config";
import { ThemeProvider } from "./providers/theme";
import type { Config } from "./types/config";
import { useState } from "react";

function App() {
  const config: Config = {
    apiKey: "",
    country: "GB",
  };
  return (
    <ThemeProvider>
      <ConfigProvider>
        <div className="container mx-auto space-y-6 p-6">
          <div className="flex flex-col items-start justify-between gap-4 sm:flex-row sm:items-center">
            <h1 className="text-2xl font-bold sm:text-3xl">
              Application Configuration
            </h1>
          </div>
          <GeneralSettings config={config} />
        </div>
      </ConfigProvider>
    </ThemeProvider>
  );
}

export default App;
