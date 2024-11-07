export function fuzzyMatch(text: string, query: string): boolean {
  const pattern = query.split('').map(char => 
    char.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')
  ).join('.*');
  const regex = new RegExp(pattern, 'i');
  return regex.test(text);
}

export function fuzzyScore(text: string, query: string): number {
  text = text.toLowerCase();
  query = query.toLowerCase();
  
  if (query.length === 0) return 0;
  if (text === query) return 1;
  
  let score = 0;
  let textIndex = 0;
  
  for (const queryChar of query) {
    while (textIndex < text.length && text[textIndex] !== queryChar) {
      textIndex++;
    }
    if (textIndex < text.length) {
      score += 1 / (textIndex + 1);
      textIndex++;
    }
  }
  
  return score / query.length;
} 