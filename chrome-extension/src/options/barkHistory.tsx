import {Box, Button, List, ListIcon, ListItem} from "@chakra-ui/react";
import {CloseIcon} from "@chakra-ui/icons";
import React, {useState} from "react";
import {Device, PullHistory, TargetDevice} from "../common/bark";
import Store from "../common/storage";

export default function BarkHistoryView() {
    const [history, setHistory] = useState([
        'Item1', 'Item2', 'Item3'
    ])
    const ds = new Store<Device>(TargetDevice)

    const handleHistory = () => {
        ds.get().then(d => {
            PullHistory(d.target, 0, 10).then(data => {
                console.log(data)
            })
        })
    }

    return (
        <Box>
            <Button onClick={handleHistory}>Get</Button>
            <List spacing={3}>
                {history.map((i, idx) => {
                    return (<ListItem key={idx}>
                        <ListIcon as={CloseIcon} color='red.300'/>
                        {i}
                    </ListItem>)
                })}
            </List>
        </Box>
    )
}
