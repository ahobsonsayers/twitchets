"use client";

import { GeneralSettings } from "./components/configGeneral";
import { GlobalSettings } from "./components/configGlobal";
import { NotificationSettings } from "./components/configNotification";
import { TicketsConfig } from "./components/configTickets";
import { ConfigProvider } from "./providers/config";
import { ThemeProvider } from "./providers/theme";

function App() {
  return (
    <ThemeProvider>
      <ConfigProvider>
        <div className="container m-6 mx-auto space-y-6">
          <h1 className="text-3xl font-bold">Twitchets Configuration</h1>
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
