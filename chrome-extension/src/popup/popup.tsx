import React from "react";
import {createRoot} from "react-dom/client";
import {Box, ChakraProvider, Link, StackDivider, VStack} from '@chakra-ui/react'
import {HashRouter, Link as ReactLink, Route, Routes} from "react-router-dom";
import CustomMessage from "./customMessage";
import BottomNav from "./bottomNav";
import DeviceOption from "./deviceOption";
import {ArrowLeftIcon} from "@chakra-ui/icons";

function Popup() {
    return (
        <VStack
            divider={<StackDivider borderColor='gray.200'/>}
            spacing={2}
            align='stretch'
        >
            <DeviceOption/>
            <CustomMessage/>
            <BottomNav/>
        </VStack>
    )
}

function Config() {
    return (
        <Box fontSize='lg'>
            <Link as={ReactLink} to="/">
                <ArrowLeftIcon/>
            </Link>
            Hello, There nothing.
        </Box>
    )
}

const PopupView = Popup;
const ConfigView = Config;

const container = document.getElementById("root");
const root = createRoot(container!);
root.render(
    <ChakraProvider>
        <HashRouter>
            <Routes>
                <Route path="/" element={
                    <PopupView/>
                }/>
                <Route path="conf" element={
                    <ConfigView/>
                }/>
            </Routes>
        </HashRouter>
    </ChakraProvider>
);


