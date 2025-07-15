import { API_BASE_URL } from '@/config/shared/env';

export function mapToFile(url: string) {
  const repeatedQuotes = url.startsWith('"') && url.endsWith('"');
  const cleanedUrl = repeatedQuotes ? url.slice(1, -1) : url;
  console.log(cleanedUrl);
  const urlObj = new URL(cleanedUrl);
  return `${API_BASE_URL}/files/${urlObj.pathname}?${urlObj.searchParams.toString()}`; // TA BIEN XD
}
