package client

type Itemseverything struct {
	XML struct {
		Version  string `json:"@version"`
		Encoding string `json:"@encoding"`
	} `json:"?xml"`
	Items struct {
		XmlnsXsi                     string `json:"@xmlns:xsi"`
		XsiNoNamespaceSchemaLocation string `json:"@xsi:noNamespaceSchemaLocation"`
		Shopcategories               struct {
			Shopcategory []struct {
				ID              string `json:"@id"`
				Value           string `json:"@value"`
				Shopsubcategory []struct {
					ID    string `json:"@id"`
					Value string `json:"@value"`
				} `json:"shopsubcategory"`
			} `json:"shopcategory"`
		} `json:"shopcategories"`
		Hideoutitem struct {
			Uniquename               string `json:"@uniquename"`
			Itemvalue                string `json:"@itemvalue"`
			Tier                     string `json:"@tier"`
			Mindistance              string `json:"@mindistance"`
			Mindistanceintunnel      string `json:"@mindistanceintunnel"`
			Placementduration        string `json:"@placementduration"`
			Primetimedurationminutes string `json:"@primetimedurationminutes"`
			Maxstacksize             string `json:"@maxstacksize"`
			Weight                   string `json:"@weight"`
			Unlockedtocraft          string `json:"@unlockedtocraft"`
			Shopcategory             string `json:"@shopcategory"`
			Shopsubcategory1         string `json:"@shopsubcategory1"`
			Uicraftsoundstart        string `json:"@uicraftsoundstart"`
			Uicraftsoundfinish       string `json:"@uicraftsoundfinish"`
			Craftingrequirements     struct {
				Silver        string `json:"@silver"`
				Time          string `json:"@time"`
				Craftingfocus string `json:"@craftingfocus"`
				Craftresource []struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"craftresource"`
			} `json:"craftingrequirements"`
		} `json:"hideoutitem"`
		Trackingitem []struct {
			Uniquename                   string `json:"@uniquename"`
			Uisprite                     string `json:"@uisprite"`
			Findtrackspell               string `json:"@findtrackspell"`
			Trackingtimereduction        string `json:"@trackingtimereduction"`
			Trackingfamebonus            string `json:"@trackingfamebonus"`
			Durability                   string `json:"@durability"`
			Itempower                    string `json:"@itempower"`
			DurabilitylossInspecttrack   string `json:"@durabilityloss_inspecttrack"`
			DurabilitylossFindtrack      string `json:"@durabilityloss_findtrack"`
			Tier                         string `json:"@tier"`
			Craftingcategory             string `json:"@craftingcategory"`
			Maxstacksize                 string `json:"@maxstacksize"`
			Weight                       string `json:"@weight"`
			Shopcategory                 string `json:"@shopcategory"`
			Shopsubcategory1             string `json:"@shopsubcategory1"`
			Uicraftsoundstart            string `json:"@uicraftsoundstart"`
			Uicraftsoundfinish           string `json:"@uicraftsoundfinish"`
			Slottype                     string `json:"@slottype"`
			Passivespellslots            string `json:"@passivespellslots"`
			Activespellslots             string `json:"@activespellslots"`
			Descriptionlocatag           string `json:"@descriptionlocatag"`
			DurabilitylossAttack         string `json:"@durabilityloss_attack"`
			DurabilitylossSpelluse       string `json:"@durabilityloss_spelluse"`
			DurabilitylossReceivedattack string `json:"@durabilityloss_receivedattack"`
			DurabilitylossReceivedspell  string `json:"@durabilityloss_receivedspell"`
			Unlockedtocraft              string `json:"@unlockedtocraft"`
			Unlockedtoequip              string `json:"@unlockedtoequip"`
			Abilitypower                 string `json:"@abilitypower"`
			Canbeovercharged             string `json:"@canbeovercharged"`
			Craftingrequirements         struct {
				Silver        string `json:"@silver"`
				Time          string `json:"@time"`
				Craftingfocus string `json:"@craftingfocus"`
				Craftresource []struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"craftresource"`
			} `json:"craftingrequirements"`
		} `json:"trackingitem"`
		Farmableitem []struct {
			Uniquename                      string `json:"@uniquename"`
			Tier                            string `json:"@tier"`
			Placefame                       string `json:"@placefame"`
			Pickupable                      string `json:"@pickupable"`
			Destroyable                     string `json:"@destroyable"`
			Unlockedtoplace                 string `json:"@unlockedtoplace"`
			Maxstacksize                    string `json:"@maxstacksize"`
			Shopcategory                    string `json:"@shopcategory"`
			Shopsubcategory1                string `json:"@shopsubcategory1"`
			Kind                            string `json:"@kind"`
			Weight                          string `json:"@weight"`
			Unlockedtocraft                 string `json:"@unlockedtocraft"`
			Animationid                     string `json:"@animationid"`
			Activefarmfocuscost             string `json:"@activefarmfocuscost,omitempty"`
			Activefarmmaxcycles             string `json:"@activefarmmaxcycles,omitempty"`
			Activefarmactiondurationseconds string `json:"@activefarmactiondurationseconds,omitempty"`
			Activefarmcyclelengthseconds    string `json:"@activefarmcyclelengthseconds,omitempty"`
			Activefarmbonus                 string `json:"@activefarmbonus,omitempty"`
			Itemvalue                       string `json:"@itemvalue,omitempty"`
			Craftingrequirements            struct {
				Silver          string `json:"@silver"`
				Time            string `json:"@time"`
				Swaptransaction string `json:"@swaptransaction"`
			} `json:"craftingrequirements,omitempty"`
			AudioInfo struct {
				Name string `json:"@name"`
			} `json:"AudioInfo"`
			Harvest struct {
				Growtime   string `json:"@growtime"`
				Lootlist   string `json:"@lootlist"`
				Lootchance string `json:"@lootchance"`
				Fame       string `json:"@fame"`
				Seed       struct {
					Chance string `json:"@chance"`
					Amount string `json:"@amount"`
				} `json:"seed"`
			} `json:"harvest,omitempty"`
			Prefabname    string `json:"@prefabname,omitempty"`
			Prefabscale   string `json:"@prefabscale,omitempty"`
			Resourcevalue string `json:"@resourcevalue,omitempty"`
			Grownitem     struct {
				Uniquename string `json:"@uniquename"`
				Growtime   string `json:"@growtime"`
				Fame       string `json:"@fame"`
				Offspring  struct {
					Chance string `json:"@chance"`
					Amount string `json:"@amount"`
				} `json:"offspring"`
			} `json:"grownitem,omitempty"`
			Consumption struct {
				Food struct {
					Nutritionmax        string `json:"@nutritionmax"`
					Secondspernutrition string `json:"@secondspernutrition"`
					Hungryoverheadicon  string `json:"@hungryoverheadicon"`
					Acceptedfood        struct {
						Foodcategory string `json:"@foodcategory"`
					} `json:"acceptedfood"`
				} `json:"food"`
			} `json:"consumption,omitempty"`
			Tile              string `json:"@tile,omitempty"`
			Uisprite          string `json:"@uisprite,omitempty"`
			Showinmarketplace string `json:"@showinmarketplace,omitempty"`
			Products          struct {
				Product struct {
					Productiontime string `json:"@productiontime"`
					Actionname     string `json:"@actionname"`
					Lootlist       string `json:"@lootlist"`
					Lootchance     string `json:"@lootchance"`
					Fame           string `json:"@fame"`
				} `json:"product"`
			} `json:"products,omitempty"`
		} `json:"farmableitem"`
		Simpleitem []struct {
			Uniquename             string `json:"@uniquename"`
			Tier                   string `json:"@tier"`
			Weight                 string `json:"@weight,omitempty"`
			Maxstacksize           string `json:"@maxstacksize,omitempty"`
			Uisprite               string `json:"@uisprite,omitempty"`
			Shopcategory           string `json:"@shopcategory"`
			Shopsubcategory1       string `json:"@shopsubcategory1"`
			Unlockedtocraft        string `json:"@unlockedtocraft,omitempty"`
			Itemvalue              string `json:"@itemvalue,omitempty"`
			Nutrition              string `json:"@nutrition,omitempty"`
			Foodcategory           string `json:"@foodcategory,omitempty"`
			Resourcetype           string `json:"@resourcetype,omitempty"`
			Famevalue              string `json:"@famevalue,omitempty"`
			Enchantmentlevel       string `json:"@enchantmentlevel,omitempty"`
			Fishingfame            string `json:"@fishingfame,omitempty"`
			Fishingminigamesetting string `json:"@fishingminigamesetting,omitempty"`
			Craftingrequirements   []struct {
				Time          string `json:"@time"`
				Amountcrafted string `json:"@amountcrafted"`
				Craftresource struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"craftresource"`
			} `json:"craftingrequirements,omitempty"`
			Descriptionlocatag      string `json:"@descriptionlocatag,omitempty"`
			Fasttravelfactor        string `json:"@fasttravelfactor,omitempty"`
			Showinmarketplace       string `json:"@showinmarketplace,omitempty"`
			Hidefromplayeroncontext string `json:"@hidefromplayeroncontext,omitempty"`
			Namelocatag             string `json:"@namelocatag,omitempty"`
			Salvageable             string `json:"@salvageable,omitempty"`
			Droponpvpknockdown      string `json:"@droponpvpknockdown,omitempty"`
			Replaceondeath          struct {
				Replacementitem struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"replacementitem"`
			} `json:"replaceondeath,omitempty"`
			Uispriteoverlay1         string `json:"@uispriteoverlay1,omitempty"`
			Tradable                 string `json:"@tradable,omitempty"`
			Descvariable0            string `json:"@descvariable0,omitempty"`
			Descvariable1            string `json:"@descvariable1,omitempty"`
			Uicraftsoundstart        string `json:"@uicraftsoundstart,omitempty"`
			Uicraftsoundfinish       string `json:"@uicraftsoundfinish,omitempty"`
			Craftingcategory         string `json:"@craftingcategory,omitempty"`
			Dontgivefameoncraft      string `json:"@dontgivefameoncraft,omitempty"`
			Canbestoredinbattlevault string `json:"@canbestoredinbattlevault,omitempty"`
		} `json:"simpleitem"`
		Consumableitem []struct {
			Uniquename             string `json:"@uniquename"`
			Itemvalue              string `json:"@itemvalue,omitempty"`
			Fishingfame            string `json:"@fishingfame,omitempty"`
			Fishingminigamesetting string `json:"@fishingminigamesetting,omitempty"`
			Descriptionlocatag     string `json:"@descriptionlocatag,omitempty"`
			Uisprite               string `json:"@uisprite,omitempty"`
			Nutrition              string `json:"@nutrition,omitempty"`
			Abilitypower           string `json:"@abilitypower"`
			Slottype               string `json:"@slottype"`
			Consumespell           string `json:"@consumespell"`
			Shopcategory           string `json:"@shopcategory"`
			Shopsubcategory1       string `json:"@shopsubcategory1"`
			Resourcetype           string `json:"@resourcetype,omitempty"`
			Tier                   string `json:"@tier"`
			Weight                 string `json:"@weight,omitempty"`
			Dummyitempower         string `json:"@dummyitempower,omitempty"`
			Maxstacksize           string `json:"@maxstacksize,omitempty"`
			Unlockedtocraft        string `json:"@unlockedtocraft"`
			Unlockedtoequip        string `json:"@unlockedtoequip"`
			Uicraftsoundstart      string `json:"@uicraftsoundstart,omitempty"`
			Uicraftsoundfinish     string `json:"@uicraftsoundfinish,omitempty"`
			Craftingcategory       string `json:"@craftingcategory,omitempty"`
			Craftingrequirements   struct {
				Silver           string `json:"@silver"`
				Amountcrafted    string `json:"@amountcrafted"`
				Forcesinglecraft string `json:"@forcesinglecraft"`
				Craftingfocus    string `json:"@craftingfocus"`
				Time             string `json:"@time"`
				Craftresource    struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"craftresource"`
			} `json:"craftingrequirements,omitempty"`
			Enchantments struct {
				Enchantment []struct {
					Enchantmentlevel     string `json:"@enchantmentlevel"`
					Consumespell         string `json:"@consumespell"`
					Dummyitempower       string `json:"@dummyitempower"`
					Abilitypower         string `json:"@abilitypower"`
					Craftingrequirements struct {
						Time          string `json:"@time"`
						Amountcrafted string `json:"@amountcrafted"`
						Craftingfocus string `json:"@craftingfocus"`
						Craftresource []struct {
							Uniquename string `json:"@uniquename"`
							Count      string `json:"@count"`
						} `json:"craftresource"`
					} `json:"craftingrequirements"`
					Upgraderequirements struct {
						Upgraderesource struct {
							Uniquename string `json:"@uniquename"`
							Count      string `json:"@count"`
						} `json:"upgraderesource"`
					} `json:"upgraderequirements"`
				} `json:"enchantment"`
			} `json:"enchantments,omitempty"`
			Famevalue        string `json:"@famevalue,omitempty"`
			Uispriteoverlay1 string `json:"@uispriteoverlay1,omitempty"`
			Tradable         string `json:"@tradable,omitempty"`
			Foodcategory     string `json:"@foodcategory,omitempty"`
		} `json:"consumableitem"`
		Consumablefrominventoryitem []struct {
			Uniquename           string `json:"@uniquename"`
			Tradable             string `json:"@tradable,omitempty"`
			Uisprite             string `json:"@uisprite,omitempty"`
			Abilitypower         string `json:"@abilitypower"`
			Consumespell         string `json:"@consumespell"`
			Shopcategory         string `json:"@shopcategory"`
			Shopsubcategory1     string `json:"@shopsubcategory1"`
			Tier                 string `json:"@tier"`
			Weight               string `json:"@weight"`
			Dummyitempower       string `json:"@dummyitempower,omitempty"`
			Maxstacksize         string `json:"@maxstacksize,omitempty"`
			Unlockedtocraft      string `json:"@unlockedtocraft"`
			Unlockedtouse        string `json:"@unlockedtouse,omitempty"`
			Uicraftsoundstart    string `json:"@uicraftsoundstart,omitempty"`
			Uicraftsoundfinish   string `json:"@uicraftsoundfinish,omitempty"`
			Uispriteoverlay1     string `json:"@uispriteoverlay1,omitempty"`
			Itemvalue            string `json:"@itemvalue,omitempty"`
			Craftingrequirements struct {
				Time            string `json:"@time"`
				Swaptransaction string `json:"@swaptransaction"`
				Currency        struct {
					Uniquename string `json:"@uniquename"`
					Amount     string `json:"@amount"`
				} `json:"currency"`
			} `json:"craftingrequirements,omitempty"`
			Allowfullstackusage          string `json:"@allowfullstackusage,omitempty"`
			Logconsumption               string `json:"@logconsumption,omitempty"`
			Estimatedmarketvalueoverride string `json:"@estimatedmarketvalueoverride,omitempty"`
			Enchantmentlevel             string `json:"@enchantmentlevel,omitempty"`
			Showinmarketplace            string `json:"@showinmarketplace,omitempty"`
			Descriptionlocatag           string `json:"@descriptionlocatag,omitempty"`
			Hidefromplayeroncontext      string `json:"@hidefromplayeroncontext,omitempty"`
			Salvageable                  string `json:"@salvageable,omitempty"`
			Namelocatag                  string `json:"@namelocatag,omitempty"`
			Uispriteoverlay2             string `json:"@uispriteoverlay2,omitempty"`
			Decayitem                    string `json:"@decayitem,omitempty"`
			Decayondatetime              string `json:"@decayondatetime,omitempty"`
			Hidedecayoverlay             string `json:"@hidedecayoverlay,omitempty"`
			AudioInfo                    struct {
				Name string `json:"@name"`
			} `json:"AudioInfo,omitempty"`
		} `json:"consumablefrominventoryitem"`
		Equipmentitem []struct {
			Uniquename                   string `json:"@uniquename"`
			Uisprite                     string `json:"@uisprite,omitempty"`
			Maxqualitylevel              string `json:"@maxqualitylevel"`
			Abilitypower                 string `json:"@abilitypower"`
			Slottype                     string `json:"@slottype"`
			Itempowerprogressiontype     string `json:"@itempowerprogressiontype,omitempty"`
			Shopcategory                 string `json:"@shopcategory"`
			Shopsubcategory1             string `json:"@shopsubcategory1"`
			Uicraftsoundstart            string `json:"@uicraftsoundstart,omitempty"`
			Uicraftsoundfinish           string `json:"@uicraftsoundfinish,omitempty"`
			Skincount                    string `json:"@skincount,omitempty"`
			Tier                         string `json:"@tier"`
			Weight                       string `json:"@weight"`
			Activespellslots             string `json:"@activespellslots"`
			Passivespellslots            string `json:"@passivespellslots"`
			Physicalarmor                string `json:"@physicalarmor"`
			Magicresistance              string `json:"@magicresistance"`
			Durability                   string `json:"@durability"`
			DurabilitylossAttack         string `json:"@durabilityloss_attack"`
			DurabilitylossSpelluse       string `json:"@durabilityloss_spelluse"`
			DurabilitylossReceivedattack string `json:"@durabilityloss_receivedattack"`
			DurabilitylossReceivedspell  string `json:"@durabilityloss_receivedspell"`
			Offhandanimationtype         string `json:"@offhandanimationtype,omitempty"`
			Unlockedtocraft              string `json:"@unlockedtocraft"`
			Unlockedtoequip              string `json:"@unlockedtoequip"`
			Hitpointsmax                 string `json:"@hitpointsmax"`
			Hitpointsregenerationbonus   string `json:"@hitpointsregenerationbonus"`
			Energymax                    string `json:"@energymax"`
			Energyregenerationbonus      string `json:"@energyregenerationbonus"`
			Crowdcontrolresistance       string `json:"@crowdcontrolresistance"`
			Itempower                    string `json:"@itempower"`
			Physicalattackdamagebonus    string `json:"@physicalattackdamagebonus"`
			Magicattackdamagebonus       string `json:"@magicattackdamagebonus"`
			Physicalspelldamagebonus     string `json:"@physicalspelldamagebonus"`
			Magicspelldamagebonus        string `json:"@magicspelldamagebonus"`
			Healbonus                    string `json:"@healbonus"`
			Bonusccdurationvsplayers     string `json:"@bonusccdurationvsplayers"`
			Bonusccdurationvsmobs        string `json:"@bonusccdurationvsmobs"`
			Threatbonus                  string `json:"@threatbonus"`
			Magiccooldownreduction       string `json:"@magiccooldownreduction"`
			Bonusdefensevsplayers        string `json:"@bonusdefensevsplayers,omitempty"`
			Bonusdefensevsmobs           string `json:"@bonusdefensevsmobs,omitempty"`
			Magiccasttimereduction       string `json:"@magiccasttimereduction,omitempty"`
			Attackspeedbonus             string `json:"@attackspeedbonus,omitempty"`
			Movespeedbonus               string `json:"@movespeedbonus,omitempty"`
			Healmodifier                 string `json:"@healmodifier,omitempty"`
			Canbeovercharged             string `json:"@canbeovercharged,omitempty"`
			Showinmarketplace            string `json:"@showinmarketplace,omitempty"`
			Energycostreduction          string `json:"@energycostreduction,omitempty"`
			Masterymodifier              string `json:"@masterymodifier,omitempty"`
			Craftingrequirements         struct {
				Silver        string `json:"@silver"`
				Time          string `json:"@time"`
				Craftingfocus string `json:"@craftingfocus"`
				Craftresource struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"craftresource"`
			} `json:"craftingrequirements,omitempty"`
			Craftingcategory      string `json:"@craftingcategory,omitempty"`
			Combatspecachievement string `json:"@combatspecachievement,omitempty"`
			SocketPreset          struct {
				Name string `json:"@name"`
			} `json:"SocketPreset,omitempty"`
			Enchantments struct {
				Enchantment []struct {
					Enchantmentlevel     string `json:"@enchantmentlevel"`
					Itempower            string `json:"@itempower"`
					Durability           string `json:"@durability"`
					Craftingrequirements struct {
						Silver        string `json:"@silver"`
						Time          string `json:"@time"`
						Craftingfocus string `json:"@craftingfocus"`
						Craftresource []struct {
							Uniquename       string `json:"@uniquename"`
							Enchantmentlevel string `json:"@enchantmentlevel"`
							Count            string `json:"@count"`
						} `json:"craftresource"`
					} `json:"craftingrequirements"`
					Upgraderequirements struct {
						Upgraderesource struct {
							Uniquename string `json:"@uniquename"`
							Count      string `json:"@count"`
						} `json:"upgraderesource"`
					} `json:"upgraderequirements,omitempty"`
				} `json:"enchantment"`
			} `json:"enchantments,omitempty"`
			Destinycraftfamefactor string `json:"@destinycraftfamefactor,omitempty"`
			AssetVfxPreset         struct {
				Name string `json:"@name"`
			} `json:"AssetVfxPreset,omitempty"`
			Famevalue         string `json:"@famevalue,omitempty"`
			Craftingspelllist struct {
				Craftspell struct {
					Uniquename string `json:"@uniquename"`
				} `json:"craftspell"`
			} `json:"craftingspelllist,omitempty"`
			Descriptionlocatag      string `json:"@descriptionlocatag,omitempty"`
			Tradable                string `json:"@tradable,omitempty"`
			Movespeed               string `json:"@movespeed,omitempty"`
			Maxload                 string `json:"@maxload,omitempty"`
			Facestate               string `json:"@facestate,omitempty"`
			Hidefromplayeroncontext string `json:"@hidefromplayeroncontext,omitempty"`
			Scalemodifier           string `json:"@scalemodifier,omitempty"`
			AudioInfo               struct {
				Name string `json:"@name"`
			} `json:"AudioInfo,omitempty"`
			Requiredaccesslevel   string `json:"@requiredaccesslevel,omitempty"`
			Uispriteoverlay1      string `json:"@uispriteoverlay1,omitempty"`
			Beardstate            string `json:"@beardstate,omitempty"`
			Mesh                  string `json:"@mesh,omitempty"`
			Omitmesh              string `json:"@omitmesh,omitempty"`
			Mainhandanimationtype string `json:"@mainhandanimationtype,omitempty"`
			Skincolormodifier     string `json:"@skincolormodifier,omitempty"`
		} `json:"equipmentitem"`
		Weapon []struct {
			Uniquename                   string `json:"@uniquename"`
			Mesh                         string `json:"@mesh,omitempty"`
			Uisprite                     string `json:"@uisprite,omitempty"`
			Maxqualitylevel              string `json:"@maxqualitylevel"`
			Abilitypower                 string `json:"@abilitypower"`
			Slottype                     string `json:"@slottype"`
			Shopcategory                 string `json:"@shopcategory"`
			Shopsubcategory1             string `json:"@shopsubcategory1"`
			Attacktype                   string `json:"@attacktype"`
			Attackdamage                 string `json:"@attackdamage,omitempty"`
			Attackspeed                  string `json:"@attackspeed,omitempty"`
			Attackrange                  string `json:"@attackrange,omitempty"`
			Twohanded                    string `json:"@twohanded"`
			Tier                         string `json:"@tier"`
			Weight                       string `json:"@weight"`
			Activespellslots             string `json:"@activespellslots"`
			Passivespellslots            string `json:"@passivespellslots"`
			Durability                   string `json:"@durability"`
			DurabilitylossAttack         string `json:"@durabilityloss_attack"`
			DurabilitylossSpelluse       string `json:"@durabilityloss_spelluse"`
			DurabilitylossReceivedattack string `json:"@durabilityloss_receivedattack"`
			DurabilitylossReceivedspell  string `json:"@durabilityloss_receivedspell"`
			Mainhandanimationtype        string `json:"@mainhandanimationtype"`
			Unlockedtocraft              string `json:"@unlockedtocraft"`
			Unlockedtoequip              string `json:"@unlockedtoequip"`
			Itempower                    string `json:"@itempower"`
			Unequipincombat              string `json:"@unequipincombat,omitempty"`
			Uicraftsoundstart            string `json:"@uicraftsoundstart"`
			Uicraftsoundfinish           string `json:"@uicraftsoundfinish"`
			Canbeovercharged             string `json:"@canbeovercharged,omitempty"`
			Canharvest                   struct {
				Resourcetype string `json:"@resourcetype"`
			} `json:"canharvest,omitempty"`
			// Craftingrequirements struct {
			// Silver        string `json:"@silver"`
			// Time          string `json:"@time"`
			// Craftingfocus string `json:"@craftingfocus"`
			// Craftresource []struct {
			// 	Uniquename string `json:"@uniquename"`
			// 	Count      string `json:"@count"`
			// } `json:"craftresource"`
			// } `json:"craftingrequirements"`
			Craftingrequirements interface{} `json:"craftingrequirements,omitempty"`
			AudioInfo            struct {
				Name string `json:"@name"`
			} `json:"AudioInfo,omitempty"`
			SocketPreset struct {
				Name string `json:"@name"`
			} `json:"SocketPreset,omitempty"`
			Craftingcategory   string `json:"@craftingcategory,omitempty"`
			Descriptionlocatag string `json:"@descriptionlocatag,omitempty"`
			Craftingspelllist  struct {
				Craftspell struct {
					Uniquename string `json:"@uniquename"`
				} `json:"craftspell"`
			} `json:"craftingspelllist,omitempty"`
			Attackbuildingdamage           string `json:"@attackbuildingdamage,omitempty"`
			Fishing                        string `json:"@fishing,omitempty"`
			Fishingspeed                   string `json:"@fishingspeed,omitempty"`
			Physicalspelldamagebonus       string `json:"@physicalspelldamagebonus,omitempty"`
			Magicspelldamagebonus          string `json:"@magicspelldamagebonus,omitempty"`
			Fxbonename                     string `json:"@fxbonename,omitempty"`
			Fxboneoffset                   string `json:"@fxboneoffset,omitempty"`
			Hitpointsmax                   string `json:"@hitpointsmax,omitempty"`
			Hitpointsregenerationbonus     string `json:"@hitpointsregenerationbonus,omitempty"`
			Itempowerprogressiontype       string `json:"@itempowerprogressiontype,omitempty"`
			Focusfireprotectionpenetration string `json:"@focusfireprotectionpenetration,omitempty"`
			Healmodifier                   string `json:"@healmodifier,omitempty"`
			Masterymodifier                string `json:"@masterymodifier,omitempty"`
			Physicalattackdamagebonus      string `json:"@physicalattackdamagebonus,omitempty"`
			Magicattackdamagebonus         string `json:"@magicattackdamagebonus,omitempty"`
			Attackvariations               struct {
				Attackchaintime string `json:"@attackchaintime"`
				Attack          []struct {
					Attackdamagefactor     string `json:"@attackdamagefactor"`
					Attackspeedfactor      string `json:"@attackspeedfactor"`
					Attackdamagetimefactor string `json:"@attackdamagetimefactor"`
					Attacktype             string `json:"@attacktype"`
				} `json:"attack"`
			} `json:"attackvariations,omitempty"`
			Projectile struct {
				Prefab      string `json:"@prefab"`
				Startsocket string `json:"@startsocket"`
				Hitsocket   string `json:"@hitsocket"`
				Flyspeed    string `json:"@flyspeed"`
				Impactvfx   struct {
					Prefab       string `json:"@prefab"`
					Impactsocket string `json:"@impactsocket"`
				} `json:"impactvfx"`
				AudioInfo struct {
					Name string `json:"@name"`
				} `json:"AudioInfo"`
			} `json:"projectile,omitempty"`
			Combatspecachievement string `json:"@combatspecachievement,omitempty"`
			Enchantments          struct {
				Enchantment []struct {
					Enchantmentlevel     string `json:"@enchantmentlevel"`
					Itempower            string `json:"@itempower"`
					Durability           string `json:"@durability"`
					Craftingrequirements struct {
						Silver        string `json:"@silver"`
						Time          string `json:"@time"`
						Craftingfocus string `json:"@craftingfocus"`
						Craftresource struct {
							Uniquename       string `json:"@uniquename"`
							Enchantmentlevel string `json:"@enchantmentlevel"`
							Count            string `json:"@count"`
						} `json:"craftresource"`
					} `json:"craftingrequirements"`
					Upgraderequirements struct {
						Upgraderesource struct {
							Uniquename string `json:"@uniquename"`
							Count      string `json:"@count"`
						} `json:"upgraderesource"`
					} `json:"upgraderequirements,omitempty"`
				} `json:"enchantment"`
			} `json:"enchantments,omitempty"`
			Destinycraftfamefactor string `json:"@destinycraftfamefactor,omitempty"`
			Showinmarketplace      string `json:"@showinmarketplace,omitempty"`
			AssetVfxPreset         struct {
				Name string `json:"@name"`
			} `json:"AssetVfxPreset,omitempty"`
			Triggershealingsickness     string `json:"@triggershealingsickness,omitempty"`
			Hidefromplayeroncontext     string `json:"@hidefromplayeroncontext,omitempty"`
			Requiredaccesslevel         string `json:"@requiredaccesslevel,omitempty"`
			Attackdamagetimefactor      string `json:"@attackdamagetimefactor,omitempty"`
			MainhandanimationtypeAttack string `json:"@mainhandanimationtype_attack,omitempty"`
			Attackvfx                   struct {
				Prefab string `json:"@prefab"`
				Target string `json:"@target"`
			} `json:"attackvfx,omitempty"`
			Tradable                 string `json:"@tradable,omitempty"`
			Skincount                string `json:"@skincount,omitempty"`
			Physicalarmor            string `json:"@physicalarmor,omitempty"`
			Magicresistance          string `json:"@magicresistance,omitempty"`
			Energymax                string `json:"@energymax,omitempty"`
			Energyregenerationbonus  string `json:"@energyregenerationbonus,omitempty"`
			Crowdcontrolresistance   string `json:"@crowdcontrolresistance,omitempty"`
			Healbonus                string `json:"@healbonus,omitempty"`
			Bonusccdurationvsplayers string `json:"@bonusccdurationvsplayers,omitempty"`
			Bonusccdurationvsmobs    string `json:"@bonusccdurationvsmobs,omitempty"`
			Threatbonus              string `json:"@threatbonus,omitempty"`
			Magiccooldownreduction   string `json:"@magiccooldownreduction,omitempty"`
			Bonusdefensevsplayers    string `json:"@bonusdefensevsplayers,omitempty"`
			Bonusdefensevsmobs       string `json:"@bonusdefensevsmobs,omitempty"`
			Magiccasttimereduction   string `json:"@magiccasttimereduction,omitempty"`
			Attackspeedbonus         string `json:"@attackspeedbonus,omitempty"`
			Movespeedbonus           string `json:"@movespeedbonus,omitempty"`
			Namelocatag              string `json:"@namelocatag,omitempty"`
		} `json:"weapon"`
		Mount []struct {
			Uniquename                          string `json:"@uniquename"`
			Mountcategory                       string `json:"@mountcategory,omitempty"`
			Maxqualitylevel                     string `json:"@maxqualitylevel"`
			Itempower                           string `json:"@itempower"`
			Abilitypower                        string `json:"@abilitypower"`
			Slottype                            string `json:"@slottype"`
			Shopcategory                        string `json:"@shopcategory"`
			Shopsubcategory1                    string `json:"@shopsubcategory1"`
			Mountedbuff                         string `json:"@mountedbuff"`
			Halfmountedbuff                     string `json:"@halfmountedbuff"`
			Tier                                string `json:"@tier"`
			Weight                              string `json:"@weight"`
			Activespellslots                    string `json:"@activespellslots"`
			Passivespellslots                   string `json:"@passivespellslots"`
			Durability                          string `json:"@durability"`
			DurabilitylossAttack                string `json:"@durabilityloss_attack"`
			DurabilitylossSpelluse              string `json:"@durabilityloss_spelluse"`
			DurabilitylossReceivedattack        string `json:"@durabilityloss_receivedattack"`
			DurabilitylossReceivedspell         string `json:"@durabilityloss_receivedspell"`
			DurabilitylossReceivedattackMounted string `json:"@durabilityloss_receivedattack_mounted"`
			DurabilitylossReceivedspellMounted  string `json:"@durabilityloss_receivedspell_mounted"`
			DurabilitylossMounting              string `json:"@durabilityloss_mounting"`
			Unlockedtocraft                     string `json:"@unlockedtocraft"`
			Unlockedtoequip                     string `json:"@unlockedtoequip"`
			Mounttime                           string `json:"@mounttime"`
			Dismounttime                        string `json:"@dismounttime"`
			Mounthitpointsmax                   string `json:"@mounthitpointsmax"`
			Mounthitpointsregeneration          string `json:"@mounthitpointsregeneration"`
			Prefabname                          string `json:"@prefabname"`
			Prefabscaling                       string `json:"@prefabscaling"`
			Despawneffect                       string `json:"@despawneffect"`
			Despawneffectscaling                string `json:"@despawneffectscaling"`
			Remountdistance                     string `json:"@remountdistance"`
			Halfmountrange                      string `json:"@halfmountrange"`
			Forceddismountcooldown              string `json:"@forceddismountcooldown"`
			Forceddismountspellcooldown         string `json:"@forceddismountspellcooldown"`
			Fulldismountcooldown                string `json:"@fulldismountcooldown"`
			Remounttime                         string `json:"@remounttime"`
			Uicraftsoundstart                   string `json:"@uicraftsoundstart"`
			Uicraftsoundfinish                  string `json:"@uicraftsoundfinish"`
			Dismountbuff                        string `json:"@dismountbuff"`
			Forceddismountbuff                  string `json:"@forceddismountbuff,omitempty"`
			Hostiledismountbuff                 string `json:"@hostiledismountbuff,omitempty"`
			Longhostiledismountbuff             string `json:"@longhostiledismountbuff,omitempty"`
			Showinmarketplace                   string `json:"@showinmarketplace,omitempty"`
			Hidefromplayeroncontext             string `json:"@hidefromplayeroncontext,omitempty"`
			VfxAddonKeyword                     string `json:"@vfxAddonKeyword"`
			Craftingrequirements                struct {
				Silver        string `json:"@silver"`
				Time          string `json:"@time"`
				Craftingfocus string `json:"@craftingfocus"`
				Craftresource []struct {
					Uniquename      string `json:"@uniquename"`
					Count           string `json:"@count"`
					Maxreturnamount string `json:"@maxreturnamount,omitempty"`
				} `json:"craftresource"`
			} `json:"craftingrequirements"`
			Craftingspelllist struct {
				Craftspell struct {
					Uniquename string `json:"@uniquename"`
				} `json:"craftspell"`
			} `json:"craftingspelllist,omitempty"`
			SocketPreset struct {
				Name string `json:"@name"`
			} `json:"SocketPreset"`
			AudioInfo struct {
				Name string `json:"@name"`
			} `json:"AudioInfo"`
			AssetVfxPreset struct {
				Name string `json:"@name"`
			} `json:"AssetVfxPreset,omitempty"`
			FootStepVfxPreset struct {
				Name string `json:"@name"`
			} `json:"FootStepVfxPreset,omitempty"`
			Canusetownportal        string `json:"@canusetownportal,omitempty"`
			Uisprite                string `json:"@uisprite,omitempty"`
			Halfmountprefaboverride string `json:"@halfmountprefaboverride,omitempty"`
			Nametagoffset           string `json:"@nametagoffset,omitempty"`
			Canuseingvg             string `json:"@canuseingvg,omitempty"`
			Mountspelllist          struct {
				Mountspell []struct {
					Uniquename string `json:"@uniquename"`
					Spellslot  string `json:"@spellslot"`
				} `json:"mountspell"`
			} `json:"mountspelllist,omitempty"`
			Enchantmentlevel       string `json:"@enchantmentlevel,omitempty"`
			Canuseinfactionwarfare string `json:"@canuseinfactionwarfare,omitempty"`
			Itemvalue              string `json:"@itemvalue,omitempty"`
			Descriptionlocatag     string `json:"@descriptionlocatag,omitempty"`
		} `json:"mount"`
		Furnitureitem []struct {
			Uniquename                 string `json:"@uniquename"`
			Shopcategory               string `json:"@shopcategory"`
			Shopsubcategory1           string `json:"@shopsubcategory1"`
			Tier                       string `json:"@tier"`
			Durability                 string `json:"@durability,omitempty"`
			Durabilitylossperdayfactor string `json:"@durabilitylossperdayfactor,omitempty"`
			Weight                     string `json:"@weight"`
			Unlockedtocraft            string `json:"@unlockedtocraft,omitempty"`
			Placeableindoors           string `json:"@placeableindoors"`
			Placeableoutdoors          string `json:"@placeableoutdoors"`
			Placeableindungeons        string `json:"@placeableindungeons"`
			Placeableinexpeditions     string `json:"@placeableinexpeditions"`
			Accessrightspreset         string `json:"@accessrightspreset,omitempty"`
			Uicraftsoundstart          string `json:"@uicraftsoundstart,omitempty"`
			Uicraftsoundfinish         string `json:"@uicraftsoundfinish,omitempty"`
			Craftingrequirements       struct {
				Time          string `json:"@time"`
				Craftingfocus string `json:"@craftingfocus"`
				Craftresource []struct {
					Uniquename string `json:"@uniquename"`
					Count      string `json:"@count"`
				} `json:"craftresource"`
			} `json:"craftingrequirements,omitempty"`
			AudioInfo struct {
				Name string `json:"@name"`
			} `json:"AudioInfo,omitempty"`
			Repairkit struct {
				Repaircostfactor string `json:"@repaircostfactor"`
				Maxtier          string `json:"@maxtier"`
			} `json:"repairkit,omitempty"`
			Customizewithguildlogo string `json:"@customizewithguildlogo,omitempty"`
			Uisprite               string `json:"@uisprite,omitempty"`
			Enchantmentlevel       string `json:"@enchantmentlevel,omitempty"`
			Tile                   string `json:"@tile,omitempty"`
			Itemvalue              string `json:"@itemvalue,omitempty"`
			Container              struct {
				Capacity    string `json:"@capacity"`
				Weightlimit string `json:"@weightlimit"`
			} `json:"container,omitempty"`
			Showinmarketplace          string      `json:"@showinmarketplace,omitempty"`
			Residencyslots             string      `json:"@residencyslots,omitempty"`
			Labourerfurnituretype      string      `json:"@labourerfurnituretype,omitempty"`
			Labourersaffected          string      `json:"@labourersaffected,omitempty"`
			Labourerhappiness          string      `json:"@labourerhappiness,omitempty"`
			Labourersperfurnitureitem  string      `json:"@labourersperfurnitureitem,omitempty"`
			Placeableonlyonislands     string      `json:"@placeableonlyonislands,omitempty"`
			Descriptionlocatag         string      `json:"@descriptionlocatag,omitempty"`
			Salvageable                string      `json:"@salvageable,omitempty"`
			Hidefromplayeroncontext    string      `json:"@hidefromplayeroncontext,omitempty"`
			Namelocatag                string      `json:"@namelocatag,omitempty"`
			Cheatprovider              interface{} `json:"cheatprovider,omitempty"`
			Maxstacksize               string      `json:"@maxstacksize,omitempty"`
			Tradable                   string      `json:"@tradable,omitempty"`
			Durabilitylossperusefactor string      `json:"@durabilitylossperusefactor,omitempty"`
		} `json:"furnitureitem"`
		Mountskin []struct {
			Uniquename           string `json:"@uniquename"`
			Prefabname           string `json:"@prefabname"`
			Prefabscaling        string `json:"@prefabscaling"`
			Despawneffect        string `json:"@despawneffect"`
			Despawneffectscaling string `json:"@despawneffectscaling"`
			VfxAddonKeyword      string `json:"@vfxAddonKeyword,omitempty"`
			SocketPreset         struct {
				Name string `json:"@name"`
			} `json:"SocketPreset"`
			AssetVfxPreset struct {
				Name string `json:"@name"`
			} `json:"AssetVfxPreset,omitempty"`
			AudioInfo struct {
				Name string `json:"@name"`
			} `json:"AudioInfo"`
			FootStepVfxPreset struct {
				Name string `json:"@name"`
			} `json:"FootStepVfxPreset,omitempty"`
			Uisprite      string `json:"@uisprite,omitempty"`
			Nametagoffset string `json:"@nametagoffset,omitempty"`
		} `json:"mountskin"`
		Journalitem []struct {
			Salvageable          string `json:"@salvageable"`
			Uniquename           string `json:"@uniquename"`
			Tier                 string `json:"@tier"`
			Maxfame              string `json:"@maxfame"`
			Baselootamount       string `json:"@baselootamount"`
			Shopcategory         string `json:"@shopcategory"`
			Shopsubcategory1     string `json:"@shopsubcategory1"`
			Weight               string `json:"@weight"`
			Unlockedtocraft      string `json:"@unlockedtocraft"`
			Fasttravelfactor     string `json:"@fasttravelfactor,omitempty"`
			Craftingrequirements struct {
				Silver          string `json:"@silver"`
				Time            string `json:"@time"`
				Swaptransaction string `json:"@swaptransaction"`
			} `json:"craftingrequirements"`
			Famefillingmissions struct {
				Gatherfame struct {
					Mintier   string `json:"@mintier"`
					Value     string `json:"@value"`
					Validitem []struct {
						ID string `json:"@id"`
					} `json:"validitem"`
				} `json:"gatherfame"`
			} `json:"famefillingmissions"`
			Lootlist struct {
				Loot struct {
					Itemname     string `json:"@itemname"`
					Itemamount   string `json:"@itemamount"`
					Weight       string `json:"@weight"`
					Labourerfame string `json:"@labourerfame"`
				} `json:"loot"`
			} `json:"lootlist"`
			Uisprite string `json:"@uisprite,omitempty"`
		} `json:"journalitem"`
		Labourercontract []struct {
			Uniquename       string `json:"@uniquename"`
			Tier             string `json:"@tier"`
			Shopcategory     string `json:"@shopcategory"`
			Shopsubcategory1 string `json:"@shopsubcategory1"`
			Unlockedtocraft  string `json:"@unlockedtocraft"`
			Weight           string `json:"@weight"`
		} `json:"labourercontract"`
		Transformationweapon []struct {
			Uniquename                     string `json:"@uniquename"`
			Transformation                 string `json:"@transformation"`
			Mesh                           string `json:"@mesh,omitempty"`
			Uisprite                       string `json:"@uisprite"`
			Maxqualitylevel                string `json:"@maxqualitylevel"`
			Abilitypower                   string `json:"@abilitypower"`
			Physicalspelldamagebonus       string `json:"@physicalspelldamagebonus"`
			Magicspelldamagebonus          string `json:"@magicspelldamagebonus"`
			Slottype                       string `json:"@slottype"`
			Shopcategory                   string `json:"@shopcategory"`
			Shopsubcategory1               string `json:"@shopsubcategory1"`
			Attacktype                     string `json:"@attacktype"`
			Mainhandanimationtype          string `json:"@mainhandanimationtype"`
			Attackdamage                   string `json:"@attackdamage"`
			Attackspeed                    string `json:"@attackspeed"`
			Attackrange                    string `json:"@attackrange"`
			Twohanded                      string `json:"@twohanded"`
			Tier                           string `json:"@tier"`
			Weight                         string `json:"@weight"`
			Activespellslots               string `json:"@activespellslots"`
			Passivespellslots              string `json:"@passivespellslots"`
			Durability                     string `json:"@durability"`
			DurabilitylossAttack           string `json:"@durabilityloss_attack"`
			DurabilitylossSpelluse         string `json:"@durabilityloss_spelluse"`
			DurabilitylossReceivedattack   string `json:"@durabilityloss_receivedattack"`
			DurabilitylossReceivedspell    string `json:"@durabilityloss_receivedspell"`
			Fxbonename                     string `json:"@fxbonename"`
			Fxboneoffset                   string `json:"@fxboneoffset"`
			Unlockedtocraft                string `json:"@unlockedtocraft"`
			Unlockedtoequip                string `json:"@unlockedtoequip"`
			Hitpointsmax                   string `json:"@hitpointsmax"`
			Itempower                      string `json:"@itempower"`
			Hitpointsregenerationbonus     string `json:"@hitpointsregenerationbonus"`
			Itempowerprogressiontype       string `json:"@itempowerprogressiontype"`
			Focusfireprotectionpenetration string `json:"@focusfireprotectionpenetration"`
			Uicraftsoundstart              string `json:"@uicraftsoundstart"`
			Uicraftsoundfinish             string `json:"@uicraftsoundfinish"`
			Craftingcategory               string `json:"@craftingcategory"`
			Healmodifier                   string `json:"@healmodifier"`
			Canbeovercharged               string `json:"@canbeovercharged"`
			Masterymodifier                string `json:"@masterymodifier"`
			Physicalattackdamagebonus      string `json:"@physicalattackdamagebonus"`
			Magicattackdamagebonus         string `json:"@magicattackdamagebonus"`
			Descriptionlocatag             string `json:"@descriptionlocatag"`
			Attackvariations               struct {
				Attackchaintime string `json:"@attackchaintime"`
				Attack          []struct {
					Attacktype             string `json:"@attacktype"`
					Attackdamagefactor     string `json:"@attackdamagefactor"`
					Attackdamagetimefactor string `json:"@attackdamagetimefactor"`
					Attackspeedfactor      string `json:"@attackspeedfactor"`
				} `json:"attack"`
			} `json:"attackvariations"`
			Projectile struct {
				Prefab      string `json:"@prefab"`
				Startsocket string `json:"@startsocket"`
				Hitsocket   string `json:"@hitsocket"`
				Flyspeed    string `json:"@flyspeed"`
				Impactvfx   struct {
					Prefab       string `json:"@prefab"`
					Impactsocket string `json:"@impactsocket"`
				} `json:"impactvfx"`
			} `json:"projectile"`
			SocketPreset struct {
				Name string `json:"@name"`
			} `json:"SocketPreset"`
			// Craftingrequirements struct {
			// 	Silver        string `json:"@silver"`
			// 	Time          string `json:"@time"`
			// 	Craftingfocus string `json:"@craftingfocus"`
			// 	Craftresource []struct {
			// 		Uniquename      string `json:"@uniquename"`
			// 		Count           string `json:"@count"`
			// 		Maxreturnamount string `json:"@maxreturnamount,omitempty"`
			// 	} `json:"craftresource"`
			// } `json:"craftingrequirements"`
			Craftingrequirements interface{} `json:"craftingrequirements,omitempty"`
			Craftingspelllist    struct {
				Craftspell []struct {
					Uniquename string `json:"@uniquename"`
					Slots      string `json:"@slots,omitempty"`
					Tag        string `json:"@tag,omitempty"`
				} `json:"craftspell"`
			} `json:"craftingspelllist"`
			AudioInfo struct {
				Name string `json:"@name"`
			} `json:"AudioInfo"`
			Combatspecachievement string `json:"@combatspecachievement,omitempty"`
			Enchantments          struct {
				Enchantment []struct {
					Enchantmentlevel     string `json:"@enchantmentlevel"`
					Itempower            string `json:"@itempower"`
					Durability           string `json:"@durability"`
					Craftingrequirements struct {
						Silver        string `json:"@silver"`
						Time          string `json:"@time"`
						Craftingfocus string `json:"@craftingfocus"`
						Craftresource []struct {
							Uniquename       string `json:"@uniquename"`
							Enchantmentlevel string `json:"@enchantmentlevel,omitempty"`
							Count            string `json:"@count"`
							Maxreturnamount  string `json:"@maxreturnamount,omitempty"`
						} `json:"craftresource"`
					} `json:"craftingrequirements"`
					Upgraderequirements struct {
						Upgraderesource struct {
							Uniquename string `json:"@uniquename"`
							Count      string `json:"@count"`
						} `json:"upgraderesource"`
					} `json:"upgraderequirements,omitempty"`
				} `json:"enchantment"`
			} `json:"enchantments,omitempty"`
			Destinycraftfamefactor string `json:"@destinycraftfamefactor,omitempty"`
		} `json:"transformationweapon"`
		Crystalleagueitem []struct {
			Uniquename               string `json:"@uniquename"`
			Uisprite                 string `json:"@uisprite"`
			Shopcategory             string `json:"@shopcategory"`
			Shopsubcategory1         string `json:"@shopsubcategory1"`
			Tier                     string `json:"@tier"`
			Enchantmentlevel         string `json:"@enchantmentlevel"`
			Resourcetype             string `json:"@resourcetype"`
			Weight                   string `json:"@weight"`
			Maxstacksize             string `json:"@maxstacksize"`
			Namelocatag              string `json:"@namelocatag"`
			Descriptionlocatag       string `json:"@descriptionlocatag"`
			Descvariable0            string `json:"@descvariable0"`
			Salvageable              string `json:"@salvageable"`
			Itemvalue                string `json:"@itemvalue"`
			Tradable                 string `json:"@tradable"`
			Unlockedtocraft          string `json:"@unlockedtocraft"`
			Canbestoredinbattlevault string `json:"@canbestoredinbattlevault"`
			Craftingrequirements     struct {
				Silver string `json:"@silver"`
			} `json:"craftingrequirements"`
			Uispriteoverlay1  string `json:"@uispriteoverlay1,omitempty"`
			Showinmarketplace string `json:"@showinmarketplace,omitempty"`
		} `json:"crystalleagueitem"`
		Killtrophy struct {
			Uniquename              string `json:"@uniquename"`
			Shopcategory            string `json:"@shopcategory"`
			Shopsubcategory1        string `json:"@shopsubcategory1"`
			Tier                    string `json:"@tier"`
			Weight                  string `json:"@weight"`
			Durability              string `json:"@durability"`
			Unlockedtocraft         string `json:"@unlockedtocraft"`
			Placeableindoors        string `json:"@placeableindoors"`
			Placeableoutdoors       string `json:"@placeableoutdoors"`
			Placeableindungeons     string `json:"@placeableindungeons"`
			Placeableinexpeditions  string `json:"@placeableinexpeditions"`
			Placeableonlyonislands  string `json:"@placeableonlyonislands"`
			Uisprite                string `json:"@uisprite"`
			Descriptionlocatag      string `json:"@descriptionlocatag"`
			Showinmarketplace       string `json:"@showinmarketplace"`
			Hidefromplayeroncontext string `json:"@hidefromplayeroncontext"`
			Craftingrequirements    struct {
				Gold string `json:"@gold"`
				Time string `json:"@time"`
			} `json:"craftingrequirements"`
		} `json:"killtrophy"`
	} `json:"items"`
}
