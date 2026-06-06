import { apiFetch } from './client';
import { Drink } from '@/types';

export const getDrinks = () => apiFetch<Drink[]>('/drinks');
export const getDrink = (id: number) => apiFetch<Drink>(`/drinks/${id}`);
export const createDrink = (body: Omit<Drink, 'id'>) =>
  apiFetch<Drink>('/drinks', { method: 'POST', body: JSON.stringify(body) });
