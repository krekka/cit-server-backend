package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "vd7mw1zpr59yjsu",
			"created": "2023-11-27 08:55:23.279Z",
			"updated": "2023-11-27 08:55:23.279Z",
			"name": "downloads",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "8yo0akhe",
					"name": "files",
					"type": "relation",
					"required": true,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "qqeiqs7tgx9gr04",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": null,
						"displayFields": null
					}
				}
			],
			"indexes": [],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("vd7mw1zpr59yjsu")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
