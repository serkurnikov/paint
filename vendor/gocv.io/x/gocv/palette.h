#ifndef _OPENCV3_PICTURE_H_
#define _OPENCV3_PICTURE_H_

#include <stdlib.h>
#include <stdio.h>

#ifdef __cplusplus
#include <opencv2/opencv.hpp>

extern "C" {
#endif

#include "core.h"

const cv::Scalar BLACK = CV_RGB(0, 0, 0);
const cv::Scalar WHITE = CV_RGB(255, 255, 255);

const cv::Scalar RED = CV_RGB(255, 0, 0);
const cv::Scalar ORANGE = CV_RGB(255, 100, 0);
const cv::Scalar YELLOW = CV_RGB(255, 255, 0);
const cv::Scalar GREEN = CV_RGB(0, 255, 0);
const cv::Scalar LIGHTBLUE = CV_RGB(60, 170, 255);
const cv::Scalar BLUE = CV_RGB(0, 0, 255);
const cv::Scalar VIOLET = CV_RGB(194, 0, 255);

const cv::Scalar GINGER = CV_RGB(215, 125, 49);
const cv::Scalar PINK = CV_RGB(255, 192, 203);
const cv::Scalar LIGHTGREEN = CV_RGB(153, 255, 153);
const cv::Scalar BROWN = CV_RGB(150, 75, 0);

typedef unsigned char uchar;
typedef unsigned int uint;

typedef struct ColorCluster {
    cv::Scalar color;
	cv::Scalar new_color;
	int count;

	ColorCluster():count(0) {}
} ColorCluster;

float RgbEuclidean(cv::Scalar p1, cv::Scalar p2);
bool ColorsSort(std::pair< int, uint > a, std::pair< int, uint > b);
void BuildPalette (char* imageFilename, char* imageOutFilename);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PICTURE_H