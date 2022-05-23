import * as React from 'react'
import Typography from '@mui/material/Typography'
import MuiLink from '@mui/material/Link'

export default function SourceCode() {
	return (
		<Typography variant="body2" color="text.secondary" align="center" mx={8} my={2}>
			{'Find source code on '}
			<MuiLink color="inherit" href="https://anuragpathak.netlify.app/"
			target="_blank">
				Github
			</MuiLink>{'.'}
		</Typography>
	)
}