import * as React from "react";
import Typography from "@mui/material/Typography";
import { PageContainer } from "@toolpad/core/PageContainer";
import { PieChart } from "@mui/x-charts/PieChart";
import { BarChart, PieValueType } from "@mui/x-charts";

type CategoryAggregate = {
  categoryId: number | null;
  categoryDescription: string | null;
  categorySymbol: string | null;
  amountSum: number;
};

function generateColors(count: number) {
  // Spread hues evenly around the color wheel
  return Array.from(
    { length: count },
    (_, i) => `hsl(${(360 / count) * i}, 60%, 55%)`
  );
}

export default function DashboardPage() {
  const [expenses, setExpenses] = React.useState<PieValueType[]>([]);

  React.useEffect(() => {
    fetch("/api/expenses")
      .then((res) => res.json())
      .then((data) => {
        const aggregates: CategoryAggregate[] = data.categoryAggregates || [];
        setExpenses(
          aggregates.map((expense) => ({
            label: `${expense.categoryDescription ?? "Uncategorized"} ${expense.categorySymbol ?? ""}`,
            value: expense.amountSum,
          }))
        );
      });
  }, []);

  const colors = generateColors(expenses.length);

  const barData = expenses.map((item) => item.value);
  const barLabels = expenses.map((item) => item.label);

  return (
    <PageContainer>
      {/* <Typography>Welcome to Cashcape!</Typography> */}
      <PieChart
        colors={colors}
        margin={{
          left: 80,
          right: 80,
          top: 80,
          bottom: 80,
        }}
        series={[
          {
            data: expenses,
            innerRadius: 75,
            outerRadius: 100,
            paddingAngle: 0,
            highlightScope: { fade: "global", highlight: "item" },
          },
        ]}
        height={260}
        width={260}
      ></PieChart>
      <BarChart
        series={[{ data: barData }]}
        xAxis={[
          {
            colorMap: { type: "ordinal", colors: colors },
            data: barLabels,
            scaleType: "band",
            tickLabelStyle: { angle: 90, fontSize: 0 },
            disableTicks: true,
          },
        ]}
        height={260}
        width={400}
      />
    </PageContainer>
  );
}
