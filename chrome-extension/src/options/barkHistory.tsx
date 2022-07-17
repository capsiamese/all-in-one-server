import {Box, Button, Center, Link, List, ListItem, Skeleton, useBoolean} from "@chakra-ui/react";
import React, {useEffect, useState} from "react";
import {Device, PullHistory, TargetDevice} from "../common/bark";
import Store from "../common/storage";
import {tab} from "../pb/compiled";
import _ from "lodash";
import BarkHistory = tab.BarkHistory;

export default function BarkHistoryView() {
    const [history, setHistory] = useState<BarkHistory[]>([])
    const ds = new Store<Device>(TargetDevice)
    const cache = new Store<BarkHistory[]>("bark-history-cache", [])

    useEffect(() => {
        cache.get().then(setHistory)
    }, [])

    const [skel, {on, off, toggle}] = useBoolean()

    const handleHistory = () => {
        on()
        ds.get().then(d => {
            let offset = history.length
            PullHistory(d.target, offset, 3).then(data => {
                off()
                history.push(...data)
                setHistory(history)
                cache.set(history).then()
            }).catch(off)
        }).catch(off)
    }

    const clearHistoryCache = () => {
        cache.set([]).then()
    }

    return (
        <Box>
            <List spacing={3}>
                {history.map((i, idx) => {
                    return (<ListItem key={idx}>
                        <Card item={i} index={idx}/>
                    </ListItem>)
                })}
                {skel ? _.times(3, () => {
                    return (<ListItem>
                        <Skeleton height='20px'/>
                    </ListItem>)
                }) : null}
            </List>
            <Center>
                <Button onClick={handleHistory}>Load more</Button>
                <Button onClick={clearHistoryCache}>Clear cache</Button>
            </Center>
        </Box>
    )
}

function Card({item, index}: { item: BarkHistory, index: number }) {
    const {title, content, params} = item
    const url = params?.url

    const bg = index % 2 == 1 ? "gray.100" : "gray.300";

    return (<Box bg={bg}>
        {title ? (<Box>
            Title: {title}
        </Box>) : null}
        {content ? (<Box>
            Content: {content}
        </Box>) : null}
        {url ? (<Box>
            Url: <Link href={url} isExternal>{url}</Link>
        </Box>) : null}
    </Box>)
}
