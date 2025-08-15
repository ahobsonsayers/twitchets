import type { Config } from "../types/config";
import { newConfig } from "../types/config";
import { createContext, useContext, useEffect, useState } from "react";

type ConfigProviderState = {
  config: Config;
  setConfig: (config: Config) => void;
};

const ConfigProviderContext = createContext<ConfigProviderState>({
  config: newConfig(),
  setConfig: () => null,
});

type ConfigProviderProps = {
  children: React.ReactNode;
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
