import React from "react"

export type Video = {
  id: string
  title: string
  thumbnail: string
  duration: string
  src: string
};

type Props = {
  movie: Video
  selectMovie: (movie: Video) => void
};

const MovieCard = ({ movie, selectMovie }: Props) => {
  return (
    <div className="movie-container" onClick={() => selectMovie(movie)}>
      {movie.thumbnail ? (
        <img className="movie-cover" src={movie.thumbnail} alt={movie.title} />
      ) : (
        <div className="movie-placeholder">No Image Found</div>
      )}
      <h2 className="movie-title">{movie.title}</h2>
      <p className="movie-vote" title={`Duration: ${movie.duration}`}>
        {movie.duration}
      </p>
    </div>
  )
}

export { MovieCard }
