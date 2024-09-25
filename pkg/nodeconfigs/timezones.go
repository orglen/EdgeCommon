// Copyright 2021 GoEdge CDN goedge.cdn@gmail.com. All rights reserved.

package nodeconfigs

import "time"

type TimeZoneGroupCode = string

const (
	TimeZoneGroupCodeAfrica       TimeZoneGroupCode = "africa"
	TimeZoneGroupCodeEurope       TimeZoneGroupCode = "europe"
	TimeZoneGroupCodeNorthAmerica TimeZoneGroupCode = "northamerica"
	TimeZoneGroupCodeSouthAmerica TimeZoneGroupCode = "southamerica"
	TimeZoneGroupCodeAntarctica   TimeZoneGroupCode = "antarctica"
	TimeZoneGroupCodeAustralasia  TimeZoneGroupCode = "australasia"
	TimeZoneGroupCodeAsia         TimeZoneGroupCode = "asia"
	TimeZoneGroupCodeUTC          TimeZoneGroupCode = "utc"
)

type TimeZoneGroup struct {
	Name string            `json:"name"`
	Code TimeZoneGroupCode `json:"code"`
}

func FindAllTimeZoneGroups() []*TimeZoneGroup {
	return []*TimeZoneGroup{
		{
			Name: "亚洲",
			Code: TimeZoneGroupCodeAsia,
		},
		{
			Name: "北美",
			Code: TimeZoneGroupCodeNorthAmerica,
		},
		{
			Name: "南美",
			Code: TimeZoneGroupCodeSouthAmerica,
		},
		{
			Name: "欧洲",
			Code: TimeZoneGroupCodeEurope,
		},
		{
			Name: "大洋洲",
			Code: TimeZoneGroupCodeAustralasia,
		},
		{
			Name: "非洲",
			Code: TimeZoneGroupCodeAfrica,
		},
		{
			Name: "南极洲",
			Code: TimeZoneGroupCodeAntarctica,
		},
		{
			Name: "UTC",
			Code: TimeZoneGroupCodeUTC,
		},
	}
}

const DefaultTimeZoneLocation = "Asia/Shanghai"

type TimeZoneLocation struct {
	Name        string            `json:"name"`
	Offset      string            `json:"offset"`
	Group       TimeZoneGroupCode `json:"group"`
	IsCanonical bool              `json:"isCanonical"`
}

var timeZones = []*TimeZoneLocation{}

