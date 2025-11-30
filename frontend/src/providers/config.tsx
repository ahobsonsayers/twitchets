import { Alert, AlertDescription, AlertTitle } from "../components/ui/alert";
import { getConfig, putConfig } from "../lib/api";
import type { Config } from "../types/config";
import { produce } from "immer";
import { TriangleAlertIcon, Save } from "lucide-react";
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
  saved: boolean;
  error: string | null;
};

const ConfigProviderContext = createContext<ConfigProviderState>({
  config: newConfig(),
  updateConfig: () => null,
  saved: false,
  error: null,
});

type ConfigProviderProps = {
  children: ReactNode;
};

export function ConfigProvider({ children, ...props }: ConfigProviderProps) {
  const [config, setConfig] = useState<Config>(newConfig());
  const [saved, setSaved] = useState(false);
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

  useEffect(() => {
    if (!saved) return;
    const timer = setTimeout(() => setSaved(false), 5000);
    return () => clearTimeout(timer);
  }, [saved]);

  useEffect(() => {
    if (!error) return;
    const timer = setTimeout(() => setError(null), 5000);
    return () => clearTimeout(timer);
  }, [error]);

  const updateConfig = async (updater: ConfigUpdater) => {
    const newConfig = produce(config, updater);
    setConfig(newConfig);

    try {
      await putConfig(newConfig);
      setError(null);
      setSaved(true);
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to update config");
      // Revert on error
      setConfig(config);
    }
  };

  const value = {
    config,
    updateConfig,
    saved,
    error,
  };

  return (
    <ConfigProviderContext.Provider {...props} value={value}>
      {children}
      {saved && (
        <div className="fixed bottom-4 right-4 z-50 max-w-sm">
          <Alert className="border-none bg-emerald-600 text-white">
            <Save />
            <AlertTitle>Saved</AlertTitle>
            <AlertDescription className="text-white">
              Changes have been saved
            </AlertDescription>
          </Alert>
        </div>
      )}
      {error && (
        <div className="fixed bottom-4 right-4 z-50 max-w-sm">
          <Alert className="bg-destructive text-destructive-foreground border-none">
            <TriangleAlertIcon />
            <AlertTitle>Error</AlertTitle>
            <AlertDescription className="text-destructive-foreground">
              {error}
            </AlertDescription>
          </Alert>
        </div>
      )}
    </ConfigProviderContext.Provider>
  );
}

export const useConfig = () => {
  const context = useContext(ConfigProviderContext);
  if (context === undefined)
    throw new Error("useConfig must be used within a ConfigProvider");

  return context;
};
