import { getFeed, likePost } from '@/api/feed';
import DrinkCard from '@/components/DrinkCard';
import { FeedPost } from '@/types';
import { useCallback, useEffect, useState } from 'react';
import { ActivityIndicator, RefreshControl, ScrollView, StyleSheet, Text, View } from 'react-native';

export default function FeedScreen() {
  const [feed, setFeed] = useState<FeedPost[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const load = useCallback(async () => {
    try {
      setError(null);
      const data = await getFeed();
      setFeed(data);
    } catch (e) {
      setError((e as Error).message);
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    void load();
  }, [load]);

  const onLike = async (id: number) => {
    setFeed((prev) => prev.map((p) => (p.id === id ? { ...p, likeCount: p.likeCount + 1 } : p)));
    try {
      const updated = await likePost(id);
      setFeed((prev) => prev.map((p) => (p.id === id ? updated : p)));
    } catch {
      await load();
    }
  };

  if (loading) return <View style={styles.center}><ActivityIndicator size="large" /></View>;

  return (
    <ScrollView
      style={styles.container}
      contentContainerStyle={styles.content}
      refreshControl={<RefreshControl refreshing={false} onRefresh={() => void load()} />}
    >
      {error ? <Text style={styles.error}>Failed to load feed: {error}</Text> : null}
      {feed.map((post) => <DrinkCard key={post.id} post={post} onLike={onLike} />)}
      {!feed.length && !error ? <Text style={styles.empty}>No feed posts yet.</Text> : null}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#F1F5F9' },
  content: { padding: 14 },
  center: { flex: 1, alignItems: 'center', justifyContent: 'center' },
  error: { color: '#B91C1C', marginBottom: 10 },
  empty: { textAlign: 'center', color: '#475569' },
});
