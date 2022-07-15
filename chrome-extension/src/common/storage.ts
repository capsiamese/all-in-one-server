import Browser, {Storage} from "webextension-polyfill";
import StorageAreaOnChangedChangesType = Storage.StorageAreaOnChangedChangesType;

const storage = Browser.storage.local;

export default class Store<T> {
    private readonly storageName: string
    private readonly defaultVal?: T
    private listener: ((arg0: StorageAreaOnChangedChangesType) => void) | undefined

    constructor(name: string, defaultValue?: T) {
        this.storageName = name;
        this.defaultVal = defaultValue
    }

    public async set(v: T) {
        return storage.set({
            [this.storageName]: v ?? this.defaultVal,
        })
    }

    public async get(): Promise<T> {
        return new Promise<T>((resolve, reject) => {
            storage.get(this.storageName).then(v => {
                resolve(v[this.storageName] ?? this.defaultVal)
            }).catch(err => {
                reject(err)
            })
        })
    }

    public addListener(f: (t: T) => void) {
        this.listener = (changes: StorageAreaOnChangedChangesType) => {
            if (changes[this.storageName]) {
                let c = changes[this.storageName] as StorageAreaOnChangedChangesType;
                f(c.newValue ?? this.defaultVal);
            }
        }
        storage.onChanged.addListener(this.listener)
    }

    public removeListener() {
        if (this.listener) {
            storage.onChanged.removeListener(this.listener)
        }
    }
}
