import { apiFetch } from './client';
import { FeedPost } from '@/types';

export const getFeed = () => apiFetch<FeedPost[]>('/feed');
export const likePost = (id: number) => apiFetch<FeedPost>(`/feed/${id}/like`, { method: 'POST' });
