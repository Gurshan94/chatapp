import React from 'react';
import { useSelector,useDispatch } from 'react-redux';
import { RootState } from '../state/store';
import { Message } from '../state/room/roomSlice';
import {Trash2 } from 'lucide-react';
import { useDeletemessageMutation } from '../state/room/roomApiSlice';
import { deleteMessageWithID } from '../state/room/roomSlice';

interface MessageCardProps {
  message: Message;
}

const MessageCard: React.FC<MessageCardProps> = ({ message }) => {
  const userId = useSelector((state: RootState) => state.user.user?.id);
  const [deleteMessage]=useDeletemessageMutation()
  const dispatch=useDispatch()
  const isOwnMessage = message.senderid === Number(userId);
  const bgColor = isOwnMessage ? 'bg-green-600' : 'bg-blue-600';
  const alignment = isOwnMessage ? 'self-end' : 'self-start';

  const handleDeleteMessage= async (messageId:number)=>{
    // Implement the delete message logic here
    try {
      const data=await deleteMessage({messageId:messageId}).unwrap()
      dispatch(deleteMessageWithID({messageId:messageId}))
    } catch (error) {
      console.error('Error deleting message:', error);
    }
  
  }

  return (
  <div key={message.id} 
  className={`group max-w-xs px-4 py-2 rounded-lg text-white ${bgColor} ${alignment} ${message.senderid==userId && 'hover:bg-red-500'}`}
  >
  <div className='flex items-center justify-between mb-1 gap-5'>
  <div className="text-xs text-grey opacity-80 mb-1">{message.username}</div>
  {message.senderid==userId && <button className="invisible group-hover:visible transition-opacity duration-200" onClick={() => handleDeleteMessage(message.id)} >
  <Trash2 size={12}/>
  </button>}
  </div>
  
  <p className="text-sm">{message.content}</p>
  </div>
  )
};

export default MessageCard;