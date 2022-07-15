import Browser, {Storage} from "webextension-polyfill";
import StorageAreaOnChangedChangesType = Storage.StorageAreaOnChangedChangesType;


class Bark {
    static DefaultDevice = "default_bark_device";
    static DeviceList = "bark_device_list";
}

export interface Device {
    name: string;
    target: string;
}

export interface BarkMessage {
    title?: string;
    body?: string;
    url?: string;
}

const storage = Browser.storage.local

async function SetDefaultDevice(device: Device) {
    return storage.set({
        [Bark.DefaultDevice]: device
    })
}

async function GetDefaultDevice() {
    let result = await storage.get(Bark.DefaultDevice);
    return result[Bark.DefaultDevice] as Device ?? {name: '', target: ''};
}

function OnDefaultDeviceChange(f: (device: Device) => void) {
    storage.onChanged.addListener(changes => {
        if (changes[Bark.DefaultDevice]) {
            let dChange = changes[Bark.DefaultDevice] as StorageAreaOnChangedChangesType
            let newValue = dChange?.newValue ?? {name: '', target: ''};
            //let {newValue} = changes[Bark.DefaultDevice] as { newValue: Device };
            f(newValue);
        }
    })
}

async function SetDeviceList(deviceList: Device[]) {
    if (deviceList) {
        switch (deviceList.length) {
            case 0:
                SetDefaultDevice({name: '', target: ''}).then();
                break
            case 1:
                SetDefaultDevice(deviceList[0]).then();
                break
        }
    }

    return storage.set({
        [Bark.DeviceList]: deviceList
    })
}

async function GetDeviceList() {
    let result = await storage.get(Bark.DeviceList);
    return result[Bark.DeviceList] ?? [];
}

// todo: add a generic function get changes[i] newValue
// <T>(change: StorageAreaOnChangedChangesType, default) => change?.newValue ?? default;

function OnDeviceListChange(f: (list: Device[]) => void) {
    storage.onChanged.addListener(changes => {
        if (changes[Bark.DeviceList]) {
            let dlChange = changes[Bark.DeviceList] as StorageAreaOnChangedChangesType
            //let {newValue} = (changes[Bark.DeviceList] ?? []) as { newValue: Device[] };
            f(dlChange.newValue ?? []);
        }
    })
}

function PushToDefault(msg: BarkMessage) {
    GetDefaultDevice().then(device => {
        Push(msg, device);
    })
}

function Push(msg: BarkMessage, device: Device) {
    if (!device) {
        return
    }
    let {name, target} = device;
    if (!name || !target) {
        return
    }
    let {title, body, url} = msg;
    console.log("push", msg, "to", device);
    try {
        fetch(target, {
            method: "POST",
            headers: {
                "Content-Type": "application/json; charset=utf-8",
            },
            body: JSON.stringify({
                'title': title,
                'body': body,
                'ext_params': {
                    'badge': 1,
                    'icon': 'none:url',
                    'group': 'none',
                    'url': url,
                },
                'category': 'category',
                'sound': "1107", //'minuet.caf'
            })
        }).catch(e => {
            console.log(`Bark fetch ${target} error ${e}`);
        })
    } catch (err) {
        console.log(`Bark send to ${name} error ${err}`)
    }
}

export {
    GetDefaultDevice,
    SetDefaultDevice,
    OnDefaultDeviceChange,

    GetDeviceList,
    SetDeviceList,
    OnDeviceListChange,

    Push,
    PushToDefault,
}
