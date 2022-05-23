import { LinearProgress, Typography } from "@mui/material"
import Box from "@mui/material/Box"
import Container from "@mui/material/Container"
import { NextPage } from "next"
import Head from "next/head"
import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import PageNotFound from "../components/PageNotFound"
import ServerError from "../components/ServerError"
import { getOriginalUrl } from "../src/server-requests"
import Status from "../src/status-types"

const PageToRedirect: NextPage = () => {
	const [status, setStatus] = useState("")
	const [message, setMessage] = useState("")

	const router = useRouter()
	useEffect(() => {
		if (!router.isReady) return
		
		const { id } = router.query
		
		getOriginalUrl(id as string)
			.then((res) => {
				console.log(res.url)
				console.log(res.message)
				if (res.url) {
					location.replace(res.url)
				}
				setStatus(res.status)
				setMessage(res.message)
			})
			.catch((_err) => {
				setStatus(Status.ServerError)
				setMessage("Something went wrong.")
			})
	}, [router])

	return (
		<div>
			<Head>
				<title>Minly - URL Shortener</title>
				<meta name="viewport" content="initial-scale=1, width=device-width" />
			</Head>
			<Container>

				<Box
					sx={{
						my: 4,
						display: 'flex',
						flexDirection: 'column',
						justifyContent: 'center',
						alignItems: 'center'
					}}
				>

					{message ? <Typography variant="h2">{message}</Typography> : <></>}
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