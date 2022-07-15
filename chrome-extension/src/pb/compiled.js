/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.tab = (function() {

    /**
     * Namespace tab.
     * @exports tab
     * @namespace
     */
    var tab = {};

    tab.Tab = (function() {

        /**
         * Properties of a Tab.
         * @memberof tab
         * @interface ITab
         * @property {string|null} [name] Tab name
         * @property {string|null} [url] Tab url
         * @property {string|null} [favicon] Tab favicon
         * @property {string|null} [uid] Tab uid
         * @property {number|null} [index] Tab index
         */

        /**
         * Constructs a new Tab.
         * @memberof tab
         * @classdesc Represents a Tab.
         * @implements ITab
         * @constructor
         * @param {tab.ITab=} [properties] Properties to set
         */
        function Tab(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Tab name.
         * @member {string} name
         * @memberof tab.Tab
         * @instance
         */
        Tab.prototype.name = "";

        /**
         * Tab url.
         * @member {string} url
         * @memberof tab.Tab
         * @instance
         */
        Tab.prototype.url = "";

        /**
         * Tab favicon.
         * @member {string} favicon
         * @memberof tab.Tab
         * @instance
         */
        Tab.prototype.favicon = "";

        /**
         * Tab uid.
         * @member {string} uid
         * @memberof tab.Tab
         * @instance
         */
        Tab.prototype.uid = "";

        /**
         * Tab index.
         * @member {number} index
         * @memberof tab.Tab
         * @instance
         */
        Tab.prototype.index = 0;

        /**
         * Creates a new Tab instance using the specified properties.
         * @function create
         * @memberof tab.Tab
         * @static
         * @param {tab.ITab=} [properties] Properties to set
         * @returns {tab.Tab} Tab instance
         */
        Tab.create = function create(properties) {
            return new Tab(properties);
        };

        /**
         * Encodes the specified Tab message. Does not implicitly {@link tab.Tab.verify|verify} messages.
         * @function encode
         * @memberof tab.Tab
         * @static
         * @param {tab.ITab} message Tab message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Tab.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
            if (message.url != null && Object.hasOwnProperty.call(message, "url"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.url);
            if (message.favicon != null && Object.hasOwnProperty.call(message, "favicon"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.favicon);
            if (message.uid != null && Object.hasOwnProperty.call(message, "uid"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.uid);
            if (message.index != null && Object.hasOwnProperty.call(message, "index"))
                writer.uint32(/* id 5, wireType 0 =*/40).int32(message.index);
            return writer;
        };

        /**
         * Encodes the specified Tab message, length delimited. Does not implicitly {@link tab.Tab.verify|verify} messages.
         * @function encodeDelimited
         * @memberof tab.Tab
         * @static
         * @param {tab.ITab} message Tab message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Tab.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Tab message from the specified reader or buffer.
         * @function decode
         * @memberof tab.Tab
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {tab.Tab} Tab
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Tab.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.tab.Tab();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.name = reader.string();
                        break;
                    }
                case 2: {
                        message.url = reader.string();
                        break;
                    }
                case 3: {
                        message.favicon = reader.string();
                        break;
                    }
                case 4: {
                        message.uid = reader.string();
                        break;
                    }
                case 5: {
                        message.index = reader.int32();
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Tab message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof tab.Tab
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {tab.Tab} Tab
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Tab.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Tab message.
         * @function verify
         * @memberof tab.Tab
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Tab.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.url != null && message.hasOwnProperty("url"))
                if (!$util.isString(message.url))
                    return "url: string expected";
            if (message.favicon != null && message.hasOwnProperty("favicon"))
                if (!$util.isString(message.favicon))
                    return "favicon: string expected";
            if (message.uid != null && message.hasOwnProperty("uid"))
                if (!$util.isString(message.uid))
                    return "uid: string expected";
            if (message.index != null && message.hasOwnProperty("index"))
                if (!$util.isInteger(message.index))
                    return "index: integer expected";
            return null;
        };

        /**
         * Creates a Tab message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof tab.Tab
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {tab.Tab} Tab
         */
        Tab.fromObject = function fromObject(object) {
            if (object instanceof $root.tab.Tab)
                return object;
            var message = new $root.tab.Tab();
            if (object.name != null)
                message.name = String(object.name);
            if (object.url != null)
                message.url = String(object.url);
            if (object.favicon != null)
                message.favicon = String(object.favicon);
            if (object.uid != null)
                message.uid = String(object.uid);
            if (object.index != null)
                message.index = object.index | 0;
            return message;
        };

        /**
         * Creates a plain object from a Tab message. Also converts values to other types if specified.
         * @function toObject
         * @memberof tab.Tab
         * @static
         * @param {tab.Tab} message Tab
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Tab.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.name = "";
                object.url = "";
                object.favicon = "";
                object.uid = "";
                object.index = 0;
            }
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.url != null && message.hasOwnProperty("url"))
                object.url = message.url;
            if (message.favicon != null && message.hasOwnProperty("favicon"))
                object.favicon = message.favicon;
            if (message.uid != null && message.hasOwnProperty("uid"))
                object.uid = message.uid;
            if (message.index != null && message.hasOwnProperty("index"))
                object.index = message.index;
            return object;
        };

        /**
         * Converts this Tab to JSON.
         * @function toJSON
         * @memberof tab.Tab
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Tab.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Tab
         * @function getTypeUrl
         * @memberof tab.Tab
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Tab.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/tab.Tab";
        };

        return Tab;
    })();

    tab.GroupOption = (function() {

        /**
         * Properties of a GroupOption.
         * @memberof tab
         * @interface IGroupOption
         * @property {Array.<string>|null} [tags] GroupOption tags
         */

        /**
         * Constructs a new GroupOption.
         * @memberof tab
         * @classdesc Represents a GroupOption.
         * @implements IGroupOption
         * @constructor
         * @param {tab.IGroupOption=} [properties] Properties to set
         */
        function GroupOption(properties) {
            this.tags = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * GroupOption tags.
         * @member {Array.<string>} tags
         * @memberof tab.GroupOption
         * @instance
         */
        GroupOption.prototype.tags = $util.emptyArray;

        /**
         * Creates a new GroupOption instance using the specified properties.
         * @function create
         * @memberof tab.GroupOption
         * @static
         * @param {tab.IGroupOption=} [properties] Properties to set
         * @returns {tab.GroupOption} GroupOption instance
         */
        GroupOption.create = function create(properties) {
            return new GroupOption(properties);
        };

        /**
         * Encodes the specified GroupOption message. Does not implicitly {@link tab.GroupOption.verify|verify} messages.
         * @function encode
         * @memberof tab.GroupOption
         * @static
         * @param {tab.IGroupOption} message GroupOption message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GroupOption.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.tags != null && message.tags.length)
                for (var i = 0; i < message.tags.length; ++i)
                    writer.uint32(/* id 1, wireType 2 =*/10).string(message.tags[i]);
            return writer;
        };

        /**
         * Encodes the specified GroupOption message, length delimited. Does not implicitly {@link tab.GroupOption.verify|verify} messages.
         * @function encodeDelimited
         * @memberof tab.GroupOption
         * @static
         * @param {tab.IGroupOption} message GroupOption message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        GroupOption.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a GroupOption message from the specified reader or buffer.
         * @function decode
         * @memberof tab.GroupOption
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {tab.GroupOption} GroupOption
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GroupOption.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.tab.GroupOption();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        if (!(message.tags && message.tags.length))
                            message.tags = [];
                        message.tags.push(reader.string());
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a GroupOption message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof tab.GroupOption
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {tab.GroupOption} GroupOption
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        GroupOption.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a GroupOption message.
         * @function verify
         * @memberof tab.GroupOption
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        GroupOption.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.tags != null && message.hasOwnProperty("tags")) {
                if (!Array.isArray(message.tags))
                    return "tags: array expected";
                for (var i = 0; i < message.tags.length; ++i)
                    if (!$util.isString(message.tags[i]))
                        return "tags: string[] expected";
            }
            return null;
        };

        /**
         * Creates a GroupOption message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof tab.GroupOption
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {tab.GroupOption} GroupOption
         */
        GroupOption.fromObject = function fromObject(object) {
            if (object instanceof $root.tab.GroupOption)
                return object;
            var message = new $root.tab.GroupOption();
            if (object.tags) {
                if (!Array.isArray(object.tags))
                    throw TypeError(".tab.GroupOption.tags: array expected");
                message.tags = [];
                for (var i = 0; i < object.tags.length; ++i)
                    message.tags[i] = String(object.tags[i]);
            }
            return message;
        };

        /**
         * Creates a plain object from a GroupOption message. Also converts values to other types if specified.
         * @function toObject
         * @memberof tab.GroupOption
         * @static
         * @param {tab.GroupOption} message GroupOption
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        GroupOption.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.tags = [];
            if (message.tags && message.tags.length) {
                object.tags = [];
                for (var j = 0; j < message.tags.length; ++j)
                    object.tags[j] = message.tags[j];
            }
            return object;
        };

        /**
         * Converts this GroupOption to JSON.
         * @function toJSON
         * @memberof tab.GroupOption
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        GroupOption.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for GroupOption
         * @function getTypeUrl
         * @memberof tab.GroupOption
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        GroupOption.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/tab.GroupOption";
        };

        return GroupOption;
    })();

    tab.Group = (function() {

        /**
         * Properties of a Group.
         * @memberof tab
         * @interface IGroup
         * @property {string|null} [name] Group name
         * @property {string|null} [uid] Group uid
         * @property {number|null} [index] Group index
         * @property {tab.IGroupOption|null} [option] Group option
         * @property {Array.<tab.ITab>|null} [tabs] Group tabs
         */

        /**
         * Constructs a new Group.
         * @memberof tab
         * @classdesc Represents a Group.
         * @implements IGroup
         * @constructor
         * @param {tab.IGroup=} [properties] Properties to set
         */
        function Group(properties) {
            this.tabs = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Group name.
         * @member {string} name
         * @memberof tab.Group
         * @instance
         */
        Group.prototype.name = "";

        /**
         * Group uid.
         * @member {string} uid
         * @memberof tab.Group
         * @instance
         */
        Group.prototype.uid = "";

        /**
         * Group index.
         * @member {number} index
         * @memberof tab.Group
         * @instance
         */
        Group.prototype.index = 0;

        /**
         * Group option.
         * @member {tab.IGroupOption|null|undefined} option
         * @memberof tab.Group
         * @instance
         */
        Group.prototype.option = null;

        /**
         * Group tabs.
         * @member {Array.<tab.ITab>} tabs
         * @memberof tab.Group
         * @instance
         */
        Group.prototype.tabs = $util.emptyArray;

        /**
         * Creates a new Group instance using the specified properties.
         * @function create
         * @memberof tab.Group
         * @static
         * @param {tab.IGroup=} [properties] Properties to set
         * @returns {tab.Group} Group instance
         */
        Group.create = function create(properties) {
            return new Group(properties);
        };

        /**
         * Encodes the specified Group message. Does not implicitly {@link tab.Group.verify|verify} messages.
         * @function encode
         * @memberof tab.Group
         * @static
         * @param {tab.IGroup} message Group message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Group.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
            if (message.uid != null && Object.hasOwnProperty.call(message, "uid"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.uid);
            if (message.index != null && Object.hasOwnProperty.call(message, "index"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.index);
            if (message.option != null && Object.hasOwnProperty.call(message, "option"))
                $root.tab.GroupOption.encode(message.option, writer.uint32(/* id 4, wireType 2 =*/34).fork()).ldelim();
            if (message.tabs != null && message.tabs.length)
                for (var i = 0; i < message.tabs.length; ++i)
                    $root.tab.Tab.encode(message.tabs[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Group message, length delimited. Does not implicitly {@link tab.Group.verify|verify} messages.
         * @function encodeDelimited
         * @memberof tab.Group
         * @static
         * @param {tab.IGroup} message Group message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Group.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Group message from the specified reader or buffer.
         * @function decode
         * @memberof tab.Group
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {tab.Group} Group
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Group.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.tab.Group();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.name = reader.string();
                        break;
                    }
                case 2: {
                        message.uid = reader.string();
                        break;
                    }
                case 3: {
                        message.index = reader.int32();
                        break;
                    }
                case 4: {
                        message.option = $root.tab.GroupOption.decode(reader, reader.uint32());
                        break;
                    }
                case 5: {
                        if (!(message.tabs && message.tabs.length))
                            message.tabs = [];
                        message.tabs.push($root.tab.Tab.decode(reader, reader.uint32()));
                        break;
                    }
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Group message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof tab.Group
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {tab.Group} Group
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Group.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Group message.
         * @function verify
         * @memberof tab.Group
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Group.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.uid != null && message.hasOwnProperty("uid"))
                if (!$util.isString(message.uid))
                    return "uid: string expected";
            if (message.index != null && message.hasOwnProperty("index"))
                if (!$util.isInteger(message.index))
                    return "index: integer expected";
            if (message.option != null && message.hasOwnProperty("option")) {
                var error = $root.tab.GroupOption.verify(message.option);
                if (error)
                    return "option." + error;
            }
            if (message.tabs != null && message.hasOwnProperty("tabs")) {
                if (!Array.isArray(message.tabs))
                    return "tabs: array expected";
                for (var i = 0; i < message.tabs.length; ++i) {
                    var error = $root.tab.Tab.verify(message.tabs[i]);
                    if (error)
                        return "tabs." + error;
                }
            }
            return null;
        };

        /**
         * Creates a Group message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof tab.Group
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {tab.Group} Group
         */
        Group.fromObject = function fromObject(object) {
            if (object instanceof $root.tab.Group)
                return object;
            var message = new $root.tab.Group();
            if (object.name != null)
                message.name = String(object.name);
            if (object.uid != null)
                message.uid = String(object.uid);
            if (object.index != null)
                message.index = object.index | 0;
            if (object.option != null) {
                if (typeof object.option !== "object")
                    throw TypeError(".tab.Group.option: object expected");
                message.option = $root.tab.GroupOption.fromObject(object.option);
            }
            if (object.tabs) {
                if (!Array.isArray(object.tabs))
                    throw TypeError(".tab.Group.tabs: array expected");
                message.tabs = [];
                for (var i = 0; i < object.tabs.length; ++i) {
                    if (typeof object.tabs[i] !== "object")
                        throw TypeError(".tab.Group.tabs: object expected");
                    message.tabs[i] = $root.tab.Tab.fromObject(object.tabs[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a Group message. Also converts values to other types if specified.
         * @function toObject
         * @memberof tab.Group
         * @static
         * @param {tab.Group} message Group
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Group.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.tabs = [];
            if (options.defaults) {
                object.name = "";
                object.uid = "";
                object.index = 0;
                object.option = null;
            }
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.uid != null && message.hasOwnProperty("uid"))
                object.uid = message.uid;
            if (message.index != null && message.hasOwnProperty("index"))
                object.index = message.index;
            if (message.option != null && message.hasOwnProperty("option"))
                object.option = $root.tab.GroupOption.toObject(message.option, options);
            if (message.tabs && message.tabs.length) {
                object.tabs = [];
                for (var j = 0; j < message.tabs.length; ++j)
                    object.tabs[j] = $root.tab.Tab.toObject(message.tabs[j], options);
            }
            return object;
        };

        /**
         * Converts this Group to JSON.
         * @function toJSON
         * @memberof tab.Group
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Group.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Group
         * @function getTypeUrl
         * @memberof tab.Group
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Group.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/tab.Group";
        };

        return Group;
    })();

    return tab;
})();

module.exports = $root;
