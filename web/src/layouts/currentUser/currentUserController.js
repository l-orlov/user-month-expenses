import React, { useState, useEffect, useCallback } from "react";
import { useDispatch, useSelector } from "react-redux";
import { sagaEventCallBegan } from "../../model/saga";
import { fetchUserError, getUser } from "../../model/user/reducer";
import { CurrentUser } from "./currentUser";

export const CurrentUserController = ({ ...props }) => {
  const [state, setState] = useState("");
  const dispatch = useDispatch();
  const user = useSelector((state) => state.user.user);
  console.log(user);
  const handleOnChange = useCallback(({ target }) => {
    setState(target.value);
  }, []);

  const handleOnSubmit = useCallback(() => {
    dispatch({
      url: `/api/v1/expenses/by-categories?size=5&userId=${state}`,
      type: sagaEventCallBegan.type,
      method: "get",
      onSuccess: getUser.type,
      onError: fetchUserError.type,
    });
  }, [dispatch, state]);

  useEffect(() => {
    dispatch({
      url: `/api/v1/expenses/by-categories?size=5&userId=${state}`,
      type: sagaEventCallBegan.type,
      method: "get",
      onSuccess: getUser.type,
      onError: fetchUserError.type,
    });
  }, [dispatch, state]);

  return (
    <CurrentUser
      {...props}
      onSubmit={handleOnSubmit}
      user={user}
      id={state}
      onChange={handleOnChange}
    />
  );
};
