import {Link} from "react-router";
import logoUrl from "~/logos/famiphoto-icon001.svg";
import {IconButton} from "@mui/material";

interface AppLogoIconProp {
    to: string
}

export default function AppLogoIcon({ to }: AppLogoIconProp) {
    return <IconButton
        edge="start"
        color="inherit"
        aria-label="FAMIPHOTO"
        component={Link}
        to={to}
        sx={{ mr: 2 }}
    >
        <img src={logoUrl} alt="FAMIPHOTO" style={{ width: 28, height: 28 }} />
    </IconButton>
}
