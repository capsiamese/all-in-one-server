import {Button, Center, Textarea} from "@chakra-ui/react";
import {useState} from "react";
import Browser from "webextension-polyfill";
import {SendToastMessage} from "../common/toastMessage";

const DebugView = Debug;
export default DebugView;

function Debug() {
    const [data, setData] = useState<any>();
    const [del, setDel] = useState("");

    const reload = () => {
        Browser.storage.local.get(null).then(result => {
            setData(JSON.stringify(result, null, 4));
            console.log(result);
        })
    }

    const bg = () => {
        SendToastMessage("Hello");
    }

    return (
        <Center>
            <Textarea rows={30} cols={30} value={data}/>
            <Button onClick={reload}>Reload</Button>
            <Button onClick={bg}>BG</Button>
            <Button onClick={() => {
                Browser.storage.local.clear().catch(e => {
                    console.log(`clear extension storage ${e}`)
                })
            }}>Clear Storage</Button>
            <input onChange={e => setDel(e.target.value)}/>
            <Button onClick={() => {
                Browser.storage.local.set({[del]: null}).catch(e => {
                    console.log(`delete ${del} ${e}`)
                });
            }}>Delete
            </Button>
        </Center>
    )
}
