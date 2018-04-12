package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type LifeGame struct {
	_2dGRID     [][]int
	_next2dGRID [][]int
}

func rule() int {
	return 1
}

func main() {
	lifeGame := LifeGame{}
	lifeGame.fileOpen()
	lifeGame.nextGeneration()
	lifeGame.print2dGRID()
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
	for i := 1; i < len(lifeGame._2dGRID)-1; i++ {
		for j := 1; j < len(lifeGame._2dGRID[i])-1; j++ {
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
	sub2DGRID := [][]int{}
	for _, _1dGRID := range lifeGame._2dGRID[pointX-1 : pointX+1+1] {
		sub2DGRID = append(sub2DGRID, _1dGRID[pointY-1:pointY+1+1])
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
