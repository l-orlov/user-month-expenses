import React from "react";
import { Route } from "react-router-dom";
import clsx from "clsx";

import "./card.scss";

export const Card = ({ className, rootPath, path, children }) => {
  const classes = clsx("card", className);
  return (
    <Route path={`${rootPath}/${path}`}>
      <div className={classes}>{children}</div>
    </Route>
  );
};
