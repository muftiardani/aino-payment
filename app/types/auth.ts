export interface User {
  id: string;
  email: string;
  full_name: string;
  role: "admin" | "user";
  created_at: string;
  updated_at: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  full_name: string;
}

export interface AuthResponse {
  token: string;
  user: User;
}
