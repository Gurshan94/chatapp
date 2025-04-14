import React from 'react';
import { useSelector } from 'react-redux';
import { RootState } from '../state/store';
import { Message } from '../state/room/roomSlice';

interface MessageCardProps {
  message: Message;
}

const MessageCard: React.FC<MessageCardProps> = ({ message }) => {
  const userId = useSelector((state: RootState) => state.user.user?.id);
  console.log(message.id)

  const isOwnMessage = message.senderid === userId;
  const bgColor = isOwnMessage ? 'bg-green-600' : 'bg-blue-600';
  const alignment = isOwnMessage ? 'self-end' : 'self-start';

  return (
    <div className={`max-w-xs px-4 py-2 rounded-lg text-white ${bgColor} ${alignment}`}>
      <p className="text-sm">{message.content}</p>
    </div>
  );
};

export default MessageCard;