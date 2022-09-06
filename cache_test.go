package cachebench_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/bluele/gcache"
	"github.com/coocood/freecache"
	hashicorp "github.com/hashicorp/golang-lru"
	koding "github.com/koding/cache"
	"github.com/muesli/cache2go"
	"github.com/patrickmn/go-cache"
)

func keySm(i int) string {
	return fmt.Sprintf("Item:%d", i)
}

func keyMd(i int) string {
	return fmt.Sprintf("ItemItemItem:%d", i)
}

func keyLg(i int) string {
	return fmt.Sprintf("ItemItemItemItemItemItemItemItemItem:%d", i)
}

func valSm(i int) string {
	// like a very big counter
	return fmt.Sprintf("%d 18446744073709551615", i)
}

func valMd(i int) string {
	// like a user profile in json
	return fmt.Sprintf(`{"id": "767ce8f5-670a-448a-913a-641c553f44ff", "email": "kazhuravlev+%d@fastmail.com", "name": "Kirill Zhuravlev", "avatar": "https://example.com/images/users/767ce8f5-670a-448a-913a-641c553f44ff.webp"}`, i)
}

func valLg(i int) string {
	// like cached image or html page
	return fmt.Sprintf(`%d [{"_id": "6317b5b42b7cd6f775c5f991","index": 0,"guid": "2b652a7d-a76a-4612-b502-06c123d796aa","isActive": false,"balance": "$1,726.28","picture": "http://placehold.it/32x32","age": 32,"eyeColor": "blue","name": "Nunez Oneill","gender": "male","company": "ENQUILITY","email": "nunezoneill@enquility.com","phone": "+1 (893) 586-2864","address": "437 Clifton Place, Marysville, Idaho, 4642",
"about": "Lorem Lorem tempor et labore do in aliquip. Culpa proident reprehenderit anim tempor magna duis nisi nostrud Lorem. Minim cillum ex exmagna nostrud irure ut sit labore sunt tempor do. Exercitation nisi incididunt aliqua culpa dolore elit.\r\n","registered": "2016-03-02T07:23:02 -04:00","latitude": 31.0147,"longitude": 100.084605,"tags": ["elit","id","ex","sit","commodo","ipsum","nostrud"],"friends": [{"id": 0,"name": "Margret Colon"},{"id": 1,"name": "Rodriquez Key"},{"id": 2,"name": "Tameka Meyer"}],"greeting": "Hello, Nunez Oneill! You have 3 unread messages.","favoriteFruit": "apple"},{"_id": "6317b5b4c32cfcf63ed0ae37","index": 1,"guid": "cc918574-49c5-4f6f-813f-c4c4cd007eba","isActive": false,"balance": "$3,640.88","picture": "http://placehold.it/32x32","age": 31,"eyeColor": "blue","name": "Enid Huffman","gender": "female","company": "ZOLARITY","email": "enidhuffman@zolarity.com","phone": "+1 (808) 565-2624","address": "548 Channel Avenue, Navarre, North Dakota, 6554",
"about": "Duis mollit cupidatat consectetur dolore ea consequat consectetur nostrud eu esse officia culpa. Duis adipisicing qui qui adipisicingmollit. Amet ex sint deserunt velit sint.\r\n","registered": "2021-07-14T10:24:18 -04:00","latitude": 32.731459,"longitude": -54.128755,"tags": ["fugiat","velit","aute","et","sit","deserunt","laboris"],"friends": [{"id": 0,"name": "Bray Ramos"},{"id": 1,"name": "Huff Lewis"},{"id": 2,"name": "Kramer Love"}],"greeting": "Hello, Enid Huffman! You have 6 unread messages.","favoriteFruit": "banana"},{"_id": "6317b5b47da52d9dad59ed72","index": 2,"guid": "26fe9e76-fdcc-4a78-94fe-208b669ad4e2","isActive": false,"balance": "$1,545.41","picture": "http://placehold.it/32x32","age": 35,"eyeColor": "blue","name": "Karen Ramsey","gender": "female","company": "ENERSOL","email": "karenramsey@enersol.com","phone": "+1 (920) 474-2377","address": "650 Jackson Street, Brownsville, Georgia, 1453",
"about": "Minim laborum excepteur eiusmod fugiat voluptate minim excepteur nulla duis amet sunt. Sint tempor excepteur minim do veniam sunt nostrud nisi voluptate. Minim Lorem ex sunt occaecat tempor deserunt Lorem excepteur ex et aute officia. Veniam in eu id eu elit excepteur consectetur ullamco reprehenderit Lorem sunt ut. Duis ea eu deserunt ut commodo dolore labore ex velit irure commodo enim. Aliquip irure tempor esse sunt nullaad sunt ullamco reprehenderit sit. Eiusmod irure id eu occaecat amet nulla cupidatat id quis et.\r\n","registered": "2020-07-26T05:23:26 -04:00","latitude": -51.12491,"longitude": -101.088261,"tags": ["sit","qui","est","labore","non","reprehenderit","duis"],"friends": [{"id": 0,"name": "Janine Terry"},{"id": 1,"name": "Earline Lang"},{"id": 2,"name": "Sherry Foreman"}],"greeting": "Hello, Karen Ramsey! You have 4 unread messages.","favoriteFruit": "banana"},{"_id": "6317b5b4131038d7611a5164","index": 3,"guid": "2b74a3c9-5d43-4be5-b6f6-75da3bded67b","isActive": true,"balance": "$1,801.90","picture": "http://placehold.it/32x32","age": 29,"eyeColor": "brown","name": "Beryl Hubbard","gender": "female","company": "BIOSPAN","email": "berylhubbard@biospan.com","phone": "+1 (926) 421-2314","address": "924 Nixon Court, Nicholson, Wyoming, 6207",
"about": "Eiusmod duis ut minim ea ad anim dolore do ullamco nulla in dolore sunt. Sit consectetur et consectetur voluptate. Est adipisicing doofficia veniam duis quis adipisicing elit magna.\r\n","registered": "2016-12-10T02:08:26 -04:00","latitude": -66.018589,"longitude": -112.140531,"tags": ["aute","qui","aliqua","dolore","nulla","ullamco","exercitation"],"friends": [{"id": 0,"name": "Burke Mann"},{"id": 1,"name": "Raymond Gordon"},{"id": 2,"name": "Russell Reynolds"}],"greeting": "Hello, Beryl Hubbard! You have 4 unread messages.","favoriteFruit": "strawberry"},{"_id": "6317b5b4e6c4c90ecb0599a5","index": 4,"guid": "3b4b35b0-3339-4814-883d-b4821aaffc50","isActive": false,"balance": "$2,437.32","picture": "http://placehold.it/32x32","age": 33,"eyeColor": "brown","name": "Allyson Buchanan","gender": "female","company": "PRIMORDIA","email": "allysonbuchanan@primordia.com","phone": "+1 (968) 464-2841","address": "730 Maple Street, Cucumber, New Jersey, 9685",
"about": "Culpa dolor est aliqua velit duis esse ea sunt aute incididunt. Labore ad in nisi pariatur minim sunt exercitation quis quis proident adirure eu aliqua. Elit minim occaecat elit irure ad excepteur qui nisi elit tempor tempor.\r\n","registered": "2016-07-12T04:39:54 -04:00","latitude": 30.899757,"longitude": 41.461522,"tags": ["velit","id","pariatur","ipsum","ex","ipsum","amet"],"friends": [{"id": 0,"name": "Nixon Washington"},{"id": 1,"name": "Conway Cervantes"},{"id": 2,"name": "Brandie Pope"}],"greeting": "Hello, Allyson Buchanan! You have 3 unread messages.","favoriteFruit": "banana"},{"_id": "6317b5b4268ae1429a08cacf","index": 5,"guid": "c72f84f9-3314-4205-9906-9b7fee27da3b","isActive": false,"balance": "$3,342.30","picture": "http://placehold.it/32x32","age": 36,"eyeColor": "blue","name": "Aileen West","gender": "female","company": "ZIDOX","email": "aileenwest@zidox.com","phone": "+1 (863) 549-3593","address": "596 Irving Street, Cochranville, Arkansas, 6274",
"about": "Anim aliqua eiusmod velit ullamco in officia incididunt. Amet do proident nisi pariatur ut duis irure dolor sunt est duis. Aliqua veniamipsum irure labore. Duis dolor et tempor aute commodo eu nulla quis dolore nisi non. Enim nulla deserunt aute Lorem labore.\r\n","registered": "2017-03-22T08:40:00 -04:00","latitude": 7.176416,"longitude": 57.586061,"tags": ["velit","consequat","in","proident","ad","id","magna"],"friends": [{"id": 0,"name": "Myrtle Knowles"},{"id": 1,"name": "Stanton Grant"},{"id": 2,"name": "Sadie Parker"}],"greeting": "Hello, Aileen West! You have 4 unread messages.","favoriteFruit": "apple"},{"_id": "6317b5b4397050f18d7bf33b","index": 6,"guid": "7994b600-2c13-4d3e-9282-609e60bf5653","isActive": true,"balance": "$2,441.63","picture": "http://placehold.it/32x32","age": 40,"eyeColor": "green","name": "Jolene Miranda","gender": "female","company": "CONCILITY","email": "jolenemiranda@concility.com","phone": "+1 (948) 513-2374","address": "820 Dekoven Court, Day, Puerto Rico, 197",
"about": "Consequat culpa non ad ipsum est sit commodo quis ut laborum eu deserunt. Amet enim et duis esse aliquip magna incididunt irure. Elit enim esse elit in cupidatat veniam pariatur labore. Irure Lorem enim quis magna ad occaecat cupidatat. Velit occaecat tempor Lorem cillum. Dolor ut culpa elit consectetur officia aliqua tempor non cupidatat cupidatat sit. Commodo id commodo Lorem exercitation laboris anim proident velit ad commodo dolaborum cillum exercitation.\r\n","registered": "2020-02-29T09:44:28 -04:00","latitude": 51.354347,"longitude": -7.084379,"tags": ["consectetur","sint","ea","nulla","veniam","labore","eu"],"friends": [{"id": 0,"name": "Macdonald Rojas"},{"id": 1,"name": "Amparo Caldwell"},{"id": 2,"name": "Kaufman Mendoza"}],"greeting": "Hello, Jolene Miranda! You have 9 unread messages.","favoriteFruit": "banana"},{"_id": "6317b5b4fc159d82aafbd209","index": 7,"guid": "719a62af-0e7b-415e-8a68-6d0c3e7b45d2","isActive": false,"balance": "$2,880.68","picture": "http://placehold.it/32x32","age": 32,"eyeColor": "blue","name": "Melody Shelton","gender": "female","company": "QUAREX","email": "melodyshelton@quarex.com","phone": "+1 (957) 577-3338","address": "389 Johnson Avenue, Lemoyne, Pennsylvania, 5019",
"about": "Aliquip est enim aute enim nulla fugiat in cupidatat in do qui. Consectetur aliquip labore minim ullamco esse est sint. Consectetur esse incididunt nostrud pariatur laboris proident non culpa fugiat nulla. Minim elit nulla ea ex qui voluptate ex anim fugiat nisi veniam velit occaecat. Mollit minim consequat sit consequat ullamco deserunt deserunt qui dolor laboris. Dolor velit officia cupidatat esse irure. Fugiat cupidatat sintpariatur exercitation fugiat anim Lorem ea amet in duis nisi nisi sint.\r\n","registered": "2014-09-20T11:43:07 -04:00","latitude": -63.123925,"longitude": 99.232033,"tags": ["est","eiusmod","nostrud","aliqua","minim","irure","quis"],"friends": [{"id": 0,"name": "Lynda Castaneda"},{"id": 1,"name": "Ware Rush"},{"id": 2,"name": "Mckinney Gross"}],"greeting": "Hello, Melody Shelton! You have 8 unread messages.","favoriteFruit": "apple"},{"_id": "6317b5b41f8bcfb810cee18d","index": 8,"guid": "5e5fcce9-66f2-4a39-97ae-369647c1831c","isActive": true,"balance": "$1,945.88","picture": "http://placehold.it/32x32","age": 21,"eyeColor": "blue","name": "Santana Snider","gender": "male","company": "DOGSPA","email": "santanasnider@dogspa.com","phone": "+1 (842) 597-3770","address": "832 Glendale Court, Sylvanite, New York, 8328",
"about": "Dolore minim culpa laborum laborum dolore irure est dolore in incididunt Lorem irure. Nostrud id ipsum reprehenderit voluptate sunt nisi. Lorem duis laboris eu dolore. Ea et id Lorem nostrud incididunt voluptate amet reprehenderit id anim et Lorem. Tempor qui excepteur ex cupidatat est laboris cillum excepteur ea tempor deserunt occaecat exercitation labore. Non culpa et duis sint ex elit eu Lorem labore deserunt labore dolor euaute. Eu irure nulla deserunt commodo culpa.\r\n","registered": "2017-01-28T03:55:37 -04:00","latitude": 32.404408,"longitude": 28.935429,"tags": ["mollit","dolor","enim","sunt","sit","labore","cupidatat"],"friends": [{"id": 0,"name": "Kendra Haley"},{"id": 1,"name": "Morris Clements"},{"id": 2,"name": "Frances Saunders"}],"greeting": "Hello, Santana Snider! You have 6 unread messages.","favoriteFruit": "strawberry"},{"_id": "6317b5b4bea3f93b01c56c70","index": 9,"guid": "03eb9736-a774-47c0-8313-1e34adcdc9c1","isActive": false,"balance": "$2,825.40","picture": "http://placehold.it/32x32","age": 34,"eyeColor": "green","name": "Shelia Singleton","gender": "female","company": "FLUM","email": "sheliasingleton@flum.com","phone": "+1 (981) 538-2366","address": "574 Juliana Place, Nipinnawasee, Arizona, 2484",
"about": "Deserunt aute amet qui aliqua aliqua esse quis do ad veniam id. Ipsum cupidatat proident labore aliquip ea duis. Velit voluptate dolore ut proident culpa culpa ex cillum. Ullamco exercitation reprehenderit tempor ex. Ullamco consequat in ullamco pariatur velit magna ut veniam ullamcoveniam qui voluptate. Reprehenderit laborum culpa irure anim consequat magna quis incididunt duis irure Lorem. Non fugiat cupidatat eu sunt.\r\n","registered": "2020-03-07T02:34:49 -04:00","latitude": 11.046685,"longitude": 25.043706,"tags": ["esse","consectetur","fugiat","magna","sint","voluptate","consectetur"],"friends": [{"id": 0,"name": "Saundra Everett"},{"id": 1,"name": "Odonnell Bradford"},{"id": 2,"name": "Linda Britt"}],"greeting": "Hello, Shelia Singleton! You have 2 unread messages.","favoriteFruit": "banana"}]`, i)
}

