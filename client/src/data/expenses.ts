import { DataModel, DataSource, DataSourceCache } from "@toolpad/core/Crud";
import { z } from "zod";

type EmployeeRole = "Market" | "Finance" | "Development";

type Category = {
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
  role: EmployeeRole;
}

const INITIAL_EMPLOYEES_STORE: Expense[] = [
  {
    id: 1,
    description: "Edward Perry",
    amount: 25,
    date: new Date().toISOString(),
    role: "Finance",
    currency: "",
  },
  {
    id: 2,
    description: "Josephine Drake",
    amount: 36,
    date: new Date().toISOString(),
    role: "Market",
    currency: "",
  },
  {
    id: 3,
    description: "Cody Phillips",
    amount: 19,
    date: new Date().toISOString(),
    role: "Development",
    currency: "",
  },
];

const getExpenses = async (): Promise<Expense[]> => {
  const res = await fetch("/api/expenses");
  const data = await res.json();
  return data.expenses;

  // const value = localStorage.getItem("employees-store");
  // return value ? JSON.parse(value) : INITIAL_EMPLOYEES_STORE;
};

const setExpensesStore = (value: Expense[]) => {
  return localStorage.setItem("employees-store", JSON.stringify(value));
};

export const expensesDataSource: DataSource<Expense> = {
  fields: [
    // { field: "id", headerName: "ID" },
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
      valueFormatter: (_, row: Expense) => {
        return new Date(row.date).toLocaleDateString("sv-SE");
      },
      width: 140,
    },
    {
      field: "category",
      headerName: "Category",
      width: 100,
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
    // Simulate loading delay
    await new Promise((resolve) => {
      setTimeout(resolve, 750);
    });

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
  getOne: async (expenseId) => {
    // Simulate loading delay
    await new Promise((resolve) => {
      setTimeout(resolve, 750);
    });

    const expensesStore = await getExpenses();

    const expenseToShow = expensesStore.find(
      (employee) => employee.id === Number(expenseId)
    );

    if (!expenseToShow) {
      throw new Error("Employee not found");
    }
    return expenseToShow;
  },
  createOne: async (data) => {
    // Simulate loading delay
    await new Promise((resolve) => {
      setTimeout(resolve, 750);
    });

    const expensesStore = await getExpenses();

    const newExpense = {
      id:
        expensesStore.reduce((max, expense) => Math.max(max, expense.id), 0) +
        1,
      ...data,
    } as Expense;

    setExpensesStore([...expensesStore, newExpense]);

    return newExpense;
  },
  updateOne: async (expenseId, data) => {
    // Simulate loading delay
    await new Promise((resolve) => {
      setTimeout(resolve, 750);
    });

    const expensesStore = await getExpenses();

    let updatedExpense: Expense | null = null;

    setExpensesStore(
      expensesStore.map((expense) => {
        if (expense.id === Number(expenseId)) {
          updatedExpense = { ...expense, ...data };
          return updatedExpense;
        }
        return expense;
      })
    );

    if (!updatedExpense) {
      throw new Error("Expense not found");
    }
    return updatedExpense;
  },
  deleteOne: async (expenseId) => {
    // Simulate loading delay
    await new Promise((resolve) => {
      setTimeout(resolve, 750);
    });

    const employeesStore = await getExpenses();

    setExpensesStore(
      employeesStore.filter((expense) => expense.id !== Number(expenseId))
    );
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
    category: z.enum(["Market", "Finance", "Development"], {
      errorMap: () => ({
        message: 'Role must be "Market", "Finance" or "Development"',
      }),
    }),
  })["~standard"].validate,
};

export const employeesCache = new DataSourceCache();
