import React from "react"
import Image from "next/image"

export type Video = {
  id: string;
  title: string;
  thumbnail: string;
  duration: string;
  src: string;
  names: string;
  tags: string;
};

type Props = {
  movie: Video;
  selectMovie: (movie: Video) => void;
};

const MovieCard = ({ movie, selectMovie }: Props) => {
  return (
    <div className="movie-container" onClick={() => selectMovie(movie)}>
      {movie.thumbnail ? (
        // <img className="movie-cover" src={movie.thumbnail} alt={movie.title} />
        <Image className="movie-cover" src={movie.thumbnail} alt={movie.title} width="200vh" height="200vh" />
      ) : (
        <div className="movie-placeholder">Không thể hiển thị ảnh</div>
      )}
      <h2 className="movie-title">{movie.title}</h2>
      <p className="movie-vote" title={`Duration: ${movie.duration}`}>
        {movie.duration}
      </p>
    </div>
  )
}

export { MovieCard }
