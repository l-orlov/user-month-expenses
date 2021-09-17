import React, { useMemo } from "react";
import PropTypes from "prop-types";
import clsx from "clsx";
import key from "weak-key";

import "./input.scss";

export const Input = React.memo(({ title, className, ...rest }) => {
  const classes = clsx("input", className, { "disabled": rest.disabled });
  const id = useMemo(() => key({ [`checkbox-${Date.now()}`]: `checkbox-${Date.now()}` }), []);
  return (
    <label htmlFor={id} className={classes}>
      <span>{title}</span>
      <input id={id} {...rest} />
    </label>
  );
});

Input.propTypes = {
  title: PropTypes.string,
  className: PropTypes.string,
};

Input.defaultProps = {
  title: "",
  className: "",
};
