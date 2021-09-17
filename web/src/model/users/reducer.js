import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  users: [],
};

const usersSlice = createSlice({
  name: "users",
  initialState,
  reducers: {
    getUsers(state, action) {
      state.users = action.payload;
    },

    fetchUsersError(state, action) {
      state.error = action.payload;
    },
  },
});

export default usersSlice.reducer;

export const { getUsers, fetchUsersError } = usersSlice.actions;
