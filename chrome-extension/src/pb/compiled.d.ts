import * as $protobuf from "protobufjs";
/** Namespace tab. */
export namespace tab {

    /** Properties of a Tab. */
    interface ITab {

        /** Tab name */
        name?: (string|null);

        /** Tab url */
        url?: (string|null);

        /** Tab favicon */
        favicon?: (string|null);

        /** Tab uid */
        uid?: (string|null);

        /** Tab index */
        index?: (number|null);
    }

    /** Represents a Tab. */
    class Tab implements ITab {

        /**
         * Constructs a new Tab.
         * @param [properties] Properties to set
         */
        constructor(properties?: tab.ITab);

        /** Tab name. */
        public name: string;

        /** Tab url. */
        public url: string;

        /** Tab favicon. */
        public favicon: string;

        /** Tab uid. */
        public uid: string;

        /** Tab index. */
        public index: number;

        /**
         * Creates a new Tab instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Tab instance
         */
        public static create(properties?: tab.ITab): tab.Tab;

        /**
         * Encodes the specified Tab message. Does not implicitly {@link tab.Tab.verify|verify} messages.
         * @param message Tab message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: tab.ITab, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Tab message, length delimited. Does not implicitly {@link tab.Tab.verify|verify} messages.
         * @param message Tab message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: tab.ITab, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Tab message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Tab
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): tab.Tab;

        /**
         * Decodes a Tab message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Tab
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): tab.Tab;

        /**
         * Verifies a Tab message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Tab message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Tab
         */
        public static fromObject(object: { [k: string]: any }): tab.Tab;

        /**
         * Creates a plain object from a Tab message. Also converts values to other types if specified.
         * @param message Tab
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: tab.Tab, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Tab to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Tab
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a GroupOption. */
    interface IGroupOption {

        /** GroupOption tags */
        tags?: (string[]|null);
    }

    /** Represents a GroupOption. */
    class GroupOption implements IGroupOption {

        /**
         * Constructs a new GroupOption.
         * @param [properties] Properties to set
         */
        constructor(properties?: tab.IGroupOption);

        /** GroupOption tags. */
        public tags: string[];

        /**
         * Creates a new GroupOption instance using the specified properties.
         * @param [properties] Properties to set
         * @returns GroupOption instance
         */
        public static create(properties?: tab.IGroupOption): tab.GroupOption;

        /**
         * Encodes the specified GroupOption message. Does not implicitly {@link tab.GroupOption.verify|verify} messages.
         * @param message GroupOption message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: tab.IGroupOption, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified GroupOption message, length delimited. Does not implicitly {@link tab.GroupOption.verify|verify} messages.
         * @param message GroupOption message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: tab.IGroupOption, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a GroupOption message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns GroupOption
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): tab.GroupOption;

        /**
         * Decodes a GroupOption message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns GroupOption
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): tab.GroupOption;

        /**
         * Verifies a GroupOption message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a GroupOption message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns GroupOption
         */
        public static fromObject(object: { [k: string]: any }): tab.GroupOption;

        /**
         * Creates a plain object from a GroupOption message. Also converts values to other types if specified.
         * @param message GroupOption
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: tab.GroupOption, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this GroupOption to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for GroupOption
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Group. */
    interface IGroup {

        /** Group name */
        name?: (string|null);

        /** Group uid */
        uid?: (string|null);

        /** Group index */
        index?: (number|null);

        /** Group option */
        option?: (tab.IGroupOption|null);

        /** Group tabs */
        tabs?: (tab.ITab[]|null);
    }

    /** Represents a Group. */
    class Group implements IGroup {

        /**
         * Constructs a new Group.
         * @param [properties] Properties to set
         */
        constructor(properties?: tab.IGroup);

        /** Group name. */
        public name: string;

        /** Group uid. */
        public uid: string;

        /** Group index. */
        public index: number;

        /** Group option. */
        public option?: (tab.IGroupOption|null);

        /** Group tabs. */
        public tabs: tab.ITab[];

        /**
         * Creates a new Group instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Group instance
         */
        public static create(properties?: tab.IGroup): tab.Group;

        /**
         * Encodes the specified Group message. Does not implicitly {@link tab.Group.verify|verify} messages.
         * @param message Group message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: tab.IGroup, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Group message, length delimited. Does not implicitly {@link tab.Group.verify|verify} messages.
         * @param message Group message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: tab.IGroup, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Group message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Group
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): tab.Group;

        /**
         * Decodes a Group message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Group
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): tab.Group;

        /**
         * Verifies a Group message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Group message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Group
         */
        public static fromObject(object: { [k: string]: any }): tab.Group;

        /**
         * Creates a plain object from a Group message. Also converts values to other types if specified.
         * @param message Group
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: tab.Group, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Group to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Group
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a Client. */
    interface IClient {

        /** Client name */
        name?: (string|null);

        /** Client extId */
        extId?: (string|null);

        /** Client uid */
        uid?: (string|null);

