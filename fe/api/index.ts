import axios from "axios"

const API_URL = "https://api.themoviedb.org/3"

export const pageTypeMap = {
  popular: getPopular,
  rated: getRated,
  upcoming: getUpcoming,
}

export async function getPopular(page = 1) {
  const endpoint = "/movie/popular"
  const {
    data: { results, total_pages },
  } = await axios.get(`${API_URL + endpoint}`, {
    params: {
      api_key: process.env.REACT_APP_TRAILER_MOVIES_WATCH_API_KEY,
      page,
    },
  })
  return { results, totalPages: total_pages }
}

export async function getRated(page = 1) {
  const endpoint = "/movie/top_rated"
  const {
    data: { results, total_pages },
  } = await axios.get(`${API_URL + endpoint}`, {
    params: {
      api_key: process.env.REACT_APP_TRAILER_MOVIES_WATCH_API_KEY,
      page,
    },
  })
  return { results, totalPages: total_pages }
}

export async function getUpcoming(page = 1) {
  const endpoint = "/movie/upcoming"
  const {
    data: { results, total_pages },
  } = await axios.get(`${API_URL + endpoint}`, {
    params: {
      api_key: process.env.REACT_APP_TRAILER_MOVIES_WATCH_API_KEY,
      page,
    },
  })
  console.log(total_pages)
  return { results, totalPages: total_pages }
}

export async function getByName(query: string) {
  const type = "/search/movie"
  const {
    data: { results },
  } = await axios.get(`${API_URL + type}`, {
    params: {
      api_key: process.env.REACT_APP_TRAILER_MOVIES_WATCH_API_KEY,
      query,
    },
  })
  return results
}

export async function getWithVideos(id: string | number) {
  const { data } = await axios.get(`${API_URL}/movie/${id}`, {
    params: {
      api_key: process.env.REACT_APP_TRAILER_MOVIES_WATCH_API_KEY,
      append_to_response: "videos",
    },
  })
  return data
}
