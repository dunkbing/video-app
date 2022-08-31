/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: ["tb-lb.sb-cd.com"],
  },
  publicRuntimeConfig: {
    apiUrl: process.env.API_URL,
  },
}

module.exports = nextConfig
