cmake_minimum_required(VERSION 3.19.0)
set(PROJECT_NAME the-cubes)
project(${PROJECT_NAME} VERSION 0.2.0)

# Optimization flags for LLVM Clang
add_compile_options("-static")
add_compile_options("-O3")
add_compile_options("-flto=full")

set(BUILD_EXAMPLES OFF CACHE BOOL "" FORCE)

add_subdirectory("deps/raylib")
find_package(raylib 3.5 QUIET)

add_executable(${PROJECT_NAME} src/main.cpp)

target_link_libraries(${PROJECT_NAME} raylib winmm)
