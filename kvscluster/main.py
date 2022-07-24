import argparse
import random
import numpy as np
from sklearn.preprocessing import StandardScaler
from hdbscan import HDBSCAN
from sklearn import metrics
import time

import matplotlib.pyplot as plt


def generate_map(num_of_keys, num_of_values_per_key, max_value):
  m = {}
  for i in range(num_of_keys):
    m[i] = []
    for j in range(num_of_values_per_key):
      m[i].append(random.randint(0, 100))
  return m

def main():
  # Parse command line
  cli = argparse.ArgumentParser(description="""Cluster the keys of a multi-valued map based
  on the set difference of their values.""")
  cli.add_argument('--nk', help='number of keys', type=int, default=100)
  cli.add_argument('--nvpk', help='number of values per key',type=int, default=2)
  cli.add_argument('--random-seed', help='use random seed', default=False)
  cli.add_argument('--seed', help='seed', default=5)
  cli.add_argument('--metric', help='metric function', type=str, default='manhattan')
  cli.add_argument('--cluster-selection-method', help='', type=str, default='leaf')
  cli.add_argument('--min-cluster-size', help='', type=int, default=5)
  cli.add_argument('--max-value', help='', type=int, default=100)

  args = cli.parse_args()

  if not args.random_seed:
    random.seed(args.seed)

  # Generate multi-valued map
  m = generate_map(args.nk, args.nvpk, args.max_value)
  # Also create reverse index to retrieve key from values "vector"
  reverse_map = {','.join([str(i) for i in v]): k for k, v in m.items()}

  # Cluster data
  # Treat value lists as vectors
  X = np.array(list(m.values()))
  X = StandardScaler().fit_transform(X)

  hdb_t1 = time.time()
  hdb = HDBSCAN(metric=args.metric, cluster_selection_method=args.cluster_selection_method, min_cluster_size=args.min_cluster_size).fit(X)
  hdb_labels = hdb.labels_
  hdb_elapsed_time = time.time() - hdb_t1

  # Number of clusters in labels, ignoring noise if present.
  n_clusters_hdb_ = len(set(hdb_labels)) - (1 if -1 in hdb_labels else 0)

  print('\n\n++ HDBSCAN Results')
  print('Estimated number of clusters: %d' % n_clusters_hdb_)
  print('Elapsed time to cluster: %.4f s' % hdb_elapsed_time)

  ##############################################################################
  # Plot result

  # Black removed and is used for noise instead.
  hdb_unique_labels = set(hdb_labels)
  hdb_colors = plt.cm.Spectral(np.linspace(0, 1, len(hdb_unique_labels)))
  fig = plt.figure(figsize=plt.figaspect(0.5))
  hdb_axis = fig.add_subplot(121)
  for k, col in zip(hdb_unique_labels, hdb_colors):
      if k == -1:
          # Black used for noise.
          col = 'k'

      hdb_axis.plot(X[hdb_labels == k, 0], X[hdb_labels == k, 1], 'o', markerfacecolor=col,
                    markeredgecolor='k', markersize=6)

  hdb_axis.set_title('HDBSCAN\nEstimated number of clusters: %d' % n_clusters_hdb_)
  fig.savefig("foo.pdf", bbox_inches='tight')

main()