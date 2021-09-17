import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import createSagaMiddleware from "redux-saga";
import { createLogger } from "redux-logger";
import { reduxBatch } from "@manaflair/redux-batch";

import { rootReducer } from "./reducers";
import { rootSaga } from "./saga";

const sagaMiddleware = createSagaMiddleware();
const logger = createLogger({
  diff: true,
  level: "log",
  duration: true,
});
const middlewares = [sagaMiddleware, logger];
const middleware = [...getDefaultMiddleware({ thunk: false }), ...middlewares];

const configureAppStore = (preloadedState) =>
  configureStore({
    reducer: rootReducer,
    enhancers: [reduxBatch],
    preloadedState,
    middleware,
  });

if (process.env.NODE_ENV !== "production" && module.hot) {
  module.hot.accept("./reducers", () => store.replaceReducer(rootReducer));
}

const store = configureAppStore({});

sagaMiddleware.run(rootSaga);

export default store;
