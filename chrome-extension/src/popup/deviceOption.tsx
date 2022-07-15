import {Flex, Select, Square} from "@chakra-ui/react";
import {HamburgerIcon} from "@chakra-ui/icons";
import React, {ChangeEvent, useState} from "react";
import {
    Device,
    GetDefaultDevice,
    GetDeviceList,
    OnDefaultDeviceChange,
    OnDeviceListChange,
    SetDefaultDevice
} from "../common/bark";

export default function DeviceOption() {
    const [def, setDef] = useState<Device>({name: "", target: ""});
    const [list, setList] = useState<Device[]>([]);

    GetDefaultDevice().then(setDef);
    OnDefaultDeviceChange(setDef);

    GetDeviceList().then(setList);
    OnDeviceListChange(setList);

    const handleChange = (event: ChangeEvent<HTMLSelectElement>) => {
        for (let i of list) {
            if (i.name === event.target.value) {
                SetDefaultDevice(i).then();
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
