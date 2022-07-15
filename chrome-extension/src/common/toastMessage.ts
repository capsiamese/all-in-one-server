import Browser, {Runtime} from "webextension-polyfill";
import Message, {MessageType} from "./message";
import MessageSender = Runtime.MessageSender;

export function SendToastMessage(message: any) {
    Browser.runtime.sendMessage({
        Type: MessageType.Toast,
        Data: message,
    } as Message).catch(e => {
        console.log('send message', e);
    })
}

export function OnToastMessageReceive(f: (message: any, sender: MessageSender) => void) {
    Browser.runtime.onMessage.addListener((message: Message, sender) => {
            if (message.Type == MessageType.Toast) {
                f(message.Data, sender)
            }
        }
    )
}
