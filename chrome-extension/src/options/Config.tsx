import {Box, Button, ButtonGroup, Input, InputGroup, InputLeftAddon, Stack} from "@chakra-ui/react";
import {ChangeEvent, useEffect, useState} from "react";
import ServerMSG, {CodeSuccess} from "../common/serverMessage"
import Store from "../common/storage";
import Browser from "webextension-polyfill";
import _ from "lodash";
import {tab} from "../pb/compiled";

export default function ConfigView() {
    const [addr, setAddr] = useState('');
    const [name, setName] = useState('');
    const [uid, setUid] = useState('');

    const uidStore = new Store<string>('client_uid', '')
    const clientName = new Store<string>('client_name', '')
    const serverURL = new Store<string>('server_url', '')

    useEffect(() => {
        uidStore.get().then(setUid)
        clientName.get().then(setName)
        serverURL.get().then(setAddr)
    }, [])

    const handlePing = () => {
        setAddr(_.trimEnd(addr, '/'))
        register(addr, name).then(msg => {
            if (msg.code != CodeSuccess) {
                alert(msg.error)
                return
            }
            setUid(msg.data.uid)
            uidStore.set(msg.data.uid).then()
            clientName.set(name).then()
            serverURL.set(addr).then()
            console.log(msg);
        })
    }

    const handlePull = () => {
        PullTabs(addr, uid).then(data => {
            if (data.code != CodeSuccess) {
                alert(data.error)
                return
            }
            console.log(data)
        })
    }

    const handleAddr = (e: ChangeEvent<HTMLInputElement>) => {
        setAddr(e.target.value)
        e.preventDefault()
    }
    const handleName = (e: ChangeEvent<HTMLInputElement>) => {
        setName(e.target.value)
        e.preventDefault()
    }

    return (
        <>
            <Stack>
                {uid ? (
                    <Box>
                        Client Uid: {uid}
                    </Box>
                ) : null}
                <InputGroup>
                    <InputLeftAddon children='address'/>
                    <Input onChange={handleAddr} value={addr}/>

                    <InputLeftAddon children='client name'/>
                    <Input onChange={handleName} value={name}/>
                </InputGroup>
            </Stack>
            <ButtonGroup>
                <Button onClick={handlePing}>Register</Button>
                <Button onClick={handlePull}>Pull</Button>
            </ButtonGroup>
        </>
    )
}

async function register(addr: string, name: string): Promise<ServerMSG> {
    let id = Browser.runtime.id
    return fetch(`${addr}/register?name=${name}&extension=${id}`, {
        method: "POST"
    }).then(res => res.json())
}

async function PullTabs(addr: string, uid: string): Promise<ServerMSG> {
    let t = new tab.Tab({name: "hello"})
    console.log(t)
    return fetch(`${addr}/${uid}/group`, {
        method: "GEt"
    }).then(res => res.json())
}
