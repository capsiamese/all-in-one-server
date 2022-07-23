import Browser, {Runtime} from "webextension-polyfill";
import MessageSender = Runtime.MessageSender;

interface Message {
    Type: MessageType
    Data: any
}

export enum MessageType {
    Toast,
}

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
