import {Flex, Select, Square} from "@chakra-ui/react";
import {HamburgerIcon} from "@chakra-ui/icons";
import React, {ChangeEvent, useState} from "react";
import {tab} from "../pb/compiled";
import Store from "../common/storage";
import {BarkDefaultDevice, BarkDeviceList} from "../common/storageKey";
import BarkDevice = tab.BarkDevice;

export default function DeviceOption() {
    const dfs = new Store<BarkDevice>(BarkDefaultDevice, new BarkDevice());
    const dl = new Store<BarkDevice[]>(BarkDeviceList, []);
    const [def, setDef] = useState<BarkDevice>(new BarkDevice({name: "", url: ""}));
    const [list, setList] = useState<BarkDevice[]>([]);

    dfs.get().then(setDef)
    dl.get().then(setList)
    dfs.addListener(setDef)
    dl.addListener(setList)

    const handleChange = (event: ChangeEvent<HTMLSelectElement>) => {
        for (let i of list) {
            if (i.name === event.target.value) {
                dfs.set(i).then()
                break
            }
        }
        event.preventDefault();
    }
    return (
        <Flex
            shadow='md'
            rounded='md'
        >
            <Square
                size='32px'
                bg='green.200'
            >
                <HamburgerIcon/>
            </Square>
            <Select
                onChange={handleChange}
                size='sm'
                fontSize='md'
                fontWeight='bold'
                value={def.name}
                variant='flushed'
            >
                {list?.map(({name}) => {
                    return <option key={name} value={name}>{name}</option>
                })}
            </Select>
        </Flex>
    )
}
