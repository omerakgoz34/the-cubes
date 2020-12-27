cmake_minimum_required(VERSION 3.0.0)
set(PROJECT_NAME the-cubes)
project(${PROJECT_NAME} VERSION 0.2.0)

set(BUILD_EXAMPLES OFF CACHE BOOL "" FORCE)

add_subdirectory("deps/raylib")

find_package(raylib 3.5 REQUIRED)

add_executable(${PROJECT_NAME} src/main.cpp)

target_link_libraries(${PROJECT_NAME} raylib)

include(CheckIPOSupported)
check_ipo_supported(RESULT result OUTPUT output)
if(result)
  set_property(TARGET ${PROJECT_NAME} PROPERTY INTERPROCEDURAL_OPTIMIZATION TRUE)
else()
  message(WARNING "IPO is not supported: ${output}")
endif()
