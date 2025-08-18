import * as React from "react";
import PieChartIcon from "@mui/icons-material/PieChart";
import ReceiptLongIcon from "@mui/icons-material/ReceiptLong";
import { Outlet } from "react-router";
import { ReactRouterAppProvider } from "@toolpad/core/react-router";
import type { Navigation } from "@toolpad/core/AppProvider";

const NAVIGATION: Navigation = [
  {
    kind: "header",
    title: "Main items",
  },
  {
    title: "Dashboard",
    icon: <PieChartIcon />,
  },
  {
    segment: "expenses",
    title: "Expenses",
    icon: <ReceiptLongIcon />,
    pattern: "expenses{/:expenseId}*",
  },
];

const BRANDING = {
  title: "Cashcape",
  logo: <img src="./cashcape.svg" alt="logo" />,
};

export default function App() {
  return (
    <ReactRouterAppProvider navigation={NAVIGATION} branding={BRANDING}>
      <Outlet />
    </ReactRouterAppProvider>
  );
}
