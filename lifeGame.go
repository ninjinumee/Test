package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type LifeGame struct {
	_2dGRID     [][]int
	_next2dGRID [][]int
}

func main() {
	lifeGame := LifeGame{}
	lifeGame.fileOpen()
	start := time.Now()
	for i := 0; i < 4000; i++ {
		lifeGame.nextGenerationMulti()
	}
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
	//lifeGame.print2dGRID()
}
func (lifeGame *LifeGame) print2dGRID() {
	for _, _1dGRID := range lifeGame._2dGRID {
		for _, element := range _1dGRID {
			fmt.Print(element)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (lifeGame *LifeGame) nextGeneration() {
	for i := 0; i < len(lifeGame._2dGRID); i++ {
		for j := 0; j < len(lifeGame._2dGRID[i]); j++ {
			sub2DGRID := lifeGame.makeSub2dGRID(i, j)
			sumValue := sumAroundCellValues(sub2DGRID)
			switch sub2DGRID[1][1] {
			case 1: //自分が生存している場合
				switch sumValue {
				case 0:
				case 1:
					//過疎
					lifeGame._next2dGRID[i][j] = 0
					break
				case 2:
				case 3:
					//生存
					break
				default:
					//過密
					lifeGame._next2dGRID[i][j] = 0
					break
				}
			case 0: //死亡している場合
				switch sumValue {
				case 3:
					//誕生
					lifeGame._next2dGRID[i][j] = 1
					break
				default:
					break
				}
			}

		}
	}
	for i := 0; i < len(lifeGame._2dGRID); i++ {
		lifeGame._2dGRID[i] = make([]int, len(lifeGame._2dGRID[i]))
		copy(lifeGame._2dGRID[i], lifeGame._next2dGRID[i])
	}
}

func (lifeGame *LifeGame) nextGenerationMulti() {
	wg := new(sync.WaitGroup)
	for i := 0; i < len(lifeGame._2dGRID); i++ {
		for j := 0; j < len(lifeGame._2dGRID[i]); j++ {
			wg.Add(1)
			go func(x int, y int) {
				defer wg.Done()
				sub2DGRID := lifeGame.makeSub2dGRID(x, y)
				sumValue := sumAroundCellValues(sub2DGRID)
				switch sub2DGRID[1][1] {
				case 1: //自分が生存している場合
					switch sumValue {
					case 0:
					case 1:
						//過疎
						lifeGame._next2dGRID[x][y] = 0
						break
					case 2:
					case 3:
						//生存
						break
					default:
						//過密
						lifeGame._next2dGRID[x][y] = 0
						break
					}
				case 0: //死亡している場合
					switch sumValue {
					case 3:
						//誕生
						lifeGame._next2dGRID[x][y] = 1
						break
					default:
						break
					}
				}
			}(i, j)

		}
	}
	wg.Wait()
	for i := 0; i < len(lifeGame._2dGRID); i++ {
		lifeGame._2dGRID[i] = make([]int, len(lifeGame._2dGRID[i]))
		copy(lifeGame._2dGRID[i], lifeGame._next2dGRID[i])
	}

}

func sumAroundCellValues(aroundCells [][]int) int {
	cellCounter := 0
	for _, _1DGRID := range aroundCells {
		for _, cellValue := range _1DGRID {
			cellCounter += cellValue
		}
	}
	return cellCounter - aroundCells[1][1]
}

func (lifeGame *LifeGame) makeSub2dGRID(pointX int, pointY int) [][]int {
	index := [3][3][2]int{}
	for i := pointX - 1; i <= pointX+1; i++ {
		for j := pointY - 1; j <= pointY+1; j++ {
			index[i-pointX+1][j-pointY+1][0] = i
			index[i-pointX+1][j-pointY+1][1] = j
		}
	}
	sub2DGRID := [][]int{}
	for i := 0; i < 3; i++ {
		sub1DGRID := []int{}
		for j := 0; j < 3; j++ {
			if index[i][j][0] < 0 {
				index[i][j][0] = len(lifeGame._2dGRID) - 1
			}
			if index[i][j][0] == len(lifeGame._2dGRID) {
				index[i][j][0] = 0
			}
			if index[i][j][1] < 0 {
				index[i][j][1] = len(lifeGame._2dGRID[i]) - 1
			}
			if index[i][j][1] == len(lifeGame._2dGRID[i]) {
				index[i][j][1] = 0
			}
			x := index[i][j][0]
			y := index[i][j][1]
			sub1DGRID = append(sub1DGRID, lifeGame._2dGRID[x][y])
		}
		sub2DGRID = append(sub2DGRID, sub1DGRID)
	}

	return sub2DGRID
}
func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func (lifeGame *LifeGame) fileOpen() {

	// 読み込み
	fr, err := os.Open("LifeGame.csv")
	failOnError(err)

	defer fr.Close()

	r := csv.NewReader(fr)
	rows, err := r.ReadAll()
	failOnError(err)
	for _, row := range rows {
		var _1DGRID []int
		for _, element := range row {
			e, fail := strconv.Atoi(element)
			failOnError(fail)
			_1DGRID = append(_1DGRID, e)
		}
		lifeGame._2dGRID = append(lifeGame._2dGRID, _1DGRID)
	}

	lifeGame._next2dGRID = make([][]int, len(lifeGame._2dGRID))
	for i := 0; i < len(lifeGame._2dGRID); i++ {
		lifeGame._next2dGRID[i] = make([]int, len(lifeGame._2dGRID[i]))
		copy(lifeGame._next2dGRID[i], lifeGame._2dGRID[i])
	}

}
