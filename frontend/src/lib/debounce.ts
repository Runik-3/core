export function debounce(callback: (...args: unknown[]) => void, wait: number = 500) {
  let timeoutId = null
  return (...args: unknown[]) => {
    if (timeoutId) {
      clearTimeout(timeoutId)
    }
    timeoutId = setTimeout(() => {
      callback(...args)
    }, wait)
  }
}

