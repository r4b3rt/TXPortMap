package common

const (
	UNKNOWN_PORT = iota
	UNKNOWN_SVC
	MySQL
	MySQL_NOT_ALLOWED
	MySQL_BLOCKED
	SSH
	FTP
	FTP_NOT_ALLOWED_OR_NOT_AVAILABLE // {服务不可用或者拒绝访问}
	SMTP
	SMTP_NOT_ALLOWED_OR_NOT_AVAILABLE // {服务不可用或者拒绝访问}
	FTP_OR_SMTP
	FTP_OR_SMTP_SERVICE_NOT_AVAILABLE // 服务不可用
	VM_AUTH_DAEMON
	POP3     // port : 110
	POPPASSD // port : 110
	IMAP4
	VNC     // port : 5900
	RDP     // port : 3389
	SSL_TLS // port : 443
	MSSQL   // port : 1433
	ORACLE  // port : 1521
	REDIS   // port : 6379
	REDIS_AUTH
	REDIS_DENIED
	MEMCACHED    // port : 11211
	TELNET       // port : 23
	HTTP         // port : 80
	PCANYWHERE   // port : 5631
	VPN_PPTP     // port : 1723
	RSYNC        // port : 873
	MSRPC        // port : 135
	NETBIOS_SSN  // port : 139
	MICROSOFT_DS // port : 445
	DECRPC
	POSTGRESQL // port : 5432
	IBMDB2
	MONGODB // port : 27017
	MONGODB_AUTH
	LDAP // port : 389
	DNS  // port : 53
	DNS_NOBANNER
	ROUTEROS // port : 8291
	RADMIN   // port : 4899
	JAVARMI
	JAVARMI_CRASHPLAN
	JABBER
	JDWP
	MMS    // port : 1755
	DTSPCD // port : 6112
	SVNSERVE
	WEBLOGIC // port : 7001
	SIP
	SVRLOC        // port : 427
	AJP13         // port : 8009
	NFS           // port : 2049
	ELASTICSEARCH // port : 9300
	RTSP          // port : 554
	LOTUSNOTES    // port : 1352
)

const (
	SOCKET_CONNECT_FAILED = iota + 10000001 // 连接失败
	SOCKET_READ_TIMEOUT                     // 读取超时错误
)

