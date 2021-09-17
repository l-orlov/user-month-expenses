import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  user: [],
};

const userSlice = createSlice({
  name: "users",
  initialState,
  reducers: {
    getUser(state, action) {
      state.user = action.payload;
    },

    fetchUserError(state, action) {
      state.error = action.payload;
    },
  },
});

export default userSlice.reducer;

export const { getUser, fetchUserError } = userSlice.actions;
