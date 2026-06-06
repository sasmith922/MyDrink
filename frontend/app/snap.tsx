import { createLog } from '@/api/logs';
import DrinkForm from '@/components/DrinkForm';
import { calculateStandardDrinks } from '@/utils/drinkCalculations';
import { useState } from 'react';
import { Pressable, ScrollView, StyleSheet, Text, View } from 'react-native';

export default function SnapScreen() {
  const [showForm, setShowForm] = useState(false);
  const [status, setStatus] = useState<string | null>(null);

  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <View style={styles.card}>
        <Text style={styles.title}>AI drink recognition coming soon.</Text>
        <Text style={styles.text}>
          We intentionally have no camera upload/computer vision implementation in this MVP.
        </Text>
        {/* Future integration point: AI image recognition workflow can be triggered from this screen. */}
        <Pressable style={styles.button} onPress={() => setShowForm((v) => !v)}>
          <Text style={styles.buttonText}>Add Drink Manually</Text>
        </Pressable>
        {status ? <Text style={styles.status}>{status}</Text> : null}
      </View>

      {showForm ? (
        <DrinkForm
          onSubmit={async (values) => {
            await createLog({
              ...values,
              standardDrinks: calculateStandardDrinks(values.servingSizeOz, values.abvPercent),
            });
            setStatus('Drink logged and added to feed.');
          }}
        />
      ) : null}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#F1F5F9' },
  content: { padding: 14 },
  card: { backgroundColor: '#fff', borderRadius: 16, padding: 16 },
  title: { fontSize: 20, fontWeight: '700', color: '#0F172A' },
  text: { marginTop: 8, color: '#334155', lineHeight: 20 },
  button: { marginTop: 12, backgroundColor: '#1D4ED8', borderRadius: 10, paddingVertical: 12, alignItems: 'center' },
  buttonText: { color: '#fff', fontWeight: '700' },
  status: { marginTop: 10, color: '#047857', fontWeight: '600' },
});
