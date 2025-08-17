import { z } from 'zod';

export const UserSchema = z.object({
  id: z.string(),
  email: z.string().email(),
  createdAt: z.string(), // ISO string
});
export type User = z.infer<typeof UserSchema>;

export const CategorySchema = z.object({
  id: z.number(),
  userId: z.string().optional(),
  description: z.string(),
  symbol: z.string(),
  createdAt: z.string(),
});
export type Category = z.infer<typeof CategorySchema>;

export const ExpenseSchema = z.object({
  id: z.number(),
  amount: z.number().min(0),
  currency: z.string(),
  description: z.string(),
  category_id: z.number().nullable().optional(),
  category: CategorySchema.nullable().optional(),
  date: z.string(), // ISO string
  createdAt: z.string(),
});
export type Expense = z.infer<typeof ExpenseSchema>;

export const ExpenseCategoryAggregateSchema = z.object({
  category_id: z.number().nullable(),
  category_description: z.string().nullable(),
  category_symbol: z.string().nullable(),
  amount_sum: z.number().min(0),
});

export type ExpenseCategoryAggregate = z.infer<
  typeof ExpenseCategoryAggregateSchema
>;

export const GetExpensesResponseSchema = z.object({
  expenses: z.array(ExpenseSchema),
  categoryAggregates: z.array(ExpenseCategoryAggregateSchema),
});
export type GetExpensesResponse = z.infer<typeof GetExpensesResponseSchema>;

export const AddExpenseParamsSchema = z.object({
  expense: ExpenseSchema,
});
export type AddExpenseParams = z.infer<typeof AddExpenseParamsSchema>;

export const AddExpenseResponseSchema = z.object({
  expense: ExpenseSchema,
});
export type AddExpenseResponse = z.infer<typeof AddExpenseResponseSchema>;

export const FieldErrorSchema = z.object({
  field: z.string().optional(),
  message: z.string(),
});
export type FieldError = z.infer<typeof FieldErrorSchema>;

export const ErrorResponseSchema = z.object({
  error: z.string(),
  details: z.array(FieldErrorSchema).optional(),
});
export type ErrorResponse = z.infer<typeof ErrorResponseSchema>;
