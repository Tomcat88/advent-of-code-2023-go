package main

import (
	"aoc23/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

type AlmanacMap struct {
	entries []AlmanacEntry
}

func (m *AlmanacMap) addEntry(destinationStart int, sourceStart int, rangeLength int) {
	m.entries = append(m.entries, AlmanacEntry{destinationStart, sourceStart, rangeLength})
}

func (m *AlmanacMap) getDestination(source int) (destination int) {
	destination = source
	for _, e := range m.entries {
		if d, found := e.getDestination(source); found {
			destination = d
		}
	}
	return
}

type AlmanacEntry struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func (e *AlmanacEntry) getDestination(source int) (int, bool) {
	if source < e.sourceStart || source > e.sourceStart+e.rangeLength {
		return -1, false
	}
	return e.destinationStart + (source - e.sourceStart), true
}

func main() {
	content, err := utils.ReadInputAsString("day5")
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", part1(content))
	fmt.Println("Part 2:", part2(content))
}

func convertStringToMap(mapEntry string) (m AlmanacMap) {
	mapLines := strings.Split(mapEntry, "\n")
	for _, line := range mapLines {
		entry := strings.Split(line, " ")
		if len(entry) < 3 {
			continue
		}
		destinationStart, _ := strconv.Atoi(entry[0])
		sourceStart, _ := strconv.Atoi(entry[1])
		rangeLength, _ := strconv.Atoi(entry[2])
		m.addEntry(destinationStart, sourceStart, rangeLength)
	}
	return
}

func part1(content string) (lowest int) {
	lowest = math.MaxInt32
	entries := strings.Split(content, "\n\n")
	seedsEntry, _ := strings.CutPrefix(entries[0], "seeds: ")
	seedsStr := strings.Split(seedsEntry, " ")
	seeds := make([]int, len(seedsStr))
	for i, s := range seedsStr {
		seeds[i], _ = strconv.Atoi(s)
	}
	seedToSoilMap := convertStringToMap(entries[1])
	soilToFertilizerMap := convertStringToMap(entries[2])
	fertilizerToWaterMap := convertStringToMap(entries[3])
	waterToLigthMap := convertStringToMap(entries[4])
	lightToTemperatureMap := convertStringToMap(entries[5])
	tempearatureToHumidityMap := convertStringToMap(entries[6])
	humidityToLocationMap := convertStringToMap(entries[7])

	for _, seed := range seeds {
		soil := seedToSoilMap.getDestination(seed)
		fertilizer := soilToFertilizerMap.getDestination(soil)
		water := fertilizerToWaterMap.getDestination(fertilizer)
		light := waterToLigthMap.getDestination(water)
		temperature := lightToTemperatureMap.getDestination(light)
		humidity := tempearatureToHumidityMap.getDestination(temperature)
		location := humidityToLocationMap.getDestination(humidity)
		if location < lowest {
			lowest = location
		}
	}
	return
}

func part2(content string) (lowest int) {
	lowest = math.MaxInt32
	entries := strings.Split(content, "\n\n")
	seedsEntry, _ := strings.CutPrefix(entries[0], "seeds: ")
	seedsStr := strings.Split(seedsEntry, " ")
	seeds := make([]int, len(seedsStr))
	for i, s := range seedsStr {
		seeds[i], _ = strconv.Atoi(s)
	}
	seedToSoilMap := convertStringToMap(entries[1])
	soilToFertilizerMap := convertStringToMap(entries[2])
	fertilizerToWaterMap := convertStringToMap(entries[3])
	waterToLigthMap := convertStringToMap(entries[4])
	lightToTemperatureMap := convertStringToMap(entries[5])
	tempearatureToHumidityMap := convertStringToMap(entries[6])
	humidityToLocationMap := convertStringToMap(entries[7])

	resultChan := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	go func() {
		for result := range resultChan {
			if result < lowest {
				fmt.Println(result)
				lowest = result
			}
		}
		done <- true
	}()
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedRange := seeds[i+1]
		for seed := seedStart; seed < seedStart+seedRange; seed++ {
			wg.Add(1)
			go func(seed int) {
				soil := seedToSoilMap.getDestination(seed)
				fertilizer := soilToFertilizerMap.getDestination(soil)
				water := fertilizerToWaterMap.getDestination(fertilizer)
				light := waterToLigthMap.getDestination(water)
				temperature := lightToTemperatureMap.getDestination(light)
				humidity := tempearatureToHumidityMap.getDestination(temperature)
				location := humidityToLocationMap.getDestination(humidity)
				resultChan <- location
				wg.Done()
			}(seed)
		}
	}
	fmt.Println("Waiting")
	wg.Wait()
	close(resultChan)
	fmt.Println("Chan closed")
	<-done
	return
}
