import { DataModel } from "@toolpad/core";

export type Category = {
  id: number;
  userId: string;
  description: string;
  symbol: string;
  createdAt: string;
};

export interface Expense extends DataModel {
  id: number;
  description: string;
  amount: number;
  currency: string;
  categoryId?: number;
  category?: Category;
  date: string;
}
