# Intro

# ChangeLog

## v0.1.0

- Cleanup unused types

## v0.0.1

- In the begining the code was without form and void and darkness was uppon the face of the terminal.

# Database Schema 

## PostgreSQL

See `./sys/schema/` for SQL scripts.

---

## Translators as a composition.

The `Event Store` will require that their Entities / Aggregates implement a serialization interface, an example can be: 

```
packaged eventstore

type IEventStoreAggregate interface {
   GetID() uuid.UUID
   SetID(uuid.UUID)
   GetVersion() int
   SetVersion(int)
   Serialize(interface{}) []byte
   UnSerialize([]byte) []byte
   GetChanges() []IEventStoreEvent
   LoadHistory([]IEventStoreEvent)
}

type IEventStoreEvent interface {
   GetType() string
   SetType(string)
   Serialize(interface{}) []byte
   UnSerialize([]byte) []byte
}  
```

This may seem that an easy win is to make your domain Entities / Aggregates implement the interfaces and you are free to go, unfortunatly this would mean that the domain aggregates would now how to save themselves on the Repository and this would break the Clean Architecture way.

Our proposed solution is that you add another layer of encapsulation, you should make your domain aggregates and events themselves follow another serialization interface, ex:

```
packaged domain

type IImportableAggregate interface {
   GetID() uuid.UUID
   SetID(uuid.UUID)
   GetVersion() int
   SetVersion(int)
   Serialize(interface{}) []byte
   UnSerialize([]byte) []byte
   SetHistory([]IImportableEvent)
}

type IImportableEvent interface {
   GetType() string
   SetType(string)
   Serialize(interface{}) []byte
   UnSerialize([]byte) []byte
}  
```

This would make both domain and repository completly unaware of each-other, and with the creation of simple translator classes isolate the layers.

An example translator class:

```
package translators

type IEventSourceAggregateTranslator interface {
  AggregateToEventSource(IImportableAggregate) IEventSourceAggregate
  EventSourceToAggregate(IEventSourceAggregate) IImportableAggregate
}

....

```

## ALTERNATIVE

Ignore the IImportable stuff, and if you want to isolate stuff, just build your own type on the translator that is able to implement the IEventStore* interfaces.

Event Simpler implement the IEventStore* on the Aggregates... I don't know if this breaks clean architecture.