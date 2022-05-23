import * as React from 'react'
import Typography from '@mui/material/Typography'
import MuiLink from '@mui/material/Link'

export default function Copyright() {
	return (
		<Typography variant="body2" color="text.secondary" mx={8} my={2} >
			{'Copyright © '}
			<MuiLink color="inherit" href="https://anuragpathak.netlify.app/"
			target="_blank">
				Anurag Pathak
			</MuiLink>{' '}
			{new Date().getFullYear()}.
		</Typography>
	)
}