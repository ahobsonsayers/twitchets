import type { Config } from "../types/config";
import { newConfig } from "../types/config";
import { produce } from "immer";
import { createContext, useContext, useState } from "react";

// ConfigUpdater is a function used to update a config
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
  children: React.ReactNode;
};

export function ConfigProvider({ children, ...props }: ConfigProviderProps) {
  const [config, setConfig] = useState<Config>(newConfig());

  const updateConfig = (updater: ConfigUpdater) => {
    setConfig(
      produce(config, (draft) => {
        updater(draft);
      }),
    );
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
