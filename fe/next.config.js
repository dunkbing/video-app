/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  images: {
    domains: [
      "picsum.photos",
      "s2.coinmarketcap.com",
      "s3.coinmarketcap.com",
      "cdn.pixabay.com",
    ],
  },
  publicRuntimeConfig: {
    apiUrl: process.env.API_URL,
  },
};

module.exports = nextConfig;