// IdentificationProtocol 用于服务识别时发送的报文
var IdentificationProtocol = []string{
	"http#474554202f20485454502f312e310d0a557365722d4167656e743a204d6f7a696c6c612f352e30202857696e646f7773204e5420362e313b20574f57363429204170706c655765624b69742f3533372e333620284b48544d4c2c206c696b65204765636b6f29204368726f6d652f35312e302e323730342e313036205361666172692f3533372e33360d0a4163636570743a202a2f2a0d0a486f73743a20",
	"rdp#030000130ee000000000000100080003000000",
	"ssl/tls(1)#16030100ae010000aa030355d1668dc7272d1242aaf6e9666fbbe21e5acc601d572b39e45a5f13609dec1600001c5a5ac02bc02fc02cc030cca9cca8c013c014009c009d002f0035000a010000656a6a0000ff010001000017000000230000000d00140012040308040401050308050501080606010201000500050100000000001200000010000e000c02683208687474702f312e3175500000000b00020100000a000a0008dada001d001700188a8a000100",
	"ssl/tls(2)#1603000069010000650303551ca7e472616e646f6d3172616e646f6d3272616e646f6d3372616e646f6d3400000c002f000a00130039000400ff01000030000d002c002a000100030002060106030602020102030202030103030302040104030402010101030102050105030502",
	"ssl/tls(3)#809e01030100750000002000006600006500006400006300006200003a00003900003800003500003400003300003200002f00001b00001a00001900001800001700001600001500001400001300001200001100000a0000090000080000060000050000040000030700c0060040040080030080020080010080000002000001e4693c2bf6d69bbbd3819fbf15c140a56f142c4d20c4c7e0b6b0b21ff929e898",
	"mssql#1201003400000000000015000601001b000102001c000c0300280004ff080001550000004d5353514c536572766572004e53464f",
	"dns#001E0006010000010000000000000776657273696F6E0462696E640000100003",
	"oracle#005a0000010000000136012c000008007fff7f08000000010020003a0000000000000000000000000000000034e600000001000000000000000028434f4e4e4543545f444154413d28434f4d4d414e443d76657273696f6e2929",
	"nfs#80000028106c8eb90000000000000002000186a3000000040000000000000000000000000000000000000000",
	"redis#2a310d0a24340d0a70696e670d0a696e666f0d0a717569740d0a",
	"weblogic#74332031322e312e320a41533a323034380a484c3a31390a0a",
	"vpn-pptp#009c00011a2b3c4d00010000010000000000000100000001ffff00016e6f6e650000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004d6963726f736f667400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
	"rsync#405253594e43443a2032392e0a0a",
	"memcached#73746174730d0a717569740d0a",
	"msrpc/netbios-ssn/microsoft-ds/postgresql#000000a4ff534d4272000000000801400000000000000000000000000000400600000100008100025043204e4554574f524b2050524f4752414d20312e3000024d4943524f534f4654204e4554574f524b5320312e303300024d4943524f534f4654204e4554574f524b5320332e3000024c414e4d414e312e3000024c4d312e3258303032000253616d626100024e54204c414e4d414e20312e3000024e54204c4d20302e313200",
	"ibm-db2#00000000444232444153202020202020010400000010397a000100000000000000000000010c0000000000000c0000000c00000004",
	"mongodb#410000004d095000ffffffffd407000000000000746573742e24636d640000000000ffffffff1b0000000173657276657253746174757300000000000000f03f00",
	"ldap#300c020101600702010204008000",
	"mikrotik#12026c6973740000000000000000000100000000",
	"ajp13#123400010a",
	"wms#01000000cefa0bb0a00000004d4d5320140000000000000000000000000000001200000001000300f0f0f0f00b0004001c0003004e00530050006c0061007900650072002f00310030002e0030002e0030002e0033003600340036003b0020007b00330033003000300041004400350030002d0032004300330039002d0034003600630030002d0041004500300041002d004200410033004500450030004300380031003300360045007d0000000000",
	"dtspcd#3030303030303032303430303064303030312020342000726f6f74000031300000",
	"jabber#3c3f786d6c2076657273696f6e3d27312e30273f3e3c73747265616d3a73747265616d20786d6c6e733a73747265616d3d27687474703a2f2f6574686572782e6a61626265722e6f72672f73747265616d732720786d6c6e733d276a61626265723a636c69656e742720786d6c3a6c616e673d2772752d52552720746f3d272e272076657273696f6e3d27312e30273e",
	"jdwp#4a4457502d48616e647368616b65",
	"sip#4f5054494f4e53207369703a6e6d205349502f322e300d0a5669613a205349502f322e302f544350206e6d3b6272616e63683d666f6f0d0a46726f6d3a203c7369703a6e6d406e6d3e3b7461673d726f6f740d0a546f3a203c7369703a6e6d32406e6d323e0d0a43616c6c2d49443a2035303030300d0a435365713a203432204f5054494f4e530d0a4d61782d466f7277617264733a2037300d0a436f6e74656e742d4c656e6774683a20300d0a436f6e746163743a203c7369703a6e6d406e6d3e0d0a4163636570743a206170706c69636174696f6e2f7364700d0a0d0a",
	"svrloc#0201000036200000000000010002656e00000015736572766963653a736572766963652d6167656e74000764656661756c7400000000",
	"pcanywhere#00000000",
	"radmin#01000000010000000808",
	"postgresql#00000042000300007573657200706f73746772657300646174616261736500706f737467726573006170706c69636174696f6e5f6e616d65004e6176696361740000",
	"rtsp/sip#4f5054494f4e53207369703a6e6d205349502f322e300d0a5669613a205349502f322e302f544350206e6d3b6272616e63683d666f6f0d0a46726f6d3a203c7369703a6e6d406e6d3e3b7461673d726f6f740d0a546f3a203c7369703a6e6d32406e6d323e0d0a43616c6c2d49443a2035303030300d0a435365713a203432204f5054494f4e530d0a4d61782d466f7277617264733a2037300d0a436f6e74656e742d4c656e6774683a20300d0a436f6e746163743a203c7369703a6e6d406e6d3e0d0a4163636570743a206170706c69636174696f6e2f7364700d0a0d0a",
	"lotusnotes#3a0000002f00000002000040020f0001003d050000000000000000000000002f000000000000000000401f0000000000000000000000000000000000",
	"onvif(1)#3c3f786d6c2076657273696f6e3d22312e302220656e636f64696e673d227574662d38223f3e3c456e76656c6f706520786d6c6e733a646e3d22687474703a2f2f7777772e6f6e7669662e6f72672f76657231302f6e6574776f726b2f7773646c2220786d6c6e733d22687474703a2f2f7777772e77332e6f72672f323030332f30352f736f61702d656e76656c6f7065223e3c4865616465723e3c7773613a4d657373616765494420786d6c6e733a7773613d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030342f30382f61646472657373696e67223e757569643a37653266353530352d393765352d343934312d613732362d3938376137656265616136303c2f7773613a4d65737361676549443e3c7773613a546f20786d6c6e733a7773613d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030342f30382f61646472657373696e67223e75726e3a736368656d61732d786d6c736f61702d6f72673a77733a323030353a30343a646973636f766572793c2f7773613a546f3e3c7773613a416374696f6e20786d6c6e733a7773613d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030342f30382f61646472657373696e67223e687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030352f30342f646973636f766572792f50726f62653c2f7773613a416374696f6e3e3c2f4865616465723e3c426f64793e3c50726f626520786d6c6e733a7873693d22687474703a2f2f7777772e77332e6f72672f323030312f584d4c536368656d612d696e7374616e63652220786d6c6e733a7873643d22687474703a2f2f7777772e77332e6f72672f323030312f584d4c536368656d612220786d6c6e733d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030352f30342f646973636f76657279223e3c54797065733e646e3a4e6574776f726b566964656f5472616e736d69747465723c2f54797065733e3c53636f706573202f3e3c2f50726f62653e3c2f426f64793e3c2f456e76656c6f70653e",
	"onvif(2)#3c3f786d6c2076657273696f6e3d22312e302220656e636f64696e673d227574662d38223f3e3c456e76656c6f706520786d6c6e733a7464733d22687474703a2f2f7777772e6f6e7669662e6f72672f76657231302f6465766963652f7773646c2220786d6c6e733d22687474703a2f2f7777772e77332e6f72672f323030332f30352f736f61702d656e76656c6f7065223e3c4865616465723e3c7773613a4d657373616765494420786d6c6e733a7773613d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030342f30382f61646472657373696e67223e757569643a31613737653736312d396361302d343939612d613835332d3463346236633935323864643c2f7773613a4d65737361676549443e3c7773613a546f20786d6c6e733a7773613d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030342f30382f61646472657373696e67223e75726e3a736368656d61732d786d6c736f61702d6f72673a77733a323030353a30343a646973636f766572793c2f7773613a546f3e3c7773613a416374696f6e20786d6c6e733a7773613d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030342f30382f61646472657373696e67223e687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030352f30342f646973636f766572792f50726f62653c3b2f7773613a416374696f6e3e3c2f4865616465723e3c426f64793e3c50726f626520786d6c6e733a7873693d22687474703a2f2f7777772e77332e6f72672f323030312f584d4c536368656d612d696e7374616e63652220786d6c6e733a7873643d22687474703a2f2f7777772e77332e6f72672f323030312f584d4c536368656d612220786d6c6e733d22687474703a2f2f736368656d61732e786d6c736f61702e6f72672f77732f323030352f30342f646973636f76657279223e3c54797065733e7464733a4465766963653c2f54797065733e3c53636f706573202f3e3c2f50726f62653e3c2f426f64793e3c2f456e76656c6f70653e",
}

