import { getMonthlyStats, getStats, getWeeklyStats } from '@/api/stats';
import StatCard from '@/components/StatCard';
import { StatsSummary } from '@/types';
import { useEffect, useState } from 'react';
import { ActivityIndicator, ScrollView, StyleSheet, Text, View } from 'react-native';

export default function StatsScreen() {
  const [allTime, setAllTime] = useState<StatsSummary | null>(null);
  const [weekly, setWeekly] = useState<StatsSummary | null>(null);
  const [monthly, setMonthly] = useState<StatsSummary | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const [a, w, m] = await Promise.all([getStats(), getWeeklyStats(), getMonthlyStats()]);
        setAllTime(a);
        setWeekly(w);
        setMonthly(m);
      } catch (e) {
        setError((e as Error).message);
      }
    })();
  }, []);

  if (!allTime || !weekly || !monthly) {
    if (error) return <View style={styles.center}><Text style={styles.error}>{error}</Text></View>;
    return <View style={styles.center}><ActivityIndicator size="large" /></View>;
  }

  return (
    <ScrollView style={styles.container} contentContainerStyle={styles.content}>
      <Text style={styles.section}>All-Time</Text>
      <View style={styles.grid}>
        <StatCard label="Total Drinks" value={allTime.totalDrinks} />
        <StatCard label="Total Calories" value={allTime.totalCalories} />
        <StatCard label="Std Drinks" value={allTime.totalStandardDrinks.toFixed(2)} />
        <StatCard label="Average ABV" value={`${allTime.averageAbv.toFixed(1)}%`} />
        <StatCard label="Most Common Type" value={allTime.mostCommonDrinkType || 'N/A'} />
        <StatCard label="Favorite" value={allTime.favoriteHighestRated || 'N/A'} />
      </View>

      <Text style={styles.section}>Weekly Summary</Text>
      <View style={styles.grid}>
        <StatCard label="Drinks" value={weekly.totalDrinks} />
        <StatCard label="Calories" value={weekly.totalCalories} />
        <StatCard label="Std Drinks" value={weekly.totalStandardDrinks.toFixed(2)} />
      </View>

      <Text style={styles.section}>Monthly Summary</Text>
      <View style={styles.grid}>
        <StatCard label="Drinks" value={monthly.totalDrinks} />
        <StatCard label="Calories" value={monthly.totalCalories} />
        <StatCard label="Std Drinks" value={monthly.totalStandardDrinks.toFixed(2)} />
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1, backgroundColor: '#F1F5F9' },
  content: { padding: 14 },
  section: { fontSize: 18, fontWeight: '700', color: '#0F172A', marginBottom: 8, marginTop: 8 },
  grid: { flexDirection: 'row', flexWrap: 'wrap', justifyContent: 'space-between' },
  center: { flex: 1, alignItems: 'center', justifyContent: 'center' },
  error: { color: '#B91C1C' },
});
