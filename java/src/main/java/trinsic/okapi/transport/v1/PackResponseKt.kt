//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: okapi/transport/v1/transport.proto

package trinsic.okapi.transport.v1;

@kotlin.jvm.JvmSynthetic
public inline fun packResponse(block: trinsic.okapi.transport.v1.PackResponseKt.Dsl.() -> kotlin.Unit): trinsic.okapi.transport.v1.Transport.PackResponse =
  trinsic.okapi.transport.v1.PackResponseKt.Dsl._create(trinsic.okapi.transport.v1.Transport.PackResponse.newBuilder()).apply { block() }._build()
public object PackResponseKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: trinsic.okapi.transport.v1.Transport.PackResponse.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: trinsic.okapi.transport.v1.Transport.PackResponse.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): trinsic.okapi.transport.v1.Transport.PackResponse = _builder.build()

    /**
     * <code>.pbmse.v1.EncryptedMessage message = 1;</code>
     */
    public var message: trinsic.okapi.pbmse.v1.Pbmse.EncryptedMessage
      @JvmName("getMessage")
      get() = _builder.getMessage()
      @JvmName("setMessage")
      set(value) {
        _builder.setMessage(value)
      }
    /**
     * <code>.pbmse.v1.EncryptedMessage message = 1;</code>
     */
    public fun clearMessage() {
      _builder.clearMessage()
    }
    /**
     * <code>.pbmse.v1.EncryptedMessage message = 1;</code>
     * @return Whether the message field is set.
     */
    public fun hasMessage(): kotlin.Boolean {
      return _builder.hasMessage()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun trinsic.okapi.transport.v1.Transport.PackResponse.copy(block: trinsic.okapi.transport.v1.PackResponseKt.Dsl.() -> kotlin.Unit): trinsic.okapi.transport.v1.Transport.PackResponse =
  trinsic.okapi.transport.v1.PackResponseKt.Dsl._create(this.toBuilder()).apply { block() }._build()
