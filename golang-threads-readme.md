Goroutine Scheduling 

4 Cores. Hyperthreading leads to 8 virtual cores.
Thread states - RUNNING, RUNNABLE, WAITING.
O/S Context Switching happens at thread level not process level. Each process is given a thread and can start other multiple threads.
Each thread has it's own state which must be preserved in context switching.

Cache lines -

Core L1 L2 L3 RAM
Complete cache line invalidated by a single change to any memory. Memory thrashing.

Context switch problem-
Switch the main thread, and schedule the new thread in the same core. 
Wait for main thread, then schedule new thread in the same core.
Schedule the new thread in the new core.

A typical context switch if takes around 1000ns (without cache coherence overhead), and the processor executes 2 instructions per nanosecond, we are looking at 2k lost instructions.
2*10^9


Golang:- 
runtime.NumCPU()
A go program will be assigned 8 Ps - Processor to execute Go code. There can be at max GOMAXPROCS Ps. Each P can be assigned to at max 1 M (worker thread, machine).
This is the thread managed by the OS and will be scheduled in any core.
Each go programs also have goroutines G, which can be considered application level threads.
Each G will be context switched on and off in M, just as OS threads are context switched on and off from core

LRQ GRQ
Each P has one LRQ, which has goroutines which will context switch in M which is assigned to respective P.
Goroutines in GRQ are not assigned yet. There is a process to move them from GRQ to LRQ.

OS schedular is preemptivee. Kernel manages it, and application running on it have no control or clue on what's happening. 
If they use some synchronization primitives like mutex/atomic instructions, then there are certain `happens-before` guarantees promised.

Go schedular runs in the user space, and is cooperative schedular. That means it needs certains events in the user space in safe-points to make scheduling decisions.
But think of it as the preemptive schedular, as nothing is in the hands of go developers, it is all taken care off in the go's runtime.

Go's thread statuses:-
Waiting
Runnable
Executing

These safe points manifest themselves in the function calls. So, if we don't have a modular program and make tight running loops, we will face latencies in the scheduling as well as garbage collection.
So it is important that function call happen within reasonable time frames. In golang 1.12 they have started to work on the non-cooperative goroutine preemption.

Events on which the schedular generally makes decisions:
1. go keyword
2. System Calls - Sometimes after the system call, M might be blocked. If that is the case, sometimes the schedular can switch the another goroutine from the LFQ to run on that M. System calls are also asynchronous in most O/S now. Using the netpoller implementation KQueue (Mac OS), Epoll (linux), iocp (Windows) these calls are executed asynchronously and won't block our M. The diagram for this can be viewed here. 
In case of a blocking, synchronous call - like calls to the file system (windows o/s has async file system API exposed as well), then it causes M to block on it and hence other G's in the LRQ will be blocked. In such a case, this M is moved out of this system along with the G attached to it. M2 replaces M and switches G2 to run on it. G keeps running on M separately. When G is finished, it is moved back to the LRQ, and M is reserved for future use. (Swapping M with a reserved thread is fast than creating a new one.)
3. garbage collection - like context switch the goroutine which will make allocations in heap
4. Using synchronization primitives.

Work Stealing - In order to keep the M busy, whenever any P goes idle, it can steal work from other Ps. It also looks at the GRQ to pick up goroutines. Refer the diagram.
Spinning is two level. 1- An idle M with an assigned P is looking for the Goroutines. 2- An M w/o any P assigned.
Idle M's of type 1 do not spin when we have idle M of type 2. There are atmost GOMAXPROCS spinning Ms. Our aimXX is to ensure that there is no RUNNABLE goroutine which can be EXECUTING, and there is no excessive blocking/unblocking at the same time.

A typical context switch in go takes about 200ns the time compared to that of an O/S. This is 5 times faster.
