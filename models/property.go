package models

import (
	"github.com/jinzhu/gorm"
	"strings"
)

var content = ` "p" , "l" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A1" , "Role_Has_Domain" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#NHC0" , "code" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P106" , "Semantic_Type" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P107" , "Display_Name" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P108" , "Preferred_Name" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P90" , "FULL_SYN" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P97" , "DEFINITION" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P383" , "term-group" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P384" , "term-source" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P378" , "def-source" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A10" , "Has_CDRH_Parent" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A11" , "Has_NICHD_Parent" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A12" , "Has_Data_Element" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A13" , "Related_To_Genetic_Biomarker" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A14" , "Neoplasm_Has_Special_Category" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A15" , "Has_CTCAE_5_Parent" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A2" , "Role_Has_Range" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A3" , "Role_Has_Parent" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A4" , "Qualifier_Applies_To" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A5" , "Has_Salt_Form" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P98" , "DesignNote" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A6" , "Has_Free_Acid_Or_Base_Form" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A7" , "Has_Target" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A8" , "Concept_In_Subset" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#A9" , "Is_Related_To_Endogenous_Product" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#Definition_Review_Date" , "Definition_Review_Date" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#Definition_Reviewer_Name" , "Definition_Reviewer_Name" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#NHC4" , "Split_From" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P100" , "OMIM_Number" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P101" , "Homologous_Gene" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P102" , "GenBank_Accession_Number" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P167" , "Image_Link" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P171" , "PubMedID_Primary_Reference" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P175" , "NSC_Code" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P200" , "OLD_PARENT" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P201" , "OLD_CHILD" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P203" , "OLD_KIND" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P204" , "OLD_ROLE" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P205" , "OLD_STATE" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P207" , "UMLS_CUI" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P208" , "NCI_META_CUI" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P210" , "CAS_Registry" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P211" , "GO_Annotation" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P215" , "KEGG_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P216" , "BioCarta_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P302" , "Accepted_Therapeutic_Use_For" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P310" , "Concept_Status" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P315" , "SNP_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P316" , "Relative_Enzyme_Activity" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P317" , "FDA_Table" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P319" , "FDA_UNII_Code" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P320" , "OID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P321" , "EntrezGene_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P322" , "Contributing_Source" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P325" , "ALT_DEFINITION" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P329" , "PDQ_Open_Trial_Search_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P330" , "PDQ_Closed_Trial_Search_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P331" , "NCBI_Taxon_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P332" , "MGI_Accession_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P333" , "Use_For" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P334" , "ICD-O-3_Code" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P350" , "Chemical_Formula" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P351" , "US_Recommended_Intake" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P352" , "Tolerable_Level" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P353" , "INFOODS" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P354" , "USDA_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P355" , "Unit" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P356" , "Essential_Amino_Acid" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P357" , "Essential_Fatty_Acid" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P358" , "Nutrient" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P359" , "Micronutrient" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P360" , "Macronutrient" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P361" , "Extensible_List" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P362" , "miRBase_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P363" , "Neoplastic_Status" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P364" , "OLD_ASSOCIATION" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P365" , "OLD_SOURCE_ASSOCIATION" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P366" , "Legacy_Concept_Name" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P367" , "PID_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P368" , "CHEBI_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P369" , "HGNC_ID" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P371" , "NICHD_Hierarchy_Term" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P372" , "Publish_Value_Set" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P376" , "Term_Browser_Value_Set_Description" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P377" , "def-definition" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P381" , "attr" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P382" , "term-name" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P385" , "source-code" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P386" , "subsource-name" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P387" , "go-id" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P388" , "go-term" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P389" , "go-evi" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P390" , "go-source" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P391" , "source-date" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P392" , "Target_Term" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P393" , "Relationship_to_Target" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P394" , "Target_Term_Type" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P395" , "Target_Code" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P396" , "Target_Terminology" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P92" , "Subsource" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P93" , "Swiss_Prot" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#P96" , "Gene_Encodes_Product" ,
 "http://ncicb.nci.nih.gov/xml/owl/EVS/Thesaurus.owl#xref-source" , "xRef Source" ,
 "http://www.geneontology.org/formats/oboInOwl#hasDbXref" , "xRef" ,
 "http://semanticscience.org/resource/SIO_000628" , "refers-to" ,
 "http://semanticscience.org/resource/SIO_000223" , "has-property" ,
 "http://semanticscience.org/resource/SIO_010078" , "encodes" ,
 "http://semanticscience.org/resource/SIO_000216" , "has-measurement-value" ,
 "http://semanticscience.org/resource/SIO_000061" , "is-located-in" ,
 "http://semanticscience.org/resource/SIO_000791" , "sequence-start-position" ,
 "http://semanticscience.org/resource/SIO_000300" , "has-value" ,
 "http://semanticscience.org/resource/SIO_000253" , "has-source" ,
 "http://semanticscience.org/resource/SIO_000772" , "has-evidence" ,
 "http://www.sequenceontology.org/miso/current_svn/term/SO:associated_with" , "associated_with" ,
`

type Property struct {
	gorm.Model
	ConceptIdentifier string `json:"identifier",gorm:"type:varchar(12),unique_index"`
	URI               string `json:"uri",gorm:"type:varchar(256),unique_index"`
	Label             string `json:"label",gorm:"type:varchar(128)"`
}

type Record struct {
	ConceptID string
	URI string
	Label string
}

func processProperties() []Record {
	var records []Record
	all := strings.Split(content, ",")
	var fore, aft string
	for i := range all{
		elm := strings.Trim(all[i], " \"\n")
		if i % 2 == 0 {
			fore = elm
		} else {
			aft = elm
			if fore == "p"{
				continue
			}
			var conceptID string
			splitURI := strings.Split(fore, "#")
			if len(splitURI) == 2{
				conceptID = splitURI[1]
			} else {
				conceptID = aft
			}
			//conceptID := strings.Split(fore, "#")[1]
			records = append(records, Record{ConceptID: conceptID, URI:fore, Label:aft})
		}
	}
	return records
}

func (p *Property) InitDatabase(db gorm.DB){
	records := processProperties()
	for i := range records{
		record := records[i]
		p.GetOrCreate(record.ConceptID, record.URI, record.Label, db)
	}
}

func (p *Property) GetOrCreate(conceptID, uri, label string, db gorm.DB) (property Property) {
	db.FirstOrCreate(&property, Property{ConceptIdentifier: conceptID, URI: uri, Label: label})
	return
}

func GetByConceptID(conceptID string, db gorm.DB) (property Property){
	db.Where("ConceptIdentifier = ?", conceptID).First(&property)
	return
}

func GetByURI(uri string, db gorm.DB) (property Property){
	db.Where("URI = ?", uri).First(&property)
	return
}

func GetByLabel(label string, db gorm.DB) (property Property){
	db.Where("Label = ?", label).First(&property)
	return
}

