import Browser from "webextension-polyfill";
import {OnDefaultDeviceChange} from "../common/bark";
import {SendToastMessage} from "../common/toastMessage";
import {createContextMenu} from "../common/contextMenu";

Browser.runtime.onStartup.addListener(() => {

});

OnDefaultDeviceChange((data) => {
    console.log(data, "background received");
});

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
