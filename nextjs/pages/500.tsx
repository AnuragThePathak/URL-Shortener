import { Container } from "@mui/material"
import { NextPage } from "next"
import ResponsiveAppBar from "../components/ResponsiveAppBar"
import ServerError from "../components/ServerError"

const Custom500: NextPage = () => {
	return (
		<Container maxWidth="xl">
			<ResponsiveAppBar />
			<ServerError />
		</Container>
	)
}

export default Custom500