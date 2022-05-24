import { LinearProgress } from "@mui/material"
import Box from "@mui/material/Box"
import Container from "@mui/material/Container"
import { NextPage } from "next"
import Head from "next/head"
import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import PageNotFound from "../components/PageNotFound"
import ResponsiveAppBar from "../components/ResponsiveAppBar"
import ServerError from "../components/ServerError"
import { getOriginalUrl } from "../src/server-requests"
import Status from "../src/status-types"

const PageToRedirect: NextPage = () => {
	const [status, setStatus] = useState("")

	const router = useRouter()
	useEffect(() => {
		if (!router.isReady) return

		const { id } = router.query

		getOriginalUrl(id as string)
			.then((res) => {
				if (res.url) {
					location.replace(res.url)
				}
				setStatus(res.status)
			})
			.catch((_err) => {
				setStatus(Status.ServerError)
			})
	}, [router])

	return (
		<div>
			<Head>
				<title>Minly - URL Shortener</title>
				<meta name="viewport" content="initial-scale=1, width=device-width" />
			</Head>
			<Container maxWidth="xl">

				<Box
					sx={{
						my: 2,
						display: 'flex',
						flexDirection: 'column',
						justifyContent: 'center',
						alignItems: 'center'
					}}
				>
					<ResponsiveAppBar />

					{!status ? <Box sx={{ width: '100%', my: 4 }}>
						<LinearProgress />
					</Box> : <></>}
					{status == Status.NotFound ? <PageNotFound /> : <></>}
					{status == Status.ServerError ? <ServerError /> : <></>}
				</Box>
			</Container>
		</div>
	)
}

export default PageToRedirect