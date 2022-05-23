import Typography from '@mui/material/Typography'
import MuiLink from '@mui/material/Link'
import Image from "next/image"
import serverErrorPic from "../public/3805090.webp"

export default function ServerError() {
	return (
		<>
			<Typography variant="body2" color="text.secondary" align="center">
				<Image src={serverErrorPic} alt="404" placeholder="blur" />
				<MuiLink href='https://www.freepik.com/vectors/server-error'
					target="_blank">
					Server error vector created by storyset - www.freepik.com
				</MuiLink>
			</Typography>
		</>
	)
}