        /** Client lastAccessTime */
        lastAccessTime?: (number|Long|null);

        /** Client groups */
        groups?: (tab.IGroup[]|null);
    }

    /** Represents a Client. */
    class Client implements IClient {

        /**
         * Constructs a new Client.
         * @param [properties] Properties to set
         */
        constructor(properties?: tab.IClient);

        /** Client name. */
        public name: string;

        /** Client extId. */
        public extId: string;

        /** Client uid. */
        public uid: string;

        /** Client lastAccessTime. */
        public lastAccessTime: (number|Long);

        /** Client groups. */
        public groups: tab.IGroup[];

        /**
         * Creates a new Client instance using the specified properties.
         * @param [properties] Properties to set
         * @returns Client instance
         */
        public static create(properties?: tab.IClient): tab.Client;

        /**
         * Encodes the specified Client message. Does not implicitly {@link tab.Client.verify|verify} messages.
         * @param message Client message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: tab.IClient, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified Client message, length delimited. Does not implicitly {@link tab.Client.verify|verify} messages.
         * @param message Client message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: tab.IClient, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a Client message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns Client
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): tab.Client;

        /**
         * Decodes a Client message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns Client
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): tab.Client;

        /**
         * Verifies a Client message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a Client message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns Client
         */
        public static fromObject(object: { [k: string]: any }): tab.Client;

        /**
         * Creates a plain object from a Client message. Also converts values to other types if specified.
         * @param message Client
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: tab.Client, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this Client to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for Client
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }

    /** Properties of a BarkHistory. */
    interface IBarkHistory {

        /** BarkHistory id */
        id?: (number|Long|null);

        /** BarkHistory ts */
        ts?: (number|Long|null);

        /** BarkHistory from */
        from?: (string|null);

        /** BarkHistory key */
        key?: (string|null);

        /** BarkHistory title */
        title?: (string|null);

        /** BarkHistory content */
        content?: (string|null);

        /** BarkHistory params */
        params?: ({ [k: string]: string }|null);
    }

    /** Represents a BarkHistory. */
    class BarkHistory implements IBarkHistory {

        /**
         * Constructs a new BarkHistory.
         * @param [properties] Properties to set
         */
        constructor(properties?: tab.IBarkHistory);

        /** BarkHistory id. */
        public id: (number|Long);

        /** BarkHistory ts. */
        public ts: (number|Long);

        /** BarkHistory from. */
        public from: string;

        /** BarkHistory key. */
        public key: string;

        /** BarkHistory title. */
        public title: string;

        /** BarkHistory content. */
        public content: string;

        /** BarkHistory params. */
        public params: { [k: string]: string };

        /**
         * Creates a new BarkHistory instance using the specified properties.
         * @param [properties] Properties to set
         * @returns BarkHistory instance
         */
        public static create(properties?: tab.IBarkHistory): tab.BarkHistory;

        /**
         * Encodes the specified BarkHistory message. Does not implicitly {@link tab.BarkHistory.verify|verify} messages.
         * @param message BarkHistory message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encode(message: tab.IBarkHistory, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Encodes the specified BarkHistory message, length delimited. Does not implicitly {@link tab.BarkHistory.verify|verify} messages.
         * @param message BarkHistory message or plain object to encode
         * @param [writer] Writer to encode to
         * @returns Writer
         */
        public static encodeDelimited(message: tab.IBarkHistory, writer?: $protobuf.Writer): $protobuf.Writer;

        /**
         * Decodes a BarkHistory message from the specified reader or buffer.
         * @param reader Reader or buffer to decode from
         * @param [length] Message length if known beforehand
         * @returns BarkHistory
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): tab.BarkHistory;

        /**
         * Decodes a BarkHistory message from the specified reader or buffer, length delimited.
         * @param reader Reader or buffer to decode from
         * @returns BarkHistory
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): tab.BarkHistory;

        /**
         * Verifies a BarkHistory message.
         * @param message Plain object to verify
         * @returns `null` if valid, otherwise the reason why it is not
         */
        public static verify(message: { [k: string]: any }): (string|null);

        /**
         * Creates a BarkHistory message from a plain object. Also converts values to their respective internal types.
         * @param object Plain object
         * @returns BarkHistory
         */
        public static fromObject(object: { [k: string]: any }): tab.BarkHistory;

        /**
         * Creates a plain object from a BarkHistory message. Also converts values to other types if specified.
         * @param message BarkHistory
         * @param [options] Conversion options
         * @returns Plain object
         */
        public static toObject(message: tab.BarkHistory, options?: $protobuf.IConversionOptions): { [k: string]: any };

        /**
         * Converts this BarkHistory to JSON.
         * @returns JSON object
         */
        public toJSON(): { [k: string]: any };

        /**
         * Gets the default type url for BarkHistory
         * @param [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns The default type url
         */
        public static getTypeUrl(typeUrlPrefix?: string): string;
    }
}
