#include <iostream>
#include <vector>
#include <string>
#include "Indexi.h"
#include "Everything.h"

using namespace std;

int main() {
    FILESYSTEM_INDEXING_HANDLER filesystem_indexing_handler;
    vector<SINGLE_FILE> files = filesystem_indexing_handler.getAllMusicFiles();

    for (auto &file : files) {
        cout << file.filename << " : " << file.filepath;
    }

    return 0;
}

vector<SINGLE_FILE> FILESYSTEM_INDEXING_HANDLER::getAllMusicFiles() {
    vector<SINGLE_FILE> filelist;
    SINGLE_FILE single_file;
    single_file.filename = "This is the file name";
    single_file.filepath = "This is the file path";
    filelist.push_back(single_file);
    return filelist;
}