func FindAllTimeZoneLocations() []*TimeZoneLocation {
	if len(timeZones) == 0 {
		var allTimeZones = []*TimeZoneLocation{{Name: "Africa/Ceuta", Offset: "+02:00", Group: "europe", IsCanonical: true}, {Name: "Africa/Kinshasa", Offset: "+01:00", Group: "africa", IsCanonical: false}, {Name: "Africa/Lagos", Offset: "+01:00", Group: "africa", IsCanonical: true}, {Name: "Africa/Lubumbashi", Offset: "+02:00", Group: "africa", IsCanonical: false}, {Name: "Africa/Maputo", Offset: "+02:00", Group: "africa", IsCanonical: true}, {Name: "America/Adak", Offset: "−09:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Anchorage", Offset: "−08:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Araguaina", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Buenos_Aires", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Catamarca", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Cordoba", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Jujuy", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/La_Rioja", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Mendoza", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Rio_Gallegos", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Salta", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/San_Juan", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/San_Luis", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Tucuman", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Argentina/Ushuaia", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Atikokan", Offset: "−05:00", Group: "northamerica", IsCanonical: false}, {Name: "America/Bahia", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Bahia_Banderas", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Belem", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Blanc-Sablon", Offset: "−04:00", Group: "northamerica", IsCanonical: false}, {Name: "America/Boa_Vista", Offset: "−04:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Boise", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Cambridge_Bay", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Campo_Grande", Offset: "−04:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Cancun", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Chicago", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Chihuahua", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Creston", Offset: "−07:00", Group: "northamerica", IsCanonical: false}, {Name: "America/Cuiaba", Offset: "−04:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Danmarkshavn", Offset: "+00:00", Group: "europe", IsCanonical: true}, {Name: "America/Dawson", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Dawson_Creek", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Denver", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Detroit", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Edmonton", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Eirunepe", Offset: "−05:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Fort_Nelson", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Fortaleza", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Glace_Bay", Offset: "−03:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Goose_Bay", Offset: "−03:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Guayaquil", Offset: "−05:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Halifax", Offset: "−03:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Hermosillo", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Indianapolis", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Knox", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Marengo", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Petersburg", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Tell_City", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Vevay", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Vincennes", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Indiana/Winamac", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Inuvik", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Iqaluit", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Juneau", Offset: "−08:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Kentucky/Louisville", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Kentucky/Monticello", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Los_Angeles", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Maceio", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Manaus", Offset: "−04:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Matamoros", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Mazatlan", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Menominee", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Merida", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Metlakatla", Offset: "−08:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Mexico_City", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Moncton", Offset: "−03:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Monterrey", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/New_York", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Nipigon", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Nome", Offset: "−08:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Noronha", Offset: "−02:00", Group: "southamerica", IsCanonical: true}, {Name: "America/North_Dakota/Beulah", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/North_Dakota/Center", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/North_Dakota/New_Salem", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Nuuk", Offset: "−02:00", Group: "europe", IsCanonical: true}, {Name: "America/Ojinaga", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Panama", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Pangnirtung", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Phoenix", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Porto_Velho", Offset: "−04:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Puerto_Rico", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Punta_Arenas", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Rainy_River", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Rankin_Inlet", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Recife", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Regina", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Resolute", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Rio_Branco", Offset: "−05:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Santarem", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Santiago", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Sao_Paulo", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "America/Scoresbysund", Offset: "+00:00", Group: "europe", IsCanonical: true}, {Name: "America/Sitka", Offset: "−08:00", Group: "northamerica", IsCanonical: true}, {Name: "America/St_Johns", Offset: "−02:30", Group: "northamerica", IsCanonical: true}, {Name: "America/Swift_Current", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Thule", Offset: "−03:00", Group: "europe", IsCanonical: true}, {Name: "America/Thunder_Bay", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Tijuana", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Toronto", Offset: "−04:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Vancouver", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Whitehorse", Offset: "−07:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Winnipeg", Offset: "−05:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Yakutat", Offset: "−08:00", Group: "northamerica", IsCanonical: true}, {Name: "America/Yellowknife", Offset: "−06:00", Group: "northamerica", IsCanonical: true}, {Name: "Antarctica/Casey", Offset: "+11:00", Group: "antarctica", IsCanonical: true}, {Name: "Antarctica/Davis", Offset: "+07:00", Group: "antarctica", IsCanonical: true}, {Name: "Antarctica/DumontDUrville", Offset: "+10:00", Group: "australasia", IsCanonical: false}, {Name: "Antarctica/Macquarie", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Antarctica/Mawson", Offset: "+05:00", Group: "antarctica", IsCanonical: true}, {Name: "Antarctica/McMurdo", Offset: "+13:00", Group: "australasia", IsCanonical: false}, {Name: "Antarctica/Palmer", Offset: "−03:00", Group: "southamerica", IsCanonical: true}, {Name: "Antarctica/Rothera", Offset: "−03:00", Group: "antarctica", IsCanonical: true}, {Name: "Antarctica/Syowa", Offset: "+03:00", Group: "asia", IsCanonical: false}, {Name: "Antarctica/Troll", Offset: "+02:00", Group: "antarctica", IsCanonical: true}, {Name: "Antarctica/Vostok", Offset: "+06:00", Group: "antarctica", IsCanonical: true}, {Name: "Asia/Almaty", Offset: "+06:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Anadyr", Offset: "+12:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Aqtau", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Aqtobe", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Atyrau", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Bangkok", Offset: "+07:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Barnaul", Offset: "+07:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Chita", Offset: "+09:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Choibalsan", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Famagusta", Offset: "+03:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Gaza", Offset: "+03:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Hebron", Offset: "+03:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Ho_Chi_Minh", Offset: "+07:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Hovd", Offset: "+07:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Irkutsk", Offset: "+08:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Jakarta", Offset: "+07:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Jayapura", Offset: "+09:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Kamchatka", Offset: "+12:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Khandyga", Offset: "+09:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Krasnoyarsk", Offset: "+07:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Kuala_Lumpur", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Kuching", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Magadan", Offset: "+11:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Makassar", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Nicosia", Offset: "+03:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Novokuznetsk", Offset: "+07:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Novosibirsk", Offset: "+07:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Omsk", Offset: "+06:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Oral", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Pontianak", Offset: "+07:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Qostanay", Offset: "+06:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Qyzylorda", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Riyadh", Offset: "+03:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Sakhalin", Offset: "+11:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Samarkand", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Shanghai", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Singapore", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Srednekolymsk", Offset: "+11:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Tashkent", Offset: "+05:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Tomsk", Offset: "+07:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Ulaanbaatar", Offset: "+08:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Urumqi", Offset: "+06:00", Group: "asia", IsCanonical: true}, {Name: "Asia/Ust-Nera", Offset: "+10:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Vladivostok", Offset: "+10:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Yakutsk", Offset: "+09:00", Group: "europe", IsCanonical: true}, {Name: "Asia/Yekaterinburg", Offset: "+05:00", Group: "europe", IsCanonical: true}, {Name: "Atlantic/Azores", Offset: "+00:00", Group: "europe", IsCanonical: true}, {Name: "Atlantic/Canary", Offset: "+01:00", Group: "europe", IsCanonical: true}, {Name: "Atlantic/Madeira", Offset: "+01:00", Group: "europe", IsCanonical: true}, {Name: "Australia/Adelaide", Offset: "+10:30", Group: "australasia", IsCanonical: true}, {Name: "Australia/Brisbane", Offset: "+10:00", Group: "australasia", IsCanonical: true}, {Name: "Australia/Broken_Hill", Offset: "+10:30", Group: "australasia", IsCanonical: true}, {Name: "Australia/Darwin", Offset: "+09:30", Group: "australasia", IsCanonical: true}, {Name: "Australia/Eucla", Offset: "+08:45", Group: "australasia", IsCanonical: true}, {Name: "Australia/Hobart", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Australia/Lindeman", Offset: "+10:00", Group: "australasia", IsCanonical: true}, {Name: "Australia/Lord_Howe", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Australia/Melbourne", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Australia/Perth", Offset: "+08:00", Group: "australasia", IsCanonical: true}, {Name: "Australia/Sydney", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Europe/Astrakhan", Offset: "+04:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Berlin", Offset: "+02:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Busingen", Offset: "+02:00", Group: "europe", IsCanonical: false}, {Name: "Europe/Kaliningrad", Offset: "+02:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Kiev", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Kirov", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Lisbon", Offset: "+01:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Madrid", Offset: "+02:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Moscow", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Samara", Offset: "+04:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Saratov", Offset: "+04:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Simferopol", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Ulyanovsk", Offset: "+04:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Uzhgorod", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Volgograd", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Zaporozhye", Offset: "+03:00", Group: "europe", IsCanonical: true}, {Name: "Europe/Zurich", Offset: "+02:00", Group: "europe", IsCanonical: true}, {Name: "Indian/Kerguelen", Offset: "+05:00", Group: "antarctica", IsCanonical: true}, {Name: "Indian/Reunion", Offset: "+04:00", Group: "africa", IsCanonical: true}, {Name: "Pacific/Auckland", Offset: "+13:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Bougainville", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Chatham", Offset: "+13:45", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Chuuk", Offset: "+10:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Easter", Offset: "−05:00", Group: "southamerica", IsCanonical: true}, {Name: "Pacific/Galapagos", Offset: "−06:00", Group: "southamerica", IsCanonical: true}, {Name: "Pacific/Gambier", Offset: "−09:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Honolulu", Offset: "−10:00", Group: "northamerica", IsCanonical: true}, {Name: "Pacific/Kanton", Offset: "+13:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Kiritimati", Offset: "+14:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Kosrae", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Kwajalein", Offset: "+12:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Majuro", Offset: "+12:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Marquesas", Offset: "−09:30", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Midway", Offset: "−11:00", Group: "australasia", IsCanonical: false}, {Name: "Pacific/Pago_Pago", Offset: "−11:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Pohnpei", Offset: "+11:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Port_Moresby", Offset: "+10:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Tahiti", Offset: "−10:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Tarawa", Offset: "+12:00", Group: "australasia", IsCanonical: true}, {Name: "Pacific/Wake", Offset: "+12:00", Group: "australasia", IsCanonical: true}}

		// UTC
		allTimeZones = append(allTimeZones, &TimeZoneLocation{
			Name:        "UTC",
			Offset:      "+00:00",
			Group:       TimeZoneGroupCodeUTC,
			IsCanonical: false,
		})

		// 检查是否存在
		var validTimeZones = []*TimeZoneLocation{}
		for _, location := range allTimeZones {
			_, err := time.LoadLocation(location.Name)
			if err == nil {
				validTimeZones = append(validTimeZones, location)
			}
		}
		timeZones = validTimeZones
	}
	return timeZones
}

func FindTimeZoneLocation(name string) *TimeZoneLocation {
	for _, location := range FindAllTimeZoneLocations() {
		if location.Name == name {
			return location
		}
	}
	return nil
}