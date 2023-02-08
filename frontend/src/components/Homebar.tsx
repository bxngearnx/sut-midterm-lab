import { slide as Menu } from "react-burger-menu";
import "./Sidebar.css";
import {
  BiHomeAlt,
  BiUser,
  BiMedal,
  BiLogInCircle,
  BiLogOutCircle
} from "react-icons/bi";

type SidemenuProps = {
  pageWrapId: string;
  outerContainerId: string;
};

export const Homebar = ({ pageWrapId, outerContainerId }: SidemenuProps) => {
  return (
    <Menu>
      <p className="memu-title">Menu</p>
      <a className="menu-item" href="/">
        <BiHomeAlt />
        Home
      </a>
      <a className="menu-item" href="/DataAdmin">
        <BiUser />
        User
      </a>
      <a className="menu-item" href="/CreateAdmin">
        <BiMedal />
        Result
      </a>
      <a className="menu-item" href="/login">
        <BiLogInCircle />
        Log in
      </a>
      
    </Menu>
  );
};