"use client";

import { GeneralSettings } from "./settings-general";
import { ConfigProvider } from "./providers/config";
import { ThemeProvider } from "./providers/theme";

function App() {
  return (
    <ThemeProvider>
      <ConfigProvider>
        <div className="container mx-auto space-y-6 p-6">
          <div className="flex flex-col items-start justify-between gap-4 sm:flex-row sm:items-center">
            <h1 className="text-2xl font-bold sm:text-3xl">
              Application Configuration
            </h1>
          </div>
          <GeneralSettings />
        </div>
      </ConfigProvider>
    </ThemeProvider>
  );
}

export default App;
