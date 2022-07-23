import {useEffect, useRef, useState} from "react";
import {Box, Button, Collapse, Flex, IconButton, Textarea, VStack} from "@chakra-ui/react";
import autosize from "autosize";
import {Push} from "../common/bark";
import {ArrowForwardIcon, SmallCloseIcon} from "@chakra-ui/icons";

export default function CustomMessage() {
    const [show, setShow] = useState(false);
    const textareaRef = useRef<HTMLTextAreaElement>(null);
    const [text, setText] = useState('');

    useEffect(() => {
        autosize<any>(textareaRef.current);
    })

    const send = () => {
        if (text.length == 0) {
            return
        }
        Push({
            content: text,
            title: "custom message",
        }).then();
    }

    return (
        <VStack spacing={0}>
            <Box bg='yellow.200' w='100%'>
                <Collapse startingHeight={'3.75rem'} in={show}>
                    <Textarea
                        ref={textareaRef}
                        style={{
                            minWidth: '10em'
                        }}
                        onChange={e => setText(e.target.value)}
                        value={text}
                        resize='none'
                        rows={5}
                        size={'sm'}
                        fontSize='md'
                        fontWeight='bold'
                        placeholder='Font Size Test'
                        onFocus={() => setShow(true)}
                        onBlur={() => setShow(false)}
                        shadow='md'
                        rounded='md'
                    >
                    </Textarea>
                </Collapse>
            </Box>
            <Flex w='100%'>
                <Button
                    w='100%'
                    size='sm'
                    onClick={send}
                    colorScheme='teal'
                    rightIcon={<ArrowForwardIcon/>}
                >
                    Send
                </Button>
                <IconButton
                    size='sm'
                    aria-label='clear-textarea'
                    colorScheme='pink'
                    onClick={() => setText('')}
                    icon={<SmallCloseIcon/>}
                />
            </Flex>
        </VStack>
    )
}
