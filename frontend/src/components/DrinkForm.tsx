import { drinkTypes } from '@/constants/drinkTypes';
import { DrinkLog } from '@/types';
import { useState } from 'react';
import { Pressable, StyleSheet, Text, TextInput, View } from 'react-native';

type Props = {
  submitLabel?: string;
  onSubmit: (values: Omit<DrinkLog, 'id' | 'standardDrinks' | 'createdAt'>) => Promise<void>;
  initialType?: string;
  initialName?: string;
  initialServingSizeOz?: number;
  initialAbvPercent?: number;
  initialCalories?: number;
};

export default function DrinkForm({
  submitLabel = 'Save Log',
  onSubmit,
  initialType = 'Beer',
  initialName = '',
  initialServingSizeOz = 12,
  initialAbvPercent = 5,
  initialCalories = 150,
}: Props) {
  const [drinkName, setDrinkName] = useState(initialName);
  const [drinkType, setDrinkType] = useState(initialType);
  const [servingSizeOz, setServingSizeOz] = useState(String(initialServingSizeOz));
  const [abvPercent, setAbvPercent] = useState(String(initialAbvPercent));
  const [calories, setCalories] = useState(String(initialCalories));
  const [location, setLocation] = useState('');
  const [notes, setNotes] = useState('');
  const [rating, setRating] = useState('4');
  const [saving, setSaving] = useState(false);

  const submit = async () => {
    setSaving(true);
    try {
      await onSubmit({
        // TODO: Replace with authenticated user context when auth is added.
        userId: 1,
        drinkName,
        drinkType,
        servingSizeOz: Number(servingSizeOz),
        abvPercent: Number(abvPercent),
        calories: Number(calories),
        location,
        notes,
        rating: Number(rating),
      });
      setLocation('');
      setNotes('');
    } finally {
      setSaving(false);
    }
  };

  return (
    <View style={styles.card}>
      <Text style={styles.sectionTitle}>Manual Drink Log</Text>
      <TextInput style={styles.input} placeholder="Drink name" value={drinkName} onChangeText={setDrinkName} />
      <Text style={styles.typesLabel}>Type</Text>
      <View style={styles.typeRow}>
        {drinkTypes.map((type) => (
          <Pressable key={type} style={[styles.chip, drinkType === type && styles.chipActive]} onPress={() => setDrinkType(type)}>
            <Text style={[styles.chipText, drinkType === type && styles.chipTextActive]}>{type}</Text>
          </Pressable>
        ))}
      </View>
      <TextInput style={styles.input} placeholder="Serving size oz" keyboardType="decimal-pad" value={servingSizeOz} onChangeText={setServingSizeOz} />
      <TextInput style={styles.input} placeholder="ABV %" keyboardType="decimal-pad" value={abvPercent} onChangeText={setAbvPercent} />
      <TextInput style={styles.input} placeholder="Calories" keyboardType="number-pad" value={calories} onChangeText={setCalories} />
      <TextInput style={styles.input} placeholder="Location" value={location} onChangeText={setLocation} />
      <TextInput style={[styles.input, styles.multiline]} placeholder="Notes" value={notes} onChangeText={setNotes} multiline />
      <TextInput style={styles.input} placeholder="Rating 0-5" keyboardType="number-pad" value={rating} onChangeText={setRating} />
      <Pressable style={styles.button} onPress={submit} disabled={saving}>
        <Text style={styles.buttonText}>{saving ? 'Saving...' : submitLabel}</Text>
      </Pressable>
    </View>
  );
}

const styles = StyleSheet.create({
  card: { backgroundColor: '#fff', borderRadius: 16, padding: 14, marginTop: 10 },
  sectionTitle: { fontSize: 17, fontWeight: '700', marginBottom: 10, color: '#0F172A' },
  input: { borderWidth: 1, borderColor: '#CBD5E1', borderRadius: 10, padding: 10, marginBottom: 10, backgroundColor: '#F8FAFC' },
  multiline: { minHeight: 60, textAlignVertical: 'top' },
  typesLabel: { marginBottom: 6, color: '#334155', fontWeight: '600' },
  typeRow: { flexDirection: 'row', flexWrap: 'wrap', gap: 8, marginBottom: 10 },
  chip: { borderWidth: 1, borderColor: '#CBD5E1', borderRadius: 999, paddingHorizontal: 10, paddingVertical: 6 },
  chipActive: { backgroundColor: '#DBEAFE', borderColor: '#60A5FA' },
  chipText: { color: '#334155' },
  chipTextActive: { color: '#1D4ED8', fontWeight: '600' },
  button: { backgroundColor: '#1D4ED8', borderRadius: 10, paddingVertical: 12, alignItems: 'center' },
  buttonText: { color: '#fff', fontWeight: '700' },
});
