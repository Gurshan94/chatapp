import { useInfiniteQuery } from '@tanstack/react-query';
import { useGetroomsMutation } from "../state/room/roomApiSlice";
import { useEffect, useRef } from 'react';
import { BackendRoom } from '../state/room/roomSlice';
import RoomCard from './roomcard';

const DiscoverRoomList = () => {
  const [fetchrooms] = useGetroomsMutation();
  const scrollContainerRef = useRef<HTMLDivElement | null>(null);
  const loaderRef = useRef<HTMLDivElement | null>(null);

  const fetchRooms = async ({ pageParam = 0 }) => {
    const res = await fetchrooms({ limit: 8, offset: pageParam }).unwrap();
    return {
      rooms: res,
      nextOffset: 8 + pageParam,
      hasMore: res.length === 8,
    };
  };

  const {
    data,
    fetchNextPage,
    hasNextPage,
    isFetchingNextPage,
    status,
    error
  } = useInfiniteQuery({
    queryKey: ['rooms'],
    queryFn: fetchRooms,
    initialPageParam: 0,
    getNextPageParam: (lastPage) =>
      lastPage.hasMore ? lastPage.nextOffset : undefined,
  });

  const allRooms = data?.pages.flatMap(page => page.rooms) || [];

  useEffect(() => {
    const scrollContainer = scrollContainerRef.current;
    const loader = loaderRef.current;
    if (!scrollContainer || !loader) return;

    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting && hasNextPage && !isFetchingNextPage) {
          fetchNextPage();
        }
      },
      {
        root: scrollContainer,
        threshold: 1.0,
      }
    );

    observer.observe(loader);

    return () => {
      if (loader) observer.unobserve(loader);
    };
  }, [hasNextPage, isFetchingNextPage, fetchNextPage]);

  return (
    <div className="flex flex-col h-full w-72 bg-gray-850 border-r border-gray-700">
      {/* Scrollable container */}
      <div
        className="flex-1 overflow-y-auto p-4 custom-scrollbar"
        ref={scrollContainerRef}
      >
        <h2 className="text-xl font-semibold mb-4">Rooms</h2>

        {status === 'pending' && <div>Loading...</div>}
        {status === 'error' && <div>Error: {error.message}</div>}

        <div className="flex flex-col gap-2">
          {allRooms.map((room: BackendRoom) => (
             <RoomCard key={room.id} room={room}/>
          ))}
        </div>

        {/* 👇 Intersection observer target */}
        <div ref={loaderRef} className="text-center py-4">
          {isFetchingNextPage && <span>Loading more...</span>}
        </div>
      </div>

      {/* 👇 Spacer so last item is not hidden behind taskbar */}
      <div className="h-20" />
    </div>
  );
};

export default DiscoverRoomList;
