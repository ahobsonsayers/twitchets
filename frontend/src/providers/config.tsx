import type { Config } from "../types/config";
import { createContext, type ReactNode, useContext, useState } from "react";
import { produce } from "immer";

function newConfig(): Config {
  return {
    apiKey: "",
    country: "GB",
    notification: {},
    global: {},
    tickets: [],
  };
}

type ConfigUpdater = (config: Config) => void;

type ConfigProviderState = {
  config: Config;
  updateConfig: (updater: ConfigUpdater) => void;
};

const ConfigProviderContext = createContext<ConfigProviderState>({
  config: newConfig(),
  updateConfig: () => null,
});

type ConfigProviderProps = {
  children: ReactNode;
};

export function ConfigProvider({ children, ...props }: ConfigProviderProps) {
  const [config, setConfig] = useState<Config>(newConfig());

  const updateConfig = (updater: ConfigUpdater) => {
    setConfig(produce(updater));
  };

  const value = {
    config,
    updateConfig,
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
