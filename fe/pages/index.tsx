import React, { useCallback, useState } from "react";
import MovieList from "@components/movie-list";
import { Header } from "@components/header";
import { Video } from "@components/movie-card";

function Home() {
  const [searchKey, setSearchKey] = useState("");
  const [selectedMovie, setSelectedMovie] = useState<Video | null>(null);
  const [playing, setPlaying] = useState(false);

  const selectMovie = useCallback(async (video: Video) => {
    setSelectedMovie(video);
    setPlaying(false);
  }, []);

  const renderVideo = (video: Video | null) => {
    if (!video) return null;
    return (
      <iframe
        title={video.title}
        width="100%"
        height="500"
        src={`https://spankbang.com/${video.id}/embed/`}
        frameBorder="0"
        scrolling="no"
        allowFullScreen
      />
    );
  };

  const searchMovies = async (e: React.FormEvent) => {
    e.preventDefault();
  };

  return (
    <div className="App">
      <Header searchMovies={searchMovies} setSearchKey={setSearchKey} />
      <div
        className="hero"
        style={
          !playing
            ? { backgroundImage: `url(${selectedMovie?.thumbnail})` }
            : {}
        }
      >
        {!playing ? (
          <div style={{ cursor: "pointer" }} onClick={() => setPlaying(true)}>
            <a id="play-video" className="video-play-button">
              <span></span>
            </a>
            <div id="video-overlay" className="video-overlay">
              <a className="video-overlay-close">&times;</a>
            </div>
          </div>
        ) : (
          renderVideo(selectedMovie)
        )}
      </div>
      <div className="content">
        <h2 className="title">{selectedMovie?.title}</h2>
        {selectedMovie?.tags && (
          <>
            <p>Thể loại</p>
            <div>
              {selectedMovie.tags.split("|").map((tag, index) => (
                <a className="video-tag" key={`${tag}-${index}`}>
                  {tag}
                </a>
              ))}
            </div>
          </>
        )}
        {selectedMovie?.names && (
          <>
            <p>Diễn viên</p>
            <div>
              {selectedMovie.names.split("|").map((tag) => (
                <a className="video-tag" key={tag}>
                  {tag}
                </a>
              ))}
            </div>
          </>
        )}
      </div>
      <MovieList selectMovie={selectMovie} query={searchKey} />
    </div>
  );
}

export default Home;
