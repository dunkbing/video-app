import React from "react"
import { Video, MovieCard } from "../movie-card"

type Props = {
  movies: any[]
  isLoadMoreAvailable: boolean
  selectMovie: (movie: Video) => void
  onLoadMore: () => void
};

export function MovieList({
  movies,
  selectMovie,
  onLoadMore,
  isLoadMoreAvailable,
}: Props) {
  const renderMovies = () => {
    return movies.map((movie) => {
      return (
        <MovieCard key={movie.id} movie={movie} selectMovie={selectMovie} />
      )
    })
  }
  return (
    <div className="container">
      {renderMovies()}
      {isLoadMoreAvailable && (
        <button type="button" className="btn-load-more" onClick={onLoadMore}>
          LOAD MORE
        </button>
      )}
    </div>
  )
}
