/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
import * as $protobuf from "protobufjs/minimal";

// Common aliases
const $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
const $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

export const protocol = $root.protocol = (() => {

    /**
     * Namespace protocol.
     * @exports protocol
     * @namespace
     */
    const protocol = {};

    protocol.Message = (function() {

        /**
         * Properties of a Message.
         * @memberof protocol
         * @interface IMessage
         * @property {number|Long|null} [type] Message type
         * @property {string|null} [content] Message content
         * @property {number|Long|null} [contentType] Message contentType
         * @property {Uint8Array|null} [file] Message file
         * @property {string|null} [fileSuffix] Message fileSuffix
         */

        /**
         * Constructs a new Message.
         * @memberof protocol
         * @classdesc Represents a Message.
         * @implements IMessage
         * @constructor
         * @param {protocol.IMessage=} [properties] Properties to set
         */
        function Message(properties) {
            if (properties)
                for (let keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Message type.
         * @member {number|Long} type
         * @memberof protocol.Message
         * @instance
         */
        Message.prototype.type = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Message content.
         * @member {string} content
         * @memberof protocol.Message
         * @instance
         */
        Message.prototype.content = "";

        /**
         * Message contentType.
         * @member {number|Long} contentType
         * @memberof protocol.Message
         * @instance
         */
        Message.prototype.contentType = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * Message file.
         * @member {Uint8Array} file
         * @memberof protocol.Message
         * @instance
         */
        Message.prototype.file = $util.newBuffer([]);

        /**
         * Message fileSuffix.
         * @member {string} fileSuffix
         * @memberof protocol.Message
         * @instance
         */
        Message.prototype.fileSuffix = "";

        /**
         * Creates a new Message instance using the specified properties.
         * @function create
         * @memberof protocol.Message
         * @static
         * @param {protocol.IMessage=} [properties] Properties to set
         * @returns {protocol.Message} Message instance
         */
        Message.create = function create(properties) {
            return new Message(properties);
        };

        /**
         * Encodes the specified Message message. Does not implicitly {@link protocol.Message.verify|verify} messages.
         * @function encode
         * @memberof protocol.Message
         * @static
         * @param {protocol.IMessage} message Message message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Message.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.type != null && Object.hasOwnProperty.call(message, "type"))
                writer.uint32(/* id 1, wireType 0 =*/8).int64(message.type);
            if (message.content != null && Object.hasOwnProperty.call(message, "content"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.content);
            if (message.contentType != null && Object.hasOwnProperty.call(message, "contentType"))
                writer.uint32(/* id 3, wireType 0 =*/24).int64(message.contentType);
            if (message.file != null && Object.hasOwnProperty.call(message, "file"))
                writer.uint32(/* id 4, wireType 2 =*/34).bytes(message.file);
            if (message.fileSuffix != null && Object.hasOwnProperty.call(message, "fileSuffix"))
                writer.uint32(/* id 5, wireType 2 =*/42).string(message.fileSuffix);
            return writer;
        };

        /**
         * Encodes the specified Message message, length delimited. Does not implicitly {@link protocol.Message.verify|verify} messages.
         * @function encodeDelimited
         * @memberof protocol.Message
         * @static
         * @param {protocol.IMessage} message Message message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Message.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Message message from the specified reader or buffer.
         * @function decode
         * @memberof protocol.Message
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {protocol.Message} Message
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Message.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            let end = length === undefined ? reader.len : reader.pos + length, message = new $root.protocol.Message();
            while (reader.pos < end) {
                let tag = reader.uint32();
                switch (tag >>> 3) {
                case 1: {
                        message.type = reader.int64();
                        break;
                    }
                case 2: {
                        message.content = reader.string();
                        break;
                    }
                case 3: {
                        message.contentType = reader.int64();
                        break;
                    }
                case 4: {
                        message.file = reader.bytes();
                        break;
                    }
                case 5: {
                        message.fileSuffix = reader.string();
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
         * Decodes a Message message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof protocol.Message
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {protocol.Message} Message
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Message.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Message message.
         * @function verify
         * @memberof protocol.Message
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Message.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.type != null && message.hasOwnProperty("type"))
                if (!$util.isInteger(message.type) && !(message.type && $util.isInteger(message.type.low) && $util.isInteger(message.type.high)))
                    return "type: integer|Long expected";
            if (message.content != null && message.hasOwnProperty("content"))
                if (!$util.isString(message.content))
                    return "content: string expected";
            if (message.contentType != null && message.hasOwnProperty("contentType"))
                if (!$util.isInteger(message.contentType) && !(message.contentType && $util.isInteger(message.contentType.low) && $util.isInteger(message.contentType.high)))
                    return "contentType: integer|Long expected";
            if (message.file != null && message.hasOwnProperty("file"))
                if (!(message.file && typeof message.file.length === "number" || $util.isString(message.file)))
                    return "file: buffer expected";
            if (message.fileSuffix != null && message.hasOwnProperty("fileSuffix"))
                if (!$util.isString(message.fileSuffix))
                    return "fileSuffix: string expected";
            return null;
        };

        /**
         * Creates a Message message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof protocol.Message
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {protocol.Message} Message
         */
        Message.fromObject = function fromObject(object) {
            if (object instanceof $root.protocol.Message)
                return object;
            let message = new $root.protocol.Message();
            if (object.type != null)
                if ($util.Long)
                    (message.type = $util.Long.fromValue(object.type)).unsigned = false;
                else if (typeof object.type === "string")
                    message.type = parseInt(object.type, 10);
                else if (typeof object.type === "number")
                    message.type = object.type;
                else if (typeof object.type === "object")
                    message.type = new $util.LongBits(object.type.low >>> 0, object.type.high >>> 0).toNumber();
            if (object.content != null)
                message.content = String(object.content);
            if (object.contentType != null)
                if ($util.Long)
                    (message.contentType = $util.Long.fromValue(object.contentType)).unsigned = false;
                else if (typeof object.contentType === "string")
                    message.contentType = parseInt(object.contentType, 10);
                else if (typeof object.contentType === "number")
                    message.contentType = object.contentType;
                else if (typeof object.contentType === "object")
                    message.contentType = new $util.LongBits(object.contentType.low >>> 0, object.contentType.high >>> 0).toNumber();
            if (object.file != null)
                if (typeof object.file === "string")
                    $util.base64.decode(object.file, message.file = $util.newBuffer($util.base64.length(object.file)), 0);
                else if (object.file.length >= 0)
                    message.file = object.file;
            if (object.fileSuffix != null)
                message.fileSuffix = String(object.fileSuffix);
            return message;
        };

        /**
         * Creates a plain object from a Message message. Also converts values to other types if specified.
         * @function toObject
         * @memberof protocol.Message
         * @static
         * @param {protocol.Message} message Message
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Message.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            let object = {};
            if (options.defaults) {
                if ($util.Long) {
                    let long = new $util.Long(0, 0, false);
                    object.type = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.type = options.longs === String ? "0" : 0;
                object.content = "";
                if ($util.Long) {
                    let long = new $util.Long(0, 0, false);
                    object.contentType = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.contentType = options.longs === String ? "0" : 0;
                if (options.bytes === String)
                    object.file = "";
                else {
                    object.file = [];
                    if (options.bytes !== Array)
                        object.file = $util.newBuffer(object.file);
                }
                object.fileSuffix = "";
            }
            if (message.type != null && message.hasOwnProperty("type"))
                if (typeof message.type === "number")
                    object.type = options.longs === String ? String(message.type) : message.type;
                else
                    object.type = options.longs === String ? $util.Long.prototype.toString.call(message.type) : options.longs === Number ? new $util.LongBits(message.type.low >>> 0, message.type.high >>> 0).toNumber() : message.type;
            if (message.content != null && message.hasOwnProperty("content"))
                object.content = message.content;
            if (message.contentType != null && message.hasOwnProperty("contentType"))
                if (typeof message.contentType === "number")
                    object.contentType = options.longs === String ? String(message.contentType) : message.contentType;
                else
                    object.contentType = options.longs === String ? $util.Long.prototype.toString.call(message.contentType) : options.longs === Number ? new $util.LongBits(message.contentType.low >>> 0, message.contentType.high >>> 0).toNumber() : message.contentType;
            if (message.file != null && message.hasOwnProperty("file"))
                object.file = options.bytes === String ? $util.base64.encode(message.file, 0, message.file.length) : options.bytes === Array ? Array.prototype.slice.call(message.file) : message.file;
            if (message.fileSuffix != null && message.hasOwnProperty("fileSuffix"))
                object.fileSuffix = message.fileSuffix;
            return object;
        };

        /**
         * Converts this Message to JSON.
         * @function toJSON
         * @memberof protocol.Message
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Message.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Gets the default type url for Message
         * @function getTypeUrl
         * @memberof protocol.Message
         * @static
         * @param {string} [typeUrlPrefix] your custom typeUrlPrefix(default "type.googleapis.com")
         * @returns {string} The default type url
         */
        Message.getTypeUrl = function getTypeUrl(typeUrlPrefix) {
            if (typeUrlPrefix === undefined) {
                typeUrlPrefix = "type.googleapis.com";
            }
            return typeUrlPrefix + "/protocol.Message";
        };

        return Message;
    })();

    return protocol;
})();

export { $root as default };
