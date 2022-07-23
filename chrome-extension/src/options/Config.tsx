import {Box, Button, ButtonGroup, Input, InputGroup, InputLeftAddon, Stack} from "@chakra-ui/react";
import {ChangeEvent, useEffect, useState} from "react";
import Store from "../common/storage";
import _ from "lodash";
import {PullTabs, Register} from "../common/webRequest";
import {ConfigStore} from "../common/storageKey";
import {tab} from "../pb/compiled";
import ClientConfig = tab.ClientConfig;


export default function ConfigView() {
    const [addr, setAddr] = useState('');
    const [name, setName] = useState('');
    const [uid, setUid] = useState('');

    const configStore = new Store<ClientConfig>(ConfigStore, new ClientConfig({name: "", url: "", uid: ""}))

    useEffect(() => {
        configStore.get().then(res => {
            setUid(res.uid)
            setName(res.name)
            setAddr(res.url)
        })
    }, [])

    const handlePing = () => {
        setAddr(_.trimEnd(addr, '/'))
        Register(addr, name).then(msg => {
            if (msg.code != 0) {
                alert(msg.error)
                return
            }
            setUid(msg.data.uid)
            configStore.set(new ClientConfig({
                name, url: addr, uid: msg.data.uid,
            })).then()
            console.log(msg);
        })
    }

    const handlePull = () => {
        PullTabs(addr, uid).then(data => {
            if (data.code != 0) {
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
