import Status from "./status-types"

export async function generateUrl(url: string) {
	const response = await fetch(process.env.NEXT_PUBLIC_SERVER_ADDRESS as string, {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ url: url })
	})

	if (response.status == 400) {
		return {
			message: "Invalid URL.",
			status: Status.InvalidRequest
		}
	}
	if (response.status == 500) {
		return {
			message: "Something went wrong.",
			status: Status.ServerError
		}
	}

	const value = await response.json()
	return {
		message: `${location.host}/${value.url}`,
		status: Status.Success
	}
}

export async function getOriginalUrl(id: string) {
	const response = await fetch(
		process.env.NEXT_PUBLIC_SERVER_ADDRESS as string + "/" + id)

	if (response.status == 404) {
		return {
			status: Status.NotFound
		}
	}
	if (response.status == 500) {
		return {
			status: Status.ServerError
		}
	}

	const value = await response.json()
	return {
		message: "",
		status: Status.Success,
		url: value.url
	}
}