import React from "react";
import PropTypes from "prop-types";
import clsx from "clsx";

import "./button.scss";

export const Button = ({ className, children, ...rest }) => {
  const classes = clsx("button", className, { "disabled": rest.disabled });
  return (
    <button className={classes} {...rest}>
      {children}
    </button>
  );
};

Button.propTypes = {
  className: PropTypes.string,
  children: PropTypes.node.isRequired,
};

Button.defaultProps = {
  className: "",
};
