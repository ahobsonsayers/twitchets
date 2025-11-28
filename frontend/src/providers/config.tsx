import type { Config } from "../types/config";
import { createContext, type ReactNode, useContext, useState } from "react";

function newConfig(): Config {
  return {
    apiKey: "",
    country: "GB",
    notification: {},
    global: {},
    tickets: [],
  };
}

type ConfigUpdater = (config: Config) => Config;

type ConfigProviderState = {
  config: Config;
  setConfig: (updater: ConfigUpdater) => void;
};

const ConfigProviderContext = createContext<ConfigProviderState>({
  config: newConfig(),
  setConfig: () => null,
});

type ConfigProviderProps = {
  children: ReactNode;
};

export function ConfigProvider({ children, ...props }: ConfigProviderProps) {
  const [config, setConfig] = useState<Config>(newConfig());

  const setConfigWithLogging = (updater: ConfigUpdater) => {
    setConfig((prevConfig) => {
      const newConfig = updater(prevConfig);
      console.log(newConfig);
      return newConfig;
    });
  };

  const value = {
    config,
    setConfig: setConfigWithLogging,
  };

  return (
    <ConfigProviderContext.Provider {...props} value={value}>
      {children}
    </ConfigProviderContext.Provider>
  );
}

export const useConfig = () => {
  const context = useContext(ConfigProviderContext);
  if (context === undefined)
    throw new Error("useConfig must be used within a ConfigProvider");

  return context;
};
