export const calculateStandardDrinks = (servingSizeOz: number, abvPercent: number): number => {
  // Approximate formula requested for MVP:
  // standardDrinks = (servingSizeOz * abvPercent * 0.6) / 12
  const result = (servingSizeOz * abvPercent * 0.6) / 12;
  return Math.round(result * 100) / 100;
};
