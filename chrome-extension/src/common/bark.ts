import {tab} from "../pb/compiled";
import _ from "lodash";
import Store from "./storage";
import {BarkDefaultDevice, ConfigStore} from "./storageKey";
import CommonResponse from "./webRequest";
import IBarkMessage = tab.IBarkMessage;
import IBarkDevice = tab.IBarkDevice;
import IBarkHistory = tab.IBarkHistory;
import BarkDevice = tab.BarkDevice;
import ClientConfig = tab.ClientConfig;

export async function Push(msg: IBarkMessage, device?: IBarkDevice) {
    if (!device) {
        const ds = new Store<IBarkDevice>(BarkDefaultDevice)
        device = await ds.get()
    }
    let {name, url} = device;
    if (!name || !url) {
        return
    }
    let {title, content, url: msgUrl} = msg;
    console.log("push", msg, "to", device);
    try {
        url = EndWithAppend(url, "/", "place_holder")
        fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            body: JSON.stringify({
                'title': title,
                'content': content,
                'ext_params': {
                    'badge': "1",
                    'icon': 'none:url',
                    'group': 'none',
                    'url': msgUrl,
                },
                'category': 'category',
                'sound': "1107", //'minuet.caf'
            })
        }).catch(e => {
            console.log(`Bark fetch ${url} error ${e}`);
        })
    } catch (err) {
        console.log(`Bark send to ${name} error ${err}`)
    }
}

export async function PullHistory(target: string, offset: number, limit: number): Promise<IBarkHistory[]> {
    let params = new URLSearchParams({
        "limit": limit.toString(),
        "offset": offset.toString(),
    })
    return new Promise<IBarkHistory[]>((resolve, reject) => {
        target = _.trimEnd(target, "/")
        fetch(`${target}?${params}`, {
            method: "GET",
        }).then(res => res.json()).then(data => {
            if (data.code != 200) {
                reject(data)
            }
            resolve(data.data ?? [])
        })
    })
}

export async function DropHistory(target: string, id: number) {
    target = EndWithAppend(target, "/", `history/${id}`)
    return fetch(target, {
        method: "DELETE"
    }).then(res => res.json())
}

export async function AddDevice(device: BarkDevice): Promise<BarkDevice> {
    const conf = new Store<ClientConfig>(ConfigStore)
    let host = await conf.get()
    if (!host) {
        device.id = 1; //TODO: get device list len
        return device
    }
    let params = new URLSearchParams({})
    return fetch(EndWithAppend(host.url, "/", `?${params}`), {
        method: "POST"
    }).then(res => res.json()).then((rsp: CommonResponse) => {
        device.id = rsp.data.id
        return device
    })
}

export async function DropDevice(device: BarkDevice) {
    const conf = new Store<ClientConfig>(ConfigStore)
    let host = await conf.get()
    if (!host) {
        return
    }
    let params = new URLSearchParams({})
    return fetch(EndWithAppend(host.url, "/", `?${params}`), {
        method: "DELETE"
    }).then(res => res.json())
}

export async function PullDevice(): Promise<BarkDevice[]> {
    const conf = new Store<ClientConfig>(ConfigStore)
    let host = await conf.get()
    return fetch(EndWithAppend(host.url, "/", host.uid), {
        method: "GET"
    }).then(res => res.json())
}

export async function EditDevice(old: BarkDevice, later: BarkDevice) {
}

function EndWithAppend(str: string, end: string, append: string): string {
    if (_.endsWith(str, end)) {
        return str + append
    } else {
        return str + end + append
    }
}
