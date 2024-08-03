import { useRadio, VisuallyHidden, cn, RadioProps } from "@nextui-org/react";

export const CustomRadio = (props: RadioProps) => {
  const {
    Component,
    children,
    description,
    getBaseProps,
    getWrapperProps,
    getInputProps,
    getLabelProps,
    getLabelWrapperProps,
    getControlProps,
  } = useRadio(props);

  return (
    <Component
      {...getBaseProps()}
      className={cn(
        "group flex items-center hover:opacity-70 active:opacity-50 justify-between flex-col-reverse tap-highlight-transparent",
        "max-w-[300px] cursor-pointer border-2 border-default rounded-lg gap-2 p-1",
        "data-[selected=true]:border-primary"
      )}
    >
      <VisuallyHidden>
        <input {...getInputProps()} />
      </VisuallyHidden>
      <span {...getWrapperProps()}>
        <span {...getControlProps()} />
      </span>
      <div
        {...getLabelWrapperProps()}
        className="flex justify-center flex-col w-full h-full"
      >
        {children && (
          <span
            {...getLabelProps()}
            className={cn("w-full h-full flex items-center ", props.className)}
          >
            {children}
          </span>
        )}
        {description && (
          <span className="text-md font-semibold text-foreground opacity-70 text-center">
            {description}
          </span>
        )}
      </div>
    </Component>
  );
};