var funcs = []struct {
	name   string
	getKey func(i int) string
	getVal func(i int) string
}{
	{
		name:   "sm",
		getKey: keySm,
		getVal: valSm,
	},
	{
		name:   "md",
		getKey: keyMd,
		getVal: valMd,
	},
	{
		name:   "lg",
		getKey: keyLg,
		getVal: valLg,
	},
}

// We will be storing many short strings as the key and value
func BenchmarkKodingCache(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c := koding.NewMemoryWithTTL(time.Duration(60) * time.Second)

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Set(toKey(i), toVal(i))
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, err := c.Get(toKey(i))
				if err == nil {
					_ = value
				}
			}
		})
	}
}

func BenchmarkHashicorpLRU(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c, _ := hashicorp.New(10)

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Add(toKey(i), toVal(i))
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, err := c.Get(toKey(i))
				if err == true {
					_ = value
				}
			}
		})
	}
}

func BenchmarkCache2Go(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c := cache2go.Cache("test")

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Add(toKey(i), 1*time.Minute, toVal(i))
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, err := c.Value(toKey(i))
				if err == nil {
					_ = value
				}
			}
		})
	}
}

func BenchmarkGoCache(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c := cache.New(1*time.Minute, 5*time.Minute)

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Add(toKey(i), toVal(i), cache.DefaultExpiration)
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, found := c.Get(toKey(i))
				if found {
					_ = value
				}
			}
		})
	}
}

