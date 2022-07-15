import React, {useEffect} from "react";
import {createRoot} from "react-dom/client";
import {HashRouter, Route, Routes} from "react-router-dom";
import HeadNav from "./nav";
import {Center, ChakraProvider, Portal, Text, useToast} from "@chakra-ui/react";
import BarkManagerView from "./barkManager";
import DebugView from "./debug";
import {OnToastMessageReceive} from "../common/toastMessage";
import TabView from "./tabs";
import ConfigView from "./Config";

function Options() {
    return (
        <div>
            <TabView/>
        </div>
    )
}

const OptionsView = Options;

function Scaffold() {
    const toast = useToast()
    useEffect(() => {
        OnToastMessageReceive((msg, sender) => {
            console.log(`receive message ${msg}`);
            toast({
                title: 'receive',
                description: `a toast message ${msg}`,
                duration: 1000,
                isClosable: true,
            })
        })
    }, [])

    return (
        <>
            <HeadNav routes={{
                "Home": "/",
                "Bark": "bark",
                "Config": "config",
                "Debug": "debug",
            }}/>
            <Routes>
                <Route path="/" element={<OptionsView/>}/>
                <Route path="bark" element={<BarkManagerView/>}/>
                <Route path="config" element={<ConfigView/>}/>
                <Route path="debug" element={<DebugView/>}/>
            </Routes>
            <Portal>
                <Center>
                    <Text fontSize='xl' color='tomato'>
                        This is footer
                    </Text>
                </Center>
            </Portal>
        </>
    )
}

const container = document.getElementById("root");
const root = createRoot(container!);
root.render(
    <ChakraProvider>
        <HashRouter>
            <Scaffold/>
        </HashRouter>
    </ChakraProvider>
);
