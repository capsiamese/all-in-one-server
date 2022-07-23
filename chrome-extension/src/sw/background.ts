import Browser from "webextension-polyfill";
import {SendToastMessage} from "../common/toastMessage";
import {createContextMenu} from "../common/contextMenu";
import Store from "../common/storage";
import {tab} from "../pb/compiled";
import {BarkDefaultDevice} from "../common/storageKey";
import BarkDevice = tab.BarkDevice;

Browser.runtime.onStartup.addListener(() => {

});

const defaultDevice = new Store<BarkDevice>(BarkDefaultDevice)
defaultDevice.addListener(d => {
    console.log(d, "background received");
})

console.log(navigator);
console.log(Notification);
Browser.notifications.getAll().then();
Browser.webNavigation.onBeforeNavigate.hasListeners();

setInterval(() => {
    //SendToastMessage("Hello");
}, 5000);

Browser.alarms.create('toast', {
    when: Date.now(),
    periodInMinutes: 12,
});
Browser.alarms.onAlarm.addListener(alarm => {
    if (alarm.name == 'toast') {
        SendToastMessage("Hello");
    }
});


const uninstallUrl = ``
Browser.runtime.setUninstallURL(uninstallUrl).then(() => {
    console.log("uninstall url set");
    console.log(Browser.runtime.lastError);
})

Browser.runtime.onInstalled.addListener(details => {
    console.log('extension install', details);
    createContextMenu();
})
