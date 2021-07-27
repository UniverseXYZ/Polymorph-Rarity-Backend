package metadata

import (
	"fmt"
	"strconv"

	"rarity-backend/constants"
	"rarity-backend/structs"
)

const POLYMORPH_IMAGE_URL string = "https://storage.googleapis.com/polymorph-images/"
const EXTERNAL_URL string = "https://universe.xyz/polymorphs/"
const GENES_COUNT = 9
const BACKGROUND_GENE_COUNT int = 12
const BASE_GENES_COUNT int = 11
const SHOES_GENES_COUNT int = 25
const PANTS_GENES_COUNT int = 33
const TORSO_GENES_COUNT int = 34
const EYEWEAR_GENES_COUNT int = 13
const HEAD_GENES_COUNT int = 31
const WEAPON_RIGHT_GENES_COUNT int = 32
const WEAPON_LEFT_GENES_COUNT int = 32

type Genome string
type Gene int

func (g Gene) toPath() string {
	if g < 10 {
		return fmt.Sprintf("0%s", strconv.Itoa(int(g)))
	}

	return strconv.Itoa(int(g))
}

func getGeneInt(g string, start, end, count int) int {
	genomeLen := len(g)
	geneStr := g[genomeLen+start : genomeLen+end]
	gene, _ := strconv.Atoi(geneStr)
	return gene % count
}

func getWeaponLeftGene(g string) int {
	return getGeneInt(g, -18, -16, WEAPON_LEFT_GENES_COUNT)
}

func getWeaponLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getWeaponLeftGene(g)
	trait := configService.WeaponLeft[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.LeftHand,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getWeaponLeftGenePath(g string) string {
	gene := getWeaponLeftGene(g)
	return Gene(gene).toPath()
}

func getWeaponRightGene(g string) int {
	return getGeneInt(g, -16, -14, WEAPON_RIGHT_GENES_COUNT)
}

func getWeaponRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getWeaponRightGene(g)
	trait := configService.WeaponRight[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.RightHand,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getWeaponRightGenePath(g string) string {
	gene := getWeaponRightGene(g)
	return Gene(gene).toPath()
}

func getHeadGene(g string) int {
	return getGeneInt(g, -14, -12, HEAD_GENES_COUNT)
}

func getHeadGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getHeadGene(g)
	trait := configService.Headwear[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Headwear,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getHeadGenePath(g string) string {
	gene := getHeadGene(g)
	return Gene(gene).toPath()
}

func getEyewearGene(g string) int {
	return getGeneInt(g, -12, -10, EYEWEAR_GENES_COUNT)
}

func getEyewearGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getEyewearGene(g)
	trait := configService.Eyewear[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Eyewear,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getEyewearGenePath(g string) string {
	gene := getEyewearGene(g)
	return Gene(gene).toPath()
}

func getShoesGene(g string) int {
	return getGeneInt(g, -10, -8, SHOES_GENES_COUNT)
}

func getShoesGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getShoesGene(g)
	trait := configService.Footwear[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Footwear,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getShoesGenePath(g string) string {
	gene := getShoesGene(g)
	return Gene(gene).toPath()
}

func getTorsoGene(g string) int {
	return getGeneInt(g, -8, -6, TORSO_GENES_COUNT)
}

func getTorsoGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getTorsoGene(g)
	trait := configService.Torso[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Torso,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getTorsoGenePath(g string) string {
	gene := getTorsoGene(g)
	return Gene(gene).toPath()
}

func getPantsGene(g string) int {
	return getGeneInt(g, -6, -4, PANTS_GENES_COUNT)
}

func getPantsGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getPantsGene(g)
	trait := configService.Pants[gene]
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Pants,
		Value:     trait.Name,
		Sets:      trait.Sets,
	}
}

func getPantsGenePath(g string) string {
	gene := getPantsGene(g)
	return Gene(gene).toPath()
}

func getBackgroundGene(g string) int {
	return getGeneInt(g, -4, -2, BACKGROUND_GENE_COUNT)
}

func getBackgroundGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBackgroundGene(g)
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Background,
		Value:     configService.Background[gene],
	}
}

func getBackgroundGenePath(g string) string {
	gene := getBackgroundGene(g)
	return Gene(gene).toPath()
}

func getBaseGene(g string) int {
	return getGeneInt(g, -2, 0, BASE_GENES_COUNT)
}

func getBaseGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
	gene := getBaseGene(g)
	return structs.Attribute{
		TraitType: constants.MorphAttriutes.Character,
		Value:     configService.Character[gene],
	}
}

func getBaseGenePath(g string) string {
	gene := getBaseGene(g)
	return Gene(gene).toPath()
}

func (g *Genome) name(configService *structs.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getBaseGene(gStr)
	return fmt.Sprintf("%v #%v", configService.Character[gene], tokenId)
}

func (g *Genome) description(configService *structs.ConfigService, tokenId string) string {
	gStr := string(*g)
	gene := getBaseGene(gStr)
	return fmt.Sprintf("The %v named %v #%v is a citizen of the Polymorph Universe and has a unique genetic code! You can scramble your Polymorph at anytime.", configService.Type[gene], configService.Character[gene], tokenId)
}

func (g *Genome) genes() []string {
	gStr := string(*g)

	res := make([]string, 0, GENES_COUNT)

	res = append(res, getWeaponRightGenePath(gStr))
	res = append(res, getWeaponLeftGenePath(gStr))
	res = append(res, getHeadGenePath(gStr))
	res = append(res, getEyewearGenePath(gStr))
	res = append(res, getTorsoGenePath(gStr))
	res = append(res, getPantsGenePath(gStr))
	res = append(res, getShoesGenePath(gStr))
	res = append(res, getBaseGenePath(gStr))
	res = append(res, getBackgroundGenePath(gStr))

	return res
}

func (g *Genome) attributes(configService *structs.ConfigService) []structs.Attribute {
	gStr := string(*g)

	res := make([]structs.Attribute, 0, GENES_COUNT)
	res = append(res, getBaseGeneAttribute(gStr, configService))
	res = append(res, getShoesGeneAttribute(gStr, configService))
	res = append(res, getPantsGeneAttribute(gStr, configService))
	res = append(res, getTorsoGeneAttribute(gStr, configService))
	res = append(res, getEyewearGeneAttribute(gStr, configService))
	res = append(res, getHeadGeneAttribute(gStr, configService))
	res = append(res, getWeaponLeftGeneAttribute(gStr, configService))
	res = append(res, getWeaponRightGeneAttribute(gStr, configService))
	res = append(res, getBackgroundGeneAttribute(gStr, configService))
	return res
}

func (g *Genome) Metadata(tokenId string, configService *structs.ConfigService) structs.Metadata {
	var m structs.Metadata
	m.Attributes = g.attributes(configService)
	m.Name = g.name(configService, tokenId)
	m.Description = g.description(configService, tokenId)
	m.ExternalUrl = fmt.Sprintf("%s%s", EXTERNAL_URL, tokenId)
	return m
}
