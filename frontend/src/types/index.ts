export type User = {
  id: number;
  username: string;
  favoriteDrink: string;
  weeklyCalorieGoal: number;
  weeklyStandardDrinkGoal: number;
  newDrinksToTryGoal: number;
  lifetimeDrinksLogged: number;
  lifetimeCalories: number;
  lifetimeStandardDrinks: number;
  goalsMetCurrentWeek: boolean;
  distinctDrinksLoggedLifetime: number;
};

export type Drink = {
  id: number;
  name: string;
  type: string;
  defaultServingSizeOz: number;
  abvPercent: number;
  calories: number;
  description: string;
};

export type DrinkLog = {
  id: number;
  userId: number;
  drinkId?: number;
  drinkName: string;
  drinkType: string;
  servingSizeOz: number;
  abvPercent: number;
  calories: number;
  standardDrinks: number;
  location: string;
  notes: string;
  rating: number;
  createdAt: string;
};

export type FeedPost = {
  id: number;
  userId: number;
  userName: string;
  drinkLogId: number;
  drinkName: string;
  drinkType: string;
  abvPercent: number;
  calories: number;
  standardDrinks: number;
  location: string;
  notes: string;
  timestamp: string;
  likeCount: number;
};

export type Goal = {
  weeklyCalorieGoal: number;
  weeklyStandardDrinkGoal: number;
  newDrinksToTryGoal: number;
};

export type Badge = {
  name: string;
  description: string;
  earned: boolean;
};

export type StatsSummary = {
  rangeLabel: string;
  totalDrinks: number;
  totalCalories: number;
  totalStandardDrinks: number;
  averageAbv: number;
  mostCommonDrinkType: string;
  favoriteHighestRated: string;
};

export type ProfileResponse = {
  user: User;
  goals: Goal;
  badges: Badge[];
};
