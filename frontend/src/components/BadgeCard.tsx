import { Badge } from '@/types';
import { StyleSheet, Text, View } from 'react-native';

type Props = {
  badge: Badge;
};

export default function BadgeCard({ badge }: Props) {
  return (
    <View style={[styles.card, badge.earned ? styles.earned : styles.locked]}>
      <Text style={styles.title}>{badge.name}</Text>
      <Text style={styles.desc}>{badge.description}</Text>
      <Text style={styles.status}>{badge.earned ? 'Unlocked' : 'Locked'}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  card: { borderRadius: 12, padding: 12, marginBottom: 10 },
  earned: { backgroundColor: '#DCFCE7' },
  locked: { backgroundColor: '#F1F5F9' },
  title: { fontSize: 15, fontWeight: '700', color: '#0F172A' },
  desc: { marginTop: 4, color: '#334155' },
  status: { marginTop: 8, fontWeight: '600', color: '#0F766E' },
});
