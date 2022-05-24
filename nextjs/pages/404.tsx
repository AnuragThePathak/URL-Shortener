import { Container } from "@mui/material"
import { NextPage } from "next"
import PageNotFound from "../components/PageNotFound"
import ResponsiveAppBar from "../components/ResponsiveAppBar"

const Custom404: NextPage = () => {
	return (
		<Container maxWidth="xl">
			<ResponsiveAppBar />
			<PageNotFound />
		</Container>
	)
}

export default Custom404