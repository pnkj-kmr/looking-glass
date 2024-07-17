export interface IP {
  ip: string;
  host: string;
  port: number;
  username: string;
  vendor: string;
}

export interface IPInfo extends IP {
  id: number;
  password: string;
}

export interface Vendor {
  id: number;
  name: string;
}

export interface Protocol {
  id: number;
  name: string;
}

export interface SrcHost {
  ip: string;
  host: string;
}

export interface Result {
  completed: boolean;
  message: string;
  is_error: boolean;
}

export interface Audit {
  access_count: number;
  query_count: number;
}
