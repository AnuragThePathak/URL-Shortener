import * as React from 'react'
import AppBar from '@mui/material/AppBar'
import Toolbar from '@mui/material/Toolbar'
import Typography from '@mui/material/Typography'
import Container from '@mui/material/Container'
import AddLinkOutlinedIcon from '@mui/icons-material/AddLinkOutlined'

const ResponsiveAppBar = () => {

	return (
		<AppBar position="static">
			<Container maxWidth="xl" >
				<Toolbar disableGutters>
					<AddLinkOutlinedIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
					<Typography
						variant="h6"
						noWrap
						component="a"
						href="/"
						sx={{
							mr: 2,
							display: { xs: 'none', md: 'flex' },
							fontFamily: 'monospace',
							fontWeight: 700,
							letterSpacing: '.1rem',
							color: 'inherit',
							textDecoration: 'none',
						}}
					>
						Minly
					</Typography>
						
					<AddLinkOutlinedIcon sx={{ display: { xs: 'flex', md: 'none' }, mr: 1 }} />
					<Typography
						variant="h5"
						noWrap
						component="a"
						href=""
						sx={{
							mr: 2,
							display: { xs: 'flex', md: 'none' },
							flexGrow: 1,
							fontFamily: 'monospace',
							fontWeight: 700,
							letterSpacing: '.1rem',
							color: 'inherit',
							textDecoration: 'none',
						}}
					>
						Minly
					</Typography>
				</Toolbar>
			</Container>
		</AppBar>
	)
}

export default ResponsiveAppBar