import Store from "./storage";

export interface TabInfo {
    Uid: string
    Title?: string
    Url: string
    FavIconUrl?: string
    Option?: TabOption
}

export interface TabOption {
}

export interface TabGroup {
    Gid: number,
    Uid: string,
    Tabs: TabInfo[],
    Date: number,
    Option?: GroupOption,
}

export interface GroupOption {

}

export async function StoreTabs(tabs: TabInfo[]) {
    tabStore.get().then(arr => {
        arr = arr ?? [];
        arr.push({
            Uid: crypto.randomUUID(),
            Gid: arr.length + 1,
            Tabs: tabs,
            Date: Date.now(),
        });
        return arr;
    }).then(arr => {
        return tabStore.set(arr);
    })
}

const tabStore = new Store<TabGroup[]>("tab-store")
