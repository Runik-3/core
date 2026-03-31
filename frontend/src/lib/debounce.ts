export function debounce<T>(
  callback: (...args: unknown[]) => Promise<T>,
): (...args: unknown[]) => Promise<T>;

export function debounce<T>(
  callback: (...args: unknown[]) => Promise<T>,
  wait: number
): (...args: unknown[]) => Promise<T>;

export function debounce<T>(
  callback: (...args: unknown[]) => T,
  wait: number = 500,
) {
  let timeoutId = null;
  return (...args: unknown[]) => {
    if (timeoutId) {
      clearTimeout(timeoutId);
    }
    timeoutId = setTimeout(() => {
      return callback(...args);
    }, wait);
  };
}
