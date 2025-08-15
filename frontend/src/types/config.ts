type Country = "GB";

export interface Config {
  apiKey: string;
  country: Country;
}

export function newConfig(): Config {
  return {
    apiKey: "",
    country: "GB",
  };
}
