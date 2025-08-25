export interface Config {
  apiKey: string;
  country: string;
}

export function newConfig(): Config {
  return {
    apiKey: "",
    country: "GB",
  };
}
