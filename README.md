## This repo contains 2 packages:
1. fuse 
2. mock


## Goals:
1. Minimal imports - Only one import (fuse) during configuration.
2. Non-intrusive - References to the package fuse should be in just startup code.
3. Small API - just enough to get work done.
4. Minimal footprint/low overhead - no complicated setup code required.

For a full usage example of these 2 packages please refer to repo <a href="https://github.com/rvauradkar1/testfuse">Guide to usage of library fuse</a>

## Package fuse

Features:

**Dependency Injection pattern** - primarily used for stateless components, all components are singletons.
1. Register components.
2. Inject stateless component dependencies.

**Resource Locator pattern** - primarily used for stateful components, all components are prototypes.

1. Register components.
2. A "Finder" function is provided to get a new copy of the component.
3. Inject stateless component dependencies.

**Constraints**:

1. Components can only be registered as pointers.
2. Components dependencies can either be through interfaces or pointers to `struct`s.


## Package mock
mock library generates mock code for all the dependencies of a component.

For a full usage example of these 2 packages please refer to repo <a href="https://github.com/rvauradkar1/testfuse">Guide to usage of library fuse</a>
