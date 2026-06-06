import { Drink } from '@/types';
import { Pressable, StyleSheet, Text, View } from 'react-native';

type Props = {
  drink: Drink;
  onLog: (drink: Drink) => void;
};

export default function CatalogueItem({ drink, onLog }: Props) {
  return (
    <View style={styles.card}>
      <Text style={styles.title}>{drink.name}</Text>
      <Text style={styles.meta}>{drink.type} · {drink.defaultServingSizeOz}oz · {drink.abvPercent}% ABV</Text>
      <Text style={styles.meta}>{drink.calories} cal</Text>
      <Text style={styles.desc}>{drink.description}</Text>
      <Pressable style={styles.button} onPress={() => onLog(drink)}>
        <Text style={styles.buttonText}>Log this drink</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  card: { backgroundColor: '#fff', borderRadius: 14, padding: 14, marginBottom: 10 },
  title: { fontSize: 16, fontWeight: '700', color: '#0F172A' },
  meta: { marginTop: 4, color: '#334155' },
  desc: { marginTop: 6, color: '#64748B' },
  button: { marginTop: 10, backgroundColor: '#E0E7FF', alignSelf: 'flex-start', paddingHorizontal: 12, paddingVertical: 8, borderRadius: 10 },
  buttonText: { color: '#3730A3', fontWeight: '600' },
});
