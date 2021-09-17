import React from "react";
import clsx from "clsx";

import "./title.scss";

export const Title = ({ className, children }) => {
  const classes = clsx("title", className);
  return <h2 className={classes}>{children}</h2>;
};
