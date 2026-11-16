cmake_minimum_required(VERSION 3.16)
project(access_log LANGUAGES CXX)

add_executable(access-log main.cpp)
target_compile_features(access-log PRIVATE cxx_std_17)
