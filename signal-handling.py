import signal
import os
import time

def receiveSignal(signalNumber, frame):
    print('Received:', signalNumber)
    raise SystemExit('Exiting')
    return

if __name__ == '__main__':
    # register the signal to be caught
    signal.signal(signal.SIGUSR1, receiveSignal)

    # register the signal to be ignored
    signal.signal(signal.SIGINT, signal.SIG_IGN)

    # output current process id
    print('My PID is:', os.getpid())

    signal.pause()

