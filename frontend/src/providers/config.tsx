import { getConfig, putConfig } from "../lib/api";
import type { Config } from "../types/config";
import { produce } from "immer";
import {
  createContext,
  type ReactNode,
  useContext,
  useEffect,
  useState,
} from "react";

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
  error: string | null;
};

const ConfigProviderContext = createContext<ConfigProviderState>({
  config: newConfig(),
  updateConfig: () => null,
  error: null,
});

type ConfigProviderProps = {
  children: ReactNode;
};

export function ConfigProvider({ children, ...props }: ConfigProviderProps) {
  const [config, setConfig] = useState<Config>(newConfig());
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const loadConfig = async () => {
      try {
        const serverConfig = await getConfig();
        setConfig(serverConfig);
      } catch (err) {
        setError(err instanceof Error ? err.message : "Failed to load config");
      }
    };

    loadConfig();
  }, []);

  const updateConfig = async (updater: ConfigUpdater) => {
    const newConfig = produce(config, updater);
    setConfig(newConfig);

    try {
      await putConfig(newConfig);
      setError(null);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to update config");
      // Revert on error
      setConfig(config);
    }
  };

  const value = {
    config,
    updateConfig,
    error,
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
