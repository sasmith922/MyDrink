import { FeedPost } from '@/types';
import { Pressable, StyleSheet, Text, View } from 'react-native';

type Props = {
  post: FeedPost;
  onLike: (id: number) => void;
};

export default function DrinkCard({ post, onLike }: Props) {
  return (
    <View style={styles.card}>
      <Text style={styles.title}>{post.userName} logged {post.drinkName}</Text>
      <Text style={styles.meta}>{post.drinkType} · {post.abvPercent}% ABV · {post.calories} cal</Text>
      <Text style={styles.meta}>Std Drinks: {post.standardDrinks.toFixed(2)}</Text>
      <Text style={styles.meta}>Location: {post.location || 'N/A'}</Text>
      <Text style={styles.notes}>{post.notes || 'No notes'}</Text>
      <Text style={styles.timestamp}>{new Date(post.timestamp).toLocaleString()}</Text>
      <Pressable style={styles.likeBtn} onPress={() => onLike(post.id)}>
        <Text style={styles.likeText}>❤️ {post.likeCount}</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  card: { backgroundColor: '#fff', borderRadius: 14, padding: 14, marginBottom: 12, shadowColor: '#000', shadowOpacity: 0.06, shadowRadius: 6, elevation: 2 },
  title: { fontSize: 16, fontWeight: '700', color: '#0F172A' },
  meta: { marginTop: 4, color: '#334155' },
  notes: { marginTop: 8, color: '#475569' },
  timestamp: { marginTop: 8, color: '#64748B', fontSize: 12 },
  likeBtn: { marginTop: 10, alignSelf: 'flex-start', backgroundColor: '#EFF6FF', borderRadius: 10, paddingHorizontal: 12, paddingVertical: 8 },
  likeText: { color: '#1D4ED8', fontWeight: '600' },
});
