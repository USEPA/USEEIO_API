"""
Extensive unit tests of the {model}/calculation endpoint
"""
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
           N = D L"""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('check the upstream impact matrix U in model %s', model_id)
            D = client.get_matrix(model_id, "D")
            L = client.get_matrix(model_id, "L")
            N = D @ L
            self.compare_matrices(N, client.get_matrix(model_id, 'N'))

    def test_direct_perspective(self):
        """Test the calculation of a result of the direct perspective."""
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
            self.compare_matrices(R, np.asarray(r['data'], dtype=np.float64))

    def test_intermediate_perspective(self):
        """Test the calculation of a result of the intermediate perspective."""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('test the intermediate perspective calculation in model %s',
                     model_id)
            # first, calculate the expected result matrix R
            L = client.get_matrix(model_id, "L")
            _, n = L.shape
            s = np.zeros(n)
            for j in range(0, n):
                s += L[:, j] * (j + 1.0)
            N = client.get_matrix(model_id, "N")
            R = N @ np.diag(s)

            # compare this with the result from the API
            demand = self.build_test_demand(model_id)
            demand['perspective'] = 'intermediate'
            r = client.calculate(model_id, demand)
            self.compare_matrices(R, np.asarray(r['data'], dtype=np.float64))

    def test_final_perspective(self):
        """Test the calculation of a result of the final perspective."""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('test the final perspective calculation in model %s',
                     model_id)

            # first, calculate the expected result matrix R
            N = client.get_matrix(model_id, "N")
            _, n = N.shape
            d = np.zeros(n)
            for j in range(0, n):
                d[j] = j + 1.0
            R = N @ np.diag(d)

            # compare this with the result from the API
            demand = self.build_test_demand(model_id)
            demand['perspective'] = 'final'
            r = client.calculate(model_id, demand)
            self.compare_matrices(R, np.asarray(r['data'], dtype=np.float64))

    def test_total_results(self):
        """All perspectives should give the same total result."""
        client = getclient()
        for model in client.get_models():
            model_id = model['id']
            log.info('test the total results in the perspectives calculation'
                     ' of model %s', model_id)

            # first, calculate the expected total result t
            N = client.get_matrix(model_id, "N")
            m, n = N.shape
            t = np.zeros(m)
            for j in range(0, n):
                t += N[:, j] * (j + 1.0)

            # compare this with the result from the API
            demand = self.build_test_demand(model_id)
            for p in ['direct', 'intermediate', 'final']:
                demand['perspective'] = p
                r = client.calculate(model_id, demand)
                totals = r['totals']
                for i in range(0, m):
                    self.assertAlmostEqual(t[i], totals[i])

    def build_test_demand(self, model_id: str):
        """Creates a test demand for the given model. For each sector with
           index i, the demand vector d gets an entry of d[i] = i + 1.0."""
        client = getclient()
        sectors = client.get_sectors(model_id)
        entries = []
        for sector in sectors:
            entries.append({
                'sector': sector['id'],
                'amount': sector['index'] + 1.0
            })
        return {'demand': entries}

    def compare_matrices(self, M1: np.ndarray, M2: np.ndarray):
        """Compares all values in two numeric matrices with AlmostEqual

        :param M1: matrix 1, np.ndarray
        :param M2: matrix 2, np.ndarray
        :return: pass/fail
        """
        m, n = M1.shape
        m_, n_ = M2.shape
        self.assertEqual(m, m_)
        self.assertEqual(n, n_)
        for i in range(0, m):
            for j in range(0, n):
                self.assertAlmostEqual(M1[i, j], M2[i, j])
