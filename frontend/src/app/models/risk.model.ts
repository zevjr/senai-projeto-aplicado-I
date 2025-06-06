export interface Risk {
  id: number;
  title: string;
  description: string;
  location?: string;
  riskLevel?: string;
  riskScale?: number;
  vehicle?: string;
  area?: string;
  person?: string;
  criticality?: string;
  solutions?: string;
  mediaUrls?: string[];
  createdAt?: Date;
  createdBy?: string;
}