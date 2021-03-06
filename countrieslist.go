package gocountrydetect

var langs map[string]Countries = map[string]Countries{
	"en": Countries{
		Country{Code: "AFG", Name: "Afghanistan"},
		Country{Code: "ALA", Name: "Åland Islands"},
		Country{Code: "ALB", Name: "Albania"},
		Country{Code: "DZA", Name: "Algeria"},
		Country{Code: "ASM", Name: "American Samoa"},
		Country{Code: "AND", Name: "Andorra"},
		Country{Code: "AGO", Name: "Angola"},
		Country{Code: "AIA", Name: "Anguilla"},
		Country{Code: "ATA", Name: "Antarctica"},
		Country{Code: "ATG", Name: "Antigua and Barbuda"},
		Country{Code: "ARG", Name: "Argentina"},
		Country{Code: "ARM", Name: "Armenia"},
		Country{Code: "ABW", Name: "Aruba"},
		Country{Code: "AUS", Name: "Australia"},
		Country{Code: "AUT", Name: "Austria"},
		Country{Code: "AZE", Name: "Azerbaijan"},
		Country{Code: "BHS", Name: "Bahamas"},
		Country{Code: "BHR", Name: "Bahrain"},
		Country{Code: "BGD", Name: "Bangladesh"},
		Country{Code: "BRB", Name: "Barbados"},
		Country{Code: "BLR", Name: "Belarus"},
		Country{Code: "BEL", Name: "Belgium"},
		Country{Code: "BLZ", Name: "Belize"},
		Country{Code: "BEN", Name: "Benin"},
		Country{Code: "BMU", Name: "Bermuda"},
		Country{Code: "BTN", Name: "Bhutan"},
		Country{Code: "BOL", Name: "Bolivia (Plurinational State of)"},
		Country{Code: "BES", Name: "Bonaire, Sint Eustatius and Saba"},
		Country{Code: "BIH", Name: "Bosnia and Herzegovina"},
		Country{Code: "BWA", Name: "Botswana"},
		Country{Code: "BVT", Name: "Bouvet Island"},
		Country{Code: "BRA", Name: "Brazil"},
		Country{Code: "IOT", Name: "British Indian Ocean Territory"},
		Country{Code: "UMI", Name: "United States Minor Outlying Islands"},
		Country{Code: "VGB", Name: "Virgin Islands (British)"},
		Country{Code: "VIR", Name: "Virgin Islands (U.S.)"},
		Country{Code: "BRN", Name: "Brunei Darussalam"},
		Country{Code: "BGR", Name: "Bulgaria"},
		Country{Code: "BFA", Name: "Burkina Faso"},
		Country{Code: "BDI", Name: "Burundi"},
		Country{Code: "KHM", Name: "Cambodia"},
		Country{Code: "CMR", Name: "Cameroon"},
		Country{Code: "CAN", Name: "Canada"},
		Country{Code: "CPV", Name: "Cabo Verde"},
		Country{Code: "CYM", Name: "Cayman Islands"},
		Country{Code: "CAF", Name: "Central African Republic"},
		Country{Code: "TCD", Name: "Chad"},
		Country{Code: "CHL", Name: "Chile"},
		Country{Code: "CHN", Name: "China"},
		Country{Code: "CXR", Name: "Christmas Island"},
		Country{Code: "CCK", Name: "Cocos (Keeling) Islands"},
		Country{Code: "COL", Name: "Colombia"},
		Country{Code: "COM", Name: "Comoros"},
		Country{Code: "COG", Name: "Congo"},
		Country{Code: "COD", Name: "Congo (Democratic Republic of the)"},
		Country{Code: "COK", Name: "Cook Islands"},
		Country{Code: "CRI", Name: "Costa Rica"},
		Country{Code: "HRV", Name: "Croatia"},
		Country{Code: "CUB", Name: "Cuba"},
		Country{Code: "CUW", Name: "Curaçao"},
		Country{Code: "CYP", Name: "Cyprus"},
		Country{Code: "CZE", Name: "Czech Republic"},
		Country{Code: "DNK", Name: "Denmark"},
		Country{Code: "DJI", Name: "Djibouti"},
		Country{Code: "DMA", Name: "Dominica"},
		Country{Code: "DOM", Name: "Dominican Republic"},
		Country{Code: "ECU", Name: "Ecuador"},
		Country{Code: "EGY", Name: "Egypt"},
		Country{Code: "SLV", Name: "El Salvador"},
		Country{Code: "GNQ", Name: "Equatorial Guinea"},
		Country{Code: "ERI", Name: "Eritrea"},
		Country{Code: "EST", Name: "Estonia"},
		Country{Code: "ETH", Name: "Ethiopia"},
		Country{Code: "FLK", Name: "Falkland Islands (Malvinas)"},
		Country{Code: "FRO", Name: "Faroe Islands"},
		Country{Code: "FJI", Name: "Fiji"},
		Country{Code: "FIN", Name: "Finland"},
		Country{Code: "FRA", Name: "France"},
		Country{Code: "GUF", Name: "French Guiana"},
		Country{Code: "PYF", Name: "French Polynesia"},
		Country{Code: "ATF", Name: "French Southern Territories"},
		Country{Code: "GAB", Name: "Gabon"},
		Country{Code: "GMB", Name: "Gambia"},
		Country{Code: "GEO", Name: "Georgia"},
		Country{Code: "DEU", Name: "Germany"},
		Country{Code: "GHA", Name: "Ghana"},
		Country{Code: "GIB", Name: "Gibraltar"},
		Country{Code: "GRC", Name: "Greece"},
		Country{Code: "GRL", Name: "Greenland"},
		Country{Code: "GRD", Name: "Grenada"},
		Country{Code: "GLP", Name: "Guadeloupe"},
		Country{Code: "GUM", Name: "Guam"},
		Country{Code: "GTM", Name: "Guatemala"},
		Country{Code: "GGY", Name: "Guernsey"},
		Country{Code: "GIN", Name: "Guinea"},
		Country{Code: "GNB", Name: "Guinea-Bissau"},
		Country{Code: "GUY", Name: "Guyana"},
		Country{Code: "HTI", Name: "Haiti"},
		Country{Code: "HMD", Name: "Heard Island and McDonald Islands"},
		Country{Code: "VAT", Name: "Holy See"},
		Country{Code: "HND", Name: "Honduras"},
		Country{Code: "HKG", Name: "Hong Kong"},
		Country{Code: "HUN", Name: "Hungary"},
		Country{Code: "ISL", Name: "Iceland"},
		Country{Code: "IND", Name: "India"},
		Country{Code: "IDN", Name: "Indonesia"},
		Country{Code: "CIV", Name: "Côte d'Ivoire"},
		Country{Code: "IRN", Name: "Iran (Islamic Republic of)"},
		Country{Code: "IRQ", Name: "Iraq"},
		Country{Code: "IRL", Name: "Ireland"},
		Country{Code: "IMN", Name: "Isle of Man"},
		Country{Code: "ISR", Name: "Israel"},
		Country{Code: "ITA", Name: "Italy"},
		Country{Code: "JAM", Name: "Jamaica"},
		Country{Code: "JPN", Name: "Japan"},
		Country{Code: "JEY", Name: "Jersey"},
		Country{Code: "JOR", Name: "Jordan"},
		Country{Code: "KAZ", Name: "Kazakhstan"},
		Country{Code: "KEN", Name: "Kenya"},
		Country{Code: "KIR", Name: "Kiribati"},
		Country{Code: "KWT", Name: "Kuwait"},
		Country{Code: "KGZ", Name: "Kyrgyzstan"},
		Country{Code: "LAO", Name: "Lao People's Democratic Republic"},
		Country{Code: "LVA", Name: "Latvia"},
		Country{Code: "LBN", Name: "Lebanon"},
		Country{Code: "LSO", Name: "Lesotho"},
		Country{Code: "LBR", Name: "Liberia"},
		Country{Code: "LBY", Name: "Libya"},
		Country{Code: "LIE", Name: "Liechtenstein"},
		Country{Code: "LTU", Name: "Lithuania"},
		Country{Code: "LUX", Name: "Luxembourg"},
		Country{Code: "MAC", Name: "Macao"},
		Country{Code: "MKD", Name: "Macedonia (the former Yugoslav Republic of)"},
		Country{Code: "MDG", Name: "Madagascar"},
		Country{Code: "MWI", Name: "Malawi"},
		Country{Code: "MYS", Name: "Malaysia"},
		Country{Code: "MDV", Name: "Maldives"},
		Country{Code: "MLI", Name: "Mali"},
		Country{Code: "MLT", Name: "Malta"},
		Country{Code: "MHL", Name: "Marshall Islands"},
		Country{Code: "MTQ", Name: "Martinique"},
		Country{Code: "MRT", Name: "Mauritania"},
		Country{Code: "MUS", Name: "Mauritius"},
		Country{Code: "MYT", Name: "Mayotte"},
		Country{Code: "MEX", Name: "Mexico"},
		Country{Code: "FSM", Name: "Micronesia (Federated States of)"},
		Country{Code: "MDA", Name: "Moldova (Republic of)"},
		Country{Code: "MCO", Name: "Monaco"},
		Country{Code: "MNG", Name: "Mongolia"},
		Country{Code: "MNE", Name: "Montenegro"},
		Country{Code: "MSR", Name: "Montserrat"},
		Country{Code: "MAR", Name: "Morocco"},
		Country{Code: "MOZ", Name: "Mozambique"},
		Country{Code: "MMR", Name: "Myanmar"},
		Country{Code: "NAM", Name: "Namibia"},
		Country{Code: "NRU", Name: "Nauru"},
		Country{Code: "NPL", Name: "Nepal"},
		Country{Code: "NLD", Name: "Netherlands"},
		Country{Code: "NCL", Name: "New Caledonia"},
		Country{Code: "NZL", Name: "New Zealand"},
		Country{Code: "NIC", Name: "Nicaragua"},
		Country{Code: "NER", Name: "Niger"},
		Country{Code: "NGA", Name: "Nigeria"},
		Country{Code: "NIU", Name: "Niue"},
		Country{Code: "NFK", Name: "Norfolk Island"},
		Country{Code: "PRK", Name: "Korea (Democratic People's Republic of)"},
		Country{Code: "MNP", Name: "Northern Mariana Islands"},
		Country{Code: "NOR", Name: "Norway"},
		Country{Code: "OMN", Name: "Oman"},
		Country{Code: "PAK", Name: "Pakistan"},
		Country{Code: "PLW", Name: "Palau"},
		Country{Code: "PSE", Name: "Palestine, State of"},
		Country{Code: "PAN", Name: "Panama"},
		Country{Code: "PNG", Name: "Papua New Guinea"},
		Country{Code: "PRY", Name: "Paraguay"},
		Country{Code: "PER", Name: "Peru"},
		Country{Code: "PHL", Name: "Philippines"},
		Country{Code: "PCN", Name: "Pitcairn"},
		Country{Code: "POL", Name: "Poland"},
		Country{Code: "PRT", Name: "Portugal"},
		Country{Code: "PRI", Name: "Puerto Rico"},
		Country{Code: "QAT", Name: "Qatar"},
		Country{Code: "KOS", Name: "Republic of Kosovo"},
		Country{Code: "REU", Name: "Réunion"},
		Country{Code: "ROU", Name: "Romania"},
		Country{Code: "RUS", Name: "Russian Federation"},
		Country{Code: "RWA", Name: "Rwanda"},
		Country{Code: "BLM", Name: "Saint Barthélemy"},
		Country{Code: "SHN", Name: "Saint Helena, Ascension and Tristan da Cunha"},
		Country{Code: "KNA", Name: "Saint Kitts and Nevis"},
		Country{Code: "LCA", Name: "Saint Lucia"},
		Country{Code: "MAF", Name: "Saint Martin (French part)"},
		Country{Code: "SPM", Name: "Saint Pierre and Miquelon"},
		Country{Code: "VCT", Name: "Saint Vincent and the Grenadines"},
		Country{Code: "WSM", Name: "Samoa"},
		Country{Code: "SMR", Name: "San Marino"},
		Country{Code: "STP", Name: "Sao Tome and Principe"},
		Country{Code: "SAU", Name: "Saudi Arabia"},
		Country{Code: "SEN", Name: "Senegal"},
		Country{Code: "SRB", Name: "Serbia"},
		Country{Code: "SYC", Name: "Seychelles"},
		Country{Code: "SLE", Name: "Sierra Leone"},
		Country{Code: "SGP", Name: "Singapore"},
		Country{Code: "SXM", Name: "Sint Maarten (Dutch part)"},
		Country{Code: "SVK", Name: "Slovakia"},
		Country{Code: "SVN", Name: "Slovenia"},
		Country{Code: "SLB", Name: "Solomon Islands"},
		Country{Code: "SOM", Name: "Somalia"},
		Country{Code: "ZAF", Name: "South Africa"},
		Country{Code: "SGS", Name: "South Georgia and the South Sandwich Islands"},
		Country{Code: "KOR", Name: "Korea (Republic of)"},
		Country{Code: "SSD", Name: "South Sudan"},
		Country{Code: "ESP", Name: "Spain"},
		Country{Code: "LKA", Name: "Sri Lanka"},
		Country{Code: "SDN", Name: "Sudan"},
		Country{Code: "SUR", Name: "SuriName"},
		Country{Code: "SJM", Name: "Svalbard and Jan Mayen"},
		Country{Code: "SWZ", Name: "Swaziland"},
		Country{Code: "SWE", Name: "Sweden"},
		Country{Code: "CHE", Name: "Switzerland"},
		Country{Code: "SYR", Name: "Syrian Arab Republic"},
		Country{Code: "TWN", Name: "Taiwan"},
		Country{Code: "TJK", Name: "Tajikistan"},
		Country{Code: "TZA", Name: "Tanzania, United Republic of"},
		Country{Code: "THA", Name: "Thailand"},
		Country{Code: "TLS", Name: "Timor-Leste"},
		Country{Code: "TGO", Name: "Togo"},
		Country{Code: "TKL", Name: "Tokelau"},
		Country{Code: "TON", Name: "Tonga"},
		Country{Code: "TTO", Name: "Trinidad and Tobago"},
		Country{Code: "TUN", Name: "Tunisia"},
		Country{Code: "TUR", Name: "Turkey"},
		Country{Code: "TKM", Name: "Turkmenistan"},
		Country{Code: "TCA", Name: "Turks and Caicos Islands"},
		Country{Code: "TUV", Name: "Tuvalu"},
		Country{Code: "UGA", Name: "Uganda"},
		Country{Code: "UKR", Name: "Ukraine"},
		Country{Code: "ARE", Name: "United Arab Emirates"},
		Country{Code: "GBR", Name: "United Kingdom of Great Britain and Northern Ireland"},
		Country{Code: "USA", Name: "United States of America"},
		Country{Code: "USA", Name: "USA"},
		Country{Code: "URY", Name: "Uruguay"},
		Country{Code: "UZB", Name: "Uzbekistan"},
		Country{Code: "VUT", Name: "Vanuatu"},
		Country{Code: "VEN", Name: "Venezuela (Bolivarian Republic of)"},
		Country{Code: "VNM", Name: "Viet Nam"},
		Country{Code: "WLF", Name: "Wallis and Futuna"},
		Country{Code: "ESH", Name: "Western Sahara"},
		Country{Code: "YEM", Name: "Yemen"},
		Country{Code: "ZMB", Name: "Zambia"},
		Country{Code: "ZWE", Name: "Zimbabwe"},
	},
	"es": Countries{
		Country{Code: "AFG", Name: "Afganistán"},
		Country{Code: "ALA", Name: "Alandia"},
		Country{Code: "ALB", Name: "Albania"},
		Country{Code: "DZA", Name: "Argelia"},
		Country{Code: "ASM", Name: "Samoa Americana"},
		Country{Code: "AND", Name: "Andorra"},
		Country{Code: "AGO", Name: "Angola"},
		Country{Code: "AIA", Name: "Anguilla"},
		Country{Code: "ATA", Name: "Antártida"},
		Country{Code: "ATG", Name: "Antigua y Barbuda"},
		Country{Code: "ARG", Name: "Argentina"},
		Country{Code: "ARM", Name: "Armenia"},
		Country{Code: "ABW", Name: "Aruba"},
		Country{Code: "AUS", Name: "Australia"},
		Country{Code: "AUT", Name: "Austria"},
		Country{Code: "AZE", Name: "Azerbaiyán"},
		Country{Code: "BHS", Name: "Bahamas"},
		Country{Code: "BHR", Name: "Bahrein"},
		Country{Code: "BGD", Name: "Bangladesh"},
		Country{Code: "BRB", Name: "Barbados"},
		Country{Code: "BLR", Name: "Bielorrusia"},
		Country{Code: "BEL", Name: "Bélgica"},
		Country{Code: "BLZ", Name: "Belice"},
		Country{Code: "BEN", Name: "Benín"},
		Country{Code: "BMU", Name: "Bermudas"},
		Country{Code: "BTN", Name: "Bután"},
		Country{Code: "BOL", Name: "Bolivia"},
		Country{Code: "BIH", Name: "Bosnia y Herzegovina"},
		Country{Code: "BWA", Name: "Botswana"},
		Country{Code: "BVT", Name: "Isla Bouvet"},
		Country{Code: "BRA", Name: "Brasil"},
		Country{Code: "IOT", Name: "Territorio Británico del Océano \u00cdndico"},
		Country{Code: "UMI", Name: "Islas Ultramarinas Menores de Estados Unidos"},
		Country{Code: "VGB", Name: "Islas Vírgenes del Reino Unido"},
		Country{Code: "VIR", Name: "Islas Vírgenes de los Estados Unidos"},
		Country{Code: "BRN", Name: "Brunei"},
		Country{Code: "BGR", Name: "Bulgaria"},
		Country{Code: "BFA", Name: "Burkina Faso"},
		Country{Code: "BDI", Name: "Burundi"},
		Country{Code: "KHM", Name: "Camboya"},
		Country{Code: "CMR", Name: "Camerún"},
		Country{Code: "CAN", Name: "Canadá"},
		Country{Code: "CPV", Name: "Cabo Verde"},
		Country{Code: "CYM", Name: "Islas Caimán"},
		Country{Code: "CAF", Name: "República Centroafricana"},
		Country{Code: "TCD", Name: "Chad"},
		Country{Code: "CHL", Name: "Chile"},
		Country{Code: "CHN", Name: "China"},
		Country{Code: "CXR", Name: "Isla de Navidad"},
		Country{Code: "CCK", Name: "Islas Cocos o Islas Keeling"},
		Country{Code: "COL", Name: "Colombia"},
		Country{Code: "COM", Name: "Comoras"},
		Country{Code: "COG", Name: "Congo"},
		Country{Code: "COD", Name: "Congo (Rep. Dem.)"},
		Country{Code: "COK", Name: "Islas Cook"},
		Country{Code: "CRI", Name: "Costa Rica"},
		Country{Code: "HRV", Name: "Croacia"},
		Country{Code: "CUB", Name: "Cuba"},
		Country{Code: "CYP", Name: "Chipre"},
		Country{Code: "CZE", Name: "República Checa"},
		Country{Code: "DNK", Name: "Dinamarca"},
		Country{Code: "DJI", Name: "Yibuti"},
		Country{Code: "DMA", Name: "Dominica"},
		Country{Code: "DOM", Name: "República Dominicana"},
		Country{Code: "ECU", Name: "Ecuador"},
		Country{Code: "EGY", Name: "Egipto"},
		Country{Code: "SLV", Name: "El Salvador"},
		Country{Code: "GNQ", Name: "Guinea Ecuatorial"},
		Country{Code: "ERI", Name: "Eritrea"},
		Country{Code: "EST", Name: "Estonia"},
		Country{Code: "ETH", Name: "Etiopía"},
		Country{Code: "FLK", Name: "Islas Malvinas"},
		Country{Code: "FRO", Name: "Islas Faroe"},
		Country{Code: "FJI", Name: "Fiyi"},
		Country{Code: "FIN", Name: "Finlandia"},
		Country{Code: "FRA", Name: "Francia"},
		Country{Code: "GUF", Name: "Guayana Francesa"},
		Country{Code: "PYF", Name: "Polinesia Francesa"},
		Country{Code: "ATF", Name: "Tierras Australes y Antárticas Francesas"},
		Country{Code: "GAB", Name: "Gabón"},
		Country{Code: "GMB", Name: "Gambia"},
		Country{Code: "GEO", Name: "Georgia"},
		Country{Code: "DEU", Name: "Alemania"},
		Country{Code: "GHA", Name: "Ghana"},
		Country{Code: "GIB", Name: "Gibraltar"},
		Country{Code: "GRC", Name: "Grecia"},
		Country{Code: "GRL", Name: "Groenlandia"},
		Country{Code: "GRD", Name: "Grenada"},
		Country{Code: "GLP", Name: "Guadalupe"},
		Country{Code: "GUM", Name: "Guam"},
		Country{Code: "GTM", Name: "Guatemala"},
		Country{Code: "GGY", Name: "Guernsey"},
		Country{Code: "GIN", Name: "Guinea"},
		Country{Code: "GNB", Name: "Guinea-Bisáu"},
		Country{Code: "GUY", Name: "Guyana"},
		Country{Code: "HTI", Name: "Haiti"},
		Country{Code: "HMD", Name: "Islas Heard y McDonald"},
		Country{Code: "VAT", Name: "Santa Sede"},
		Country{Code: "HND", Name: "Honduras"},
		Country{Code: "HKG", Name: "Hong Kong"},
		Country{Code: "HUN", Name: "Hungría"},
		Country{Code: "ISL", Name: "Islandia"},
		Country{Code: "IND", Name: "India"},
		Country{Code: "IDN", Name: "Indonesia"},
		Country{Code: "CIV", Name: "Costa de Marfil"},
		Country{Code: "IRN", Name: "Iran"},
		Country{Code: "IRQ", Name: "Irak"},
		Country{Code: "IRL", Name: "Irlanda"},
		Country{Code: "IMN", Name: "Isla de Man"},
		Country{Code: "ISR", Name: "Israel"},
		Country{Code: "ITA", Name: "Italia"},
		Country{Code: "JAM", Name: "Jamaica"},
		Country{Code: "JPN", Name: "Japón"},
		Country{Code: "JEY", Name: "Jersey"},
		Country{Code: "JOR", Name: "Jordania"},
		Country{Code: "KAZ", Name: "Kazajistán"},
		Country{Code: "KEN", Name: "Kenia"},
		Country{Code: "KIR", Name: "Kiribati"},
		Country{Code: "KWT", Name: "Kuwait"},
		Country{Code: "KGZ", Name: "Kirguizistán"},
		Country{Code: "LAO", Name: "Laos"},
		Country{Code: "LVA", Name: "Letonia"},
		Country{Code: "LBN", Name: "Líbano"},
		Country{Code: "LSO", Name: "Lesotho"},
		Country{Code: "LBR", Name: "Liberia"},
		Country{Code: "LBY", Name: "Libia"},
		Country{Code: "LIE", Name: "Liechtenstein"},
		Country{Code: "LTU", Name: "Lituania"},
		Country{Code: "LUX", Name: "Luxemburgo"},
		Country{Code: "MAC", Name: "Macao"},
		Country{Code: "MKD", Name: "Macedonia"},
		Country{Code: "MDG", Name: "Madagascar"},
		Country{Code: "MWI", Name: "Malawi"},
		Country{Code: "MYS", Name: "Malasia"},
		Country{Code: "MDV", Name: "Maldivas"},
		Country{Code: "MLI", Name: "Mali"},
		Country{Code: "MLT", Name: "Malta"},
		Country{Code: "MHL", Name: "Islas Marshall"},
		Country{Code: "MTQ", Name: "Martinica"},
		Country{Code: "MRT", Name: "Mauritania"},
		Country{Code: "MUS", Name: "Mauricio"},
		Country{Code: "MYT", Name: "Mayotte"},
		Country{Code: "MEX", Name: "México"},
		Country{Code: "FSM", Name: "Micronesia"},
		Country{Code: "MDA", Name: "Moldavia"},
		Country{Code: "MCO", Name: "Mónaco"},
		Country{Code: "MNG", Name: "Mongolia"},
		Country{Code: "MNE", Name: "Montenegro"},
		Country{Code: "MSR", Name: "Montserrat"},
		Country{Code: "MAR", Name: "Marruecos"},
		Country{Code: "MOZ", Name: "Mozambique"},
		Country{Code: "MMR", Name: "Myanmar"},
		Country{Code: "NAM", Name: "Namibia"},
		Country{Code: "NRU", Name: "Nauru"},
		Country{Code: "NPL", Name: "Nepal"},
		Country{Code: "NLD", Name: "Países Bajos"},
		Country{Code: "NCL", Name: "Nueva Caledonia"},
		Country{Code: "NZL", Name: "Nueva Zelanda"},
		Country{Code: "NIC", Name: "Nicaragua"},
		Country{Code: "NER", Name: "Níger"},
		Country{Code: "NGA", Name: "Nigeria"},
		Country{Code: "NIU", Name: "Niue"},
		Country{Code: "NFK", Name: "Isla de Norfolk"},
		Country{Code: "PRK", Name: "Corea del Norte"},
		Country{Code: "MNP", Name: "Islas Marianas del Norte"},
		Country{Code: "NOR", Name: "Noruega"},
		Country{Code: "OMN", Name: "Omán"},
		Country{Code: "PAK", Name: "Pakistán"},
		Country{Code: "PLW", Name: "Palau"},
		Country{Code: "PSE", Name: "Palestina"},
		Country{Code: "PAN", Name: "Panamá"},
		Country{Code: "PNG", Name: "Papúa Nueva Guinea"},
		Country{Code: "PRY", Name: "Paraguay"},
		Country{Code: "PER", Name: "Perú"},
		Country{Code: "PHL", Name: "Filipinas"},
		Country{Code: "PCN", Name: "Islas Pitcairn"},
		Country{Code: "POL", Name: "Polonia"},
		Country{Code: "PRT", Name: "Portugal"},
		Country{Code: "PRI", Name: "Puerto Rico"},
		Country{Code: "QAT", Name: "Catar"},
		Country{Code: "KOS", Name: "Kosovo"},
		Country{Code: "REU", Name: "Reunión"},
		Country{Code: "ROU", Name: "Rumania"},
		Country{Code: "RUS", Name: "Rusia"},
		Country{Code: "RWA", Name: "Ruanda"},
		Country{Code: "BLM", Name: "San Bartolomé"},
		Country{Code: "SHN", Name: "Santa Helena"},
		Country{Code: "KNA", Name: "San Cristóbal y Nieves"},
		Country{Code: "LCA", Name: "Santa Lucía"},
		Country{Code: "MAF", Name: "Saint Martin"},
		Country{Code: "SPM", Name: "San Pedro y Miquelón"},
		Country{Code: "VCT", Name: "San Vicente y Granadinas"},
		Country{Code: "WSM", Name: "Samoa"},
		Country{Code: "SMR", Name: "San Marino"},
		Country{Code: "STP", Name: "Santo Tomé y Príncipe"},
		Country{Code: "SAU", Name: "Arabia Saudí"},
		Country{Code: "SEN", Name: "Senegal"},
		Country{Code: "SRB", Name: "Serbia"},
		Country{Code: "SYC", Name: "Seychelles"},
		Country{Code: "SLE", Name: "Sierra Leone"},
		Country{Code: "SGP", Name: "Singapur"},
		Country{Code: "SVK", Name: "República Eslovaca"},
		Country{Code: "SVN", Name: "Eslovenia"},
		Country{Code: "SLB", Name: "Islas Salomón"},
		Country{Code: "SOM", Name: "Somalia"},
		Country{Code: "ZAF", Name: "República de Sudáfrica"},
		Country{Code: "SGS", Name: "Islas Georgias del Sur y Sandwich del Sur"},
		Country{Code: "KOR", Name: "Corea del Sur"},
		Country{Code: "SSD", Name: "Sudán del Sur"},
		Country{Code: "ESP", Name: "España"},
		Country{Code: "LKA", Name: "Sri Lanka"},
		Country{Code: "SDN", Name: "Sudán"},
		Country{Code: "SUR", Name: "Surinam"},
		Country{Code: "SJM", Name: "Islas Svalbard y Jan Mayen"},
		Country{Code: "SWZ", Name: "Suazilandia"},
		Country{Code: "SWE", Name: "Suecia"},
		Country{Code: "CHE", Name: "Suiza"},
		Country{Code: "SYR", Name: "Siria"},
		Country{Code: "TWN", Name: "Taiwán"},
		Country{Code: "TJK", Name: "Tayikistán"},
		Country{Code: "TZA", Name: "Tanzania"},
		Country{Code: "THA", Name: "Tailandia"},
		Country{Code: "TLS", Name: "Timor Oriental"},
		Country{Code: "TGO", Name: "Togo"},
		Country{Code: "TKL", Name: "Islas Tokelau"},
		Country{Code: "TON", Name: "Tonga"},
		Country{Code: "TTO", Name: "Trinidad y Tobago"},
		Country{Code: "TUN", Name: "Túnez"},
		Country{Code: "TUR", Name: "Turquía"},
		Country{Code: "TKM", Name: "Turkmenistán"},
		Country{Code: "TCA", Name: "Islas Turks y Caicos"},
		Country{Code: "TUV", Name: "Tuvalu"},
		Country{Code: "UGA", Name: "Uganda"},
		Country{Code: "UKR", Name: "Ucrania"},
		Country{Code: "ARE", Name: "Emiratos Árabes Unidos"},
		Country{Code: "GBR", Name: "Reino Unido"},
		Country{Code: "USA", Name: "Estados Unidos"},
		Country{Code: "USA", Name: "EEUU"},
		Country{Code: "URY", Name: "Uruguay"},
		Country{Code: "UZB", Name: "Uzbekistán"},
		Country{Code: "VUT", Name: "Vanuatu"},
		Country{Code: "VEN", Name: "Venezuela"},
		Country{Code: "VNM", Name: "Vietnam"},
		Country{Code: "WLF", Name: "Wallis y Futuna"},
		Country{Code: "ESH", Name: "Sahara Occidental"},
		Country{Code: "YEM", Name: "Yemen"},
		Country{Code: "ZMB", Name: "Zambia"},
		Country{Code: "ZWE", Name: "Zimbabue"},
	},
}

func getCountries(lang string) Countries {
	for k, l := range langs	{
		if k == lang {
			return l
		}
	}

	return nil
}