import React from "react";
import { NavLink } from "react-router-dom";
import clsx from "clsx";
import key from "weak-key";

import { Card } from "./card";

import "./pages.scss";

const NavTag = ({ className, children, rootPath, path }) => {
  const classes = clsx("nav-tag", className);
  return (
    <NavLink activeClassName="selected" to={`${rootPath}/${path}`} className={classes}>
      <div className="nav-tag__text">{children}</div>
    </NavLink>
  );
};

const Pages = ({ className, rootPath, navigations, children }) => {
  const classes = clsx("pages", className);
  return (
    <div className={classes}>
      <div className="pages-navigation">
        {navigations.map((el) => (
          <NavTag key={key(el)} rootPath={rootPath} path={el.path}>
            {el.title}
          </NavTag>
        ))}
      </div>
      <div className="pages-content">{children}</div>
    </div>
  );
};

export default Pages;

Pages.Card = Card;
