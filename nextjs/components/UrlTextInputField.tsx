import CloseIcon from '@mui/icons-material/Close'
import ContentCopyOutlinedIcon from '@mui/icons-material/ContentCopyOutlined'
import TextField from '@mui/material/TextField'
import { Alert, AlertColor, Button, Container, IconButton, Snackbar } from "@mui/material"
import { FormEvent, Fragment, SyntheticEvent, useState } from "react"
import { generateUrl } from "../src/server-requests"
import Status from "../src/status-types"

export default function UrlTextInputField() {
	const [open, setOpen] = useState(false)
	const [url, setUrl] = useState("")
	const [status, setStatus] = useState("")
	const [message, setMessage] = useState("")

	function handleSubmit(e: FormEvent) {
		e.preventDefault()

		if (url) {
			setUrl("")
			generateUrl(url)
				.then((res) => {
					setStatus(res.status)
					setMessage(res.message)
				})
				.catch((_err) => {
					setStatus(Status.ServerError)
					setMessage("Can't connect to server.")
				})
		} else {
			setOpen(true)
		}
	}

	const handleClose = (_event: SyntheticEvent | Event, reason?: string) => {
		if (reason === 'clickaway') {
			return
		}

		setOpen(false)
	}

	const action = (
		<Fragment>
			<IconButton
				size="small"
				aria-label="close"
				color="inherit"
				onClick={handleClose}
			>
				<CloseIcon fontSize="small" />
			</IconButton>
		</Fragment>
	)

	return (
		<Container>
			<Container
				sx={{
					display: "flex",
					m: 2,
					flexWrap: { sm: "nowrap", xs: "wrap" },
					justifyContent: "center"
				}}>

				<TextField fullWidth label="Type the URL to shorten" id="url-input"
					type="url" sx={{ m: 1 }} value={url}
					onChange={(e) => setUrl(e.target.value)} />

				<Button variant="contained" size="large" sx={{ m: 1 }}
					onClick={handleSubmit}>Submit</Button>
			</Container>

			{status ? <Alert severity={status as AlertColUndef} sx={{
				my: 2,
				mx: 6
			}} action={
				<Button color="inherit" size="small"
					onClick={(_e) => {
						navigator.clipboard.writeText(message)
					}}>
					<ContentCopyOutlinedIcon
						sx={{ display: { xs: 'none', md: 'flex' } }}
					/>
				</Button>
			}>
				{message}
			</Alert> : <></>}

			<Snackbar
				open={open}
				autoHideDuration={4000}
				onClose={handleClose}
				message="URL can't be empty."
				action={action}
			/>
		</Container>
	)
}

type AlertColUndef = AlertColor | undefined
