package handler

const(

	// True Names
	AmarManohar = "Amar Manohar"
	PremBhopali = "Prem Bhopali"
	RaveenBajaj = "Raveen Bajaj"
	Karishma = "Karishma"
	RamGopalBajaj = "Ram Gopal Bajaj"
	ShyamGopalBajaj = "Shyam Gopal Bajaj"
	CrimeMasterGogo = "Crime Master Gogo"
	Robert = "Robert"
	VinodBhalla = "Vonod Bhalla"
	BankelalBhopali = "Bankelal Bhopali"
	MurliManohar = "Murli Manohar"
	Johnny = "Johnny"
	AnandAkela = "Anand Akela"
	InspectorPandey = "Inspector Pandey"
	Sevaram = "Sevaram"
	JuhiChawla = "Juhi Chawla"
	Govinda = "Govinda"

	// Aliases
	AmarManoharAlias = "Amar Singh"
	PremBhopaliAlias = "Prem Khurrana"
	RaveenaBajajAlias = "Karishma"
	KarishmaAlias = "Raveena"
	RamGopalBajajAlias = "Teja"
)

// GetAllNames gets all character names of the movie.
func GetAllNames()[]string{
	return []string{
		AmarManohar,
		PremBhopali,
		RaveenBajaj,
		Karishma,
		RamGopalBajaj,
		ShyamGopalBajaj,
		CrimeMasterGogo,
		Robert,
		VinodBhalla,
		BankelalBhopali,
		MurliManohar,
		Johnny,
		AnandAkela,
		InspectorPandey,
		Sevaram,
		JuhiChawla,
		Govinda,
	}
}

// GetNameVsAliasMap gets a map of character true names and their alias names.
func GetNameVsAliasMap() map[string]string{
	return map[string]string{
		AmarManohar: AmarManoharAlias,
		PremBhopali: PremBhopaliAlias,
		RaveenBajaj: RaveenaBajajAlias,
		Karishma: KarishmaAlias,
		RamGopalBajaj: RamGopalBajajAlias,
	}
}
