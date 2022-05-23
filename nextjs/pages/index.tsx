import Box from "@mui/material/Box"
import Container from "@mui/material/Container"
import Typography from "@mui/material/Typography"
import type { NextPage } from 'next'
import Link from "next/link"
import ResponsiveAppBar from "../components/ResponsiveAppBar"
import Copyright from "../src/Copyright"

const Home: NextPage = () => {
  return (
    <Container maxWidth="lg">
      <Box
        sx={{
          my: 4,
          display: 'flex',
          flexDirection: 'column',
          justifyContent: 'center',
          alignItems: 'center',
        }}
      >
        <ResponsiveAppBar />
        <Typography variant="h4" component="h1" gutterBottom>
          MUI v5 + Next.js with TypeScript example
        </Typography>
        <Link href="/about" color="secondary">
          Go to the about page
        </Link>
        <Copyright />
      </Box>
    </Container>
  )
}

export default Home