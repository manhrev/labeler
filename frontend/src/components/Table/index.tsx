import React from 'react'
import { ColumnDef, InitialTableState, getCoreRowModel, getSortedRowModel, flexRender, useReactTable } from '@tanstack/react-table'
import { Table as TableRoot, TableBody, TableColumn, TableHeader, TableRow, TableCell } from '@nextui-org/react'
import { getCommonPinningStyles } from './util'

interface TableProps<TData> extends React.HTMLAttributes<HTMLDivElement> {
    data: TData[]
    columns: ColumnDef<TData>[]
    initialState?: InitialTableState
    tableContainerClassName?: string
}

function Table<TData>({ data, columns, className, tableContainerClassName, initialState }: TableProps<TData>) {
    const table = useReactTable({
        data: data,
        columns: columns,
        getCoreRowModel: getCoreRowModel(),
        getSortedRowModel: getSortedRowModel(),
        initialState
    })

    return (
        <div className={tableContainerClassName}>
            <TableRoot className={className}>
                <TableHeader className="sticky top-0">
                    {table.getHeaderGroups().map((headerGroup) => (
                        <TableRow key={headerGroup.id}>
                            {headerGroup.headers.map((header) => {
                                return (
                                    <TableColumn
                                        key={header.id}
                                        //   className={header.column.columnDef.meta?.headerClassName}
                                        style={getCommonPinningStyles(header.column)}
                                    >
                                        {header.isPlaceholder ? null : flexRender(header.column.columnDef.header, header.getContext())}
                                    </TableColumn>
                                )
                            })}
                        </TableRow>
                    ))}
                </TableHeader>
                <TableBody>
                    {table.getRowModel().rows.map((row) => (
                        <TableRow key={row.id} data-state={row.getIsSelected() && 'selected'}>
                            {row.getVisibleCells().map((cell) => {
                                return (
                                    <TableCell style={getCommonPinningStyles(cell.column)} key={cell.id}>
                                        {flexRender(cell.column.columnDef.cell, cell.getContext())}
                                    </TableCell>
                                )
                            })}
                        </TableRow>
                    ))}
                </TableBody>
            </TableRoot>
        </div>
    )
}

export default Table
