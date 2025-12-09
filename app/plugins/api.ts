import type { ApiResponse } from "~/types/api";

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig();
  const authStore = useAuthStore();

  const api = async <T = any>(
    endpoint: string,
    options: RequestInit & { params?: Record<string, any> } = {}
  ): Promise<ApiResponse<T>> => {
    const { params, ...fetchOptions } = options;

    // Build URL with query params
    let url = `${config.public.apiBase}${endpoint}`;
    if (params) {
      const query = new URLSearchParams(params).toString();
      url += `?${query}`;
    }

    // Add auth token if available
    const headers: HeadersInit = {
      "Content-Type": "application/json",
      ...fetchOptions.headers,
    };

    if (authStore.token) {
      headers["Authorization"] = `Bearer ${authStore.token}`;
    }

    try {
      const response = await fetch(url, {
        ...fetchOptions,
        headers,
      });

      const data = await response.json();

      if (!response.ok) {
        // Handle 401 Unauthorized
        if (response.status === 401) {
          authStore.logout();
        }

        return {
          success: false,
          error: data.error || "An error occurred",
        };
      }

      return data as ApiResponse<T>;
    } catch (error) {
      console.error("API Error:", error);
      return {
        success: false,
        error: error instanceof Error ? error.message : "Network error",
      };
    }
  };

  return {
    provide: {
      api,
    },
  };
});
