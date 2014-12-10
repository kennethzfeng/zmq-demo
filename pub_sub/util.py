import zmq


def display_version():
    print('pyzmq version: %s' % zmq.__version__)
    print('ZeroMQ version: %s' % zmq.zmq_version())
