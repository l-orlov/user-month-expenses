import { combineReducers } from "redux";

import users from "./users/reducer";
import user from "./user/reducer";

const appReducer = combineReducers({ users, user });

export function rootReducer(state, action) {
  if (action.type === "CLEAR_STORE") {
    state = undefined;
  }

  return appReducer(state, action);
}
