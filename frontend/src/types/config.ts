export interface Config {
  apiKey: string;
  country: string;
  notification: NotificationConfig;
  global: GlobalConfig;
}

export interface NotificationConfig {
  ntfy: NtfyConfig;
}

export interface NtfyConfig {
  url?: string;
  topic?: string;
  username?: string;
  password?: string;
}

export interface GlobalConfig {
  regions?: string[];
  eventSimilarity?: number;
  numTickets?: number; // Must be an integer
  maxTicketPrice?: number;
  discount?: number;
}

export function newConfig(): Config {
  return {
    apiKey: "",
    country: "GB",
    notification: {
      ntfy: {},
    },
    global: {},
  };
}
