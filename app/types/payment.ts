export interface Payment {
  id: string;
  user_id: string;
  amount: number;
  status: "pending" | "completed" | "failed" | "refunded";
  payment_method_id: string;
  payment_method?: PaymentMethod;
  category_id: string;
  category?: Category;
  description: string;
  transaction_date: string;
  created_at: string;
  updated_at: string;
}

export interface PaymentMethod {
  id: string;
  name: string;
  code: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface Category {
  id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

export interface CreatePaymentRequest {
  amount: number;
  payment_method_id: string;
  category_id: string;
  description: string;
  transaction_date: string;
}

export interface UpdatePaymentRequest {
  amount: number;
  status: "pending" | "completed" | "failed" | "refunded";
  payment_method_id: string;
  category_id: string;
  description: string;
  transaction_date: string;
}

export interface PaymentListResponse {
  payments: Payment[];
  total: number;
  page: number;
  limit: number;
}

export interface DashboardStats {
  total_payments: number;
  completed_count: number;
  pending_count: number;
  total_amount: number;
}
