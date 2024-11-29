import { CSSProperties } from "react";
import { Column } from "@tanstack/react-table";

export const getCommonPinningStyles = <TData>(
  column: Column<TData>
): CSSProperties => {
  const isPinned = column.getIsPinned();
  const isLastLeftPinnedColumn = column.getIsLastColumn("left");
  const isFirstRightPinnedColumn = column.getIsFirstColumn("right");
  return {
    boxShadow: isLastLeftPinnedColumn
      ? "-4px 0 4px -4px #E6E6E8 inset"
      : isFirstRightPinnedColumn
      ? "4px 0 4px -4px #E6E6E8 inset"
      : undefined,
    left: isPinned === "left" ? `${column.getStart("left")}px` : undefined,
    right: isPinned === "right" ? `${column.getAfter("right")}px` : undefined,
    position: isPinned ? "sticky" : "relative",
    minWidth: column.getSize(),
    // maxWidth: column.getSize(),
    // width: column.getSize(),
    zIndex: isPinned ? 1 : 0,
  };
};
