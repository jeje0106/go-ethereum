// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://3f87316c80ee6ec015e21ce7e2cd5b1067d946a52e87ae83aacc03fa050379c1c3c5412fc31a7092e210636848e2d57db28453779aa1ce7da7113bcd80060adf@65.109.62.117:30303",
        "enode://ddfa3047c9e30644aa9681b7e06138e95164257902664dba912a6fd6c34aa63c266623f25c02789b77632be7690c764d021c7b21146ec44fe565cafe781fc919@5.9.108.123:30305",
        "enode://1d95597cf2da3fa32946440b58378eb3f1d1577265cbc72561ab87316f0131afaac3f137f7b0371129cad740b8b790e3894015ff119d2b190ef06f1e45682aec@141.94.74.221:30303",
        "enode://2a4ef033eca445442b99c6f61e8c163c1fdbb109474a873e708fbc868cbdd8b2b9058b28ad494b012599cf1a5d930084a94ba9cfb9ea0ad70d5a014a0eb168e2@65.109.36.117:30303",
        "enode://bef55cef017293b1fb6811d87c4810718f25f37c37a5a14764c3f605261d35a5af4209ad0fe028ac445732f45df09b54b82d7681f277bbb271cd5a3bc203db7d@51.161.208.55:30303",
        "enode://7e4f0f1a9765be7a1fc06f813d7e6ab9f264ad0bbbfff400c2ca56195244f3da1cc9d4eca3806dc9c2c99b17bff027a1f6a19be997b0adb4ed311781940ef2e9@148.113.159.39:30303",
        "enode://7c616e782bff5cbfab8b198fda7b6d0d3c5dd7aa767efc3b03f5016578bbdfd85b3d390f48853713f8225455bfc2fe3f32045a71d33a84fb66f6f2e2a34497f4@51.195.88.16:20304",
        "enode://9322074f5d2580b878de6b097532de4282f9fc6aba69e071b3697d2599742297ef8f7fea892daeae00321f8791b44a50284057f195107c41144d91d0268b3186@15.235.183.162:30303",
        "enode://3629e4f85435c63c999ec62b6c5cda49f57c778372bb299f344326604b1507244a0ce73a1f176d77e0881fd485ba2593047e290eb519ff3288a77277006d4639@81.183.53.190:34006",
        "enode://c92ca90c449ffca020902fa5c85c5e9c3007f1e46d2ae44a6501d1ac2d951d3aff602df8c498bd6898c06bded6dbf5e9a85473b966990b2ab6f4550de8de931d@170.39.103.34:30303",
        "enode://9f5c7fc1221bc3e210d338709aab9398131122e8de8bcdeb95031f7519f365171b4614fa5ca7e9e208525377b198c9557f283b6d5c012cfd97bf526e20c48672@23.88.2.85:30303",
        "enode://5c771ff0c5a0400edaeb12dfb9b405b35a3b6fcb1e780cc223eebb050563f67e122fbbbd0d617f89c2163f97062a9bc89b8ba4c03d039b2d09ba7ce8d2c3c979@35.82.33.142:30303",
        "enode://697b2b48edda05bedfc3721a66061f727aa76c186a0b8c0b639481883ed9fc856ca4fde8590118a60382f2e8ccbb4bc0c1773e395f0e9f42f041cbfd09b7c319@15.235.13.47:30303",
        "enode://72e71d4e445a67997bb519ffa92845dad43e16d65985c3a5e68ee8c59a5356ed9703dc01700235fc7d48346c56525c262e45e77ea1dbeeb2d97e5b026040656f@81.183.53.190:30008",
        "enode://86c25a2ff11acfffc0792d0e6e880592e7f75c1c01fd43e53addbeb0fcdabde6a9401d72a1c6d1f41efc27130db5ef39b28dd20e57fcc01e1a6478b94fc29131@81.91.189.51:30303",
        "enode://6add27b7c948c66e32823eeae87ce62e0d4e3cdfd10684f2d902856295e154de40f5f7bac0f5a50168f2cc0073c7613935a274706837c85cc458cc95382b7fe3@51.210.219.138:30303",
        "enode://7e18627788dba8158ffb65039beb2a6fdf9a2d5908648b04397a31cab4cedfb045085e71ff95bf8cdd01b08e8d6ba0502a1f7996da84551e2aa7c681cf8aab48@81.183.53.190:39005",
        "enode://cd047287abe643affb85457908176f713e184103538d23a9e97e8d8da2eab3ec86b374d286c232ffb377628301dc34b1dd98dca4897d74a34d0b3cc2bb3afea6@116.203.240.7:30303",
        "enode://832b857a6951a293a7d89a1e82a361ce91243cce9859bb9f6e925cff651e50f584262ed337787763f498db3736ae0de6eb9e64b1b8c23b3433d72ece2749cc60@81.183.53.190:34008",
        "enode://6e316b06b5264446ba46a981671f8963b6ca147a6d7eabb24d4cda18e6b3dd7d9caf05ca5bfd3ed32b2bee0f663e077749eeea23946f057ecff6f4ef317d1c89@147.135.97.187:30303",
        "enode://72f6747eb1e3250f0f14055e489f02c56b43a34934277987da63dadbc9aeec20e435c3c5dc9e3af794f5612053fe0b4de2eccabd62e397661d42249835bb94e3@162.19.169.122:30303",
        "enode://873631a6313fe26e309b566fa16b8b12b9e37bf019180e7c18e982290cc3df489664d57480e52a0acb6f8c33322010379bcf11173bcfb018a91f06b475c4bac1@144.76.17.180:30303",
        "enode://b55fde475b19a8cd07c7ab8ffb15be40950db3525113872d12cd5624a6942a77ad19033f88737969cfd7bfbe52f06ceeafb4728a93d1f7eb8b7bdb4c7a50504f@81.183.53.190:30009",
        "enode://fdf1fdfa5c0e581bc6e718411a35c37e5485836777eb36e99276191efc0e66169c82abb8908bf9e7e2bcb8f07476ac13ec82fa08812baed1ec7a4e6b529968f0@51.38.41.102:30303",
        "enode://aae79ff5076aecaef5e1748bb50f02bb14703823200f184c7beb9b57b12ebe58416dfc96dd41a459fcd91c7931f9c05323b124be03ee2f5a19febaee0d15678a@65.109.62.118:30303",
        "enode://326263381c98db85d71d486f20bbf538e1d3416c6e1600231e9b0ea4c39a322a9dbe8af4a820fc1968ab07eb833e799a9e5a455d924bef7ca5fceb3a0b725818@78.46.47.220:30303",
        "enode://10414cf1a0756ab9182c87d444ec666263a45c0dd449a14513ea6f779419061b2ae480c919560553cdc4644ae91bd28631d80adf46d4caf12771272d56eab056@81.183.53.190:30003",
        "enode://284ceb92d59145a8ac1658b6e87dd6ae4f6ab66d0e472d69bd09a1686de488c964c925d3b2dfd968bc12521c3d29875e83794e3a023ab5d39330a5e123258d25@15.235.183.163:30303",
        "enode://140898c33cf1b5cd3b6f50af2cd97cd993be074ac3345032719f4b2def39e8946fc81352b7e1e3e1c5715e2990a98dc29d926aceef8481eceae14ac8940a61c3@81.183.53.190:30007",
        "enode://bf09f1fde281ed2825165acd727bd93e9c1b98677f3dae73fdce7907e22e63010114c09a44306045232111a0f27cdedd7eda0db66ee7fbc9a6c26b7b0f657f3b@65.108.234.2:30303",
        "enode://ebcc59ae1ac2baf263352593f87550a6ccb0a5633d4eab89a29121b71d088725ab08af065bc140d44024ec79c8a88bd041955299c106cc90905c3767f647b931@176.9.4.4:30305",
        "enode://40aaed4348bd1b24d818a974a3442db484852e6a0be4dfd0d12e17c447ad26dfd3fc5475816b45a3ad41feaca9e770cab1d6ec75615b9dd8f5322407a3d40aa1@35.161.79.164:30303",
        "enode://f4831645b1fef436acaab038da8eb406dc396bcfaa24204ba5ff2ffb31cafd02ba1cd49f2dc65cd701e4ce7b2ecf4ff7e7dd30086296185d36ef98792c4a6a3d@51.81.154.11:30303",
        "enode://77cd733b915ece08c5b9c555cdc546ccd17781ec134d8bba4eed2769ea83cdc3d497335add0a472f8745f0b64ce16244ce2b364c149b2b7ed7c21710895b9ddf@51.195.62.151:30305",
        "enode://22f4a668ff68b9104d67adc411e034b260e2d2712a60f14ac01a0f7716419464cc055dfb8d8b5071831b28eae5d35d6c06d2f6f4f8e320be7251604dca8ee17e@109.108.71.107:30307",
        "enode://b60d58d11a03a80330e38ca83694003541cb3b5b6861ddb52c41ceca86ff264caa04c0f232ae869b38b925ca95e0816eb52ce681b45c867049e7c964d4c1bcd1@37.221.197.208:30303",
        "enode://a11c1339206754b04353e56c30411b2aecf35b28ae4de397c20f2560731ebd7926e0590c8d31889202ebba24efb57d961ed4c168fb4505389067a2c761eaa4c7@45.43.30.3:38001",
        "enode://22d0e1d9673bf48149d3355346be056c2161bcd6cb5982a4fac32224c06cc0d41810465dbe848fa34080d257178b2ffba73cda6ec5ae3423e928a27a185c9dc4@81.183.53.190:34009",
        "enode://4ac7a279cf0787a92ab1660bce95dd7271366b2fa182f9f8e2cef90702765e0705e026a205cb9590c19c0b981037efaac54d6b61e1aa42a3208e7eedfd98e64b@18.141.220.138:30303",
        "enode://d99c2d33a23c06a92f9696009fa133ad011ce5b073c919dc7a1a5dc24102eeafa4c0b7e90dcf3118e06aa77dd4dae8aedd3448ba0e58584c0cc1252024b4aff8@89.232.109.242:33309",
        "enode://c46a74769a8fb669619d8493789ad920a169a4cddc390e08995eb281106e441882b46140f9cc5d7482de6b3e955630e074368321817540f522bcca4f0df61248@81.183.53.190:34004",
        "enode://23e7f1bc5b8c0e92569b1e7abedfc97f58c8612670a8a77ed3e3f909aad8c7c3758df40f1a2f426a437ba05dd18f3469b0031bb1118e371b7467d54aaaf5e015@81.183.53.190:30001",
        "enode://8a66b37212a72f338e86c63cc05fa7a1faa954bd384a48f3656032239388fa889c9b59b01879078a44776ca0c4562137a929d3a78662cfc3de36af154937f81d@176.9.38.209:30303",
        "enode://7afb0231d620d8b1f063ed470dbc4f972b8e18e8b3785992d2220c44a2d5b18cd87e6098ede85dcd1758a98e1f5242a036aa1891c6d11321761ee4bfdffb5896@3.74.145.51:30303",
        "enode://339d36a0db73327133a978dd55072f6a6f5ecde272b835028e830170906731d7b315566b7f72b927b51fe507334b5226e0f0fcec166191a9a92adbb49b264fa2@81.183.53.190:30002",
        "enode://2ba5f404dad7c35fbd4432394a867f24c47430ace619fa4c6694e88d7ad9563806e672fd5243f2f6f9461ce2808aef501c101176515dafdebfed7cac498f722b@81.183.53.190:34002",
        "enode://3ab9c8854b689c8405bf1e341ef9367dec699f3a61939b6a0a6c7b8de01d10f4c75fbdbded01ae39994deae247df3115cbde1fb8af5fced7bf1aa2e34e3b6fa7@65.108.236.6:30303",
        "enode://2ad5b86795fd46ea8dd77af92346bcbaa8db018189cadfa3006d2bf42dae545f65f9a0a5ab3a30c7c8cf5f81f50130b20a1eff5f904f1aca07590b59f9359e3e@5.9.136.199:30304",
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// geth
	"enode://9246d00bc8fd1742e5ad2428b80fc4dc45d786283e05ef6edbd9002cbc335d40998444732fbe921cb88e1d2c73d1b1de53bae6a2237996e9bfe14f871baf7066@18.168.182.86:30303",
	// besu
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://d4f764a48ec2a8ecf883735776fdefe0a3949eb0ca476bd7bc8d0954a9defe8fea15ae5da7d40b5d2d59ce9524a99daedadf6da6283fca492cc80b53689fb3b3@46.4.99.122:32109",
	"enode://d2b720352e8216c9efc470091aa91ddafc53e222b32780f505c817ceef69e01d5b0b0797b69db254c586f493872352f5a022b4d8479a00fc92ec55f9ad46a27e@88.99.70.182:30303",
}

var KilnBootnodes = []string{
	"enode://c354db99124f0faf677ff0e75c3cbbd568b2febc186af664e0c51ac435609badedc67a18a63adb64dacc1780a28dcefebfc29b83fd1a3f4aa3c0eb161364cf94@164.92.130.5:30303",
	"enode://d41af1662434cad0a88fe3c7c92375ec5719f4516ab6d8cb9695e0e2e815382c767038e72c224e04040885157da47422f756c040a9072676c6e35c5b1a383cce@138.68.66.103:30303",
	"enode://91a745c3fb069f6b99cad10b75c463d527711b106b622756e9ef9f12d2631b6cb885f831d1c8731b9bc7177cae5e1ea1f1be087f86d7d30b590a91f22bc041b0@165.232.180.230:30303",
	"enode://b74bd2e8a9f0c53f0c93bcce80818f2f19439fd807af5c7fbc3efb10130c6ee08be8f3aaec7dc0a057ad7b2a809c8f34dc62431e9b6954b07a6548cc59867884@164.92.140.200:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
