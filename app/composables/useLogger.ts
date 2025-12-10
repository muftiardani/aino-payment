import { LogLevel, formatLogMessage, createLogEntry } from '~/utils/logFormatter'

/**
 * Logger composable for centralized logging
 * @param context - Context/module name for the logger
 */
export function useLogger(context: string = 'App') {
  // Get log level from config or environment
  const currentLogLevel = import.meta.dev ? LogLevel.DEBUG : LogLevel.WARN

  /**
   * Check if log level should be logged
   */
  const shouldLog = (level: LogLevel): boolean => {
    return level >= currentLogLevel
  }

  /**
   * Internal log function
   */
  const log = (level: LogLevel, levelName: string, message: string, data?: unknown) => {
    if (!shouldLog(level)) return

    // Create structured log entry
    const entry = createLogEntry(levelName, context, message, data)

    // Format message for console
    const { formatted, styles } = formatLogMessage(levelName, context, message)

    // Log to console with appropriate method
    switch (level) {
      case LogLevel.DEBUG:
        if (data) {
          console.debug(formatted, ...styles, '\n', entry)
        } else {
          console.debug(formatted, ...styles)
        }
        break

      case LogLevel.INFO:
        if (data) {
          console.info(formatted, ...styles, '\n', entry)
        } else {
          console.info(formatted, ...styles)
        }
        break

      case LogLevel.WARN:
        if (data) {
          console.warn(formatted, ...styles, '\n', entry)
        } else {
          console.warn(formatted, ...styles)
        }
        break

      case LogLevel.ERROR:
        if (data) {
          console.error(formatted, ...styles, '\n', entry)
        } else {
          console.error(formatted, ...styles)
        }
        break
    }

    // TODO: Send to external logging service (Sentry, Datadog, etc.)
    // if (import.meta.client && level >= LogLevel.ERROR) {
    //   sendToExternalService(entry)
    // }
  }

  /**
   * Debug level logging
   */
  const debug = (message: string, data?: unknown) => {
    log(LogLevel.DEBUG, 'DEBUG', message, data)
  }

  /**
   * Info level logging
   */
  const info = (message: string, data?: unknown) => {
    log(LogLevel.INFO, 'INFO', message, data)
  }

  /**
   * Warning level logging
   */
  const warn = (message: string, data?: unknown) => {
    log(LogLevel.WARN, 'WARN', message, data)
  }

  /**
   * Error level logging
   */
  const error = (message: string, data?: unknown) => {
    log(LogLevel.ERROR, 'ERROR', message, data)
  }

  /**
   * Start a console group
   */
  const group = (label: string) => {
    if (shouldLog(LogLevel.DEBUG)) {
      console.group(`[${context}] ${label}`)
    }
  }

  /**
   * End a console group
   */
  const groupEnd = () => {
    if (shouldLog(LogLevel.DEBUG)) {
      console.groupEnd()
    }
  }

  /**
   * Display data as table
   */
  const table = (data: unknown) => {
    if (shouldLog(LogLevel.DEBUG)) {
      console.table(data)
    }
  }

  /**
   * Time a function execution
   */
  const time = (label: string) => {
    if (shouldLog(LogLevel.DEBUG)) {
      console.time(`[${context}] ${label}`)
    }
  }

  /**
   * End timing
   */
  const timeEnd = (label: string) => {
    if (shouldLog(LogLevel.DEBUG)) {
      console.timeEnd(`[${context}] ${label}`)
    }
  }

  return {
    debug,
    info,
    warn,
    error,
    group,
    groupEnd,
    table,
    time,
    timeEnd,
  }
}
