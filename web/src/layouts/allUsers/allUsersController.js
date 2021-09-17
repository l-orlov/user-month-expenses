import React, { useEffect, useCallback } from "react";
import { useDispatch, useSelector } from "react-redux";
import { sagaEventCallBegan } from "../../model/saga";
import { fetchUsersError, getUsers } from "../../model/users/reducer";
import { AllUsers } from "./allUsers";

export const AllUsersController = ({ ...props }) => {
  const dispatch = useDispatch();
  const users = useSelector((state) => state.users.users);

  const handleOnUpdate = useCallback(() => {
    dispatch({
      url: `/api/v1/expenses/by-categories?size=5`,
      type: sagaEventCallBegan.type,
      method: "get",
      onSuccess: getUsers.type,
      onError: fetchUsersError.type,
    });
  }, [dispatch]);

  useEffect(() => {
    dispatch({
      url: `/api/v1/expenses/by-categories?size=5`,
      type: sagaEventCallBegan.type,
      method: "get",
      onSuccess: getUsers.type,
      onError: fetchUsersError.type,
    });
  }, [dispatch]);

  return <AllUsers {...props} onUpdate={handleOnUpdate} users={users} />;
};
