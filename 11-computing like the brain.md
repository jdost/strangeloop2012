#Computing like the Brain
##Jeff Hawkins

###Neocortex
* "Wrinkly" part of the brain
* Predictive modeling system
* Builds online models from streaming data (predictions, anomaly detection, actions)
* Principles
    1. Hierarchy
    1. Sequence memory - perception is created by patterns of sensory data
    1. Sparse distributed representations - only a small percent is active

###Sparse Distributed Representation
* Dense Representations
    * Few bits (8-128)
    * All combinations of 1's and 0's
    * Individual bits have no inherent meaning
    * Represenation is assigned by programmer
* SDRs
    * Many bits (thousands)
    * Few 1's, mostly 0's
    * Each bit has semantic meaning
    * Meaning of each bit is learned, not assigned
* Properties
    1. Similarity (shared bits = semantic similarity)
    1. Store and compare (subsampling)
    1. Union membership

###Sequence Memory
* Neurons have a **lot** of dendrites
* Act as "coincidence detectors" (OR of many detectors)
* When a cell becomes "active" makes connections to previously active cells
    (subsample)
* 2 dimensions only a first order prediction
    (ABAC would be borked as both A's would predict B and C)
* Adding a 3rd dimensions, you can define long sequences
    (So A would exist twice in the sequence, but preceding B and C in different
    places in the sequence)
* Variable order
* Distributed
* Multiple simultaneous predictions
* High capacity
* Semantic generalization

###Online Learning
* Train on every new input
* If pattern does not repeat, forget it
* If pattern repeats, reinforce it

###Predictive Analytics (Today)
* We have a lot of data stored, allowing for visualization and predictive models
* Challenges: Prep, obsolescence, people
* Future: streams through online models, generating actions
* Criteria: Automated model creation, Continuous learning

###Future of Machine Intelligence
* More theory (Attention, Sensory-motor integration)
* Embodiment
    (Today: Cloud based service,
    Options: embedded, distributed with servers around the world connected)
* Hardware (Speed, cost, power)
    1. Memory (Natural fault tolerance, losing a chunk of bits just degrades the
        prediction, doesn't destroy it)
    1. Interconnects (Sparsity and hierarchy, sub-sampling)
* Applications
    Today: Prediction/anomaly detection
    Classics: vision, language, speech?
    Big wins? Brains are slow, (5ms response)
