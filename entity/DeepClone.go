/*
 *  @Author : huangzj
 *  @Time : 2020/5/15 9:54
 *  @Description：
 */

package entity

type SynthesisRuleCmp struct {
	*SynthesisRule
	RemainingSynthesisTime int
	Sort                   int
}

type SynthesisRule struct {
	SynthesisId    string  `json:"SynthesisId"`
	SynthesisNum   int     `json:"SynthesisNum"`
	CostItems      [][]int `json:"CostItems"`
	SynthesisItems [][]int `json:"SynthesisItems"`
}
