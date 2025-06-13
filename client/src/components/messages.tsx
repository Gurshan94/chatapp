import { useInfiniteQuery } from "@tanstack/react-query"
import { useEffect,useRef } from 'react';
import { Message } from "../state/room/roomSlice";
import { useFetchmessagesMutation } from "../state/room/roomApiSlice";
import { useDispatch,useSelector } from "react-redux";
import { appendMessages } from "../state/room/roomSlice";
import { RootState } from "../state/store";
import MessageCard from "./messagecard";


interface MyComponentProps {
  roomid: number|null;
}

const MessageList: React.FC<MyComponentProps> = ({roomid}) => {
  const [fetchmessages]=useFetchmessagesMutation()
  const scrollContainerRef = useRef<HTMLDivElement | null>(null);
  const loaderRef = useRef<HTMLDivElement | null>(null);
  const dispatch=useDispatch()

  const fetchMessages = async ({ pageParam=0 }) => {

    const res:Message[] = await fetchmessages({roomid: roomid, limit: 8, offset: pageParam}).unwrap();
     return {
      messages: res,
      nextOffset: 8+pageParam,
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
      queryKey: ['messages',roomid],
      queryFn: fetchMessages,
      initialPageParam: 0,
      getNextPageParam: (lastPage) =>
       lastPage.hasMore ? lastPage.nextOffset : undefined,
    });

    useEffect(() => {
      const messages = data?.pages.flatMap(page => page.messages) || [];
    
      if (messages.length > 0) {
        dispatch(appendMessages({ roomId: roomid, messages }));
      }
    }, [data, dispatch, roomid]);


    const displayMessage = useSelector((state: RootState) => {
      return state.room.messages
    });

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
      <div className="flex flex-col h-full  bg-gray-850 border-r border-gray-700">
      {/* Scrollable container */}
      <div
        className="flex-1 overflow-y-auto p-4 custom-scrollbar"
        ref={scrollContainerRef}
      >

        {status === 'error' && <div>Error: {error.message}</div>}

        <div className="flex flex-col gap-2">
          {displayMessage && displayMessage.map((message: Message) => 
             <MessageCard key={message.id} message={message} />
          )}
        </div>

        {/* 👇 Intersection observer target */}
        <div ref={loaderRef} className="text-center py-4">
          {isFetchingNextPage && <span>Loading more...</span>}
        </div>
      </div>
    </div>
    );
  };
  
export default MessageList