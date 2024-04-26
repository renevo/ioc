# Inversion Of Control Container

[![GoDoc](https://godoc.org/github.com/renevo/ioc?status.svg)](https://godoc.org/github.com/renevo/ioc)
[![Go Report Card](https://goreportcard.com/badge/github.com/renevo/ioc)](https://goreportcard.com/report/github.com/renevo/ioc)
[![Test](https://github.com/renevo/ioc/actions/workflows/test.yml/badge.svg)](https://github.com/renevo/ioc/actions/workflows/test.yml)

This package provides a simple utility container for faciliting Inversion of Control software pattern. For clarity, this is not a dependency injection framework, merely a go routine safe container to share interfaces across an application.

## Tradeoffs

Go has generics, but they aren't as extensive as other programming languages. Because of that, the `ioc.Container` struct works directly with `reflect.Type` and `reflect.Value`. However, there are two access patterns to work with the underlying container with type safety. These are technically optional, but make for a much better developer experience.

## Global Access

While it is not ideal to use a global, a global option is provided with static functions for interacting with a singleton `ioc.Container` (see example).

This package contains no `init()` functions.

## Context Access

Functions to interact with the container from `context.Context` are provided so that the container can be infered from a parent context. (see example).

## Custom Access

A generic container is provided to allow for access to the underlying `ioc.Container` in a type safe way. You can either use the `ioc.GenericContainer` directly, or implement the functions from within the generic container in your own custom implementation.
