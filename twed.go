//
//	Package - transpiled by c4go
//
//	If you have found any issues, please raise an issue at:
//	https://github.com/Konstantin8105/c4go/
//

package main

import "math"
import "github.com/Konstantin8105/c4go/noarch"

// CTWED - transpiled function from  /home/pascal/go/src/github.com/kylhuk/masterthesis/dev/src/twed.c:74
func CTWED(ta []float64, la []int32, tsa []float64, tb []float64, lb []int32, tsb []float64, nu []float64, lambda []float64, degree []int32, dist []float64, dimension []int32) {
	if la[0] < 0 || lb[0] < 0 {
		// Note the original code was extended to cover inputs of R^N.
		//
		//   This requires an extra argument (int *dimension),
		//   and in that case expects ta and tb to be c-ordered shape
		//   (la, N) and (lb, N) respectively.
		//
		//   I also add the root into the norm calls, because it
		//   is more consistent with implementations that followed this.
		//
		//   Otherwise up to white space and some formatting,
		//   this is taken right from Marteau as a reference implementation.
		//
		//   Garrett Wright, 2020
		//
		//
		//  Filename: twed.c
		//  source code for the Time Warp Edit Distance in ANSI C.
		//  Author: Pierre-Francois Marteau, adapted for R integration by Jörg Schaber
		//  Version: V1.2.a du 25/08/2014, radix addition line 101, thanks to Benjamin Herwig from University of Kassel, Germany
		//  Licence: GPL
		//  ******************************************************************
		//  This software and description is free delivered "AS IS" with no
		//  guaranties for work at all. Its up to you testing it modify it as
		//  you like, but no help could be expected from me due to lag of time
		//  at the moment. I will answer short relevant questions and help as
		//  my time allow it. I have tested it played with it and found no
		//  problems in stability or malfunctions so far.
		//  Have fun.
		//  *****************************************************************
		//  Please cite as:
		//  @article{Marteau:2009:TWED,
		//  author = {Marteau, Pierre-Francois},
		//  title = {Time Warp Edit Distance with Stiffness Adjustment for Time Series Matching},
		//  journal = {IEEE Trans. Pattern Anal. Mach. Intell.},
		//  issue_date = {February 2009},
		//  volume = {31},
		//  number = {2},
		//  month = feb,
		//  year = {2009},
		//  issn = {0162-8828},
		//  pages = {306--318},
		//  numpages = {13},
		//  url = {http://dx.doi.org/10.1109/TPAMI.2008.76},
		//  doi = {10.1109/TPAMI.2008.76},
		//  acmid = {1496043},
		//  publisher = {IEEE Computer Society},
		//  address = {Washington, DC, USA},
		//  keywords = {Dynamic programming, Pattern recognition, Pattern recognition, time series, algorithms, similarity measures,
		//  Similarity measures, algorithms, similarity measures., time series},
		//  }
		//
		//
		//  INPUTS
		//  double ta[]: array containing the first time series; ta[i] is the i^th sample for i in {0, .., la-1}
		//  int *la: length of the first time series
		//  double tsa[]: array containing the time stamps for time series ta; tsa[i] is the time stamp for sample ta[i]. The length of tsa array is expected to be la.
		//  double tb[]: array containing the second time series; tb[i] is the i^th sample for i in {0, .., lb-1}
		//  int *lb: length of the second time series
		//  double tsb[]: array containing the time stamps for time series tb; tsb[j] is the time stamp for sample tb[j]. The length of tsb array is expected to be lb.
		//  double *nu: value for parameter nu
		//  double *lambda: value for parameter lambda
		//  int *degree: power degree for the evaluation of the local distance between samples: degree>0 required
		//  OUTPUT
		//  double: the TWED distance between time series ta and tb.
		//
		//
		//    TWED PAMI
		//
		noarch.Fprintf(noarch.Stderr, []byte("twed: the lengths of the input timeseries should be greater or equal to 0\n\x00"))
		noarch.Exit(-1)
	}
	var r int32 = la[0]
	var c int32 = lb[0]
	var deg int32 = noarch.Abs(degree[0])
	var disti1 float64
	var distj1 float64
	var i int32
	var j int32
	var n int32
	var dim int32 = dimension[0]
	var D [][]float64 = make([][]float64, uint32(r+1))
	var Di1 []float64 = make([]float64, uint32(r+1))
	var Dj1 []float64 = make([]float64, uint32(c+1))
	var dmin float64
	var htrans float64
	var dist0 float64
	{
		// allow a hack to repro authors original results.
		// allocations
		for i = 0; i <= r; i++ {
			D[i] = make([]float64, uint32(c+1))
		}
	}
	{
		// local costs initializations
		for j = 1; j <= c; j++ {
			distj1 = 0
			for n = 0; n < dim; n++ {
				if j > 1 {
					distj1 += math.Pow(math.Abs(tb[(j-2)*dim+n]-tb[(j-1)*dim+n]), float64(deg))
				} else {
					distj1 += math.Pow(math.Abs(tb[(j-1)*dim+n]), float64(deg))
				}
			}
			if degree[0] < 0 {
				// NOTE original author did not nth-root
				// I provide "no root" as negative degree, to match any prior results
				// Consider it a semi hidden feature.
				Dj1[j] = distj1
			} else if deg == 2 {
				Dj1[j] = math.Sqrt(distj1)
			} else {
				Dj1[j] = math.Pow(distj1, 1/float64(deg))
			}
		}
	}
	for i = 1; i <= r; i++ {
		disti1 = 0
		for n = 0; n < dim; n++ {
			if i > 1 {
				disti1 += math.Pow(math.Abs(ta[(i-2)*dim+n]-ta[(i-1)*dim+n]), float64(deg))
			} else {
				disti1 += math.Pow(math.Abs(ta[(i-1)*dim+n]), float64(deg))
			}
		}
		if degree[0] < 0 {
			// NOTE original author did not nth-root
			// I provide "no root" as negative degree, to match any prior results
			// Consider it a semi hidden feature.
			Di1[i] = disti1
		} else if deg == 2 {
			Di1[i] = math.Sqrt(disti1)
		} else {
			Di1[i] = math.Pow(disti1, 1/float64(deg))
		}
		for j = 1; j <= c; j++ {
			dist[0] = 0
			dist0 = 0
			for n = 0; n < dim; n++ {
				dist[0] += math.Pow(math.Abs(ta[(i-1)*dim+n]-tb[(j-1)*dim+n]), float64(deg))
				if i > 1 && j > 1 {
					dist0 += math.Pow(math.Abs(ta[(i-2)*dim+n]-tb[(j-2)*dim+n]), float64(deg))
				}
			}
			if degree[0] < 0 {
				// NOTE original author did not nth-root
				// I provide "no root" as negative degree, to match any prior results
				// Consider it a semi hidden feature.
				D[i][j] = dist[0] + dist0
			} else if deg == 2 {
				D[i][j] = math.Sqrt(dist[0]) + math.Sqrt(dist0)
			} else {
				D[i][j] = math.Pow(dist[0], 1/float64(deg)) + math.Pow(dist0, 1/float64(deg))
			}
		}
	}
	// for i
	// border of the cost matrix initialization
	D[0][0] = 0
	for i = 1; i <= r; i++ {
		D[i][0] = float64(__builtin_inff())
	}
	for j = 1; j <= c; j++ {
		D[0][j] = float64(__builtin_inff())
	}
	for i = 1; i <= r; i++ {
		for j = 1; j <= c; j++ {
			htrans = math.Abs(tsa[i-1] - tsb[j-1])
			if j > 1 && i > 1 {
				htrans += math.Abs(tsa[i-2] - tsb[j-2])
			}
			dist0 = D[i-1][j-1] + D[i][j] + nu[0]*htrans
			dmin = dist0
			if i > 1 {
				htrans = tsa[i-1] - tsa[i-2]
			} else {
				htrans = tsa[i-1]
			}
			dist[0] = Di1[i] + D[i-1][j] + lambda[0] + nu[0]*htrans
			if dmin > dist[0] {
				dmin = dist[0]
			}
			if j > 1 {
				htrans = tsb[j-1] - tsb[j-2]
			} else {
				htrans = tsb[j-1]
			}
			dist[0] = Dj1[j] + D[i][j-1] + lambda[0] + nu[0]*htrans
			if dmin > dist[0] {
				dmin = dist[0]
			}
			D[i][j] = dmin
		}
	}
	dist[0] = D[r][c]
	{
		// freeing
		for i = 0; i <= r; i++ {
			_ = D[i]
		}
	}
	_ = D
	_ = Di1
	_ = Dj1
}

// __builtin_inff from math.h
// c function : float __builtin_inff()
// dep pkg    : math
// dep func   :
func __builtin_inff() float32 {
	return float32(math.Inf(0))
}
