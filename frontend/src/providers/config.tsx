import type { Config } from "../types/config";
import { newConfig } from "../types/config";
import { createContext, type ReactNode, useContext, useState } from "react";

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

  const value = {
    config,
    setConfig,
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
