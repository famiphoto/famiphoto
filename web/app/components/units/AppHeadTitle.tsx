import {Typography} from "@mui/material";

export default function AppHeadTitle({children}: {children: React.ReactNode}) {
    return <Typography variant="h6" component="h1" sx={{ flexGrow: 1 }}>
        {children}
    </Typography>
}
