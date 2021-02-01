#include "palette.h"

float RgbEuclidean(cv::Scalar p1, cv::Scalar p2) {
   float val = sqrtf( (p1.val[0]-p2.val[0])*(p1.val[0]-p2.val[0]) +
   		(p1.val[1]-p2.val[1])*(p1.val[1]-p2.val[1]) +
   		(p1.val[2]-p2.val[2])*(p1.val[2]-p2.val[2]) +
   		(p1.val[3]-p2.val[3])*(p1.val[3]-p2.val[3]));

   	return val;
}

bool ColorsSort(std::pair< int, uint > a, std::pair< int, uint > b) {
	return (a.second > b.second);
}

void BuildPalette(char* imageFilename, char* imageOutFilename) {
    cv::Mat src = cv::imread(imageFilename, cv::IMREAD_COLOR);

    cv::resize(src, src, cv::Size(), 0.5, 0.5);

    cv::Mat cluster_indexes = cv::Mat::zeros(cv::Size(src.cols, src.rows), CV_8UC3);

    #define CLUSTER_COUNT 10
    	int cluster_count = CLUSTER_COUNT;
    	ColorCluster clusters[CLUSTER_COUNT];

    	int i=0, j=0, k=0, x=0, y=0;

    #if 0
    	clusters[0].new_color = RED;
    	clusters[1].new_color = ORANGE;
    	clusters[2].new_color = YELLOW;
    #elif 0
    	clusters[0].new_color = BLACK;
    	clusters[1].new_color = GREEN;
    	clusters[2].new_color = WHITE;
    #else
    	cv::RNG rng = cv::RNG(-1);
    	for(k=0; k<cluster_count; k++) {}
    	    //clusters[k].new_color = cv::Scalar((CV_RandInt(&rng)%255), (CV_RandInt(&rng)%255), (CV_RandInt(&rng)%255), 0)
    #endif

    float min_rgb_euclidean = 0, old_rgb_euclidean=0;

    while(1) {
    		for(k=0; k<cluster_count; k++) {
    			clusters[k].count = 0;
    			clusters[k].color = clusters[k].new_color;
    			clusters[k].new_color = cv::Scalar(0);
    		}

    		for (y=0; y<src.cols; y++) {
            			for (x=0; x<src.rows; x++) {

            			    cv::Vec3b intensity = src.at<cv::Vec3b>(x, y);

            			    uchar B = intensity.val[0];
            			    uchar G = intensity.val[1];
            			    uchar R = intensity.val[2];

            				min_rgb_euclidean = 255*255*255;
            				int cluster_index = -1;
            				for(k=0; k<cluster_count; k++) {
            					float euclid = RgbEuclidean(cv::Scalar(B, G, R, 0), clusters[k].color);
            					if(  euclid < min_rgb_euclidean ) {
            						min_rgb_euclidean = euclid;
            						cluster_index = k;
            					}
            				}

            				cv::Vec3b cluster_intensity = cluster_indexes.at<cv::Vec3b>(x, y);
            				cluster_index = cluster_intensity.val[0];

            				clusters[cluster_index].count++;
            				clusters[cluster_index].new_color.val[0] += B;
            				clusters[cluster_index].new_color.val[1] += G;
            				clusters[cluster_index].new_color.val[2] += R;
            			}
            }

            min_rgb_euclidean = 0;

            for(k=0; k<cluster_count; k++) {
            			clusters[k].new_color.val[0] /= clusters[k].count;
            			clusters[k].new_color.val[1] /= clusters[k].count;
            			clusters[k].new_color.val[2] /= clusters[k].count;
            			float ecli = RgbEuclidean(clusters[k].new_color, clusters[k].color);
            			if(ecli > min_rgb_euclidean) {
            			    min_rgb_euclidean = ecli;
            			}
            }

            if( fabs(min_rgb_euclidean - old_rgb_euclidean)<1 )
            			break;

            old_rgb_euclidean = min_rgb_euclidean;
    }

    std::vector< std::pair< int, uint > > colors;
    colors.reserve(CLUSTER_COUNT);

    int colors_count = 0;
    for(i=0; i<CLUSTER_COUNT; i++) {
    		std::pair< int, uint > color;
    		color.first = i;
    		color.second = clusters[i].count;
    		colors.push_back( color );
    		if(clusters[i].count>0)
    			colors_count++;
    }

    std::sort( colors.begin(), colors.end(), ColorsSort );
}