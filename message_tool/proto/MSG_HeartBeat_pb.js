/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.MSG_HeartBeat.CS_HeartBeat_Req', null, global);
goog.exportSymbol('proto.MSG_HeartBeat.ErrorCode', null, global);
goog.exportSymbol('proto.MSG_HeartBeat.SC_HeartBeat_Rsp', null, global);
goog.exportSymbol('proto.MSG_HeartBeat.SUBMSG', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.MSG_HeartBeat.CS_HeartBeat_Req, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.MSG_HeartBeat.CS_HeartBeat_Req.displayName = 'proto.MSG_HeartBeat.CS_HeartBeat_Req';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.prototype.toObject = function(opt_includeInstance) {
  return proto.MSG_HeartBeat.CS_HeartBeat_Req.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.MSG_HeartBeat.CS_HeartBeat_Req} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.toObject = function(includeInstance, msg) {
  var f, obj = {
    svrpoint: jspb.Message.getFieldWithDefault(msg, 1, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.MSG_HeartBeat.CS_HeartBeat_Req}
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.MSG_HeartBeat.CS_HeartBeat_Req;
  return proto.MSG_HeartBeat.CS_HeartBeat_Req.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.MSG_HeartBeat.CS_HeartBeat_Req} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.MSG_HeartBeat.CS_HeartBeat_Req}
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readUint32());
      msg.setSvrpoint(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.MSG_HeartBeat.CS_HeartBeat_Req.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.MSG_HeartBeat.CS_HeartBeat_Req} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSvrpoint();
  if (f !== 0) {
    writer.writeUint32(
      1,
      f
    );
  }
};


/**
 * optional uint32 SvrPoint = 1;
 * @return {number}
 */
proto.MSG_HeartBeat.CS_HeartBeat_Req.prototype.getSvrpoint = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.MSG_HeartBeat.CS_HeartBeat_Req.prototype.setSvrpoint = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.MSG_HeartBeat.SC_HeartBeat_Rsp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.MSG_HeartBeat.SC_HeartBeat_Rsp.displayName = 'proto.MSG_HeartBeat.SC_HeartBeat_Rsp';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp.prototype.toObject = function(opt_includeInstance) {
  return proto.MSG_HeartBeat.SC_HeartBeat_Rsp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.MSG_HeartBeat.SC_HeartBeat_Rsp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.MSG_HeartBeat.SC_HeartBeat_Rsp}
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.MSG_HeartBeat.SC_HeartBeat_Rsp;
  return proto.MSG_HeartBeat.SC_HeartBeat_Rsp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.MSG_HeartBeat.SC_HeartBeat_Rsp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.MSG_HeartBeat.SC_HeartBeat_Rsp}
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.MSG_HeartBeat.SC_HeartBeat_Rsp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.MSG_HeartBeat.SC_HeartBeat_Rsp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.MSG_HeartBeat.SC_HeartBeat_Rsp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


/**
 * @enum {number}
 */
proto.MSG_HeartBeat.SUBMSG = {
  BEGIN: 0,
  CS_HEARTBEAT: 1,
  SC_HEARTBEAT: 2
};

/**
 * @enum {number}
 */
proto.MSG_HeartBeat.ErrorCode = {
  INVALID: 0,
  SUCCESS: 1,
  FAIL: 2
};

goog.object.extend(exports, proto.MSG_HeartBeat);
