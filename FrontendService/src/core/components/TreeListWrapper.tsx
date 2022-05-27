import { FC } from "react";

type Props = {
  wrappers: FC[];
}

export const TreeListWrapper: FC<Props> = ({ wrappers, children }) => {
  if(!wrappers.length) return <>{children}</>;

  const Component = wrappers[0];
  return (
    <Component>
      <TreeListWrapper wrappers={wrappers.slice(1)}>
        {children}
      </TreeListWrapper>
    </Component>
  )
}