cmake_minimum_required(VERSION 3.10)
project(Indexi)

set(CMAKE_CXX_STANDARD 11)

add_executable(Indexi Indexi.cpp Everything.h Indexi.h json.hpp)
target_link_libraries(Indexi libeverything.a)