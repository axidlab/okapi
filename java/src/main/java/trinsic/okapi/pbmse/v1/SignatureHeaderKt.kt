//Generated by the protocol buffer compiler. DO NOT EDIT!
// source: pbmse/v1/pbmse.proto

package trinsic.okapi.pbmse.v1;

@kotlin.jvm.JvmSynthetic
public inline fun signatureHeader(block: trinsic.okapi.pbmse.v1.SignatureHeaderKt.Dsl.() -> kotlin.Unit): trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader =
  trinsic.okapi.pbmse.v1.SignatureHeaderKt.Dsl._create(trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader.newBuilder()).apply { block() }._build()
public object SignatureHeaderKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
      @kotlin.PublishedApi
      internal fun _create(builder: trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
    internal fun _build(): trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader = _builder.build()

    /**
     * <code>string algorithm = 1;</code>
     */
    public var algorithm: kotlin.String
      @JvmName("getAlgorithm")
      get() = _builder.getAlgorithm()
      @JvmName("setAlgorithm")
      set(value) {
        _builder.setAlgorithm(value)
      }
    /**
     * <code>string algorithm = 1;</code>
     */
    public fun clearAlgorithm() {
      _builder.clearAlgorithm()
    }

    /**
     * <code>string key_id = 2;</code>
     */
    public var keyId: kotlin.String
      @JvmName("getKeyId")
      get() = _builder.getKeyId()
      @JvmName("setKeyId")
      set(value) {
        _builder.setKeyId(value)
      }
    /**
     * <code>string key_id = 2;</code>
     */
    public fun clearKeyId() {
      _builder.clearKeyId()
    }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader.copy(block: trinsic.okapi.pbmse.v1.SignatureHeaderKt.Dsl.() -> kotlin.Unit): trinsic.okapi.pbmse.v1.Pbmse.SignatureHeader =
  trinsic.okapi.pbmse.v1.SignatureHeaderKt.Dsl._create(this.toBuilder()).apply { block() }._build()
