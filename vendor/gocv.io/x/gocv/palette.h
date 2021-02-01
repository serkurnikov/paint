#ifndef _OPENCV3_PICTURE_H_
#define _OPENCV3_PICTURE_H_

#ifdef __cplusplus
#include <opencv2/opencv.hpp>

extern "C" {
#endif

#include "core.h"


void BuildPalette (char* imageFilename, char* imageOutFilename);

#ifdef __cplusplus
}
#endif

#endif //_OPENCV3_PICTURE_H