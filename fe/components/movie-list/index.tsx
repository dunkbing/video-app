import React, { useEffect, useState } from "react";
import ReactPaginate from "react-paginate";
import { Video, MovieCard } from "../movie-card";
import { apiHelper, WithPagination } from "../../api";

type Props = {
  query?: string;
  selectMovie: (movie: Video) => void;
};

const MovieList: React.FC<Props> = ({ query, selectMovie }: Props) => {
  const [videos, setVideos] = useState<WithPagination<Video>>({
    items: [],
    totalPages: 0,
    currentPage: 0,
    total: 0,
  });
  const [page, setPage] = useState(1);
  useEffect(() => {
    async function callApi() {
      const params = query ? { query, page } : { page };
      const { data } = await apiHelper.get<WithPagination<Video>>(
        "/videos",
        params
      );
      console.log("videos", data);
      selectMovie(data.items[0]);
      setVideos(data);
    }
    void callApi();
  }, [page, query]);

  return (
    <>
      <div className="container">
        {videos.items.map((video) => (
          <MovieCard
            key={`${video.id}-${video.title}`}
            movie={video}
            selectMovie={selectMovie}
          />
        ))}
      </div>
      <div id="react-paginate">
        <ReactPaginate
          previousLabel={"<"}
          nextLabel={">"}
          breakLabel={"..."}
          breakClassName={"break-me"}
          pageCount={videos.totalPages}
          marginPagesDisplayed={2}
          pageRangeDisplayed={5}
          onPageChange={({ selected }) => {
            console.log(selected);
            setPage(selected + 1);
          }}
          containerClassName={"pagination"}
          activeClassName={"active"}
        />
      </div>
    </>
  );
};

export default React.memo(MovieList);
