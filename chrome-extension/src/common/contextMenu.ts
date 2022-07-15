import Browser, {Menus, Tabs} from "webextension-polyfill";
import {PushToDefault} from "./bark";
import {StoreTabs, TabInfo} from "./tabStore";
import Tab = Tabs.Tab;
import OnClickData = Menus.OnClickData;

export function createContextMenu() {
    items.forEach(i => {
        i();
    });
    Browser.contextMenus.onClicked.addListener(onContextMenuClick);
}

function onContextMenuClick(info: OnClickData, tab: Tab | undefined) {
    //handler[info.menuItemId]!(info, tab);
    handler[info.menuItemId]?.(info, tab);
}

type fun = () => void;
const items: fun[] = [];
items.push(() => {
    Browser.contextMenus.create({
        type: "normal",
        contexts: ['all'],
        id: BarkSender,
        title: 'Send',
    });
})

items.push(() => {
    Browser.contextMenus.create({
        type: "normal",
        contexts: ['all'],
        id: TabStoreThis,
        title: 'Store This Tab',
    });
})

const BarkSender = "bark-sender"
const TabStoreThis = "tab-store-this"
const TabStoreAll = "tab-store-all"
const TabStoreLeft = "tab-store-left"
const TabStoreRight = "tab-store-right"
const OpenOptionPage = "open-option-page"

const handler: any = {}
handler[BarkSender] = (info: OnClickData, tab: Tab | undefined) => {
    // TODO: 支持发送更多的类型, eg. selected text
    let u = info.pageUrl
        ?? info.frameUrl
        ?? info.linkUrl
        ?? info.srcUrl
    if (!u) {
        return
    }
    PushToDefault({url: u,});
}

handler[TabStoreThis] = (info: OnClickData, tab: Tab | undefined) => {
    tabOptions(TabStoreThis, info, tab);
}

const urlScheme = new RegExp(/^https?:\/\//);

const tabOptions = (op: string, info: OnClickData, tab: Tab | undefined) => {
    /*
    tab?.favIconUrl
    tab?.id
    tab?.index
    tab?.title
    tab?.url
     */
    Browser.tabs.query({currentWindow: true}).then((tabs: Tab[]) => {
        StoreTabs(tabs.filter(t => {
            switch (op) {
                case TabStoreThis:
                    return t.id == tab?.id
                case TabStoreAll:
                    return true
                case TabStoreLeft:
                    return t.index < (tab?.index ?? 0)
                case TabStoreRight:
                    return t.index > (tab?.index ?? t.index + 1)
                default:
                    return false
            }
        }).map<TabInfo>((t) => {
            if (t.id) {
                Browser.tabs.remove(t.id).then(() => {
                    console.log(`tab ${t.id} remove`)
                })
            }
            return {
                Uid: crypto.randomUUID(),
                Url: t.url ?? "",
                Title: t.title,
                FavIconUrl: t.favIconUrl,
            }
        })).then()
    })
}
