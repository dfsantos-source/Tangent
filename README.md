# Tangent

A backend project, currently in progress.

The current architecture is separated into subpackages. The root package holds all domain types, interfaces, and structs. The subpackages, such as /http and /postgres contain the actual implementation of the domain types and serve as tools to be used throughout the application. The /cmd directory is what ties these subpackages together, using dependency injection and best design principles. 