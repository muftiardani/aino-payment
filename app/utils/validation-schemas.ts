import { z } from 'zod'

// ============================================
// Auth Validation Schemas
// ============================================

export const loginSchema = z.object({
  email: z.string().min(1, 'Email is required').email('Invalid email format'),
  password: z.string().min(1, 'Password is required'),
})

export const registerSchema = z
  .object({
    full_name: z
      .string()
      .min(2, 'Full name must be at least 2 characters')
      .max(100, 'Full name must not exceed 100 characters'),
    email: z.string().min(1, 'Email is required').email('Invalid email format'),
    password: z
      .string()
      .min(6, 'Password must be at least 6 characters')
      .max(100, 'Password must not exceed 100 characters'),
    confirm_password: z.string().min(1, 'Please confirm your password'),
  })
  .refine(data => data.password === data.confirm_password, {
    message: 'Passwords do not match',
    path: ['confirm_password'],
  })

export const forgotPasswordSchema = z.object({
  email: z.string().min(1, 'Email is required').email('Invalid email format'),
})

export const resetPasswordSchema = z
  .object({
    token: z.string().min(1, 'Reset token is required'),
    new_password: z
      .string()
      .min(6, 'Password must be at least 6 characters')
      .max(100, 'Password must not exceed 100 characters'),
    confirm_password: z.string().min(1, 'Please confirm your password'),
  })
  .refine(data => data.new_password === data.confirm_password, {
    message: 'Passwords do not match',
    path: ['confirm_password'],
  })

// ============================================
// Payment Validation Schemas
// ============================================

export const createPaymentSchema = z.object({
  amount: z
    .union([z.string(), z.number()])
    .transform(val => (typeof val === 'string' ? Number.parseFloat(val) : val))
    .refine(val => !Number.isNaN(val) && val > 0, {
      message: 'Amount must be greater than 0',
    })
    .refine(val => val <= 999999999, {
      message: 'Amount must not exceed 999,999,999',
    }),
  category_id: z.string().min(1, 'Please select a category'),
  payment_method_id: z.string().min(1, 'Please select a payment method'),
  transaction_date: z.string().min(1, 'Transaction date is required'),
  description: z.string().max(500, 'Description must not exceed 500 characters').optional(),
})

export const updatePaymentSchema = z.object({
  amount: z
    .union([z.string(), z.number()])
    .transform(val => (typeof val === 'string' ? Number.parseFloat(val) : val))
    .refine(val => !Number.isNaN(val) && val > 0, {
      message: 'Amount must be greater than 0',
    })
    .refine(val => val <= 999999999, {
      message: 'Amount must not exceed 999,999,999',
    }),
  status: z.enum(['pending', 'completed', 'failed', 'refunded'], {
    message: 'Please select a valid status',
  }),
  category_id: z.string().min(1, 'Please select a category'),
  payment_method_id: z.string().min(1, 'Please select a payment method'),
  transaction_date: z.string().min(1, 'Transaction date is required'),
  description: z.string().max(500, 'Description must not exceed 500 characters').optional(),
})

// ============================================
// Type Exports (for TypeScript inference)
// ============================================

export type LoginFormData = z.infer<typeof loginSchema>
export type RegisterFormData = z.infer<typeof registerSchema>
export type ForgotPasswordFormData = z.infer<typeof forgotPasswordSchema>
export type ResetPasswordFormData = z.infer<typeof resetPasswordSchema>
export type CreatePaymentFormData = z.infer<typeof createPaymentSchema>
export type UpdatePaymentFormData = z.infer<typeof updatePaymentSchema>
