import { DataGrid, GridColDef } from '@mui/x-data-grid';
import { Expense } from '../../types';

export default function CustomizedDataGrid({
  expenses,
}: {
  expenses: Expense[];
}) {
  // TODO: Remove isMobile if not needed
  // const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

  // const columnsToShow = isMobile ? columns.slice(0, 3) : columns;
  const columns: GridColDef[] = [
    {
      field: 'date',
      headerName: 'Date',
      width: 150,
      valueFormatter: (_, row: Expense) => {
        return new Date(row.date).toLocaleDateString('sv-SE');
      },
    },
    { field: 'description', headerName: 'Description', width: 150 },
    { field: 'amount', headerName: 'Amount', width: 100 },
    { field: 'currency', headerName: 'Currency', width: 100 },
    {
      field: 'category',
      headerName: 'Category',
      width: 100,
      valueGetter: (_, row: Expense) => {
        console.log('test', row.category?.symbol);
        return `${row.category?.symbol || ''} ${
          row.category?.description || ''
        } `;
      },
    },
  ];

  const rows = expenses;

  return (
    <DataGrid
      // checkboxSelection
      disableColumnMenu
      rows={rows}
      columns={columns}
      getRowClassName={(params) =>
        params.indexRelativeToCurrentPage % 2 === 0 ? 'even' : 'odd'
      }
      initialState={{
        pagination: { paginationModel: { pageSize: 20 } },
      }}
      pageSizeOptions={[10, 20, 50]}
      disableColumnResize
      density="compact"
      slotProps={{
        filterPanel: {
          filterFormProps: {
            logicOperatorInputProps: {
              variant: 'outlined',
              size: 'small',
            },
            columnInputProps: {
              variant: 'outlined',
              size: 'small',
              sx: { mt: 'auto' },
            },
            operatorInputProps: {
              variant: 'outlined',
              size: 'small',
              sx: { mt: 'auto' },
            },
            valueInputProps: {
              InputComponentProps: {
                variant: 'outlined',
                size: 'small',
              },
            },
          },
        },
      }}
    />
  );
}