// Top 100常见端口
var Top100Ports = []string{
	"21-23", "25", "53", "69", "79", "80-89", "110-111", "135", "139", "143", "161", "322", "389", "443", "445", "465", "512-515", "524", "587", "873", "993", "995", "999",
	"1080", "1158", "1352", "1433", "1521", "1863", "2049", "2100", "2181", "2222", "2323", "3128", "3306", "3389", "4848", "4899", "5000", "5061", "5432", "5631-5632", "5800", "5900", "6379", "7001", "7080", "7090", "8000", "8008",
	"8009", "8069", "8080-8090", "8141", "8161", "8291", "8443", "8888", "8899", "8880", "9001", "9080", "9090", "9200", "9300", "9443", "9898", "9900",
	"17001-17003", "11211", "20080", "27017",
}

// Top 1000常见端口
var Top1000Ports = []string{
	"1", "3-4", "6-7", "9", "13", "17", "19-26", "30", "32-33", "37", "42-43", "49", "53", "70", "79-85", "88-90", "99-100", "106", "109-111", "113", "119", "125", "135", "139", "143-144", "146", "161", "163", "179", "199", "211-212", "222", "254-256", "259", "264", "280", "301", "306", "311", "340", "366", "389", "406-407", "416-417", "425", "427", "443-445", "458", "464-465", "481", "497", "500", "512-515", "524", "541", "543-545", "548", "554-555", "563", "587", "593", "616-617", "625", "631", "636", "646", "648", "666-668", "683", "687", "691", "700", "705", "711", "714", "720", "722", "726", "749", "765", "777", "783", "787", "800-801", "808", "843", "873", "880", "888", "898", "900-903", "911-912", "981", "987", "990", "992-993", "995", "999-1002",
	"1007", "1009-1011", "1021-1100", "1102", "1104-1108", "1110-1114", "1117", "1119", "1121-1124", "1126", "1130-1132", "1137-1138", "1141", "1145", "1147", "1148-1149", "1151-1152", "1154", "1163-1166", "1169", "1174-1175", "1183", "1185-1187", "1192", "1198-1199", "1201", "1213", "1216-1218", "1233-1234", "1236", "1244", "1247-1248", "1259", "1271-1272", "1277", "1287", "1296", "1300-1301", "1309-1311", "1322", "1328", "1334", "1352", "1417", "1433-1434", "1443", "1455", "1461", "1494", "1500-1501", "1503", "1521", "1524", "1533", "1556", "1580", "1583", "1594", "1600", "1641", "1658", "1666", "1687-1688", "1700", "1717-1721", "1723", "1755", "1761", "1782-1783", "1801", "1805", "1812", "1839-1840", "1862-1864", "1875", "1900", "1914", "1935", "1947", "1971-1972", "1974", "1984", "1998-2010",
	"2013", "2020-2022", "2030", "2033-2035", "2038", "2040-2043", "2045-2049", "2065", "2068", "2099-2100", "2103", "2105-2107", "2111", "2119", "2121", "2126", "2135", "2144", "2160-2161", "2170", "2179", "2190", "2191", "2196", "2200", "2222", "2251", "2260", "2288", "2301", "2323", "2366", "2381-2383", "2393-2394", "2399", "2401", "2492", "2500", "2522", "2525", "2557", "2601-2602", "2604-2605", "2607-2608", "2638", "2701-2702", "2710", "2717-2718", "2725", "2800", "2809", "2811", "2869", "2875", "2909-2910", "2920", "2967-2968", "2998", "3000-3001", "3003", "3005-3007", "3011", "3013", "3017", "3030-3031", "3050", "3052", "3071", "3077", "3128", "3168", "3211", "3221", "3260-3261", "3268-3269", "3283", "3300-3301", "3306", "3322-3325", "3333", "3351", "3367", "3369-3372", "3389-3390", "3404", "3476", "3493", "3517", "3527", "3546", "3551", "3580", "3659", "3689", "3690", "3703", "3737", "3766", "3784", "3800-3801", "3809", "3814", "3826-3828", "3851", "3869", "3871", "3878", "3880", "3889", "3905", "3914", "3918", "3920", "3945", "3971", "3986", "3995", "3998", "4000-4006", "4045", "4111", "4125-4126", "4129", "4224", "4242", "4279", "4321", "4343", "4443-4446", "4449-4550", "4567", "4662", "4848", "4899-4900", "4998",
	"5000-5004", "5009", "5030", "5033", "5050-5051", "5054", "5060-5061", "5080", "5087", "5100-5102", "5120", "5190", "5200", "5214", "5221-5222", "5225-5226", "5269", "5280", "5298", "5357", "5405", "5414", "5431-5432", "5440", "5500", "5510", "5544", "5550", "5555", "5560", "5566", "5631", "5633", "5666", "5678-5679", "5718", "5730", "5800-5802", "5810-5811", "5815", "5822", "5825", "5850", "5859", "5862", "5877", "5900-5904", "5906-5907", "5910-5911", "5915", "5922", "5925", "5950", "5952", "5959-5963", "5987-5989", "5998-6007", "6009", "6025", "6059", "6100-6101", "6106", "6112", "6123", "6129", "6156", "6346", "6389", "6502", "6510", "6543", "6547", "6565-6567", "6580", "6646", "6666-6669", "6689", "6692", "6699", "6779", "6788-6789", "6792", "6839", "6881", "6901", "6969", "7000-7002", "7004", "7007", "7019", "7025", "7070", "7100", "7103", "7106", "7200-7201", "7402", "7435", "7443", "7496", "7512", "7625", "7627", "7676", "7741", "7777-7778", "7800", "7911", "7920-7921", "7937-7938", "7999-8002", "8007-8011", "8021-8022", "8031", "8042", "8045", "8080-8090", "8093", "8099-8100", "8180-8181", "8192-8194", "8200", "8222", "8254", "8290-8292", "8300", "8333", "8383", "8400", "8402", "8443", "8500", "8600", "8649", "8651-8652", "8654", "8701", "8800", "8873", "8888", "8899", "8880", "8994",
	"9000-9003", "9009-9011", "9040", "9050", "9060", "9043", "9071", "9080-9081", "9090-9091", "9099-9103", "9110-9111", "9200", "9207", "9220", "9290", "9415", "9418", "9485", "9500", "9443", "9502-9503", "9535", "9575", "9593-9595", "9618", "9666", "9876-9878", "9898", "9900", "9917", "9943-9944", "9968", "9998-10004", "10009-10010", "10012", "10024-10025", "10082", "10180", "10215", "10243", "10566", "10616-10617", "10621", "10626", "10628-10629", "10778", "11110-11111", "11967", "12000", "12174", "12265", "12345", "13456", "13722", "13782-13783", "14000", "14238", "14441-14442", "15000", "15002-15004", "15660", "15742", "16000-16001", "16012", "16016", "16018", "16080", "16113", "16992-16993", "17877", "17988", "18040", "18101", "18988", "19101", "19283", "19315", "19350", "19780", "19801", "19842", "20000", "20005", "20031", "20221-20222", "20828", "21571", "22939", "23502", "24444", "24800", "25734-25735", "26214", "27000", "27352-27353", "27355-27356", "27715", "28201", "30000", "30718", "30951", "31038", "31337", "32768-32785", "33354", "33899", "34571-34573", "35500", "38292", "40193", "40911", "41511", "42510", "44176", "44442-44443", "44501", "45100", "48080", "49152-49161", "49163", "49165", "49167", "49175-49176", "49400", "49999-50003", "50006", "50300", "50389", "50500", "50636", "50800", "51103", "51493", "52673", "52822", "52848", "52869", "54045", "54328", "55055-55056", "55555", "55600", "56737-56738", "57294", "57797", "58080", "60020", "60443", "61532", "61900", "62078", "63331", "64623", "64680", "65000", "65129", "65389",
}

