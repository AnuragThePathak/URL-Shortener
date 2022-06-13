export default function checkUrlValidity(url: string): boolean {
	try {
		const urlObject = new URL(url)

		if (urlObject.protocol == "http:" || urlObject.protocol == "https:") {
			return true
		}

		return false
	} catch (e) {
		return false
	}
}