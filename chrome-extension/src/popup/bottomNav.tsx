import React from "react";
import {Button, Center, Flex, Link, Spacer} from "@chakra-ui/react";
import {ArrowRightIcon, SettingsIcon} from "@chakra-ui/icons";
import {Link as ReactLink} from "react-router-dom";
import Browser from "webextension-polyfill";

export default function BottomNav() {
    return (
        <Flex>
            <Button
                leftIcon={<SettingsIcon/>}
                colorScheme='pink'
                variant='solid'
                size='sm'
                onClick={() => Browser.runtime.openOptionsPage().then()}
            >
                Settings
            </Button>
            {/*
            <Link href={'options.html'} target="_blank">
                <AttachmentIcon/>
            </Link>
            */}
            <Spacer/>
            <Center>
                <Link as={ReactLink} to="conf" size='sm'>
                    <ArrowRightIcon/>
                </Link>
            </Center>
        </Flex>
    )
}