type Identification_Port struct {
	SzSvcName             string // 服务名
	Port                  uint64 // 端口
	Identification_RuleId int    // 针对规则条目序号
}

var St_Identification_Port = []Identification_Port{
	{"rdp", 3389, 1},
	{"mssql", 1433, 5},
	{"dns", 53, 6},
	{"oracle", 1521, 7},
	{"nfs", 2049, 8},
	{"redis", 6379, 9},
	{"weblogic", 7001, 10},
	{"weblogic", 7002, 10},
	{"vpn", 1723, 11},
	{"rsync", 873, 12},
	{"memcached", 11211, 13},
	{"msrcp", 135, 14},
	{"netbios-ssn", 139, 14},
	{"microsoft-ds", 445, 14},
	{"mongodb", 27017, 16},
	{"ldap", 389, 17},
	{"mikrotik", 8291, 18},
	{"ajp13", 8009, 19},
	{"wms", 1755, 20},
	{"dtspcd", 6112, 21},
	{"svrloc", 427, 25},
	{"pcanywhere", 5631, 26},
	{"pcanywhere", 5632, 26},
	{"radmin", 4899, 27},
	{"postgresql", 5432, 28}, // 模拟Navicat登陆
	{"postgresql", 5432, 14}, // 通过rpc包判断
	{"rtsp", 554, 29},
	{"sip", 5060, 29},
	{"lotusnotes", 1352, 30},
	{"onvif",3702,31},
	{"onvif",3702,32},
}

