import * as React from "react";
import { Crud } from "@toolpad/core/Crud";
import { useParams } from "react-router";
import { expensesDataSource, employeesCache } from "../data/expenses";
import { Expense } from "../data/types";

export default function ExpensesCrudPage() {
  const { expenseId } = useParams();

  return (
    <Crud<Expense>
      dataSource={expensesDataSource}
      // dataSourceCache={employeesCache}
      rootPath="/expenses"
      initialPageSize={25}
      defaultValues={{ currency: "SEK" }}
      pageTitles={{
        show: `Expense ${expenseId}`,
        create: "New Expense",
        edit: `Expense ${expenseId}`,
      }}
    />
  );
}
