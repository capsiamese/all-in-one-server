export default interface ServerMSG {
    code: number
    error?: string
    data?: any
}

export const CodeSuccess = 0
