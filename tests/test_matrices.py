"""
Tests {model}/{matrix} and {model}/{matrix}?row= and {model}/{matrix}?col= endpoints
"""

import logging as log
import random
import unittest

from client import Client


class MatrixTest(unittest.TestCase):

    def test_matrices(self):
        """For each numeric matrix this test first fetches the complete matrix
           and a random row and column from the server and checks that the
           data of the matrix, row, and column data match."""
        client = Client.get()
        for model in client.get_models():
            model_id = model['id']
            for matrix in ['A', 'B', 'C', 'D', 'L', 'N']:
                log.info('check matrix %s in model %s', matrix, model_id)
                M = client.get_matrix(model_id, matrix)
                rows, cols = M.shape
                row = random.randint(0, rows - 1)
                log.info('check row %i in matrix %s/%s', row, matrix, model_id)
                r = client.get_matrix_row(model_id, matrix, row)
                for col in range(0, cols):
                    self.assertAlmostEqual(M[row, col], r[col])
                col = random.randint(0, cols - 1)
                log.info('check column %i in matrix %s/%s',
                         col, matrix, model_id)
                c = client.get_matrix_column(model_id, matrix, col)
                for row in range(0, rows):
                    self.assertAlmostEqual(M[row, col], c[row])
