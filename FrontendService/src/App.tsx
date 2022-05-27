import { Core, InitParams } from "./core/Core";
import { registerPlugins } from "./plugins/registerPlugins";
import { ToastContainer, ToastContainerProps } from "react-toastify";

import "react-toastify/dist/ReactToastify.css";
import axios from "axios";

const toastrOptions: ToastContainerProps = {
  position: "top-center",
  autoClose: 5000,
  hideProgressBar: true,
  newestOnTop: false,
  closeOnClick: true,
  rtl: false,
  pauseOnFocusLoss: true,
  draggable: true,
  pauseOnHover: true,
};

axios.defaults.withCredentials = true;

const App = () => {
  const init = ({ registerPlugin }: InitParams) => {
    registerPlugins(registerPlugin);
  };

  return (
    <>
      <Core initialize={init} />
      <ToastContainer {...toastrOptions} />
    </>
  );
};

export default App;
