export interface Config {
  apiKey: string;
  country: string;
  notification: NotificationConfig;
  global: CommonConfig;
  tickets: TicketConfig[];
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

export interface CommonConfig {
  regions?: string[];
  eventSimilarity?: number;
  numTickets?: number; // Must be an integer
  maxTicketPrice?: number;
  discount?: number;
}

export interface TicketConfig extends CommonConfig {
  event?: string;
}

export function newConfig(): Config {
  return {
    apiKey: "",
    country: "GB",
    notification: {
      ntfy: {},
    },
    global: {},
    tickets: [],
  };
}
