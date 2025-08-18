import { DataModel, DataSource, DataSourceCache } from "@toolpad/core/Crud";
import { z } from "zod";
import dayjs from "dayjs";
import { Expense } from "./types";

// TODO: Fetch from server instead
const standardCategories = [
  { id: 1, description: "Food", symbol: "🍔" },
  { id: 2, description: "Transport", symbol: "🚗" },
  { id: 3, description: "Shopping", symbol: "🛍️" },
  { id: 4, description: "Utilities & Bills", symbol: "💡" },
  { id: 5, description: "Health & Fitness", symbol: "💊" },
];

const getExpenses = async (): Promise<Expense[]> => {
  const res = await fetch("/api/expenses");
  const data = await res.json();
  return data.expenses;
};

const setExpensesStore = (value: Expense[]) => {
  return localStorage.setItem("employees-store", JSON.stringify(value));
};

export const expensesDataSource: DataSource<Expense> = {
  fields: [
    { field: "description", headerName: "Description", width: 140 },
    {
      field: "amount",
      headerName: "Amount",
      type: "number",
      valueFormatter: (_, row: Expense) => {
        return `${row.amount} ${row.currency || ""}`;
      },
    },
    {
      field: "date",
      headerName: "Date",
      type: "date",
      valueFormatter: (_, row: Expense) => dayjs(row.date).format("YYYY-MM-DD"),
      width: 140,
    },
    {
      field: "categoryId",
      headerName: "Category",
      width: 100,
      type: "singleSelect",
      valueOptions: standardCategories.map((cat) => ({
        value: cat.id,
        label: `${cat.symbol} ${cat.description}`,
      })),
      sortable: false,
      filterable: false,
      valueFormatter: (_, row: Expense) => {
        return `${row.category?.symbol || ""} ${
          row.category?.description || ""
        } `;
      },
    },
  ],
  getMany: async ({ paginationModel, filterModel, sortModel }) => {
    const expensesStore = await getExpenses();

    let filteredExpenses = [...expensesStore];

    // Apply filters (example only)
    if (filterModel?.items?.length) {
      filterModel.items.forEach(({ field, value, operator }) => {
        if (!field || value == null) {
          return;
        }

        filteredExpenses = filteredExpenses.filter((expense) => {
          const expenseValue = expense[field];

          switch (operator) {
            case "contains":
              return String(expenseValue)
                .toLowerCase()
                .includes(String(value).toLowerCase());
            case "equals":
              return expenseValue === value;
            case "startsWith":
              return String(expenseValue)
                .toLowerCase()
                .startsWith(String(value).toLowerCase());
            case "endsWith":
              return String(expenseValue)
                .toLowerCase()
                .endsWith(String(value).toLowerCase());
            case ">":
              return (expenseValue as number) > value;
            case "<":
              return (expenseValue as number) < value;
            default:
              return true;
          }
        });
      });
    }

    // Apply sorting
    if (sortModel?.length) {
      filteredExpenses.sort((a, b) => {
        for (const { field, sort } of sortModel) {
          if ((a[field] as number) < (b[field] as number)) {
            return sort === "asc" ? -1 : 1;
          }
          if ((a[field] as number) > (b[field] as number)) {
            return sort === "asc" ? 1 : -1;
          }
        }
        return 0;
      });
    }

    // Apply pagination
    const start = paginationModel.page * paginationModel.pageSize;
    const end = start + paginationModel.pageSize;
    const paginatedExpenses = filteredExpenses.slice(start, end);

    return {
      items: paginatedExpenses,
      itemCount: filteredExpenses.length,
    };
  },
  // TODO Implement fetch one expense
  getOne: async (expenseId) => {
    throw new Error("Feature coming soon");
  },
  createOne: async (inputData) => {
    const response = await fetch("/api/expenses", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ expense: inputData }),
    });

    if (!response.ok) {
      throw new Error("Failed to create expense in database");
    }

    const newExpense = await response.json();
    console.log(newExpense);

    return newExpense;
  },
  // TODO Implement Update API endpoint
  updateOne: async (expenseId, data) => {
    throw new Error("Feature coming soon");
  },
  deleteOne: async (expenseId) => {
    const response = await fetch(`/api/expenses/${expenseId}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Failed to delete expense");
    }
  },
  validate: z.object({
    description: z
      .string({ required_error: "Description is required" })
      .nonempty("Description is required"),
    amount: z
      .number({ required_error: "Amount is required" })
      .min(1, "Amount must be at least 1"),
    date: z
      .string({ required_error: "Date is required" })
      .nonempty("Date is required"),
    //TODO: Fix category input
    categoryId: z.number().optional(),
  })["~standard"].validate,
};

export const employeesCache = new DataSourceCache();
