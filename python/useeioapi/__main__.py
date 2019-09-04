import os
import sys
import useeioapi


def main():
    datadir = './data'
    port = '8080'

    # parse possible command line args
    args = sys.argv
    flag = None
    for arg in args[1:]:  # type: str
        if arg.startswith('-'):
            flag = arg
            continue
        if flag is None:
            print('Unknown argument', arg)
            continue
        if flag == '-data':
            datadir = arg
            flag = None
            continue
        if flag == '-port':
            port = arg
            flag = None
            continue
        print('Unknown flag', flag)
        flag = None

    # check if the data folder exists
    if not os.path.isdir(datadir) or not os.path.exists(datadir):
        print('ERROR: %s is not a folder or does not exist' % datadir)
        exit(-1)

    # check if the port number is an integer
    try:
        int(port)
    except ValueError:
        print('ERROR: %s is not a valid port number' % port)
        exit(-1)

    print('serve data from %s at http://localhost:%s' % (datadir, port))
    useeioapi.serve(datadir, port)


main()
