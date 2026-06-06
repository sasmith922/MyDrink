import { createDrink, getDrinks } from '@/api/drinks';
import { createLog } from '@/api/logs';
import CatalogueItem from '@/components/CatalogueItem';
import { Drink } from '@/types';
import { calculateStandardDrinks } from '@/utils/drinkCalculations';
import { useEffect, useMemo, useState } from 'react';
import { ActivityIndicator, Alert, Pressable, ScrollView, StyleSheet, Text, TextInput, View } from 'react-native';

export default function CatalogueScreen() {
  const [drinks, setDrinks] = useState<Drink[]>([]);
  const [query, setQuery] = useState('');
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [name, setName] = useState('');
  const [type, setType] = useState('Beer');
  const [servingSize, setServingSize] = useState('12');
  const [abv, setAbv] = useState('5');
  const [calories, setCalories] = useState('150');
  const [description, setDescription] = useState('');

  const load = async () => {
    try {
      setError(null);
      setDrinks(await getDrinks());
    } catch (e) {
      setError((e as Error).message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    void load();
  }, []);

  const filtered = useMemo(
    () => drinks.filter((d) => d.name.toLowerCase().includes(query.toLowerCase()) || d.type.toLowerCase().includes(query.toLowerCase())),
    [drinks, query],
  );

  const addDrink = async () => {
    try {
      await createDrink({
        name,
        type,
        defaultServingSizeOz: Number(servingSize),
        abvPercent: Number(abv),
        calories: Number(calories),
        description,
      });
      setName('');
      setDescription('');
      await load();
    } catch (e) {
      Alert.alert('Unable to add drink', (e as Error).message);
    }
  };

  const logDrink = async (drink: Drink) => {
    try {
      await createLog({
        // TODO: Replace with authenticated user context when auth is added.
        userId: 1,
        drinkId: drink.id,
        drinkName: drink.name,
        drinkType: drink.type,
        servingSizeOz: drink.defaultServingSizeOz,
        abvPercent: drink.abvPercent,
        calories: drink.calories,
        standardDrinks: calculateStandardDrinks(drink.defaultServingSizeOz, drink.abvPercent),
        location: 'Catalogue Quick Log',
        notes: `Logged from catalogue: ${drink.name}`,
        rating: 4,
      });
      Alert.alert('Logged', `${drink.name} added to feed.`);
    } catch (e) {
      Alert.alert('Unable to log drink', (e as Error).message);
    }
  };

  if (loading) return <View style={styles.center}><ActivityIndicator size="large" /></View>;

  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <TextInput style={styles.input} placeholder="Search drinks" value={query} onChangeText={setQuery} />
      {error ? <Text style={styles.error}>{error}</Text> : null}
      {filtered.map((drink) => <CatalogueItem key={drink.id} drink={drink} onLog={logDrink} />)}

      <View style={styles.sectionCard}>
        <Text style={styles.sectionTitle}>Add Custom Drink</Text>
        <TextInput style={styles.input} placeholder="Name" value={name} onChangeText={setName} />
        <TextInput style={styles.input} placeholder="Type" value={type} onChangeText={setType} />
        <TextInput style={styles.input} placeholder="Serving size oz" value={servingSize} onChangeText={setServingSize} keyboardType="decimal-pad" />
        <TextInput style={styles.input} placeholder="ABV %" value={abv} onChangeText={setAbv} keyboardType="decimal-pad" />
        <TextInput style={styles.input} placeholder="Calories" value={calories} onChangeText={setCalories} keyboardType="number-pad" />
        <TextInput style={styles.input} placeholder="Description" value={description} onChangeText={setDescription} />
        <Pressable style={styles.button} onPress={addDrink}>
          <Text style={styles.buttonText}>Save Drink</Text>
        </Pressable>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#F1F5F9' },
  content: { padding: 14 },
  sectionCard: { marginTop: 16, backgroundColor: '#fff', borderRadius: 14, padding: 14 },
  sectionTitle: { fontSize: 18, fontWeight: '700', color: '#0F172A', marginBottom: 10 },
  input: { borderWidth: 1, borderColor: '#CBD5E1', borderRadius: 10, padding: 10, marginBottom: 10, backgroundColor: '#fff' },
  button: { backgroundColor: '#1D4ED8', borderRadius: 10, paddingVertical: 12, alignItems: 'center' },
  buttonText: { color: '#fff', fontWeight: '700' },
  center: { flex: 1, alignItems: 'center', justifyContent: 'center' },
  error: { color: '#B91C1C', marginBottom: 10 },
});
