import Typography from '@mui/material/Typography'
import MuiLink from '@mui/material/Link'
import Image from "next/image"
import serverErrorPic from "../public/3805090.webp"

export default function ServerError() {
	return (
		<>
			<Typography variant="h2">Something went wrong.</Typography>
				<Image src={serverErrorPic} alt="404" placeholder="blur" />
			<Typography variant="body2" color="text.secondary" align="center">
				<MuiLink href='https://www.freepik.com/vectors/server-error'
					target="_blank">
					Server error vector created by storyset - www.freepik.com
				</MuiLink>
			</Typography>
		</>
	)
}