package helpers

import (
	"log"
	"rarity-backend/constants"
	"rarity-backend/metadata"
	"rarity-backend/structs"
)

// GetAttribute calcualtes the old and new attributes.
//
// This is later used to create a history snapshot of the polymorph
func GetAttribute(newGene string, oldGene string, geneIdx int, configService *structs.ConfigService) (structs.Attribute, structs.Attribute) {
	geneIdx = -geneIdx
	newAttribute := structs.Attribute{}
	oldAttribute := structs.Attribute{}

	if geneIdx > constants.CHARACTER_GENE_START_IDX && geneIdx <= constants.CHARACTER_GENE_END_IDX {
		newAttribute = metadata.GetBaseGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBaseGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.BACKGROUND_GENE_START_IDX && geneIdx <= constants.BACKGROUND_GENE_END_IDX {
		newAttribute = metadata.GetBackgroundGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetBackgroundGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.PANTS_GENE_START_IDX && geneIdx <= constants.PANTS_GENE_END_IDX {
		newAttribute = metadata.GetPantsGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetPantsGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.TORSO_GENE_START_IDX && geneIdx <= constants.TORSO_GENE_END_IDX {
		newAttribute = metadata.GetTorsoGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetTorsoGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.SHOES_GENE_START_IDX && geneIdx <= constants.SHOES_GENE_END_IDX {
		newAttribute = metadata.GetShoesGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetShoesGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.EYEWEAR_GENE_START_IDX && geneIdx <= constants.EYEWEAR_GENE_END_IDX {
		newAttribute = metadata.GetEyewearGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetEyewearGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.HEAD_GENE_START_IDX && geneIdx <= constants.HEAD_GENE_END_IDX {
		newAttribute = metadata.GetHeadGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetHeadGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.RIGHT_HAND_GENE_START_IDX && geneIdx <= constants.RIGHT_HAND_GENE_END_IDX {
		newAttribute = metadata.GetWeaponRightGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetWeaponRightGeneAttribute(oldGene, configService)
	} else if geneIdx > constants.LEFT_HAND_GENE_START_IDX && geneIdx <= constants.LEFT_HAND_GENE_END_IDX {
		newAttribute = metadata.GetWeaponLeftGeneAttribute(newGene, configService)
		oldAttribute = metadata.GetWeaponLeftGeneAttribute(oldGene, configService)
	} else {
		log.Println("Neshto se precaka da mu eba mamata ..")
	}

	return newAttribute, oldAttribute
}
