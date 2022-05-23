import TextField from '@mui/material/TextField';
import { Button, Container } from "@mui/material"

export default function UrlTextInputField() {
  return (
		<Container
			sx={{
				display: "flex",
				m: 2,
				flexWrap: { sm: "nowrap", xs: "wrap"},
				justifyContent: "center"
			}}>

      <TextField fullWidth label="Type the URL to shorten" id="url-input"
			 required type="url" sx={{m: 1}}/>
			 <Button variant="contained" size="large" sx={{m: 1}}>Submit</Button>
		</Container>
  );
}
