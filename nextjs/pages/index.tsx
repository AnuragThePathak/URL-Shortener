import Box from "@mui/material/Box"
import Container from "@mui/material/Container"
import MuiLink from '@mui/material/Link'
import Typography from "@mui/material/Typography"
import type { NextPage } from 'next'
import Head from "next/head"
import ResponsiveAppBar from "../components/ResponsiveAppBar"
import UrlTextInputField from "../components/UrlTextInputField"
import Copyright from "../components/Copyright"
import SourceCode from "../components/SourceCode"
import heroImage from "../public/66224.webp"
import Image from "next/image"

const Home: NextPage = () => {
  return (
    <div>
      <Head>
        <title>Minly - URL Shortener</title>
        <meta name="viewport" content="initial-scale=1, width=device-width" />
      </Head>
      <Container maxWidth="xl">
        <Box
          sx={{
            my: 1,
            display: 'flex',
            flexDirection: 'column',
            justifyContent: 'center',
            alignItems: 'center',
          }}
        >
          <ResponsiveAppBar />
          <Typography variant="h3" component="h1" gutterBottom sx={{
            m: 6
          }}>
            Short URLs, Easy to Share
          </Typography>
          <UrlTextInputField />

          <Image src={heroImage} alt="share-image" placeholder="blur" 
          height={450} width={600} />

          <Typography variant="body2" >
          <MuiLink color="inherit" href="https://www.freepik.com/vectors/data-sharing"
			target="_blank">
            Data sharing vector created by rawpixel.com - www.freepik.com
          </MuiLink>
          </Typography>

          <Typography variant="h4" m={4} color="InfoText">{"This website was created by "}
          <MuiLink href="https://twitter.com/AnuragThePathak"
			target="_blank">
            Anurag Pathak
          </MuiLink>{" for learning purpose."}
          </Typography>

          <Container sx={{
            // m: 5,
            display: "flex",
            justifyContent: "space-evenly",
            flexWrap: {xs: "wrap", sm: "nowrap"},
            bgcolor: "whitesmoke"
          }}>

            <SourceCode />
            <Copyright />
          </Container>
        </Box>
      </Container>
    </div>
  )
}

export default Home
