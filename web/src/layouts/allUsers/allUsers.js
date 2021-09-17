import React from "react";
import clsx from "clsx";
import key from "weak-key";

import InfoBlock from "../../components/infoBlock";
import Pages from "../../components/pages";
import { Title } from "../../components/title";

import "./allUsers.scss";
import { Button } from "../../components/button";

const { Card } = Pages;

export const AllUsers = ({ className, rootPath, path, onUpdate, users }) => {
  const classes = clsx("all-users", className);
  return (
    <Card className={classes} rootPath={rootPath} path={path}>
      <Title>Данные по всем пользователям</Title>
      <div className="all-users-content">
        <Button className="all-users-update" onClick={onUpdate}>
          Обновить данные
        </Button>
        <div className="all-users-list">
          {users.map((el) => (
            <InfoBlock key={key(el)} title={el.category}>
              {el.amount}
            </InfoBlock>
          ))}
        </div>
      </div>
    </Card>
  );
};
