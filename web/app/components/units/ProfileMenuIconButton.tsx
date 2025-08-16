import { AccountCircle } from "@mui/icons-material";
import { IconButton } from "@mui/material";

export default function ProfileMenuIconButton () {
    return <IconButton
        size="large"
        aria-label="account of current user"
        aria-controls="menu-appbar"
        aria-haspopup="true"
        color="inherit"
    >
        <AccountCircle />
    </IconButton>
}