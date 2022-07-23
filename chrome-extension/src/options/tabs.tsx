import React, {useEffect, useState} from "react";
import Store from "../common/storage";
import {DragDropContext, Draggable, Droppable, DropResult} from "react-beautiful-dnd";
import {Accordion, AccordionButton, AccordionIcon, AccordionItem, AccordionPanel, Box} from "@chakra-ui/react";
import {tab} from "../pb/compiled";
import {TabStore} from "../common/storageKey";
import {PullTabGroup} from "../common/tabStore";
import IGroup = tab.IGroup;
import Group = tab.Group;
import Tab = tab.Tab;

const TabView = TabManager
export default TabView;

function TabManager() {
    const [groups, setGroups] = useState<IGroup[]>([]);
    const s = new Store<IGroup[]>(TabStore);
    useEffect(() => {
        s.get().then(setGroups);
        s.addListener(setGroups);
        PullTabGroup().then()
    }, []);

    const onDragEnd = (result: DropResult) => {
        if (!result.destination) {
            return
        }
        let sourceGIdx = groups.findIndex(g => {
            return g.uid == result.source.droppableId
        })
        let destGIdx = groups.findIndex(g => {
            return g.uid == result.destination?.droppableId
        })
        let destIdx = result.destination.index
        let srcIdx = result.source.index

        let item = groups[sourceGIdx].tabs?.splice(srcIdx, 1)[0];
        item ? groups[destGIdx].tabs?.splice(destIdx, 0, item) : null;
        s.set(groups).then();
    }

    return (
        <DragDropContext onDragEnd={onDragEnd}>
            {groups?.map((groupItem, idx) => (
                <TabGroupView item={new Group(groupItem)} key={groupItem.uid} idx={idx}/>
            ))}
        </DragDropContext>
    )
}

function TabGroupView({item, idx}: { item: Group, idx: number }) {
    return (
        <Accordion allowMultiple={true}>
            <AccordionItem>
                <AccordionButton>
                    <Box flex='1' textAlign='left'>
                        {item.uid}{`@${idx}`}
                    </Box>
                    <AccordionIcon/>
                </AccordionButton>
                <AccordionPanel>
                    <Droppable droppableId={item.uid}>
                        {provided => (
                            <div ref={provided.innerRef} {...provided.droppableProps}>
                                {item.tabs.map((t, i) => (
                                    <TabItem item={new Tab(t)} idx={i} key={t.uid/*这个key必须唯一且不变*/}/>
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

function TabItem({item, idx}: { item: Tab, idx: number }) {
    return (
        <Draggable draggableId={item.uid} index={idx}>
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
                        {item.name}{`@${idx}`}
                    </Box>
                </div>
            )}
        </Draggable>
    )
}
