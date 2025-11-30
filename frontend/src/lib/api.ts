import type { Config } from "../types/config";
import type { paths } from "../types/openapi";
import createClient from "openapi-fetch";

const client = createClient<paths>({
  baseUrl: `${window.location.origin}/api`,
});

export async function getConfig(): Promise<Config> {
  const { data, error } = await client.GET("/config", {});
  if (error) {
    throw new Error(`Failed to fetch config: ${error}`);
  }
  if (!data) {
    throw new Error("No config data received");
  }
  return data;
}

export async function putConfig(config: Config): Promise<void> {
  const { error } = await client.PUT("/config", {
    body: config,
  });
  if (error) {
    throw new Error(`Failed to update config: ${error}`);
  }
}
