import { getProfile, updateGoals } from '@/api/profile';
import BadgeCard from '@/components/BadgeCard';
import { Goal, ProfileResponse } from '@/types';
import { useEffect, useState } from 'react';
import { ActivityIndicator, Pressable, ScrollView, StyleSheet, Text, TextInput, View } from 'react-native';

export default function ProfileScreen() {
  const [profile, setProfile] = useState<ProfileResponse | null>(null);
  const [goals, setGoals] = useState<Goal>({
    weeklyCalorieGoal: 0,
    weeklyStandardDrinkGoal: 0,
    newDrinksToTryGoal: 0,
  });
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const data = await getProfile();
        setProfile(data);
        setGoals(data.goals);
      } catch (e) {
        setError((e as Error).message);
      }
    })();
  }, []);

  if (!profile) {
    if (error) return <View style={styles.center}><Text style={styles.error}>{error}</Text></View>;
    return <View style={styles.center}><ActivityIndicator size="large" /></View>;
  }

  const saveGoals = async () => {
    setSaving(true);
    try {
      const updated = await updateGoals(goals);
      setProfile(updated);
      setGoals(updated.goals);
    } catch (e) {
      setError((e as Error).message);
    } finally {
      setSaving(false);
    }
  };

  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <View style={styles.card}>
        <Text style={styles.username}>{profile.user.username}</Text>
        <Text style={styles.meta}>Favorite drink: {profile.user.favoriteDrink}</Text>
        <Text style={styles.meta}>Lifetime drinks: {profile.user.lifetimeDrinksLogged}</Text>
        <Text style={styles.meta}>Lifetime calories: {profile.user.lifetimeCalories}</Text>
        <Text style={styles.meta}>Lifetime std drinks: {profile.user.lifetimeStandardDrinks.toFixed(2)}</Text>
      </View>

      <View style={styles.card}>
        <Text style={styles.sectionTitle}>Goals</Text>
        <TextInput
          style={styles.input}
          keyboardType="number-pad"
          value={String(goals.weeklyCalorieGoal)}
          onChangeText={(t) => setGoals((g) => ({ ...g, weeklyCalorieGoal: Number(t || 0) }))}
          placeholder="Weekly calorie goal"
        />
        <TextInput
          style={styles.input}
          keyboardType="decimal-pad"
          value={String(goals.weeklyStandardDrinkGoal)}
          onChangeText={(t) => setGoals((g) => ({ ...g, weeklyStandardDrinkGoal: Number(t || 0) }))}
          placeholder="Weekly standard drink goal"
        />
        <TextInput
          style={styles.input}
          keyboardType="number-pad"
          value={String(goals.newDrinksToTryGoal)}
          onChangeText={(t) => setGoals((g) => ({ ...g, newDrinksToTryGoal: Number(t || 0) }))}
          placeholder="New drinks to try goal"
        />
        <Pressable style={styles.button} onPress={saveGoals}>
          <Text style={styles.buttonText}>{saving ? 'Saving...' : 'Save Goals'}</Text>
        </Pressable>
      </View>

      <Text style={styles.badgesTitle}>Badges</Text>
      {profile.badges.map((badge) => <BadgeCard key={badge.name} badge={badge} />)}
      {error ? <Text style={styles.error}>{error}</Text> : null}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#F1F5F9' },
  content: { padding: 14 },
  card: { backgroundColor: '#fff', borderRadius: 14, padding: 14, marginBottom: 12 },
  username: { fontSize: 22, fontWeight: '700', color: '#0F172A' },
  meta: { marginTop: 6, color: '#334155' },
  sectionTitle: { fontSize: 18, fontWeight: '700', marginBottom: 10, color: '#0F172A' },
  input: { borderWidth: 1, borderColor: '#CBD5E1', borderRadius: 10, padding: 10, marginBottom: 10 },
  button: { backgroundColor: '#1D4ED8', borderRadius: 10, paddingVertical: 12, alignItems: 'center' },
  buttonText: { color: '#fff', fontWeight: '700' },
  badgesTitle: { fontSize: 18, fontWeight: '700', marginBottom: 8, color: '#0F172A' },
  center: { flex: 1, alignItems: 'center', justifyContent: 'center' },
  error: { color: '#B91C1C', marginTop: 8 },
});
