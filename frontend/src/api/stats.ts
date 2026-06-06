import { apiFetch } from './client';
import { StatsSummary } from '@/types';

export const getStats = () => apiFetch<StatsSummary>('/stats');
export const getWeeklyStats = () => apiFetch<StatsSummary>('/stats/weekly');
export const getMonthlyStats = () => apiFetch<StatsSummary>('/stats/monthly');
