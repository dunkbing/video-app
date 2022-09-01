import React, {ChangeEvent, useCallback} from "react"

function debounce<T extends unknown[], U>(wait: number, callback: (...args: T) => PromiseLike<U> | U) {
  let timer: NodeJS.Timeout | undefined

  return (...args: T): Promise<U> => {
    clearTimeout(timer as NodeJS.Timeout)
    return new Promise((resolve) => {
      timer = setTimeout(() => resolve(callback(...args)), wait)
    })
  }
}

type Props = {
  searchMovies: React.FormEventHandler<HTMLFormElement>;
  onSearch: (val: string) => void;
};

export function Header({ searchMovies, onSearch }: Props) {
  const searchDebounce = useCallback(
    debounce(500, (newQuery: string) => {
      onSearch?.(newQuery)
    }),
    [onSearch],
  )

  function handleTextChange(e: ChangeEvent<HTMLInputElement>) {
    void searchDebounce(e.target.value)
  }

  return (
    <header className="header">
      <h1>Xem phim hay</h1>
      <form onSubmit={searchMovies}>
        <input type="text" onChange={handleTextChange} placeholder="Dien vien, the loai, ten phim" />
        <button type="submit">Search</button>
      </form>
    </header>
  )
}
