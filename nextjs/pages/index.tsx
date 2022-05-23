import Box from "@mui/material/Box"
import Container from "@mui/material/Container"
import Typography from "@mui/material/Typography"
import type { NextPage } from 'next'
import Head from "next/head"
import Link from "../src/Link"
import ResponsiveAppBar from "../components/ResponsiveAppBar"
import UrlTextInputField from "../components/UrlTextInputField"
import Copyright from "../src/Copyright"

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
          
          <Link href="/about" color="secondary">
            Go to the about page
          </Link>
          <Copyright />
        </Box>
      </Container>
    </div>
  )
}

export default Home
