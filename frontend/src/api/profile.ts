import { apiFetch } from './client';
import { Goal, ProfileResponse } from '@/types';

export const getProfile = () => apiFetch<ProfileResponse>('/profile');
export const updateGoals = (goals: Goal) => apiFetch<ProfileResponse>('/profile/goals', {
  method: 'PUT',
  body: JSON.stringify(goals),
});