func BenchmarkFreecache(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c := freecache.NewCache(1024 * 1024 * 5)

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Set([]byte(toKey(i)), []byte(toVal(i)), 60)
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, err := c.Get([]byte(toKey(i)))
				if err == nil {
					_ = value
				}
			}
		})
	}
}

func BenchmarkBigCache(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c, _ := bigcache.NewBigCache(bigcache.Config{
			Shards:             1024,
			LifeWindow:         10 * time.Minute,
			MaxEntriesInWindow: 1000 * 10 * 60,
			MaxEntrySize:       500,
			HardMaxCacheSize:   10,
		})

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Set(toKey(i), []byte(toVal(i)))
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, err := c.Get(toKey(i))
				if err == nil {
					_ = value
				}
			}
		})
	}
}

func BenchmarkGCache(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		c := gcache.New(b.N).LRU().Build()

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c.Set(toKey(i), toVal(i))
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, err := c.Get(toKey(i))
				if err == nil {
					_ = value
				}
			}
		})
	}
}

func BenchmarkSyncMap(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		var m sync.Map

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.Store(toKey(i), toVal(i))
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, found := m.Load(toKey(i))
				if found {
					_ = value
				}
			}
		})
	}
}

func BenchmarkMap(b *testing.B) {
	for _, fn := range funcs {
		toKey := fn.getKey
		toVal := fn.getVal
		m := make(map[string]string)

		b.Run("Set_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m[toKey(i)] = toVal(i)
			}
		})

		b.Run("Get_"+fn.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				value, found := m[toKey(i)]
				if found {
					_ = value
				}
			}
		})
	}
}
