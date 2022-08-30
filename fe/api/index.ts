import axios from "axios";
import getConfig from "next/config";
import { Video } from "@components/movie-card";

export interface ApiResponse<T> {
  data: T;
}

export interface WithPagination<T> {
  items: T[];
  totalPages: number;
  total: number;
  currentPage: number;
}

const { publicRuntimeConfig } = getConfig() as RuntimeConfig;

export const apiUrl = publicRuntimeConfig.apiUrl;

const axiosInstance = axios.create({
  baseURL: `${apiUrl}`,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
});

export const apiHelper = {
  get: async function <TRes, TReq = any>(
    url: string,
    params: TReq
  ): Promise<ApiResponse<TRes>> {
    const response = await axiosInstance.get<ApiResponse<TRes>>(url, {
      params,
    });
    return response.data;
  },
  post: async function <TRequest, TResponse>(
    url: string,
    data: TRequest,
    config?: any
  ): Promise<ApiResponse<TResponse>> {
    const response = await axiosInstance.post<ApiResponse<TResponse>>(
      url,
      data,
      config
    );
    return response.data;
  },
  put: async function <TRequest, TResponse>(
    url: string,
    data: TRequest
  ): Promise<ApiResponse<TResponse>> {
    const response = await axiosInstance.put<ApiResponse<TResponse>>(url, data);
    return response.data;
  },
  patch: async function <TRequest, TResponse>(
    url: string,
    data: TRequest
  ): Promise<ApiResponse<TResponse>> {
    const response = await axiosInstance.patch<ApiResponse<TResponse>>(
      url,
      data
    );
    return response.data;
  },
  delete: async function <TResponse>(
    url: string
  ): Promise<ApiResponse<TResponse>> {
    const response = await axiosInstance.delete<ApiResponse<TResponse>>(url);
    return response.data;
  },
};

export async function getPopular(page = 1) {
  const res = await axios.get<any, ApiResponse<Video[]>>(
    `${publicRuntimeConfig.apiUrl}/videos`,
    {
      params: {
        page,
      },
    }
  );
  return res.data;
}
