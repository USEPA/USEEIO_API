import logging as log
import numpy as np
import unittest

from config import getclient


class CalculationTest(unittest.TestCase):

    def test_check_leontief_inverse(self):
        """Test the reproducibility of the Leontief inverse:
           L = (I - A)^{-1}
        """
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('check the Leontief inverse L in model %s', model_id)
            A = client.get_matrix(model_id, "A")
            _, n = A.shape
            L = np.linalg.inv(np.eye(n) - A)
            self.compare_matrices(L, client.get_matrix(model_id, 'L'))

    def test_direct_impacts(self):
        """Test the reproducibility of the direct impact matrix:
           D = C B"""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('check the direct impact matrix D in model %s', model_id)
            B = client.get_matrix(model_id, "B")
            C = client.get_matrix(model_id, "C")
            D = C @ B
            self.compare_matrices(D, client.get_matrix(model_id, 'D'))

    def test_upstream_impacts(self):
        """Test the reproducibility of the upstream impact matrix:
           U = D L"""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('check the upstream impact matrix U in model %s', model_id)
            D = client.get_matrix(model_id, "D")
            L = client.get_matrix(model_id, "L")
            U = D @ L
            self.compare_matrices(U, client.get_matrix(model_id, 'U'))

    def test_direct_perspective(self):
        """Test the calculation of a result of the direct perspecitve."""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('test the direct perspective calculation in model %s',
                     model_id)
            # first, calculate the expected result matrix R
            L = client.get_matrix(model_id, "L")
            _, n = L.shape
            s = np.zeros(n)
            for j in range(0, n):
                s += L[:, j] * (j + 1.0)
            D = client.get_matrix(model_id, "D")
            R = D @ np.diag(s)

            # compare this with the result from the API
            demand = self.build_test_demand(model_id)
            demand['perspective'] = 'direct'
            r = client.calculate(model_id, demand)
            self.compare_matrices(R, np.asarray(r['data'], dtype=np.float))

    def build_test_demand(self, model_id: str):
        client = getclient()
        sectors = client.get_sectors(model_id)
        entries = []
        for sector in sectors:
            entries.append({
                'sector': sector['id'],
                'amount': sector['index'] + 1.0
            })
        return {'entries': entries}

    def compare_matrices(self, M1: np.ndarray, M2: np.ndarray):
        m, n = M1.shape
        m_, n_ = M2.shape
        self.assertEqual(m, m_)
        self.assertEqual(n, n_)
        for i in range(0, m):
            for j in range(0, n):
                self.assertAlmostEqual(M1[i, j], M2[i, j])
