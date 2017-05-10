package electiondata

// Winning parties only
const (
	ALL   = "ALL"
	CON   = "CON"
	DUP   = "DUP"
	GRN   = "GRN"
	IND   = "IND"
	LAB   = "LAB"
	LABCO = "LABCO"
	LIB   = "LIB"
	PC    = "PC"
	SDLP  = "SDLP"
	SF    = "SF"
	SNP   = "SNP"
	SPK   = "SPK"
	UKIP  = "UKIP"
	UUP   = "UUP"
)

func (r *Result) SanitizeParty() {
	// Will set the party if is matches one of the following, otherwise will
	// leave it alone.
	//
	// This is the 2015 set of parties
	switch r.Party {
	case `30-50 Coalition`:
		r.Party = `30-50 Coalition`
	case `AB`:
		r.Party = `Above and Beyond`
	case `Above and Beyond`:
		r.Party = `Above and Beyond`
	case `AD`:
		r.Party = `Apolitical Democrats`
	case `AGS`:
		r.Party = `Alliance for Green Socialism`
	case `Al-Zebabist Nation of Ooog`:
		r.Party = `Al-Zebabist Nation of Ooog`
	case `ALC`:
		r.Party = `Alter Change`
	case `All People's Party`:
		r.Party = `All People's Party`
	case `Alter Change`:
		r.Party = `Alter Change`
	case `Animal Welfare Party`:
		r.Party = `Animal Welfare Party`
	case `ANWP`:
		r.Party = `Animal Welfare Party`
	case `APNI`:
		r.Party = `Alliance Party`
	case `Apolitical Dem`:
		r.Party = `Apolitical Democrats`
	case `APPL`:
		r.Party = `All People's Party`
	case `AZNO`:
		r.Party = `Al-Zebabist Nation of Ooog`
	case `BD`:
		r.Party = `British Democrats`
	case `Beer, Baccy and Scratchings`:
		r.Party = `Beer, Baccy and Scratchings`
	case `BEER`:
		r.Party = `Beer, Baccy and Scratchings`
	case `BI`:
		r.Party = `British Independents`
	case `BIA`:
		r.Party = `Bournemouth Independent Alliance`
	case `Birthday Party`:
		r.Party = `Birthday Party`
	case `BNP`:
		r.Party = `British National Party`
	case `Bournemouth Independent Alliance`:
		r.Party = `Bournemouth Independent Alliance`
	case `British Democrats`:
		r.Party = `British Democrats`
	case `British Independents`:
		r.Party = `British Independents`
	case `C`:
		r.Party = `Conservative Party`
	case `CAM`:
		r.Party = `Campaign`
	case `Campaign`:
		r.Party = `Campaign`
	case `Cannabis is Safer than Alcohol`:
		r.Party = `Cannabis is Safer than Alcohol`
	case `CFU`:
		r.Party = `Citizens for Uttlesford`
	case `Ch P`:
		r.Party = `Christian Party`
	case `Children of the Atom`:
		r.Party = `Children of the Atom`
	case `CHM`:
		r.Party = `Christian Movement for Great Britain`
	case `CHP`:
		r.Party = `Christian Party`
	case `Christian Move`:
		r.Party = `Christian Movement for Great Britain`
	case `Christian Party`:
		r.Party = `Christian Party`
	case `Christian Peoples Alliance`:
		r.Party = `Christian Peoples Alliance`
	case `Citizens for Uttlesford`:
		r.Party = `Citizens for Uttlesford`
	case `Class War`:
		r.Party = `Class War`
	case `CMSP`:
		r.Party = `Common Sense Party`
	case `CMU`:
		r.Party = `Communities United Party`
	case `COML`:
		r.Party = `Communist League`
	case `Comm. League`:
		r.Party = `Communist League`
	case `Common Sense Party`:
		r.Party = `Common Sense Party`
	case `Communist`:
		r.Party = `Communist Party`
	case `Communities United Party`:
		r.Party = `Communities United Party`
	case `COMP`:
		r.Party = `Communist Party`
	case `CON`:
		r.Party = `Conservative Party`
	case `Con`:
		r.Party = `Conservative Party`
	case `COTA`:
		r.Party = `Children of the Atom`
	case `CPA`:
		r.Party = `Christian Peoples Alliance`
	case `CSA`:
		r.Party = `Cannabis is Safer than Alcohol`
	case `CW`:
		r.Party = `Class War`
	case `DD`:
		r.Party = `Digital Democracy`
	case `DEM`:
		r.Party = `Democratic Party`
	case `Democratic Party`:
		r.Party = `Democratic Party`
	case `Democratic Reform Party`:
		r.Party = `Democratic Reform Party`
	case `Digital Democracy`:
		r.Party = `Digital Democracy`
	case `DRP`:
		r.Party = `Democratic Reform Party`
	case `DUP`:
		r.Party = `Democratic Unionist Party`
	case `Eccentric Party of Great Britain`:
		r.Party = `Eccentric Party of Great Britain`
	case `ED`:
		r.Party = `English Democrats`
	case `Eng Dem`:
		r.Party = `English Democrats`
	case `Eng Ind`:
		r.Party = `English Independence Party`
	case `EPGB`:
		r.Party = `Eccentric Party of Great Britain`
	case `EUR`:
		r.Party = `Europeans Party`
	case `Europeans Party`:
		r.Party = `Europeans Party`
	case `Evolution Party`:
		r.Party = `Evolution Party`
	case `FPTP`:
		r.Party = `Free Public Transport Party`
	case `Free Public Transport Party`:
		r.Party = `Free Public Transport Party`
	case `Free United Kingdom Party`:
		r.Party = `Free United Kingdom Party`
	case `FREE`:
		r.Party = `Free United Kingdom Party`
	case `GGG`:
		r.Party = `Guildford Greenbelt Group`
	case `Green Soc`:
		r.Party = `Alliance for Green Socialism`
	case `Green`:
		r.Party = `Green Party`
	case `GRN`:
		r.Party = `Green Party`
	case `Guildford Greenbelt Group`:
		r.Party = `Guildford Greenbelt Group`
	case `Hoi Polloi`:
		r.Party = `Hoi Polloi`
	case `HOPO`:
		r.Party = `Hoi Polloi`
	case `HUM`:
		r.Party = `Humanity`
	case `Humanity`:
		r.Party = `Humanity`
	case `ICHC`:
		r.Party = `Independent Community and Health Concern`
	case `IFB`:
		r.Party = `Independents for Bristol`
	case `IFE`:
		r.Party = `Independence from Europe`
	case `Ind`:
		r.Party = `Independent`
	case `Independence from Europe`:
		r.Party = `Independence from Europe`
	case `Independent Political Alliance Party`:
		r.Party = `Independent Political Alliance Party`
	case `Independents for Bristol`:
		r.Party = `Independents for Bristol`
	case `IPAP`:
		r.Party = `Independent Political Alliance Party`
	case `Islam Zinda Baad Platform`:
		r.Party = `Islam Zinda Baad Platform`
	case `IZB`:
		r.Party = `Islam Zinda Baad Platform`
	case `JAC`:
		r.Party = `Justice & Anti-Corruption Party`
	case `JMB`:
		r.Party = `Justice for Men and Boys`
	case `Justice Anti-Corruption`:
		r.Party = `Justice & Anti-Corruption Party`
	case `Justice for Men and Boys`:
		r.Party = `Justice for Men and Boys`
	case `KIR`:
		r.Party = `Let's Keep it Real`
	case `LAB`:
		r.Party = `Labour Party`
	case `Lab`:
		r.Party = `Labour Party`
	case `Land Party`:
		r.Party = `Land Party`
	case `LCBP`:
		r.Party = `Let Every Child Have Both Parents`
	case `LD`:
		r.Party = `Liberal Democrat`
	case `Left Unity`:
		r.Party = `Left Unity`
	case `Let Every Child Have Both Parents`:
		r.Party = `Let Every Child Have Both Parents`
	case `Let's Keep it Real`:
		r.Party = `Let's Keep it Real`
	case `LHSC`:
		r.Party = `Locally Informed Health and Social Care`
	case `LIB`:
		r.Party = `Liberal`
	case `Liberal`:
		r.Party = `Liberal`
	case `Liberty GB`:
		r.Party = `Liberty GB`
	case `LIGB`:
		r.Party = `Liberty GB`
	case `Lincolnshire Independents`:
		r.Party = `Lincolnshire Independents`
	case `LNIN`:
		r.Party = `Lincolnshire Independents`
	case `Locally Informed Health and Social Care`:
		r.Party = `Locally Informed Health and Social Care`
	case `Loony`:
		r.Party = `Monster Raving Loony Party`
	case `LP`:
		r.Party = `Land Party`
	case `LU`:
		r.Party = `Left Unity`
	case `MAD`:
		r.Party = `Movement for Active Democracy`
	case `Magna Carta`:
		r.Party = `Magna Carta Party`
	case `Magna Carter Conservation Party`:
		r.Party = `Magna Carter Conservation Party`
	case `MAIN`:
		r.Party = `Mainstream`
	case `Mainstream`:
		r.Party = `Mainstream`
	case `MAIP`:
		r.Party = `Manston Airport Independent Party`
	case `Manston Airport Independent Party`:
		r.Party = `Manston Airport Independent Party`
	case `MCCP`:
		r.Party = `Magna Carter Conservation Party`
	case `MCP`:
		r.Party = `Magna Carta Party`
	case `Meb Ker`:
		r.Party = `Mebyon Kernow`
	case `MK`:
		r.Party = `Mebyon Kernow`
	case `Movement for Active Democracy`:
		r.Party = `Movement for Active Democracy`
	case `MRLP`:
		r.Party = `Monster Raving Loony Party`
	case `National Front`:
		r.Party = `National Front`
	case `National Health Action Party`:
		r.Party = `National Health Action Party`
	case `National Liberal Party`:
		r.Party = `National Liberal Party`
	case `NEP`:
		r.Party = `North East Party`
	case `New Independent Centralists`:
		r.Party = `New Independent Centralists`
	case `New Society of Worth`:
		r.Party = `New Society of Worth`
	case `NF`:
		r.Party = `National Front`
	case `NHSA`:
		r.Party = `National Health Action Party`
	case `NIC`:
		r.Party = `New Independent Centralists`
	case `NLP`:
		r.Party = `National Liberal Party`
	case `NORP`:
		r.Party = `Northern Party`
	case `North East Party`:
		r.Party = `North East Party`
	case `Northern Party`:
		r.Party = `Northern Party`
	case `NSOW`:
		r.Party = `New Society of Worth`
	case `Party for a United Thanet`:
		r.Party = `Party for a United Thanet`
	case `PATR`:
		r.Party = `Patria`
	case `Patria`:
		r.Party = `Patria`
	case `Patriotic Socialist Party`:
		r.Party = `Patriotic Socialist Party`
	case `PBP`:
		r.Party = `People Before Profit`
	case `PC`:
		r.Party = `Plaid Cymru`
	case `PCF`:
		r.Party = `Putting Croydon First`
	case `PDP`:
		r.Party = `UK Progressive Democracy Party`
	case `Peace`:
		r.Party = `Peace Party, non-violence, justice, environment`
	case `People before Profit`:
		r.Party = `People Before Profit`
	case `People First`:
		r.Party = `People First`
	case `PF`:
		r.Party = `People First`
	case `Pilgrim Party`:
		r.Party = `Pilgrim Party`
	case `PILP`:
		r.Party = `Pilgrim Party`
	case `Pirate`:
		r.Party = `Pirate Party UK`
	case `PIRT`:
		r.Party = `Pirate Party UK`
	case `PLP`:
		r.Party = `Pluralist Party`
	case `Pluralist Party`:
		r.Party = `Pluralist Party`
	case `POOL`:
		r.Party = `Poole People`
	case `Poole People`:
		r.Party = `Poole People`
	case `POPP`:
		r.Party = `Principles of Politics Party`
	case `Population Party UK`:
		r.Party = `Population Party UK`
	case `PPNV`:
		r.Party = `Peace Party, non-violence, justice, environment`
	case `PPUK`:
		r.Party = `Population Party UK`
	case `Principles of Politics Party`:
		r.Party = `Principles of Politics Party`
	case `PSP`:
		r.Party = `Patriotic Socialist Party`
	case `PUT`:
		r.Party = `Party for a United Thanet`
	case `Putting Croydon First`:
		r.Party = `Putting Croydon First`
	case `RBD`:
		r.Party = `Rebooting Democracy`
	case `Realists' Party`:
		r.Party = `Realists' Party`
	case `Rebooting Democracy`:
		r.Party = `Rebooting Democracy`
	case `Red Flag ? Anti-Corruption`:
		r.Party = `Red Flag ? Anti-Corruption`
	case `Reduce VAT in Sport`:
		r.Party = `Reduce VAT in Sport`
	case `Republican Socialist`:
		r.Party = `Republican Socialist`
	case `RES`:
		r.Party = `Respect Party`
	case `Respect`:
		r.Party = `Respect Party`
	case `RFAC`:
		r.Party = `Red Flag ? Anti-Corruption`
	case `RFP`:
		r.Party = `Rochdale First`
	case `Rochdale First`:
		r.Party = `Rochdale First`
	case `Roman Party`:
		r.Party = `Roman Party`
	case `RS`:
		r.Party = `Republican Socialist`
	case `RVIS`:
		r.Party = `Reduce VAT in Sport`
	case `SACL`:
		r.Party = `Scotland Against Crooked Lawyers`
	case `SDLP`:
		r.Party = `Social Democratic & Labour Party`
	case `SDP`:
		r.Party = `Social Democratic Party`
	case `SECA`:
		r.Party = `Stop Emotional Child Abuse, Vote Elmo`
	case `SEP`:
		r.Party = `Socialist Equality Party`
	case `SF`:
		r.Party = `Sinn Fein`
	case `SHP`:
		r.Party = `Your Vote Could Save Our Hospital`
	case `SLP`:
		r.Party = `Socialist Labour Party`
	case `SN`:
		r.Party = `Something New`
	case `SNP`:
		r.Party = `Scottish National Party`
	case `SOC`:
		r.Party = `Socialist Party`
	case `Social Democratic`:
		r.Party = `Social Democratic Party`
	case `Socialist Equality`:
		r.Party = `Socialist Equality Party`
	case `Socialist Labour`:
		r.Party = `Socialist Labour Party`
	case `Socialist`:
		r.Party = `Socialist Party`
	case `Something New`:
		r.Party = `Something New`
	case `Southport Party`:
		r.Party = `Southport Party`
	case `SPP`:
		r.Party = `Sustainable Population Party`
	case `Spk`:
		r.Party = `Speaker`
	case `Speaker`:
		r.Party = `Speaker`
	case `SSP`:
		r.Party = `Scottish Socialist Party`
	case `Stop Emotional Child Abuse, Vote Elmo`:
		r.Party = `Stop Emotional Child Abuse, Vote Elmo`
	case `Sustainable Population Party`:
		r.Party = `Sustainable Population Party`
	case `T35C`:
		r.Party = `30-50 Coalition`
	case `TBP`:
		r.Party = `Birthday Party`
	case `TEP`:
		r.Party = `Evolution Party`
	case `Trade Unionist`:
		r.Party = `Trade Unionist and Socialist Coalition`
	case `Traditional Unionist Voice`:
		r.Party = `Traditional Unionist Voice`
	case `TRP`:
		r.Party = `Realists' Party`
	case `TRPA`:
		r.Party = `Roman Party`
	case `TSP`:
		r.Party = `Southport Party`
	case `TUP`:
		r.Party = `Ubuntu Party`
	case `TUSC`:
		r.Party = `Trade Unionist and Socialist Coalition`
	case `TUV`:
		r.Party = `Traditional Unionist Voice`
	case `U Party`:
		r.Party = `U Party`
	case `Ubuntu Party`:
		r.Party = `Ubuntu Party`
	case `UCUNF`:
		r.Party = `Ulster Conservatives and Unionists - New Force`
	case `UK Progressive Democracy Party`:
		r.Party = `UK Progressive Democracy Party`
	case `UKIP`:
		r.Party = `United Kingdom Independence Party`
	case `UP`:
		r.Party = `U Party`
	case `UUP`:
		r.Party = `Ulster Unionist Party`
	case `VAP`:
		r.Party = `Vapers in Power`
	case `Vapers in Power`:
		r.Party = `Vapers in Power`
	case `War Veteran's Pro-Traditional Family Party`:
		r.Party = `War Veteran's Pro-Traditional Family Party`
	case `WARP`:
		r.Party = `We Are The Reality Party`
	case `We Are The Reality Party`:
		r.Party = `We Are The Reality Party`
	case `Wessex Regionalist`:
		r.Party = `Wessex Regionalist`
	case `WGP`:
		r.Party = `Wigan Independent`
	case `Whig Party`:
		r.Party = `Whig Party`
	case `WHIG`:
		r.Party = `Whig Party`
	case `Wigan Independent`:
		r.Party = `Wigan Independent`
	case `Workers' Revolution`:
		r.Party = `Workers' Revolutionary Party`
	case `Workers`:
		r.Party = `Workers' Party`
	case `World Peace Through Song`:
		r.Party = `World Peace Through Song`
	case `WP`:
		r.Party = `Workers' Party`
	case `WPTS`:
		r.Party = `World Peace Through Song`
	case `WR`:
		r.Party = `Wessex Regionalist`
	case `WRP`:
		r.Party = `Workers' Revolutionary Party`
	case `WVP`:
		r.Party = `War Veteran's Pro-Traditional Family Party`
	case `YF`:
		r.Party = `Yorkshire First`
	case `Yorkshire First Party`:
		r.Party = `Yorkshire First`
	case `Young People's Party`:
		r.Party = `Young People's Party`
	case `Your Vote Could Save Our Hospital`:
		r.Party = `Your Vote Could Save Our Hospital`
	case `YPP`:
		r.Party = `Young People's Party`
	}
}
