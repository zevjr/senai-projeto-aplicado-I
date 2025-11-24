export interface User {
  uid: string;
  email: string;
  username?: string;
  role: 'operator' | 'admin' | string;
  created_at?: string;
  deleted_at?: string | null;
}
