"use client";

import { ConfigProvider } from "./providers/config";
import { ThemeProvider } from "./providers/theme";
import { GeneralSettings } from "./settings-general";
import { GlobalSettings } from "./settings-global";
import { NotificationSettings } from "./settings-notification";

function App() {
  return (
    <ThemeProvider>
      <ConfigProvider>
        <div className="container m-6 mx-auto space-y-6">
          <h1 className="text-3xl font-bold">Application Configuration</h1>
          <GeneralSettings />
          <NotificationSettings />
          <GlobalSettings />
        </div>
      </ConfigProvider>
    </ThemeProvider>
  );
}

export default App;
