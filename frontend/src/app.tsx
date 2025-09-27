"use client";

import { GeneralSettings } from "./configGeneral";
import { GlobalSettings } from "./configGlobal";
import { NotificationSettings } from "./configNotification";
import { TicketsConfig } from "./configTickets";
import { ConfigProvider } from "./providers/config";
import { ThemeProvider } from "./providers/theme";

function App() {
  return (
    <ThemeProvider>
      <ConfigProvider>
        <div className="container m-6 mx-auto space-y-6">
          <h1 className="text-3xl font-bold">Application Configuration</h1>
          <GeneralSettings />
          <NotificationSettings />
          <GlobalSettings />
          <TicketsConfig />
        </div>
      </ConfigProvider>
    </ThemeProvider>
  );
}

export default App;
