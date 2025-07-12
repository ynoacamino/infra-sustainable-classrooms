import { cn } from '@/lib/shared/utils';

/**
 * Table component
 *
 * A styled table component with consistent design system integration.
 * Includes responsive overflow handling and border styling.
 *
 * @param props - Table component props
 * @param props.className - Additional CSS classes
 * @returns The rendered table component
 *
 * @example
 * ```tsx
 * <Table>
 *   <TableHeader>
 *     <TableRow>
 *       <TableHead>Name</TableHead>
 *       <TableHead>Email</TableHead>
 *     </TableRow>
 *   </TableHeader>
 *   <TableBody>
 *     <TableRow>
 *       <TableCell>John Doe</TableCell>
 *       <TableCell>john@example.com</TableCell>
 *     </TableRow>
 *   </TableBody>
 * </Table>
 * ```
 */
function Table({ className, ...props }: React.ComponentProps<'table'>) {
  return (
    <div
      data-slot="table-container"
      className="relative w-full overflow-x-auto border rounded-md"
    >
      <table
        data-slot="table"
        className={cn('w-full caption-bottom text-sm', className)}
        {...props}
      />
    </div>
  );
}

/**
 * TableHeader component
 *
 * Header section of a table, typically containing column headers.
 *
 * @param props - Table header props
 * @param props.className - Additional CSS classes
 * @returns The rendered table header component
 *
 * @example
 * ```tsx
 * <TableHeader>
 *   <TableRow>
 *     <TableHead>Column 1</TableHead>
 *     <TableHead>Column 2</TableHead>
 *   </TableRow>
 * </TableHeader>
 * ```
 */
function TableHeader({ className, ...props }: React.ComponentProps<'thead'>) {
  return (
    <thead
      data-slot="table-header"
      className={cn('[&_tr]:border-b', className)}
      {...props}
    />
  );
}

/**
 * TableBody component
 *
 * Body section of a table containing the main data rows.
 *
 * @param props - Table body props
 * @param props.className - Additional CSS classes
 * @returns The rendered table body component
 *
 * @example
 * ```tsx
 * <TableBody>
 *   <TableRow>
 *     <TableCell>Data 1</TableCell>
 *     <TableCell>Data 2</TableCell>
 *   </TableRow>
 * </TableBody>
 * ```
 */
function TableBody({ className, ...props }: React.ComponentProps<'tbody'>) {
  return (
    <tbody
      data-slot="table-body"
      className={cn('[&_tr:last-child]:border-0', className)}
      {...props}
    />
  );
}

/**
 * TableFooter component
 *
 * Footer section of a table, typically containing summary information.
 *
 * @param props - Table footer props
 * @param props.className - Additional CSS classes
 * @returns The rendered table footer component
 *
 * @example
 * ```tsx
 * <TableFooter>
 *   <TableRow>
 *     <TableCell>Total</TableCell>
 *     <TableCell>$100.00</TableCell>
 *   </TableRow>
 * </TableFooter>
 * ```
 */
function TableFooter({ className, ...props }: React.ComponentProps<'tfoot'>) {
  return (
    <tfoot
      data-slot="table-footer"
      className={cn(
        'bg-muted/50 border-t font-medium [&>tr]:last:border-b-0',
        className,
      )}
      {...props}
    />
  );
}

/**
 * TableRow component
 *
 * Individual row within a table. Supports hover states and selection styling.
 *
 * @param props - Table row props
 * @param props.className - Additional CSS classes
 * @returns The rendered table row component
 *
 * @example
 * ```tsx
 * <TableRow>
 *   <TableCell>Cell 1</TableCell>
 *   <TableCell>Cell 2</TableCell>
 * </TableRow>
 *
 * <TableRow data-state="selected">
 *   <TableCell>Selected row</TableCell>
 * </TableRow>
 * ```
 */
function TableRow({ className, ...props }: React.ComponentProps<'tr'>) {
  return (
    <tr
      data-slot="table-row"
      className={cn(
        'hover:bg-muted/50 data-[state=selected]:bg-muted border-b transition-colors',
        className,
      )}
      {...props}
    />
  );
}

/**
 * TableHead component
 *
 * Header cell component for table columns. Used within TableHeader rows.
 *
 * @param props - Table head props
 * @param props.className - Additional CSS classes
 * @returns The rendered table head component
 *
 * @example
 * ```tsx
 * <TableHead>Column Name</TableHead>
 * <TableHead className="text-right">Price</TableHead>
 * ```
 */
function TableHead({ className, ...props }: React.ComponentProps<'th'>) {
  return (
    <th
      data-slot="table-head"
      className={cn(
        'text-foreground h-10 px-2 text-left align-middle font-medium whitespace-nowrap [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]',
        className,
      )}
      {...props}
    />
  );
}

/**
 * TableCell component
 *
 * Data cell component for table content. Used within TableBody rows.
 *
 * @param props - Table cell props
 * @param props.className - Additional CSS classes
 * @returns The rendered table cell component
 *
 * @example
 * ```tsx
 * <TableCell>Cell content</TableCell>
 * <TableCell className="text-right">$99.99</TableCell>
 * ```
 */
function TableCell({ className, ...props }: React.ComponentProps<'td'>) {
  return (
    <td
      data-slot="table-cell"
      className={cn(
        'p-2 align-middle whitespace-nowrap [&:has([role=checkbox])]:pr-0 [&>[role=checkbox]]:translate-y-[2px]',
        className,
      )}
      {...props}
    />
  );
}

/**
 * TableCaption component
 *
 * Caption component for providing a table description or title.
 *
 * @param props - Table caption props
 * @param props.className - Additional CSS classes
 * @returns The rendered table caption component
 *
 * @example
 * ```tsx
 * <Table>
 *   <TableCaption>A list of recent transactions</TableCaption>
 *   <TableHeader>
 *     <!-- header content -->
 *   </TableHeader>
 *   <!-- body content -->
 * </Table>
 * ```
 */
function TableCaption({
  className,
  ...props
}: React.ComponentProps<'caption'>) {
  return (
    <caption
      data-slot="table-caption"
      className={cn('text-muted-foreground mt-4 text-sm', className)}
      {...props}
    />
  );
}

export {
  Table,
  TableHeader,
  TableBody,
  TableFooter,
  TableHead,
  TableRow,
  TableCell,
  TableCaption,
};
