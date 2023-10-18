<div align="center">
<p>
    <img width="100" src="https://raw.githubusercontent.com/julelang/resources/master/jule_icon.svg?sanitize=true">
</p>
<h1>The Jule Programming Language</h1>

An effective programming language to build efficient, fast, reliable and safe software.

[Website](https://jule.dev) |
[Manual](https://manual.jule.dev) |
[Contributing](https://jule.dev/contribute) |
[Community](https://jule.dev/community)

[![Build](https://github.com/julelang/jule/actions/workflows/build.yml/badge.svg)](https://github.com/julelang/jule/actions/workflows/build.yml)
[![Tests](https://github.com/julelang/jule/actions/workflows/tests.yml/badge.svg)](https://github.com/julelang/jule/actions/workflows/tests.yml)
[![Jule Community](https://dcbadge.vercel.app/api/server/ReWQgPDnP6?style=flat)](https://discord.gg/ReWQgPDnP6)

</strong>

</div>

## Introduction

This repository is the main source tree of the Jule. \
It contains the reference compiler, API, and standard library.

Jule does not have a stable version yet and is still being developed to become more stable.
Some commits may not be fully honored due to some compiler/API errors.
Please report it with the [Jule Issue Tracker](https://github.com/julelang/jule/issues) if you come across something like this.
You can also [join the Discord community](https://discord.gg/ReWQgPDnP6) to discuss, helping, and ask more questions about Jule with the community.

## Our Mission

The mission is to develop a reliable and safe programming language that focuses on systems programming. While doing this, performance should not be compromised too much. Jule should be a balanced language that keeping the cost to a minimum for safety purposes. Even if it is a new language design, the popular C and C++ languages, which are frequently used in system programming, should not be ignored. Therefore, good interoperability support should be provided for these languages.

To be safe, Jule should adopt approaches that serve the purpose in a balanced way, rather than having a design that contains a lot of complexity in the language. It should be a language that is significantly safe as well as being satisfactorily efficient and performant for developers.

- **Simplicity and Maintainability**: Jule's design should to be such that it avoids imposing constraints that increase time costs of developers, or having to think about the rules they must frequently follow when designing their algorithms. As much as possible, Jule's safety design and experience should be such that it prevents developers from thinking the way Jule wants them to and not the way they do. Design choices for safety or performance must be made without adding significant complexity to the Jule at the same time.

- **Safety**: Jule code must be safe at both compile time and runtime. This requires measures such as bounds checking and nil dereferencing checking at compile time and runtime. Jule also should adopt immutability and disallow shadowing by default as safety measures. Anything that might be unsafe should be blocked or checked. If unsafe behavior is desired, this should be clearly stated and awareness should be made.

- **Efficiency and Performance**: To be efficient, Jule should try to behave in a way that uses the least memory consumption. In addition, it should be high-performance and be able to create fast software. Instead of relying entirely on the backend compiler that will be used when the IR is created, it should have an optimizing reference compiler to generate the best IR.

- **High C/C++ Interoperability**: C and C++ are languages ​​that are popular in system programming, widely used and have a wide ecosystem. Therefore, Jule should embrace and collaborate with the existing ecosystem rather than completely ignoring it. Jule developers should also be able to use and interoperate with C/C++ within Jule if they wish. Therefore, Jule is compiled by creating C++ IR. It supports different compilers such as Clang, GNU GCC. It has an API developed in C++ to make things easier for Jule developers and make things more standardized.

- **Compile Time**: Jule should have a good compile time that serves the mission. For this, it must have capabilities such as having compile-time constants, avoiding repetitive checks by checking generics once for each instance, and performing safety-related checks whenever possible. The compiler should support developers with safety checks at compile-time and give some optimization options such as constant evaluations.

- **Brand-New**: Jule should be like a truly brand-new language. For example, it should have UTF-8 strings, support Unicode and type inference, and a cross-platform portable standard library. Unless the developer really wants to do platform specific programming, shouldn't worry about the portability of the standard library, should have Unicode support and a common standard like UTF-8 should be easy to handle.


## Future of Jule
Jule is in beta versions, changing and evolving rapidly. Our main priority right now is for Jule to become more stable and have a robust standard library. To make it easier for the community and official developers to develop any tools for Jule, a significant portion of Jule's official compiler is included in the standard library. The standard library has important stages such as lexer, parser and semantic analyzer and is suitable for use in tool development.

The syntax and language design of the Jule programming language has emerged and is not expected to undergo major changes. Many parts of reference compiler, are included in the standard library such as a lexer, parser, and semantic analyzer. This will also allow developers to quickly develop tools for the language by leveraging Jule's standard library.

There is also the idea of ​​adding a package manager that ships with the official compiler at some stage. Jule's modern understanding of language and convenience suggest that there should be a package manager that comes with the compiler. This package manager will provide management of non-standard library packages developed and published by the community. Jule's standard library only gets updates with compiler releases.

In addition to the currently supported C++ compilers, the official compiler may get support or improve support for more compilers and development environments in the future. Expected that Jule's platform support will expand further in areas such as backend compiler, CPU architecture, and operating system.

## Community

Contribute and get involved in our community.

Join Julenours to support Jule, explore and interact with the community.\
Our main community platforms:

- [Official Discord Server of The Jule Community](https://discord.gg/CZhK7dyh9X)
- [GitHub Discussions](https://github.com/jule-lang/jule/discussions)

## Compile from Source

If you want to get Jule from the source, there are many ways to do so.
Jule has a bootstrapped compiler, so you'll need to get one first if you don't have one.
There are two options to do this: obtain the release or use IR.
However, it is recommended to use IR as it is always more up to date and ensures there is enough left to compile the [master](https://github.com/julelang/jule/tree/master) branch.
Officially, the recommended method to always get the most up-to-date build of compiler from the latest source code is to use IR.

If you already have a compiler, you can use build scripts designed for developers by obtaining the latest source code.
But remember, these are for developers and they compile the compiler for debugging new source code, not for production use. So you can get an inefficient and slow version.

- Learn about: [compile from IR](https://manual.jule.dev/getting-started/install-from-source/compile-from-ir.html)
- Learn about: [build Scripts](https://manual.jule.dev/getting-started/install-from-source/build-scripts.html)

## Contributing

Any contribution to Jule is greatly appreciated, whether it's a typo fix, a brand new compiler feature, or a bug report.

The Jule project only uses issues for things like proposals, bug reports, and vulnerabilities.
If you want to discuss anything, [discussions](https://github.com/julelang/jule/discussions) is a better place for that.
If you are interested in reporting a security vulnerability, please read the out [security policy](https://github.com/julelang/jule/security/policy) first.

- Please read [Julenour Code of Conduct](https://jule.dev/code-of-conduct) before contributing anything
- To contribute, please read the [contribution guidelines](https://jule.dev/contribute)
- For discussions and asking questions, please use [discussions](https://github.com/julelang/jule/discussions)
- Regarding security, please refer to the [security policy](https://github.com/julelang/jule/security/policy)

## License

The reference compiler, API, and standard library are distributed under the terms of the BSD 3-Clause license. <br>
[See License Details](./LICENSE)

