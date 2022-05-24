import Typography from '@mui/material/Typography'
import MuiLink from '@mui/material/Link'
import Image from "next/image"
import notFoundPic from "../public/2696450.webp"

export default function PageNotFound() {
	return (
		<>
			<Typography variant="h2">URL not found.</Typography>
			<Image src={notFoundPic} alt="404" placeholder="blur" />
			<Typography variant="body2" color="text.secondary" align="center">
				<MuiLink color="inherit" href="https://www.freepik.com/vectors/404-page"
					target="_blank">
					404 page vector created by pikisuperstar - www.freepik.com
				</MuiLink>
			</Typography>
		</>
	)
}