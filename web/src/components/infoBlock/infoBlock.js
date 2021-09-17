import React from "react";
import clsx from "clsx";

import "./infoBlock.scss";

const InfoBlock = ({ className, title, children }) => {
  const classes = clsx("info-block", className);
  return (
    <div className={classes}>
      <span>{title}:</span>
      <span>{children}</span>
    </div>
  );
};

export default InfoBlock;
