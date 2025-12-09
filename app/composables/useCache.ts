import { useStorage } from '@vueuse/core'

interface CacheEntry<T> {
  data: T
  timestamp: number
}

export const useCache = <T>(key: string, ttl: number = 300000) => {
  // Use localStorage with 'ainopay_' prefix
  const storageKey = `ainopay_cache_${key}`

  const cache = useStorage<CacheEntry<T> | null>(
    storageKey,
    null,
    import.meta.client ? localStorage : undefined,
    { mergeDefaults: true }
  )

  const isExpired = (): boolean => {
    if (!cache.value) return true
    return Date.now() - cache.value.timestamp > ttl
  }

  const get = (): T | null => {
    if (!cache.value) return null

    if (isExpired()) {
      clear()
      return null
    }

    return cache.value.data
  }

  const set = (data: T): void => {
    cache.value = {
      data,
      timestamp: Date.now(),
    }
  }

  const clear = (): void => {
    cache.value = null
  }

  return { get, set, clear, isExpired }
}
