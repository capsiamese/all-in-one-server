import React, {useEffect, useState} from "react";
import {TabGroup, TabInfo} from "../common/tabStore";
import Store from "../common/storage";
import {DragDropContext, Draggable, Droppable, DropResult} from "react-beautiful-dnd";
import {Accordion, AccordionButton, AccordionIcon, AccordionItem, AccordionPanel, Box} from "@chakra-ui/react";

const TabView = TabManager
export default TabView;

function TabManager() {
    const [groups, setGroups] = useState<TabGroup[]>([]);
    const s = new Store<TabGroup[]>('tab-store');
    useEffect(() => {
        s.get().then(setGroups);
        s.addListener(setGroups);
    }, []);

    const onDragEnd = (result: DropResult) => {
        if (!result.destination) {
            return
        }
        let sourceGIdx = groups.findIndex(g => {
            return g.Uid == result.source.droppableId
        })
        let destGIdx = groups.findIndex(g => {
            return g.Uid == result.destination?.droppableId
        })
        let destIdx = result.destination.index
        let srcIdx = result.source.index

        let item = groups[sourceGIdx].Tabs.splice(srcIdx, 1)[0];
        groups[destGIdx].Tabs.splice(destIdx, 0, item);
        s.set(groups).then();
    }

    return (
        <DragDropContext onDragEnd={onDragEnd}>
            {groups?.map((groupItem, idx) => (
                <TabGroupView item={groupItem} key={groupItem.Uid} idx={idx}/>
            ))}
        </DragDropContext>
    )
}

function TabGroupView({item, idx}: { item: TabGroup, idx: number }) {
    return (
        <Accordion allowMultiple={true}>
            <AccordionItem>
                <AccordionButton>
                    <Box flex='1' textAlign='left'>
                        {item.Uid}{`@${idx}`}
                    </Box>
                    <AccordionIcon/>
                </AccordionButton>
                <AccordionPanel>
                    <Droppable droppableId={item.Uid}>
                        {provided => (
                            <div ref={provided.innerRef} {...provided.droppableProps}>
                                {item?.Tabs?.map((t, i) => (
                                    <TabItem item={t} idx={i} key={t.Uid/*这个key必须唯一且不变*/}/>
                                ))}
                                {provided.placeholder}
                            </div>
                        )}
                    </Droppable>
                </AccordionPanel>
            </AccordionItem>
        </Accordion>
    )
}

function TabItem({item, idx}: { item: TabInfo, idx: number }) {
    return (
        <Draggable draggableId={item.Uid} index={idx}>
            {provided => (
                <div ref={provided.innerRef}
                     {...provided.draggableProps}
                     {...provided.dragHandleProps}>
                    <Box
                        p='40px'
                        color='white'
                        mt='4'
                        bg='teal.500'
                        rounded='md'
                        shadow='md'
                    >
                        {item.Title}{`@${idx}`}
                    </Box>
                </div>
            )}
        </Draggable>
    )
}
