import "./App.scss";
import Pages from "./components/pages";
import { AllUsersController } from "./layouts/allUsers/";
import { CurrentUserController } from "./layouts/currentUser";
const navigations = [
  { path: "all", title: "По всем" },
  { path: "current", title: "По одному" },
];

function App() {
  return (
    <div className="App">
      <Pages rootPath="/statistics" navigations={navigations}>
        <AllUsersController rootPath="/statistics" path="all" />
        <CurrentUserController rootPath="/statistics" path="current" />
      </Pages>
    </div>
  );
}

export default App;
