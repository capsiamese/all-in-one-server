import Store from "./storage";
import {tab} from "../pb/compiled";
import {TabStore} from "./storageKey";
import ITab = tab.ITab;
import Group = tab.Group;

export async function StoreTabs(tabs: ITab[]) {
    const tabStore = new Store<Group[]>(TabStore)
    tabStore.get().then(arr => {
        arr = arr ?? [];
        arr.push(new Group({
            uid: crypto.randomUUID(),
            index: arr.length + 1,
            tabs: tabs,
            ts: Date.now() / 1000,
            name: "group",
        }));
        return arr;
    }).then(arr => {
        return tabStore.set(arr);
    })
    // todo: send to server
}

export async function PullTabGroup(): Promise<Group[]> {
    return []
}
