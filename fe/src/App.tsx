import React, {useCallback, useEffect, useState} from 'react';
import { MovieList } from './components/movie-list';
import { Header } from './components/header';
import { getByName, getWithVideos, pageTypeMap } from './api';
import { Video } from "./components/movie-card";
import './App.css';

function App() {
  const [movies, setMovies] = useState([]);
  const [searchKey, setSearchKey] = useState('');
  const [selectedMovie, setSelectedMovie] = useState<Video | null>(null);
  const [playTrailer, setPlayTrailer] = useState(false);
  const [availableTrailer, setAvailableTrailer] = useState(true);
  const [page, setPage] = useState(1);
  const [pageType, setPageType] = useState(Object.keys(pageTypeMap)[0]);
  const [isLoadMoreAvailable, setIsLoadMoreAvailable] = useState(true);

  const selectMovie = useCallback(async (movie) => {
    setPlayTrailer(false);
    const data = await getWithVideos(movie.id);
    setSelectedMovie(data);
    checkAvailableTrailer(data);
  }, [])

  const checkAvailableTrailer = (movie) => {
    const isAvailable = Boolean(movie.videos.results.length);
    setAvailableTrailer(isAvailable);
  }

  useEffect(() => {
    async function callAPI() {
      const {results} = await pageTypeMap[pageType](); // getPopular();
      await selectMovie(results[0]);
      setMovies(results);
      setPage(1);
    }
    void callAPI();
  }, [pageType, selectMovie]);

  useEffect(() => {
    async function callAPI() {
      const {results, totalPages} = await pageTypeMap[pageType](page);
      setIsLoadMoreAvailable(totalPages > page);
      await selectMovie(results[0]);
      setMovies([...movies, ...results]);
    }
    void callAPI();
  }, [movies, page, pageType, selectMovie]);

  const renderVideo = (video: Video) => {
    return (
      <iframe
        title={video.title}
        width="560"
        height="315"
        src={`https://spankbang.com/${video.id}/embed/`}
        frameBorder="0"
        scrolling="no"
        allowFullScreen
      />
    )
  }

  const searchMovies = async (e: React.FormEvent) => {
    e.preventDefault();
    const results = await getByName(searchKey);
    setMovies(results)
    await selectMovie(results[0]);
  }

  return (
    <div className="App">
      <Header
        searchMovies={searchMovies}
        setSearchKey={setSearchKey}
      />
      {Object.keys(pageTypeMap).map(item => {
        return <button className={item === pageType && 'page-type-current'} type='button' onClick={() => setPageType(item)}>{item}</button>
      })}
      <div className="hero" style={{backgroundImage: `url(${selectedMovie?.thumbnail})`}}>
        <div className='content'>
          {selectedMovie?.src && playTrailer && renderVideo(selectedMovie)}
          {
            availableTrailer &&
            <button className='button' onClick={() => setPlayTrailer(true)}>Play Trailer</button>
          }
          <h2 className='title'>{selectedMovie?.title}</h2>
          {/*<p className='overview'>{selectedMovie.overview}</p>*/}
        </div>
      </div>
      <MovieList
        movies={movies}
        selectMovie={selectMovie}
        onLoadMore={() => setPage(page + 1)}
        isLoadMoreAvailable={isLoadMoreAvailable}
      />
    </div>
  );
}

export default App;