// rdp_receive_packet
var os___xrdp_1 = []byte{0x03, 0x00, 0x00, 0x09, 0x02, 0xf0, 0x80, 0x21, 0x80}                                                             // xrdp
var os___xrdp_2 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x01, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00} // xrdp
var os___xrdp_3 = []byte{0x03, 0x00, 0x00, 0x0b, 0x06, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00}                                                 // xrdp
var os___xrdp_4 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x01, 0x08, 0x00, 0x01, 0x00, 0x00, 0x00} // xrdp

var os______old = []byte{0x03, 0x00, 0x00, 0x0b, 0x06, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00} // Windows 2000 Advanced Server || Windoes XP Professional || Windows Embedded POSReady 2009 || Windows Embedded Standard

var os___2008_1 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x00, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Windows Server 2008 R2 Datacenter
var os___2008_2 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x01, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Windows Server 2008 R2 Datacenter
var os___2008_3 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x09, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Windows Server 2008 R2 Standard

var os___2012_1 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x00, 0x08, 0x00, 0x01, 0x00, 0x00, 0x00} // Windows Server 2012 R2
var os___2012_2 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x07, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Windows Server 2012
var os__2012_r2 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x0f, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Windows Server 2012 R2

var os__Vista_1 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x02, 0x1f, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Vista以后的操作系统

var os_Multiple_1 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x03, 0x00, 0x08, 0x00, 0x02, 0x00, 0x00, 0x00} // Windows 2003 / 2008 / 2012
var os_Multiple_2 = []byte{0x03, 0x00, 0x00, 0x13, 0x0e, 0xd0, 0x00, 0x00, 0x12, 0x34, 0x00, 0x03, 0x00, 0x08, 0x00, 0x03, 0x00, 0x00, 0x00} // Windows 2003 / 2008 / 2012

var iPacketMask = len(IdentificationProtocol)
