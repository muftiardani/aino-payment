/**
 * Log levels enum
 */
export enum LogLevel {
  DEBUG = 0,
  INFO = 1,
  WARN = 2,
  ERROR = 3,
  NONE = 4,
}

/**
 * Log entry interface
 */
export interface LogEntry {
  timestamp: string
  level: string
  context: string
  message: string
  data?: unknown
  url?: string
}

/**
 * Format timestamp for logs
 */
export const formatTimestamp = (): string => {
  return new Date().toISOString()
}

/**
 * Get color for log level (console styling)
 */
export const getColorForLevel = (level: string): string => {
  const colors: Record<string, string> = {
    DEBUG: '#6B7280', // gray
    INFO: '#3B82F6', // blue
    WARN: '#F59E0B', // yellow
    ERROR: '#EF4444', // red
  }
  return colors[level] || '#6B7280'
}

/**
 * Sanitize sensitive data from logs
 */
export const sanitizeData = (data: unknown): unknown => {
  if (!data || typeof data !== 'object') {
    return data
  }

  const sanitized = Array.isArray(data) ? [...data] : { ...data }
  const sensitiveFields = [
    'password',
    'token',
    'apiKey',
    'secret',
    'authorization',
    'auth',
    'accessToken',
    'refreshToken',
  ]

  const sanitizeObject = (obj: unknown): unknown => {
    if (!obj || typeof obj !== 'object') return obj

    if (Array.isArray(obj)) {
      return obj.map(item => sanitizeObject(item))
    }

    const result: Record<string, unknown> = { ...obj }

    for (const key in result) {
      // Check if key is sensitive
      const lowerKey = key.toLowerCase()
      if (sensitiveFields.some(field => lowerKey.includes(field))) {
        result[key] = '***REDACTED***'
      }
      // Recursively sanitize nested objects
      else if (typeof result[key] === 'object' && result[key] !== null) {
        result[key] = sanitizeObject(result[key])
      }
    }

    return result
  }

  return sanitizeObject(sanitized)
}

/**
 * Truncate large objects for logging
 */
export const truncateData = (data: unknown, maxLength = 1000): unknown => {
  const str = JSON.stringify(data)
  if (str.length <= maxLength) {
    return data
  }

  return {
    _truncated: true,
    _originalLength: str.length,
    _preview: str.substring(0, maxLength) + '...',
  }
}

/**
 * Format log message for console output
 */
export const formatLogMessage = (
  level: string,
  context: string,
  message: string
): { formatted: string; styles: string[] } => {
  const timestamp = formatTimestamp()
  const color = getColorForLevel(level)

  const formatted = `%c[${timestamp}]%c %c[${level}]%c %c[${context}]%c ${message}`

  const styles = [
    'color: #9CA3AF', // timestamp - gray
    '',
    `color: ${color}; font-weight: bold`, // level - colored
    '',
    'color: #8B5CF6; font-weight: bold', // context - purple
    '',
  ]

  return { formatted, styles }
}

/**
 * Create structured log entry
 */
export const createLogEntry = (
  level: string,
  context: string,
  message: string,
  data?: unknown
): LogEntry => {
  return {
    timestamp: formatTimestamp(),
    level,
    context,
    message,
    data: data ? sanitizeData(truncateData(data)) : undefined,
    url: typeof window !== 'undefined' ? window.location.href : undefined,
  }
}
