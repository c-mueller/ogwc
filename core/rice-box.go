package core

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "entities.json",
		FileModTime: time.Unix(1540526959, 0),
		Content:     string("{\n  \"202\": {\n    \"id\": 202,\n    \"name\": \"Kleiner Transporter\",\n    \"cargo_capacity\": 5000,\n    \"cost\": {\n      \"metal\": 2000,\n      \"crystal\": 2000,\n      \"deuterium\": 0\n    }\n  },\n  \"203\": {\n    \"id\": 203,\n    \"name\": \"Großer Transporter\",\n    \"cargo_capacity\": 25000,\n    \"cost\": {\n      \"metal\": 6000,\n      \"crystal\": 6000,\n      \"deuterium\": 0\n    }\n  },\n  \"204\": {\n    \"id\": 204,\n    \"name\": \"Leichter Jäger\",\n    \"cargo_capacity\": 50,\n    \"cost\": {\n      \"metal\": 3000,\n      \"crystal\": 1000,\n      \"deuterium\": 0\n    }\n  },\n  \"205\": {\n    \"id\": 205,\n    \"name\": \"Schwerer Jäger\",\n    \"cargo_capacity\": 100,\n    \"cost\": {\n      \"metal\": 6000,\n      \"crystal\": 4000,\n      \"deuterium\": 0\n    }\n  },\n  \"206\": {\n    \"id\": 206,\n    \"name\": \"Kreuzer\",\n    \"cargo_capacity\": 800,\n    \"cost\": {\n      \"metal\": 20000,\n      \"crystal\": 7000,\n      \"deuterium\": 2000\n    }\n  },\n  \"207\": {\n    \"id\": 207,\n    \"name\": \"Schlachtschiff\",\n    \"cargo_capacity\": 1500,\n    \"cost\": {\n      \"metal\": 45000,\n      \"crystal\": 15000,\n      \"deuterium\": 0\n    }\n  },\n  \"208\": {\n    \"id\": 208,\n    \"name\": \"Kolonieschiff\",\n    \"cargo_capacity\": 7500,\n    \"cost\": {\n      \"metal\": 10000,\n      \"crystal\": 20000,\n      \"deuterium\": 10000\n    }\n  },\n  \"209\": {\n    \"id\": 209,\n    \"name\": \"Recycler\",\n    \"cargo_capacity\": 20000,\n    \"cost\": {\n      \"metal\": 10000,\n      \"crystal\": 6000,\n      \"deuterium\": 2000\n    }\n  },\n  \"211\": {\n    \"id\": 211,\n    \"name\": \"Bomber\",\n    \"cargo_capacity\": 500,\n    \"cost\": {\n      \"metal\": 50000,\n      \"crystal\": 25000,\n      \"deuterium\": 15000\n    }\n  },\n  \"212\": {\n    \"id\": 212,\n    \"name\": \"Solarsatelit\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 0,\n      \"crystal\": 2000,\n      \"deuterium\": 500\n    }\n  },\n  \"213\": {\n    \"id\": 213,\n    \"name\": \"Zerstörer\",\n    \"cargo_capacity\": 2000,\n    \"cost\": {\n      \"metal\": 60000,\n      \"crystal\": 50000,\n      \"deuterium\": 15000\n    }\n  },\n  \"214\": {\n    \"id\": 214,\n    \"name\": \"Todesstern\",\n    \"cargo_capacity\": 1000000,\n    \"cost\": {\n      \"metal\": 5000000,\n      \"crystal\": 4000000,\n      \"deuterium\": 1000000\n    }\n  },\n  \"215\": {\n    \"id\": 215,\n    \"name\": \"Schlachtkreuzer\",\n    \"cargo_capacity\": 750,\n    \"cost\": {\n      \"metal\": 30000,\n      \"crystal\": 40000,\n      \"deuterium\": 15000\n    }\n  },\n  \"401\": {\n    \"id\": 401,\n    \"name\": \"Raketenwerfer\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 2000,\n      \"crystal\": 0,\n      \"deuterium\": 0\n    }\n  },\n  \"402\": {\n    \"id\": 402,\n    \"name\": \"Leichtes Lasergeschütz\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 1500,\n      \"crystal\": 500,\n      \"deuterium\": 0\n    }\n  },\n  \"403\": {\n    \"id\": 403,\n    \"name\": \"Schweres Lasergeschütz\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 6000,\n      \"crystal\": 2000,\n      \"deuterium\": 0\n    }\n  },\n  \"404\": {\n    \"id\": 404,\n    \"name\": \"Gaußkanone\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 20000,\n      \"crystal\": 15000,\n      \"deuterium\": 2000\n    }\n  },\n  \"405\": {\n    \"id\": 405,\n    \"name\": \"Ionenkanone\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 2000,\n      \"crystal\": 6000,\n      \"deuterium\": 0\n    }\n  },\n  \"406\": {\n    \"id\": 406,\n    \"name\": \"Plasmawerfer\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 50000,\n      \"crystal\": 50000,\n      \"deuterium\": 30000\n    }\n  },\n  \"407\": {\n    \"id\": 407,\n    \"name\": \"Kleine Schlidkuppel\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 10000,\n      \"crystal\": 10000,\n      \"deuterium\": 0\n    }\n  },\n  \"408\": {\n    \"id\": 408,\n    \"name\": \"Große Schlidkuppel\",\n    \"cargo_capacity\": 0,\n    \"cost\": {\n      \"metal\": 50000,\n      \"crystal\": 50000,\n      \"deuterium\": 0\n    }\n  }\n}"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1540526959, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "entities.json"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`config`, &embedded.EmbeddedBox{
		Name: `config`,
		Time: time.Unix(1540526959, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"entities.json": file2,
		},
	})
}
