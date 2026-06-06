import { apiFetch } from './client';
import { DrinkLog } from '@/types';

export const getLogs = () => apiFetch<DrinkLog[]>('/logs');

export const createLog = (body: Omit<DrinkLog, 'id' | 'standardDrinks' | 'createdAt'> & { standardDrinks?: number }) =>
  apiFetch<DrinkLog>('/logs', { method: 'POST', body: JSON.stringify(body) });
