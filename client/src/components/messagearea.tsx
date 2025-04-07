import Message from "./messages"

const Messagearea=()=>{

    return (
        <div className="h-[calc(100vh-68px)] flex flex-col flex-1">
            
        <div className="flex-1 overflow-y-auto p-4" id="scrollable-message-container">
            <Message />
        </div>

        <div className="p-4 border-t border-gray-700 bg-gray-800">
            <div className="flex gap-2">
            <input
                type="text"
                placeholder="Type your message..."
                className="flex-1 p-3 rounded-lg bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
            <button className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700">
                Send
            </button>
            </div>
        </div>
        </div>

    )

}

export default Messagearea