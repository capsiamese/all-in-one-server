import React, {useEffect, useState} from "react";
import {
    Button,
    ButtonGroup,
    Center,
    Editable,
    EditableInput,
    EditablePreview,
    Input,
    Table,
    TableContainer,
    Tbody,
    Td,
    Th,
    Thead,
    Tooltip,
    Tr,
    VStack
} from "@chakra-ui/react";
import Browser from "webextension-polyfill";
import BarkHistoryView from "./barkHistory";
import Store from "../common/storage";
import {BarkDefaultDevice, BarkDeviceList} from "../common/storageKey";
import {tab} from "../pb/compiled";
import BarkDevice = tab.BarkDevice;
import IBarkDevice = tab.IBarkDevice;

const BarkManagerView = BarkManager;
export default BarkManagerView;

function BarkManager() {
    return (
        <>
            <Center>
                <VStack>
                    <DeviceStack/>
                    <Button onClick={() => {
                        Browser.storage.local.clear().catch(e => {
                            console.log(`clear extension storage ${e}`)
                        })
                    }}>Clear Storage</Button>

                    <BarkHistoryView/>
                </VStack>
            </Center>
        </>
    )
}

function DeviceStack() {
    const defaultDeviceStore = new Store<BarkDevice>(BarkDefaultDevice)
    const [device, setDefaultDevice] = useState<BarkDevice>();
    defaultDeviceStore.get().then(d => {
        setDefaultDevice(d)
    })

    const deviceListStore = new Store<BarkDevice[]>(BarkDeviceList, [])
    const [deviceList, setDeviceList] = useState<BarkDevice[]>([]);
    deviceListStore.get().then(list => {
        setDeviceList(list)
    })

    useEffect(() => {
        defaultDeviceStore.addListener(setDefaultDevice);
        deviceListStore.addListener(setDeviceList);
    }, []);

    const handleChangeDevice = (newDevice: string) => {
        let newDefault = deviceList.find((d) => {
            return d.name == newDevice
        });
        if (newDefault) {
            defaultDeviceStore.set(newDefault).then()
        }
    }
    const editDevice = (idx: number, d: IBarkDevice) => {
        let item = new BarkDevice(d);
        let old = deviceList[idx];
        deviceList[idx] = item
        if (old.name == device?.name) {
            defaultDeviceStore.set(item).then();
        }
        deviceListStore.set(deviceList).then();
    }

    const handleRemoveDevice = (deviceName: string) => {
        if (deviceName == device?.name) {
            return
        }
        let idx = deviceList.findIndex(d => {
            return d.name == deviceName
        })
        deviceList.splice(idx, 1)
        deviceListStore.set(deviceList).then()
    }

    const [newName, setNewName] = useState('');
    const [newTarget, setNewTarget] = useState('');

    const AddNewDevice = () => {
        if (newName == '' || newTarget == '') {
            return
        }
        for (const i of deviceList) {
            if (i.name == newName) {
                // todo: show error
                return
            }
        }
        deviceList.push(new BarkDevice({
            name: newName,
            url: newTarget,
        }));
        deviceListStore.set(deviceList).then(() => {
            setNewName('');
            setNewTarget('');
        });
    }

    return (
        <TableContainer>
            <Table variant='striped' colorScheme='gray'>
                <Thead>
                    <Tr>
                        <Th>设备名</Th>
                        <Th>URL</Th>
                        <Th></Th>
                    </Tr>
                </Thead>
                <Tbody>
                    {deviceList.map((d, i) => {
                        return (
                            <TbodyDeviceItem
                                idx={i}
                                device={d}
                                key={i}
                                isDefault={device?.name == d.name}
                                editDevice={editDevice}
                                setDefault={handleChangeDevice}
                                remove={handleRemoveDevice}
                            />
                        )
                    })}
                    <Tr>
                        <Td> <Input value={newName} onChange={e => setNewName(e.target.value)}/> </Td>
                        <Td> <Input value={newTarget} onChange={e => setNewTarget(e.target.value)}/> </Td>
                        <Td>
                            <Tooltip hasArrow label='Add new device'>
                                <Button onClick={AddNewDevice}
                                        isDisabled={newName == "" || newTarget == ""}
                                >Add</Button>
                            </Tooltip>
                        </Td>
                    </Tr>
                </Tbody>
            </Table>
        </TableContainer>
    )
}

function TbodyDeviceItem({idx, device, editDevice, setDefault, isDefault, remove}: {
    idx: number,
    device: BarkDevice,
    isDefault: boolean,
    editDevice: (idx: number, device: IBarkDevice) => void,
    setDefault: (name: string) => void
    remove: (name: string) => void
}) {
    const [edit, setEdit] = useState(false);
    const [name, setN] = useState(device.name);
    const [target, setT] = useState(device.url);

    return (
        <Tr>
            <EditableTableItem
                defaultValue={device.name}
                onChange={setN}
                onBlur={() => setEdit(false)}
                onEdit={() => setEdit(true)}
            />
            <EditableTableItem
                defaultValue={device.url}
                onChange={setT}
                onBlur={() => setEdit(false)}
                onEdit={() => setEdit(true)}
            />
            <Td>
                {edit ? (
                    <ButtonGroup>
                        <Button onClick={() => {
                            editDevice(idx, {name, url: target})
                            setEdit(false)
                        }}>Ok</Button>
                    </ButtonGroup>
                ) : (
                    <ButtonGroup>
                        <Tooltip hasArrow label='Set default' bg='blue.400' isDisabled={isDefault}>
                            <Button
                                isDisabled={isDefault}
                                onClick={() => setDefault(device.name)}
                            >*</Button>
                        </Tooltip>
                        <Tooltip hasArrow label='Remove device' bg='red.400'>
                            <Button
                                isDisabled={isDefault}
                                onClick={() => remove(device.name)}
                            >x</Button>
                        </Tooltip>
                    </ButtonGroup>
                )}
            </Td>
        </Tr>
    )
}

function EditableTableItem(props: {
    defaultValue: string,
    onEdit: () => void,
    onBlur: () => void,
    onChange: (v: string) => void,
}) {
    return (
        <Td>
            <Editable
                {...props}
                submitOnBlur={false}
            >
                <EditablePreview/>
                <EditableInput/>
            </Editable>
        </Td>
    )
}
