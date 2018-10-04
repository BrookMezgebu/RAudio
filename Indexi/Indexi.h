//
// Created by BrookMG on 10/3/2018.
//
#include <iostream>
#include <string>
#include <vector>

using namespace std;

struct SINGLE_FILE {
    wstring filename;
    wstring filepath;
};

class FILESYSTEM_INDEXING_HANDLER {

public:
    vector<SINGLE_FILE> getAllMusicFiles ();
    wstring getCsvContent(vector<SINGLE_FILE> array);
    void writeToCSVFile (string filename);
};

