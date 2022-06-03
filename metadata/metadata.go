package metadata

import (
	"fmt"
	"strconv"
	"strings"

	"rarity-backend/constants"
	"rarity-backend/structs"
)

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
	return getGeneInt(g, constants.LEFT_HAND_GENE_START_IDX, constants.LEFT_HAND_GENE_END_IDX, constants.WEAPON_LEFT_GENES_COUNT)
}

func GetWeaponLeftGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.RIGHT_HAND_GENE_START_IDX, constants.RIGHT_HAND_GENE_END_IDX, constants.WEAPON_RIGHT_GENES_COUNT)
}

func GetWeaponRightGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.HEAD_GENE_START_IDX, constants.HEAD_GENE_END_IDX, constants.HEAD_GENES_COUNT)
}

func GetHeadGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, -12, -10, constants.EYEWEAR_GENES_COUNT)
}

func GetEyewearGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.SHOES_GENE_START_IDX, constants.SHOES_GENE_END_IDX, constants.SHOES_GENES_COUNT)
}

func GetShoesGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.TORSO_GENE_START_IDX, constants.TORSO_GENE_END_IDX, constants.TORSO_GENES_COUNT)
}

func GetTorsoGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.PANTS_GENE_START_IDX, constants.PANTS_GENE_END_IDX, constants.PANTS_GENES_COUNT)
}

func GetPantsGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.BACKGROUND_GENE_START_IDX, constants.BACKGROUND_GENE_END_IDX, constants.BACKGROUND_GENE_COUNT)
}

func GetBackgroundGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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
	return getGeneInt(g, constants.CHARACTER_GENE_START_IDX, constants.CHARACTER_GENE_END_IDX, constants.BASE_GENES_COUNT)
}

func GetBaseGeneAttribute(g string, configService *structs.ConfigService) structs.Attribute {
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

func (g *Genome) Genes() []string {
	gStr := string(*g)

	res := make([]string, 0, constants.GENES_COUNT)

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

	res := make([]structs.Attribute, 0, constants.GENES_COUNT)
	res = append(res, GetBaseGeneAttribute(gStr, configService))
	res = append(res, GetShoesGeneAttribute(gStr, configService))
	res = append(res, GetPantsGeneAttribute(gStr, configService))
	res = append(res, GetTorsoGeneAttribute(gStr, configService))
	res = append(res, GetEyewearGeneAttribute(gStr, configService))
	res = append(res, GetHeadGeneAttribute(gStr, configService))
	res = append(res, GetWeaponLeftGeneAttribute(gStr, configService))
	res = append(res, GetWeaponRightGeneAttribute(gStr, configService))
	res = append(res, GetBackgroundGeneAttribute(gStr, configService))
	return res
}

func (g *Genome) Metadata(tokenId string, configService *structs.ConfigService) structs.Metadata {
	var m structs.Metadata
	m.Attributes = g.attributes(configService)
	m.Name = g.name(configService, tokenId)
	m.Description = g.description(configService, tokenId)
	m.ExternalUrl = fmt.Sprintf("%s%s", constants.EXTERNAL_URL, tokenId)

	genes := g.Genes()

	imageUrl := strings.Builder{}
	imageUrl.WriteString(constants.POLYMORPH_IMAGE_URL)

	imageUrl3D := strings.Builder{}
	imageUrl3D.WriteString(constants.POLYMORPH_IMAGE_URL_3D)

	for _, gene := range genes {
		imageUrl.WriteString(gene)
		imageUrl3D.WriteString(gene)
	}

	imageUrl.WriteString(".jpg")
	imageUrl3D.WriteString(".jpg")

	m.Image = imageUrl.String()
	m.Image3D = imageUrl3D.String()

	return m
}
