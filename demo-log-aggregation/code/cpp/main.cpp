#include <algorithm>
#include <array>
#include <cstdint>
#include <iostream>
#include <map>
#include <string>
#include <vector>

struct Request {
    std::uint64_t status;
    std::uint64_t bytes;
    std::uint64_t latency;
};

std::array<std::string, 5> split_fields(const std::string& line) {
    std::array<std::string, 5> fields;
    std::size_t begin = 0;
    for (std::size_t i = 0; i + 1 < fields.size(); ++i) {
        const auto comma = line.find(',', begin);
        fields[i] = line.substr(begin, comma - begin);
        begin = comma + 1;
    }
    fields.back() = line.substr(begin);
    return fields;
}

int main() {
    std::ios::sync_with_stdio(false);
    std::cin.tie(nullptr);

    // Keep each request so reading and summarizing are separate steps.
    std::map<std::string, std::vector<Request>> groups;
    std::string line;
    while (std::getline(std::cin, line)) {
        const auto fields = split_fields(line);
        groups[fields[1]].push_back({
            std::stoull(fields[2]),
            std::stoull(fields[3]),
            std::stoull(fields[4]),
        });
    }

    for (const auto& [endpoint, requests] : groups) {
        std::uint64_t errors = 0;
        std::uint64_t bytes = 0;
        std::uint64_t max_latency = 0;
        for (const auto& request : requests) {
            if (request.status >= 500) {
                ++errors;
            }
            bytes += request.bytes;
            max_latency = std::max(max_latency, request.latency);
        }
        std::cout << endpoint << ' ' << requests.size() << ' ' << errors << ' '
                  << bytes << ' ' << max_latency << '\n';
    }
    return std::cin.bad() || !std::cout;
}
