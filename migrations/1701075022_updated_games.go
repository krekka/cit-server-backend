package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("91u3l1khgdrul8r")
		if err != nil {
			return err
		}

		// add
		new_news := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ne4a2mwi",
			"name": "news",
			"type": "relation",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"collectionId": "1e7mqwyoj2olz9d",
				"cascadeDelete": false,
				"minSelect": null,
				"maxSelect": null,
				"displayFields": null
			}
		}`), new_news)
		collection.Schema.AddField(new_news)

		// update
		edit_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ogs6naez",
			"name": "name",
			"type": "text",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_name)
		collection.Schema.AddField(edit_name)

		// update
		edit_bannerImages := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "o7ncyyos",
			"name": "bannerImages",
			"type": "json",
			"required": true,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_bannerImages)
		collection.Schema.AddField(edit_bannerImages)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("91u3l1khgdrul8r")
		if err != nil {
			return err
		}

		// remove
		collection.Schema.RemoveField("ne4a2mwi")

		// update
		edit_name := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "ogs6naez",
			"name": "name",
			"type": "text",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"min": null,
				"max": null,
				"pattern": ""
			}
		}`), edit_name)
		collection.Schema.AddField(edit_name)

		// update
		edit_bannerImages := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "o7ncyyos",
			"name": "bannerImages",
			"type": "json",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_bannerImages)
		collection.Schema.AddField(edit_bannerImages)

		return dao.SaveCollection(collection)
	})
}
