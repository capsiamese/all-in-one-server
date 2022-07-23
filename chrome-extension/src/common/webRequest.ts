import Browser from "webextension-polyfill";
import {tab} from "../pb/compiled";
import Group = tab.Group;

export default interface CommonResponse {
    code: number
    error: string
    data: any
}

export interface BarkResponse {
    code: number
    message: string
    timestamp: number
    data: any
}

export async function Register(addr: string, name: string): Promise<CommonResponse> {
    let id = Browser.runtime.id
    return fetch(`${addr}/register?name=${name}&extension=${id}`, {
        method: "POST"
    }).then(res => res.json())
}

export async function PullTabs(addr: string, uid: string): Promise<CommonResponse> {
    let t = new tab.Tab({name: "hello"})
    console.log(t)
    return fetch(`${addr}/${uid}/group`, {
        method: "GEt"
    }).then(res => res.json())
}
