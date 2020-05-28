export enum Role {
  BASIC = 'BASIC',
  MEMBER = 'MEMBER',
  REGULAR = 'REGULAR',
  MODERATOR = 'MODERATOR',
  ADMIN = 'ADMIN',
}

export interface LockedUser {
  name: string;
  reason: string;
  expiration: number;
  expirationInHour: number;
}

export interface LockedIP {
  ip: string;
  reason: string;
  expiration: number;
  expirationInHour: number;
}
