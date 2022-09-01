import axios from "axios"
import getConfig from "next/config"

export interface ApiResponse<T> {
  data: T;
}

export interface WithPagination<T> {
  items: T[];
  totalPages: number;
  total: number;
  currentPage: number;
}

const { publicRuntimeConfig } = getConfig() as RuntimeConfig

export const apiUrl = publicRuntimeConfig.apiUrl

const axiosInstance = axios.create({
  baseURL: `${apiUrl}`,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
})

export const apiHelper = {
  get: async function <TRes, TReq = any>(
    url: string,
    params: TReq
  ): Promise<ApiResponse<TRes>> {
    const response = await axiosInstance.get<ApiResponse<TRes>>(url, {
      params,
    })
    return response.data
  },
}
