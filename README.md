# Minly - URL Shortener

## How is it useful

- Minly is a Full-Stack web application that allows you to generate short URLs which you can use in place of the original ones after generation.
- It can be very useful in like Tweets, Linkedin connection request notes, or such purposes, where there are character limits and you can't use long URLs.

## Tech specification

- For the frontend, Nextjs is used with Material UI.
- Backend is created using Golang (no framework used). For the database, we are using MongoDB atlas.
- Frontend is hosted in [Vercel](https://mly.vercel.app/) and the backend on Heroku.
- For CI/CD, Github action is being used. No manual deploy/integration tests are required.
