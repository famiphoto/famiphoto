import { Outlet, Link } from "react-router";
import {AppBar, Toolbar, Typography, IconButton} from "@mui/material";
import AppLogoIcon from "~/components/units/AppLogoIcon";
import AppHeadTitle from "~/components/units/AppHeadTitle";
import ProfileMenuIconButton from "~/components/units/ProfileMenuIconButton";

export default function photoLayout() {
  return (
    <>
      <AppBar position="static" elevation={0}>
        <Toolbar>
          <AppLogoIcon to="photos/listing" />
          <AppHeadTitle>写真</AppHeadTitle>
            <ProfileMenuIconButton></ProfileMenuIconButton>
        </Toolbar>
      </AppBar>
      <Outlet />
    </>
  );
}
