#include <iostream>
#include <vector>
#include <string>
#include <iomanip>
#include <fstream>
#include <codecvt>
#include <io.h>
#include "Indexi.h"
#include "Everything.h"
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

int main() {
    FILESYSTEM_INDEXING_HANDLER filesystem_indexing_handler;
    filesystem_indexing_handler.writeToCSVFile("allMusic");
    return 0;
}

vector<SINGLE_FILE> FILESYSTEM_INDEXING_HANDLER::getAllMusicFiles() {
    vector<SINGLE_FILE> filelist;
    Everything_SetSearchW(L".mp3");
    Everything_SetRequestFlags(EVERYTHING_REQUEST_FILE_NAME|EVERYTHING_REQUEST_PATH|EVERYTHING_REQUEST_SIZE);
    Everything_SetSort(EVERYTHING_SORT_NAME_ASCENDING);
    Everything_QueryW(TRUE);

    {
        DWORD i = 0;
        DWORD numberOfItems = Everything_GetNumResults();
        for( i = 0; i < numberOfItems ; i++) {
            LARGE_INTEGER size;
            Everything_GetResultSize(i,&size);

            wstring ws = Everything_GetResultFileNameW(i);
            wstring wsp = Everything_GetResultPathW(i);

            SINGLE_FILE single_file;
            single_file.filename = ws;
            single_file.filepath = wsp;
            filelist.push_back(single_file);
        }
    }

    return filelist;
}

wstring FILESYSTEM_INDEXING_HANDLER::getCsvContent(vector<SINGLE_FILE> array) {
    wstring csv;
    for (auto &item : array) {
        csv.append(L"\"");
        csv.append(item.filename);
        csv.append(L"\"");
        csv.append(L",");
        csv.append(L"\"");
        csv.append(item.filepath);
        csv.append(L"\"");
        csv.append(L"\n");
    }
    return csv;
}

void FILESYSTEM_INDEXING_HANDLER::writeToCSVFile(string filename) {
    filename += ".csv";
    const locale utf8_locale = locale(locale() , new codecvt_utf8_utf16<wchar_t>());
    wfstream file (filename , ios::out | ios::binary); //open file stream with the filename and flag output
    if(file.fail()) { //if it fails to open the file.
        cout << "Error Occurred.";
        return;
    }
    file.imbue(utf8_locale);

    (file << this->getCsvContent(this->getAllMusicFiles()));
    file.close(); //finally close the file

}
