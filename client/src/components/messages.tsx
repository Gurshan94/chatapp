import { useInfiniteQuery } from "@tanstack/react-query"
import { fetchItems } from "./ietms";
import { useEffect } from 'react';
import { useInView } from 'react-intersection-observer';

const Message = () => {
    const { data, error, status, fetchNextPage, isFetchingNextPage } =
      useInfiniteQuery({
        queryKey: ['items'],
        queryFn: fetchItems,
        initialPageParam: 0,
        getNextPageParam: (lastPage) => lastPage.nextPage,
      });
  
    const { ref, inView } = useInView({
      root: document.querySelector("#scrollable-message-container"),
      rootMargin: '0px',
      threshold: 1.0,
    });
  
    useEffect(() => {
      if (inView) {
        fetchNextPage();
      }
    }, [inView, fetchNextPage]);
  
    if (status === 'pending') return <div>Loading...</div>;
    if (status === 'error') return <div>{error.message}</div>;
  
    return (
      <div className="flex flex-col gap-2">
        {data.pages.map((page) => (
          <div key={page.currentPage} className="flex flex-col gap-2">
            {page.data.map((item) => (
              <div key={item.id} className="inline-block max-w-md bg-gray-700 text-white rounded-lg px-4 py-2 break-words">
                {item.name}
              </div>
            ))}
          </div>
        ))}
  
        <div ref={ref}>{isFetchingNextPage && 'Loading more...'}</div>
      </div>
    );
  };
  
export default Message