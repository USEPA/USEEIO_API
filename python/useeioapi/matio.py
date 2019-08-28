import struct
import numpy


def read_shape(file_path: str):
    """ Reads and returns the shape (rows, columns) from the matrix stored in
        the given file.
    """
    with open(file_path, 'rb') as f:
        rows = struct.unpack('<i', f.read(4))[0]
        cols = struct.unpack('<i', f.read(4))[0]
        return rows, cols


def read_matrix(file_path: str):
    shape = read_shape(file_path)
    return numpy.memmap(file_path, mode='c', dtype='<f8',
                        shape=shape, offset=8, order='F')


def write_matrix(M, file_path: str):
    with open(file_path, 'wb') as f:
        rows, cols = M.shape
        f.write(struct.pack("<i", rows))
        f.write(struct.pack("<i", cols))
        for col in range(0, cols):
            for row in range(0, rows):
                val = M[row, col]
                f.write(struct.pack("<d", val))
