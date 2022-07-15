export default interface Message {
    Type: MessageType
    Data: any
}

export enum MessageType {
    Toast,
}
