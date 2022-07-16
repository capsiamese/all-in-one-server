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

    tab.Client = (function() {

        /**
         * Properties of a Client.
         * @memberof tab
         * @interface IClient
         * @property {string|null} [name] Client name
         * @property {string|null} [extId] Client extId
         * @property {string|null} [uid] Client uid
         * @property {number|Long|null} [lastAccessTime] Client lastAccessTime
         * @property {Array.<tab.IGroup>|null} [groups] Client groups
         */

        /**
         * Constructs a new Client.
         * @memberof tab
         * @classdesc Represents a Client.
         * @implements IClient
         * @constructor
         * @param {tab.IClient=} [properties] Properties to set
         */
        function Client(properties) {
            this.groups = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Client name.
         * @member {string} name
         * @memberof tab.Client
         * @instance
         */
        Client.prototype.name = "";

        /**
         * Client extId.
         * @member {string} extId
         * @memberof tab.Client
         * @instance
         */
        Client.prototype.extId = "";

        /**
         * Client uid.
         * @member {string} uid
         * @memberof tab.Client
         * @instance
         */
        Client.prototype.uid = "";

        /**
         * Client lastAccessTime.
         * @member {number|Long} lastAccessTime
         * @memberof tab.Client
         * @instance
         */
        Client.prototype.lastAccessTime = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Client groups.
         * @member {Array.<tab.IGroup>} groups
         * @memberof tab.Client
         * @instance
         */
        Client.prototype.groups = $util.emptyArray;

        /**
         * Creates a new Client instance using the specified properties.
         * @function create
         * @memberof tab.Client
         * @static
         * @param {tab.IClient=} [properties] Properties to set
         * @returns {tab.Client} Client instance
         */
        Client.create = function create(properties) {
            return new Client(properties);
        };

        /**
         * Encodes the specified Client message. Does not implicitly {@link tab.Client.verify|verify} messages.
         * @function encode
         * @memberof tab.Client
         * @static
         * @param {tab.IClient} message Client message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Client.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.name != null && Object.hasOwnProperty.call(message, "name"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
            if (message.extId != null && Object.hasOwnProperty.call(message, "extId"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.extId);
            if (message.uid != null && Object.hasOwnProperty.call(message, "uid"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.uid);
            if (message.lastAccessTime != null && Object.hasOwnProperty.call(message, "lastAccessTime"))
                writer.uint32(/* id 4, wireType 0 =*/32).int64(message.lastAccessTime);
            if (message.groups != null && message.groups.length)
                for (var i = 0; i < message.groups.length; ++i)
                    $root.tab.Group.encode(message.groups[i], writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Client message, length delimited. Does not implicitly {@link tab.Client.verify|verify} messages.
         * @function encodeDelimited
         * @memberof tab.Client
         * @static
         * @param {tab.IClient} message Client message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Client.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Client message from the specified reader or buffer.
         * @function decode
         * @memberof tab.Client
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {tab.Client} Client
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Client.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.tab.Client();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.name = reader.string();
                        break;
                    }
                case 2: {
                        message.extId = reader.string();
                        break;
                    }
                case 3: {
                        message.uid = reader.string();
                        break;
                    }
                case 4: {
                        message.lastAccessTime = reader.int64();
                        break;
                    }
                case 5: {
                        if (!(message.groups && message.groups.length))
                            message.groups = [];
                        message.groups.push($root.tab.Group.decode(reader, reader.uint32()));
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
         * Decodes a Client message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof tab.Client
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {tab.Client} Client
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Client.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Client message.
         * @function verify
         * @memberof tab.Client
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Client.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.name != null && message.hasOwnProperty("name"))
                if (!$util.isString(message.name))
                    return "name: string expected";
            if (message.extId != null && message.hasOwnProperty("extId"))
                if (!$util.isString(message.extId))
                    return "extId: string expected";
            if (message.uid != null && message.hasOwnProperty("uid"))
                if (!$util.isString(message.uid))
                    return "uid: string expected";
            if (message.lastAccessTime != null && message.hasOwnProperty("lastAccessTime"))
                if (!$util.isInteger(message.lastAccessTime) && !(message.lastAccessTime && $util.isInteger(message.lastAccessTime.low) && $util.isInteger(message.lastAccessTime.high)))
                    return "lastAccessTime: integer|Long expected";
            if (message.groups != null && message.hasOwnProperty("groups")) {
                if (!Array.isArray(message.groups))
                    return "groups: array expected";
                for (var i = 0; i < message.groups.length; ++i) {
                    var error = $root.tab.Group.verify(message.groups[i]);
                    if (error)
                        return "groups." + error;
                }
            }
            return null;
        };

        /**
         * Creates a Client message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof tab.Client
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {tab.Client} Client
         */
        Client.fromObject = function fromObject(object) {
            if (object instanceof $root.tab.Client)
                return object;
            var message = new $root.tab.Client();
            if (object.name != null)
                message.name = String(object.name);
            if (object.extId != null)
                message.extId = String(object.extId);
            if (object.uid != null)
                message.uid = String(object.uid);
            if (object.lastAccessTime != null)
                if ($util.Long)
                    (message.lastAccessTime = $util.Long.fromValue(object.lastAccessTime)).unsigned = false;
                else if (typeof object.lastAccessTime === "string")
                    message.lastAccessTime = parseInt(object.lastAccessTime, 10);
                else if (typeof object.lastAccessTime === "number")
                    message.lastAccessTime = object.lastAccessTime;
                else if (typeof object.lastAccessTime === "object")
                    message.lastAccessTime = new $util.LongBits(object.lastAccessTime.low >>> 0, object.lastAccessTime.high >>> 0).toNumber();
            if (object.groups) {
                if (!Array.isArray(object.groups))
                    throw TypeError(".tab.Client.groups: array expected");
                message.groups = [];
                for (var i = 0; i < object.groups.length; ++i) {
                    if (typeof object.groups[i] !== "object")
                        throw TypeError(".tab.Client.groups: object expected");
                    message.groups[i] = $root.tab.Group.fromObject(object.groups[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a Client message. Also converts values to other types if specified.
         * @function toObject
         * @memberof tab.Client
         * @static
         * @param {tab.Client} message Client
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Client.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.groups = [];
            if (options.defaults) {
                object.name = "";
                object.extId = "";
                object.uid = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.lastAccessTime = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.lastAccessTime = options.longs === String ? "0" : 0;
            }
            if (message.name != null && message.hasOwnProperty("name"))
                object.name = message.name;
            if (message.extId != null && message.hasOwnProperty("extId"))
                object.extId = message.extId;
            if (message.uid != null && message.hasOwnProperty("uid"))
                object.uid = message.uid;
            if (message.lastAccessTime != null && message.hasOwnProperty("lastAccessTime"))
                if (typeof message.lastAccessTime === "number")
                    object.lastAccessTime = options.longs === String ? String(message.lastAccessTime) : message.lastAccessTime;
                else
                    object.lastAccessTime = options.longs === String ? $util.Long.prototype.toString.call(message.lastAccessTime) : options.longs === Number ? new $util.LongBits(message.lastAccessTime.low >>> 0, message.lastAccessTime.high >>> 0).toNumber() : message.lastAccessTime;
            if (message.groups && message.groups.length) {
                object.groups = [];
                for (var j = 0; j < message.groups.length; ++j)
                    object.groups[j] = $root.tab.Group.toObject(message.groups[j], options);
            }
            return object;
        };

        /**
         * Converts this Client to JSON.
         * @function toJSON
         * @memberof tab.Client
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Client.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Client
         * @function getTypeUrl
         * @memberof tab.Client
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Client.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/tab.Client";
        };

        return Client;
    })();

    tab.BarkHistory = (function() {

        /**
         * Properties of a BarkHistory.
         * @memberof tab
         * @interface IBarkHistory
         * @property {number|Long|null} [id] BarkHistory id
         * @property {number|Long|null} [ts] BarkHistory ts
         * @property {string|null} [from] BarkHistory from
         * @property {string|null} [key] BarkHistory key
         * @property {Object.<string,string>|null} [data] BarkHistory data
         */

        /**
         * Constructs a new BarkHistory.
         * @memberof tab
         * @classdesc Represents a BarkHistory.
         * @implements IBarkHistory
         * @constructor
         * @param {tab.IBarkHistory=} [properties] Properties to set
         */
        function BarkHistory(properties) {
            this.data = {};
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * BarkHistory id.
         * @member {number|Long} id
         * @memberof tab.BarkHistory
         * @instance
         */
        BarkHistory.prototype.id = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * BarkHistory ts.
         * @member {number|Long} ts
         * @memberof tab.BarkHistory
         * @instance
         */
        BarkHistory.prototype.ts = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * BarkHistory from.
         * @member {string} from
         * @memberof tab.BarkHistory
         * @instance
         */
        BarkHistory.prototype.from = "";

        /**
         * BarkHistory key.
         * @member {string} key
         * @memberof tab.BarkHistory
         * @instance
         */
        BarkHistory.prototype.key = "";

        /**
         * BarkHistory data.
         * @member {Object.<string,string>} data
         * @memberof tab.BarkHistory
         * @instance
         */
        BarkHistory.prototype.data = $util.emptyObject;

        /**
         * Creates a new BarkHistory instance using the specified properties.
         * @function create
         * @memberof tab.BarkHistory
         * @static
         * @param {tab.IBarkHistory=} [properties] Properties to set
         * @returns {tab.BarkHistory} BarkHistory instance
         */
        BarkHistory.create = function create(properties) {
            return new BarkHistory(properties);
        };

        /**
         * Encodes the specified BarkHistory message. Does not implicitly {@link tab.BarkHistory.verify|verify} messages.
         * @function encode
         * @memberof tab.BarkHistory
         * @static
         * @param {tab.IBarkHistory} message BarkHistory message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        BarkHistory.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.id != null && Object.hasOwnProperty.call(message, "id"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.id);
            if (message.ts != null && Object.hasOwnProperty.call(message, "ts"))
                writer.uint32(/* id 2, wireType 0 =*/16).int64(message.ts);
            if (message.from != null && Object.hasOwnProperty.call(message, "from"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.from);
            if (message.key != null && Object.hasOwnProperty.call(message, "key"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.key);
            if (message.data != null && Object.hasOwnProperty.call(message, "data"))
                for (var keys = Object.keys(message.data), i = 0; i < keys.length; ++i)
                    writer.uint32(/* id 5, wireType 2 =*/42).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]).uint32(/* id 2, wireType 2 =*/18).string(message.data[keys[i]]).ldelim();
            return writer;
        };

        /**
         * Encodes the specified BarkHistory message, length delimited. Does not implicitly {@link tab.BarkHistory.verify|verify} messages.
         * @function encodeDelimited
         * @memberof tab.BarkHistory
         * @static
         * @param {tab.IBarkHistory} message BarkHistory message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        BarkHistory.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a BarkHistory message from the specified reader or buffer.
         * @function decode
         * @memberof tab.BarkHistory
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {tab.BarkHistory} BarkHistory
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        BarkHistory.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.tab.BarkHistory(), key, value;
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.id = reader.int64();
                        break;
                    }
                case 2: {
                        message.ts = reader.int64();
                        break;
                    }
                case 3: {
                        message.from = reader.string();
                        break;
                    }
                case 4: {
                        message.key = reader.string();
                        break;
                    }
                case 5: {
                        if (message.data === $util.emptyObject)
                            message.data = {};
                        var end2 = reader.uint32() + reader.pos;
                        key = "";
                        value = "";
                        while (reader.pos < end2) {
                            var tag2 = reader.uint32();
                            switch (tag2 >>> 3) {
                            case 1:
                                key = reader.string();
                                break;
                            case 2:
                                value = reader.string();
                                break;
                            default:
                                reader.skipType(tag2 & 7);
                                break;
                            }
                        }
                        message.data[key] = value;
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
         * Decodes a BarkHistory message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof tab.BarkHistory
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {tab.BarkHistory} BarkHistory
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        BarkHistory.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a BarkHistory message.
         * @function verify
         * @memberof tab.BarkHistory
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        BarkHistory.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.id != null && message.hasOwnProperty("id"))
                if (!$util.isInteger(message.id) && !(message.id && $util.isInteger(message.id.low) && $util.isInteger(message.id.high)))
                    return "id: integer|Long expected";
            if (message.ts != null && message.hasOwnProperty("ts"))
                if (!$util.isInteger(message.ts) && !(message.ts && $util.isInteger(message.ts.low) && $util.isInteger(message.ts.high)))
                    return "ts: integer|Long expected";
            if (message.from != null && message.hasOwnProperty("from"))
                if (!$util.isString(message.from))
                    return "from: string expected";
            if (message.key != null && message.hasOwnProperty("key"))
                if (!$util.isString(message.key))
                    return "key: string expected";
            if (message.data != null && message.hasOwnProperty("data")) {
                if (!$util.isObject(message.data))
                    return "data: object expected";
                var key = Object.keys(message.data);
                for (var i = 0; i < key.length; ++i)
                    if (!$util.isString(message.data[key[i]]))
                        return "data: string{k:string} expected";
            }
            return null;
        };

        /**
         * Creates a BarkHistory message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof tab.BarkHistory
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {tab.BarkHistory} BarkHistory
         */
        BarkHistory.fromObject = function fromObject(object) {
            if (object instanceof $root.tab.BarkHistory)
                return object;
            var message = new $root.tab.BarkHistory();
            if (object.id != null)
                if ($util.Long)
                    (message.id = $util.Long.fromValue(object.id)).unsigned = false;
                else if (typeof object.id === "string")
                    message.id = parseInt(object.id, 10);
                else if (typeof object.id === "number")
                    message.id = object.id;
                else if (typeof object.id === "object")
                    message.id = new $util.LongBits(object.id.low >>> 0, object.id.high >>> 0).toNumber();
            if (object.ts != null)
                if ($util.Long)
                    (message.ts = $util.Long.fromValue(object.ts)).unsigned = false;
                else if (typeof object.ts === "string")
                    message.ts = parseInt(object.ts, 10);
                else if (typeof object.ts === "number")
                    message.ts = object.ts;
                else if (typeof object.ts === "object")
                    message.ts = new $util.LongBits(object.ts.low >>> 0, object.ts.high >>> 0).toNumber();
            if (object.from != null)
                message.from = String(object.from);
            if (object.key != null)
                message.key = String(object.key);
            if (object.data) {
                if (typeof object.data !== "object")
                    throw TypeError(".tab.BarkHistory.data: object expected");
                message.data = {};
                for (var keys = Object.keys(object.data), i = 0; i < keys.length; ++i)
                    message.data[keys[i]] = String(object.data[keys[i]]);
            }
            return message;
        };

        /**
         * Creates a plain object from a BarkHistory message. Also converts values to other types if specified.
         * @function toObject
         * @memberof tab.BarkHistory
         * @static
         * @param {tab.BarkHistory} message BarkHistory
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        BarkHistory.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.objects || options.defaults)
                object.data = {};
            if (options.defaults) {
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.id = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.id = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.ts = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.ts = options.longs === String ? "0" : 0;
                object.from = "";
                object.key = "";
            }
            if (message.id != null && message.hasOwnProperty("id"))
                if (typeof message.id === "number")
                    object.id = options.longs === String ? String(message.id) : message.id;
                else
                    object.id = options.longs === String ? $util.Long.prototype.toString.call(message.id) : options.longs === Number ? new $util.LongBits(message.id.low >>> 0, message.id.high >>> 0).toNumber() : message.id;
            if (message.ts != null && message.hasOwnProperty("ts"))
                if (typeof message.ts === "number")
                    object.ts = options.longs === String ? String(message.ts) : message.ts;
                else
                    object.ts = options.longs === String ? $util.Long.prototype.toString.call(message.ts) : options.longs === Number ? new $util.LongBits(message.ts.low >>> 0, message.ts.high >>> 0).toNumber() : message.ts;
            if (message.from != null && message.hasOwnProperty("from"))
                object.from = message.from;
            if (message.key != null && message.hasOwnProperty("key"))
                object.key = message.key;
            var keys2;
            if (message.data && (keys2 = Object.keys(message.data)).length) {
                object.data = {};
                for (var j = 0; j < keys2.length; ++j)
                    object.data[keys2[j]] = message.data[keys2[j]];
            }
            return object;
        };

        /**
         * Converts this BarkHistory to JSON.
         * @function toJSON
         * @memberof tab.BarkHistory
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        BarkHistory.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for BarkHistory
         * @function getTypeUrl
         * @memberof tab.BarkHistory
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        BarkHistory.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/tab.BarkHistory";
        };

        return BarkHistory;
    })();

    return tab;
})();

module.exports = $root;
