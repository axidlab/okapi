# Generated by the protocol buffer compiler.  DO NOT EDIT!
# sources: okapi/hashing/v1/hashing.proto
# plugin: python-betterproto
# This file has been @generated
from dataclasses import dataclass

import betterproto


@dataclass(eq=False, repr=False)
class Blake3HashRequest(betterproto.Message):
    data: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class Blake3HashResponse(betterproto.Message):
    digest: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class Blake3KeyedHashRequest(betterproto.Message):
    data: bytes = betterproto.bytes_field(1)
    key: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class Blake3KeyedHashResponse(betterproto.Message):
    digest: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class Blake3DeriveKeyRequest(betterproto.Message):
    context: bytes = betterproto.bytes_field(1)
    key_material: bytes = betterproto.bytes_field(2)


@dataclass(eq=False, repr=False)
class Blake3DeriveKeyResponse(betterproto.Message):
    digest: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class Sha256HashRequest(betterproto.Message):
    data: bytes = betterproto.bytes_field(1)


@dataclass(eq=False, repr=False)
class Sha256HashResponse(betterproto.Message):
    digest: bytes = betterproto.bytes_field(1)
