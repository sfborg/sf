PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE version (sf__id TEXT NOT NULL) STRICT;
INSERT INTO version VALUES('v0.5.1');
CREATE TABLE metadata (
  col__id INTEGER PRIMARY KEY AUTOINCREMENT,
  col__doi TEXT DEFAULT '',
  col__title TEXT NOT NULL,
  col__alias TEXT DEFAULT '',
  col__description TEXT DEFAULT '',
  col__issued TEXT DEFAULT '',
  col__version TEXT DEFAULT '',
  col__keywords TEXT DEFAULT '',
  col__geographic_scope TEXT DEFAULT '',
  col__taxonomic_scope TEXT DEFAULT '',
  col__temporal_scope TEXT DEFAULT '',
  col__confidence INTEGER DEFAULT NULL,
  col__completeness INTEGER DEFAULT NULL,
  col__license TEXT DEFAULT '',
  col__url TEXT DEFAULT '',
  col__logo TEXT DEFAULT '',
  col__label TEXT DEFAULT '',
  col__citation TEXT DEFAULT '',
  col__private INTEGER DEFAULT NULL -- bool 
) STRICT;
CREATE TABLE contact (
  col__id INTEGER PRIMARY KEY AUTOINCREMENT,
  col__metadata_id INTEGER DEFAULT 1,
  col__orcid TEXT DEFAULT '',
  col__given TEXT NOT NULL,
  col__family TEXT NOT NULL,
  col__rorid TEXT DEFAULT '',
  col__department TEXT DEFAULT '',
  col__organisation TEXT DEFAULT '',
  col__city TEXT DEFAULT '',
  col__state TEXT DEFAULT '',
  col__country TEXT DEFAULT '',
  col__email TEXT NOT NULL,
  col__url TEXT DEFAULT '',
  col__note TEXT DEFAULT ''
) STRICT;
CREATE TABLE editor (
  col__id INTEGER PRIMARY KEY AUTOINCREMENT,
  col__metadata_id INTEGER DEFAULT 1,
  col__orcid TEXT DEFAULT '',
  col__given TEXT NOT NULL,
  col__family TEXT NOT NULL,
  col__rorid TEXT DEFAULT '',
  col__department TEXT DEFAULT '',
  col__organisation TEXT DEFAULT '',
  col__city TEXT DEFAULT '',
  col__state TEXT DEFAULT '',
  col__country TEXT DEFAULT '',
  col__email TEXT DEFAULT '',
  col__url TEXT DEFAULT '',
  col__note TEXT DEFAULT ''
) STRICT;
CREATE TABLE creator (
  col__id INTEGER PRIMARY KEY AUTOINCREMENT,
  col__metadata_id INTEGER DEFAULT 1,
  col__orcid TEXT DEFAULT '',
  col__given TEXT NOT NULL,
  col__family TEXT NOT NULL,
  col__rorid TEXT DEFAULT '',
  col__department TEXT DEFAULT '',
  col__organisation TEXT DEFAULT '',
  col__city TEXT DEFAULT '',
  col__state TEXT DEFAULT '',
  col__country TEXT DEFAULT '',
  col__email TEXT DEFAULT '',
  col__url TEXT DEFAULT '',
  col__note TEXT DEFAULT ''
) STRICT;
CREATE TABLE publisher (
  col__id INTEGER PRIMARY KEY AUTOINCREMENT,
  col__metadata_id INTEGER DEFAULT 1,
  col__orcid TEXT DEFAULT '',
  col__given TEXT DEFAULT '',
  col__family TEXT DEFAULT '',
  col__rorid TEXT DEFAULT '',
  col__department TEXT DEFAULT '',
  col__organisation TEXT DEFAULT '',
  col__city TEXT DEFAULT '',
  col__state TEXT DEFAULT '',
  col__country TEXT DEFAULT '',
  col__email TEXT DEFAULT '',
  col__url TEXT DEFAULT '',
  col__note TEXT DEFAULT ''
) STRICT;
CREATE TABLE contributor (
  col__id INTEGER PRIMARY KEY AUTOINCREMENT,
  col__metadata_id INTEGER DEFAULT 1,
  col__orcid TEXT DEFAULT '',
  col__given TEXT NOT NULL,
  col__family TEXT NOT NULL,
  col__rorid TEXT DEFAULT '',
  col__department TEXT DEFAULT '',
  col__organisation TEXT DEFAULT '',
  col__city TEXT DEFAULT '',
  col__state TEXT DEFAULT '',
  col__country TEXT DEFAULT '',
  col__email TEXT DEFAULT '',
  col__url TEXT DEFAULT '',
  col__note TEXT DEFAULT ''
) STRICT;
CREATE TABLE source (
  col__id TEXT PRIMARY KEY,
  col__metadata_id INTEGER DEFAULT 1,
  col__type TEXT DEFAULT '',
  col__title TEXT DEFAULT '',
  col__authors TEXT DEFAULT '',
  col__issued TEXT DEFAULT '',
  col__isbn TEXT DEFAULT ''
) STRICT;
CREATE TABLE author (
  col__id TEXT PRIMARY KEY,
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__alternative_id TEXT DEFAULT '', -- sep by ','
  col__given TEXT DEFAULT '',
  col__family TEXT NOT NULL,
  -- f. for filius,  Jr., etc
  col__suffix TEXT DEFAULT '',
  col__abbreviation_botany TEXT DEFAULT '',
  col__alternative_names TEXT DEFAULT '', -- separated by '|'
  col__sex_id TEXT REFERENCES sex(col__id) DEFAULT '',
  col__country TEXT DEFAULT '',
  col__birth TEXT DEFAULT '',
  col__birth_place TEXT DEFAULT '',
  col__death TEXT DEFAULT '',
  col__affiliation TEXT DEFAULT '',
  col__interest TEXT DEFAULT '',
  col__reference_id TEXT DEFAULT '', -- sep by ','
  -- url
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE reference (
  col__id TEXT PRIMARY KEY,
  col__alternative_id TEXT DEFAULT '', -- sep by ',', scope:id, id, URI/URN
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__citation TEXT DEFAULT '',
  col__type_id TEXT REFERENCES reference_type(col__id) DEFAULT '',
  -- author/s in format of either
  -- family1, given1; family2, given2; ..
  -- or
  -- given1 family1, given2 family2, ...
  col__author TEXT DEFAULT '',
  col__author_id TEXT DEFAULT '', -- 'ref' author, sep ','
  col__editor TEXT DEFAULT '', -- 'ref' author, sep ','
  col__editor_id TEXT DEFAULT '', -- 'ref' author, sep ','
  col__title TEXT DEFAULT '',
  col__title_short TEXT DEFAULT '',
  -- container_author is an author or a parent volume (book, journal) 
  col__container_author TEXT DEFAULT '',
  -- container_title of the parent container
  col__container_title TEXT DEFAULT '',
  -- container_title_short of the parent container
  col__container_title_short TEXT DEFAULT '',
  col__issued TEXT DEFAULT '', -- yyyy-mm-dd
  col__accessed TEXT DEFAULT '', -- yyyy-mm-dd
  -- collection_title of the parent volume
  col__collection_title TEXT DEFAULT '',
  -- collection_editor of the parent volume
  col__collection_editor TEXT DEFAULT '',
  col__volume TEXT DEFAULT '',
  col__issue TEXT DEFAULT '',
  -- edition number
  col__edition TEXT DEFAULT '',
  -- page number
  col__page TEXT DEFAULT '',
  col__publisher TEXT DEFAULT '',
  col__publisher_place TEXT DEFAULT '',
  -- version of the reference
  col__version TEXT DEFAULT '',
  col__isbn TEXT DEFAULT '',
  col__issn TEXT DEFAULT '',
  col__doi TEXT DEFAULT '',
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE name (
  col__id TEXT PRIMARY KEY,
  col__alternative_id TEXT DEFAULT '',
  col__source_id TEXT DEFAULT '',
  -- basionym_id TEXT DEFAULT '', -- use name_relation instead
  tw__taxon_name_id TEXT DEFAULT '', -- TaxonWorks taxon_names.id
  gn__scientific_name_string TEXT NOT NULL, -- full name with authorship (if given)
  gn__parse_quality INTEGER DEFAULT NULL,
  gn__canonical_simple TEXT DEFAULT '',
  gn__canonical_full TEXT DEFAULT '',
  gn__canonical_stemmed TEXT DEFAULT '',
  gn__cardinality INTEGER DEFAULT NULL,
  gn__virus INTEGER DEFAULT NULL, -- bool
  gn__hybrid TEXT DEFAULT '',
  gn__surrogate TEXT DEFAULT '',
  gn__authors TEXT DEFAULT '', -- separated by '|'
  gn__id TEXT DEFAULT '', -- UUID v5 generated for GN from name-string
  col__scientific_name TEXT NOT NULL, -- full canonical form
  col__authorship TEXT DEFAULT '', -- verbatim authorship
  col__rank_id TEXT REFERENCES rank(col__id) DEFAULT '',
  col__uninomial TEXT DEFAULT '',
  col__genus TEXT DEFAULT '',
  col__infrageneric_epithet TEXT DEFAULT '',
  col__specific_epithet TEXT DEFAULT '',
  col__infraspecific_epithet TEXT DEFAULT '',
  col__cultivar_epithet TEXT DEFAULT '',
  col__notho_id TEXT DEFAULT '', -- ref name_part
  col__original_spelling INTEGER DEFAULT NULL, -- bool
  col__combination_authorship TEXT DEFAULT '', -- separated by '|'
  col__combination_authorship_id TEXT DEFAULT '', -- separated by '|'
  col__combination_ex_authorship TEXT DEFAULT '', -- separated by '|'
  col__combination_ex_authorship_id TEXT DEFAULT '', -- separated by '|'
  col__combination_authorship_year TEXT DEFAULT '',
  col__basionym_authorship TEXT DEFAULT '', -- separated by '|'
  col__basionym_authorship_id TEXT DEFAULT '', -- separated by '|'
  col__basionym_ex_authorship TEXT DEFAULT '', -- separated by '|'
  col__basionym_ex_authorship_id TEXT DEFAULT '', -- separated by '|'
  col__basionym_authorship_year TEXT DEFAULT '',
  col__code_id TEXT REFERENCES nom_code(col__id) DEFAULT '',
  col__status_id TEXT REFERENCES nom_status(col__id) DEFAULT '',
  col__reference_id TEXT DEFAULT '', -- refs about taxon sep ','
  col__published_in_year TEXT DEFAULT '',
  col__published_in_page TEXT DEFAULT '',
  col__published_in_page_link TEXT DEFAULT '',
  col__gender_id TEXT REFERENCES gender(col__id) DEFAULT '',
  col__gender_agreement INTEGER DEFAULT NULL, -- bool
  col__etymology TEXT DEFAULT '',
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
INSERT INTO name VALUES('a4','','','','Rhea pennata',1,'Rhea pennata','Rhea pennata','Rhea pennat',2,0,'','','','c9fb794c-28dc-5213-af87-90a39016ea48','Rhea pennata','','UNRANKED','','Rhea','','pennata','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a1','','','','Rhea americana',1,'Rhea americana','Rhea americana','Rhea american',2,0,'','','','7223006c-735f-5669-8931-afa190462dfe','Rhea americana','','UNRANKED','','Rhea','','americana','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a2','','','','Struthio camelus syryacus',1,'Struthio camelus syryacus','Struthio camelus syryacus','Struthio camel syryac',3,0,'','','','b680d131-9c7d-5cd1-b480-e754f2e1f254','Struthio camelus syryacus','','UNRANKED','','Struthio','','camelus','syryacus','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a5','','','','Struthio camelus camelus',1,'Struthio camelus camelus','Struthio camelus camelus','Struthio camel camel',3,0,'','','','610661db-8036-50f1-ba7e-bc3903b6a26c','Struthio camelus camelus','','UNRANKED','','Struthio','','camelus','camelus','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a3','','','','Rhea americana nobilis',1,'Rhea americana nobilis','Rhea americana nobilis','Rhea american nobil',3,0,'','','','85c94df5-33a0-5550-89d2-216a7e75e564','Rhea americana nobilis','','UNRANKED','','Rhea','','americana','nobilis','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a6','','','','Nothocercus bonapartei intercedens',1,'Nothocercus bonapartei intercedens','Nothocercus bonapartei intercedens','Nothocercus bonaparte intercedens',3,0,'','','','2f0b82e4-4a24-5a46-abd2-7fb81ee4cf15','Nothocercus bonapartei intercedens','','UNRANKED','','Nothocercus','','bonapartei','intercedens','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a11','','','','Rhea pennata garleppi',1,'Rhea pennata garleppi','Rhea pennata garleppi','Rhea pennat garlepp',3,0,'','','','f3962ff5-0be0-5fb1-805c-651e08d44891','Rhea pennata garleppi','','UNRANKED','','Rhea','','pennata','garleppi','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a8','','','','Nothocercus bonaparteus',1,'Nothocercus bonaparteus','Nothocercus bonaparteus','Nothocercus bonaparte',2,0,'','','','9c69886f-4c9c-5ce2-94fe-c61aa40fbd71','Nothocercus bonaparteus','','UNRANKED','','Nothocercus','','bonaparteus','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a9','','','','Rhea pennata tarapacensis/garleppi',4,'','','',NULL,0,'','','','8f3ba7d1-df68-5e76-8639-6b517d381e0b','Rhea pennata tarapacensis/garleppi','','UNRANKED','','','','','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a12','','','','Nothocercus bonapartei',1,'Nothocercus bonapartei','Nothocercus bonapartei','Nothocercus bonaparte',2,0,'','','','776d4797-2d42-535d-a269-32d7b05d48e6','Nothocercus bonapartei','','UNRANKED','','Nothocercus','','bonapartei','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a7','','','','Rhea americana intermedia',1,'Rhea americana intermedia','Rhea americana intermedia','Rhea american intermed',3,0,'','','','7a3ce1b2-27fa-5b16-92d9-44aa54956b46','Rhea americana intermedia','','UNRANKED','','Rhea','','americana','intermedia','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a15','','','','Struthio camelus',1,'Struthio camelus','Struthio camelus','Struthio camel',2,0,'','','','538c17b3-9329-5cb5-adde-494851fdd3fd','Struthio camelus','','UNRANKED','','Struthio','','camelus','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a14','','','','Struthio molybdophanes',1,'Struthio molybdophanes','Struthio molybdophanes','Struthio molybdophan',2,0,'','','','7fd7d843-ce33-5721-9e97-f7aae55668eb','Struthio molybdophanes','','UNRANKED','','Struthio','','molybdophanes','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a13','','','','Rhea pennata tarapacensis',1,'Rhea pennata tarapacensis','Rhea pennata tarapacensis','Rhea pennat tarapacens',3,0,'','','','71366414-25df-5265-a1b8-1edd684bb4b1','Rhea pennata tarapacensis','','UNRANKED','','Rhea','','pennata','tarapacensis','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a17','','','','Nothocercus julius',1,'Nothocercus julius','Nothocercus julius','Nothocercus iul',2,0,'','','','42940433-a03e-5c3d-bbf7-a0c27887da59','Nothocercus julius','','UNRANKED','','Nothocercus','','julius','','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a18','','','','Struthio camelus massaieus',1,'Struthio camelus massaieus','Struthio camelus massaieus','Struthio camel massaie',3,0,'','','','76f73fd3-663b-50f9-a4e5-d0a3f51e0777','Struthio camelus massaieus','','UNRANKED','','Struthio','','camelus','massaieus','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a10','','','','Rhea pennata pennata',1,'Rhea pennata pennata','Rhea pennata pennata','Rhea pennat pennat',3,0,'','','','7f16f80d-3558-599f-abbd-228f1359de33','Rhea pennata pennata','','UNRANKED','','Rhea','','pennata','pennata','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a16','','','','Nothocercus bonapartei discrepans',1,'Nothocercus bonapartei discrepans','Nothocercus bonapartei discrepans','Nothocercus bonaparte discrepans',3,0,'','','','a27f9d99-4e0e-5e3d-b3bf-2538c25f5db2','Nothocercus bonapartei discrepans','','UNRANKED','','Nothocercus','','bonapartei','discrepans','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
INSERT INTO name VALUES('a19','','','','Rhea americana araneipes',1,'Rhea americana araneipes','Rhea americana araneipes','Rhea american araneip',3,0,'','','','75897b53-20fb-5356-823e-66bdadbe2ab9','Rhea americana araneipes','','UNRANKED','','Rhea','','americana','araneipes','','',NULL,'','','','','','','','','','','ZOOLOGICAL','ESTABLISHED','','','','','',NULL,'','','','','');
CREATE TABLE taxon (
  col__id TEXT PRIMARY KEY,
  col__alternative_id TEXT DEFAULT '', -- scope:id, id sep ','
  gn__local_id TEXT DEFAULT '', -- internal ID from the source
  gn__global_id TEXT DEFAULT '', -- GUID attached to the record.
  tw__otu_id TEXT DEFAULT '', -- TaxonWorks OTU ID for round-tripping
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__parent_id TEXT REFERENCES taxon(col__id) DEFAULT '',
  col__ordinal INTEGER DEFAULT NULL, -- for sorting
  col__branch_length INTEGER DEFAULT NULL, --length of 'bread crumbs'
  col__name_id TEXT NOT NULL REFERENCES name(col__id) DEFAULT '',
  col__name_phrase TEXT DEFAULT '', -- eg `sensu stricto` and other annotations
  col__according_to_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__according_to_page TEXT DEFAULT '',
  col__according_to_page_link TEXT DEFAULT '',
  col__scrutinizer TEXT DEFAULT '',
  col__scrutinizer_id TEXT DEFAULT '', -- ORCID usually
  col__scrutinizer_date TEXT DEFAULT '',
  col__status_id TEXT REFERENCES taxonomic_status(col__id) DEFAULT '',
  col__reference_id TEXT DEFAULT '', -- list of references about the taxon hypothesis
  col__extinct INTEGER DEFAULT NULL, -- bool
  col__temporal_range_start_id TEXT REFERENCES geo_time(col__id) DEFAULT '',
  col__temporal_range_end_id TEXT REFERENCES geo_time(col__id) DEFAULT '',
  col__environment_id TEXT DEFAULT '', -- environment ids sep by ','
  col__species TEXT DEFAULT '',
  col__section TEXT DEFAULT '',
  col__subgenus TEXT DEFAULT '',
  col__genus TEXT DEFAULT '',
  col__subtribe TEXT DEFAULT '',
  col__tribe TEXT DEFAULT '',
  col__subfamily TEXT DEFAULT '',
  col__family TEXT DEFAULT '',
  col__superfamily TEXT DEFAULT '',
  col__suborder TEXT DEFAULT '',
  col__order TEXT DEFAULT '',
  col__subclass TEXT DEFAULT '',
  col__class TEXT DEFAULT '',
  col__subphylum TEXT DEFAULT '',
  col__phylum TEXT DEFAULT '',
  col__kingdom TEXT DEFAULT '',
  sf__realm TEXT DEFAULT '',
  sf__species_id TEXT DEFAULT '', -- for flat classification IDs
  sf__section_id TEXT DEFAULT '',
  sf__subgenus_id TEXT DEFAULT '',
  sf__genus_id TEXT DEFAULT '',
  sf__subtribe_id TEXT DEFAULT '',
  sf__tribe_id TEXT DEFAULT '',
  sf__subfamily_id TEXT DEFAULT '',
  sf__family_id TEXT DEFAULT '',
  sf__superfamily_id TEXT DEFAULT '',
  sf__suborder_id TEXT DEFAULT '',
  sf__order_id TEXT DEFAULT '',
  sf__subclass_id TEXT DEFAULT '',
  sf__class_id TEXT DEFAULT '',
  sf__subphylum_id TEXT DEFAULT '',
  sf__phylum_id TEXT DEFAULT '',
  sf__kingdom_id TEXT DEFAULT '',
  sf__realm_id TEXT DEFAULT '',
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
INSERT INTO taxon VALUES('a4','','','','','','',NULL,NULL,'a4','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a1','','','','','','',NULL,NULL,'a1','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a2','','','','','','',NULL,NULL,'a2','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Struthionidae (Ostriches)','','','Struthioniformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a5','','','','','','',NULL,NULL,'a5','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Struthionidae (Ostriches)','','','Struthioniformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a3','','','','','','',NULL,NULL,'a3','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a6','','','','','','',NULL,NULL,'a6','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Tinamidae (Tinamous)','','','Tinamiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a11','','','','','','',NULL,NULL,'a11','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a8','','','','','','',NULL,NULL,'a8','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Tinamidae (Tinamous)','','','Tinamiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a9','','','','','','',NULL,NULL,'a9','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a12','','','','','','',NULL,NULL,'a12','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Tinamidae (Tinamous)','','','Tinamiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a7','','','','','','',NULL,NULL,'a7','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a15','','','','','','',NULL,NULL,'a15','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Struthionidae (Ostriches)','','','Struthioniformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a14','','','','','','',NULL,NULL,'a14','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Struthionidae (Ostriches)','','','Struthioniformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a13','','','','','','',NULL,NULL,'a13','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a17','','','','','','',NULL,NULL,'a17','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Tinamidae (Tinamous)','','','Tinamiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a18','','','','','','',NULL,NULL,'a18','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Struthionidae (Ostriches)','','','Struthioniformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a10','','','','','','',NULL,NULL,'a10','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a16','','','','','','',NULL,NULL,'a16','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Tinamidae (Tinamous)','','','Tinamiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
INSERT INTO taxon VALUES('a19','','','','','','',NULL,NULL,'a19','','','','','','','','ACCEPTED','',NULL,'','','','','','','','','','','Rheidae (Rheas)','','','Rheiformes','','Aves','','Chordata','Animalia','','','','','','','','','','','','','','','','','','','','','','');
CREATE TABLE synonym (
  col__id TEXT NOT NULL DEFAULT '', -- optional
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__name_id TEXT NOT NULL REFERENCES name(col__id) DEFAULT '',
  col__name_phrase TEXT DEFAULT '', -- annotation (eg `sensu lato` etc)
  col__according_to_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__status_id TEXT REFERENCES taxonomic_status(col__id) DEFAULT '',
  col__reference_id TEXT DEFAULT '', -- ids, sep by ',' about this synonym
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT '',
  CHECK (col__id = '' OR col__id != col__taxon_id)
) STRICT;
CREATE TABLE vernacular (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__name TEXT NOT NULL,
  col__transliteration TEXT DEFAULT '',
  col__language TEXT DEFAULT '',
  col__preferred INTEGER DEFAULT NULL, -- bool
  col__country TEXT DEFAULT '',
  col__area TEXT DEFAULT '',
  col__sex_id TEXT REFERENCES sex(col__id) DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE name_relation (
  col__name_id TEXT NOT NULL REFERENCES name(col__id) DEFAULT '',
  col__related_name_id TEXT NOT NULL REFERENCES name(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  -- nom_rel_type enum
  col__type_id TEXT NOT NULL REFERENCES nom_rel_type(col__id) DEFAULT '',
  tw__name_relationship_type TEXT DEFAULT '', -- canonical NOMEN ontology URI
  -- starting page number for the nomenclatural event
  col__page TEXT DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE type_material (
  col__id TEXT DEFAULT '', -- optional
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__name_id TEXT NOT NULL REFERENCES name(col__id) DEFAULT '',
  col__citation TEXT DEFAULT '',
  col__status_id TEXT REFERENCES type_status(col__id) DEFAULT '',
  col__institution_code TEXT DEFAULT '',
  col__catalog_number TEXT DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__locality TEXT DEFAULT '',
  col__country TEXT DEFAULT '',
  col__latitude REAL DEFAULT 0,
  col__longitude REAL DEFAULT 0,
  col__altitude int DEFAULT 0,
  col__host TEXT DEFAULT '',
  col__sex_id TEXT REFERENCES sex(col__id) DEFAULT '',
  col__date TEXT DEFAULT '',
  col__collector TEXT DEFAULT '',
  col__associated_sequences TEXT DEFAULT '',
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE distribution (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__area TEXT DEFAULT '',
  col__area_id TEXT DEFAULT '',
  col__gazetteer_id TEXT REFERENCES gazetteer(col__id) DEFAULT '',
  col__status_id TEXT REFERENCES distribution_status(col__id) DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE media (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__url TEXT NOT NULL, -- in CoLDP media is always a link
  col__type TEXT DEFAULT '', -- MIME type
  col__format TEXT DEFAULT '',
  col__title TEXT DEFAULT '',
  col__created TEXT DEFAULT '',
  col__creator TEXT DEFAULT '',
  col__license TEXT DEFAULT '',
  col__link TEXT DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE treatment (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__document TEXT NOT NULL,
  col__format TEXT DEFAULT '', -- HTML, XML, TXT
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE species_estimate (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__estimate INTEGER NOT NULL, -- estimated number of species
  col__type_id TEXT NOT NULL REFERENCES estimate_type(col__id) DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE taxon_property (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__property TEXT NOT NULL, -- name of the property
  col__value TEXT NOT NULL,
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__page TEXT DEFAULT '',
  col__ordinal INTEGER DEFAULT NULL, -- sorting value
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE species_interaction (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__related_taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__related_taxon_scientific_name TEXT DEFAULT '', -- id or hardcoded name?
  col__type_id TEXT NOT NULL REFERENCES species_interaction_type(col__id) DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE taxon_concept_relation (
  col__taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__related_taxon_id TEXT NOT NULL REFERENCES taxon(col__id) DEFAULT '',
  col__source_id TEXT REFERENCES source(col__id) DEFAULT '',
  col__type_id TEXT REFERENCES taxon_concept_rel_type(col__id) DEFAULT '',
  col__reference_id TEXT REFERENCES reference(col__id) DEFAULT '',
  col__remarks TEXT DEFAULT '',
  col__modified TEXT DEFAULT '',
  col__modified_by TEXT DEFAULT ''
) STRICT;
CREATE TABLE name_match (
  col__name_id TEXT REFERENCES name(col__id) DEFAULT '',
  gn__scientific_name_string TEXT DEFAULT '',
  ref_col__name_id TEXT DEFAULT '',
  ref_gn__scientific_name_string TEXT DEFAULT '',
  ref_col__scientific_name TEXT DEFAULT '',
  ref_col__authorship TEXT DEFAULT '',
  gn__match_id TEXT REFERENCES match_type(col__id) DEFAULT ''
) STRICT;
CREATE TABLE match_type (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO match_type VALUES('');
INSERT INTO match_type VALUES('EXACT');
INSERT INTO match_type VALUES('EXACT_PARTIAL');
INSERT INTO match_type VALUES('FUZZY');
INSERT INTO match_type VALUES('FUZZY_PARTIAL');
CREATE TABLE nom_code (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO nom_code VALUES('');
INSERT INTO nom_code VALUES('BACTERIAL');
INSERT INTO nom_code VALUES('BOTANICAL');
INSERT INTO nom_code VALUES('CULTIVARS');
INSERT INTO nom_code VALUES('PHYTOSOCIOLOGICAL');
INSERT INTO nom_code VALUES('VIRUS');
INSERT INTO nom_code VALUES('ZOOLOGICAL');
CREATE TABLE name_part (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO name_part VALUES('');
INSERT INTO name_part VALUES('GENERIC');
INSERT INTO name_part VALUES('INFRAGENERIC');
INSERT INTO name_part VALUES('SPECIFIC');
INSERT INTO name_part VALUES('INFRASPECIFIC');
CREATE TABLE gender (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO gender VALUES('');
INSERT INTO gender VALUES('MASCULINE');
INSERT INTO gender VALUES('FEMININE');
INSERT INTO gender VALUES('NEUTRAL');
CREATE TABLE sex (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO sex VALUES('');
INSERT INTO sex VALUES('MALE');
INSERT INTO sex VALUES('FEMALE');
INSERT INTO sex VALUES('HERMAPHRODITE');
CREATE TABLE estimate_type (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO estimate_type VALUES('');
INSERT INTO estimate_type VALUES('SPECIES_EXTINCT');
INSERT INTO estimate_type VALUES('SPECIES_LIVING');
INSERT INTO estimate_type VALUES('ESTIMATED_SPECIES');
CREATE TABLE distribution_status (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO distribution_status VALUES('');
INSERT INTO distribution_status VALUES('NATIVE');
INSERT INTO distribution_status VALUES('DOMESTICATED');
INSERT INTO distribution_status VALUES('ALIEN');
INSERT INTO distribution_status VALUES('UNCERTAIN');
CREATE TABLE type_status (
  col__id TEXT PRIMARY KEY,
  col__name TEXT,
  col__root TEXT REFERENCES type_status(col__id),
  col__primary INTEGER, -- bool
  col__codes TEXT -- nom codes sep ',' 
) STRICT;
INSERT INTO type_status VALUES('','','',0,'');
INSERT INTO type_status VALUES('OTHER','other','OTHER',0,'');
INSERT INTO type_status VALUES('HOMOEOTYPE','homoeotype','HOMOEOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('PLESIOTYPE','plesiotype','PLESIOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('PLASTOTYPE','plastotype','PLASTOTYPE',0,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('PLASTOSYNTYPE','plastosyntype','SYNTYPE',0,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('PLASTOPARATYPE','plastoparatype','PARATYPE',0,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('PLASTONEOTYPE','plastoneotype','NEOTYPE',0,'');
INSERT INTO type_status VALUES('PLASTOLECTOTYPE','plastolectotype','LECTOTYPE',0,'');
INSERT INTO type_status VALUES('PLASTOISOTYPE','plastoisotype','HOLOTYPE',0,'');
INSERT INTO type_status VALUES('PLASTOHOLOTYPE','plastoholotype','HOLOTYPE',0,'');
INSERT INTO type_status VALUES('ALLOTYPE','allotype','PARATYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('ALLONEOTYPE','alloneotype','NEOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('ALLOLECTOTYPE','allolectotype','LECTOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('PARANEOTYPE','paraneotype','NEOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('PARALECTOTYPE','paralectotype','LECTOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('ISOSYNTYPE','isosyntype','SYNTYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('ISOPARATYPE','isoparatype','PARATYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('ISONEOTYPE','isoneotype','NEOTYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('ISOLECTOTYPE','isolectotype','LECTOTYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('ISOEPITYPE','isoepitype','EPITYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('ISOTYPE','isotype','HOLOTYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('TOPOTYPE','topotype','TOPOTYPE',0,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('SYNTYPE','syntype','SYNTYPE',1,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('PATHOTYPE','pathotype','PATHOTYPE',0,'BACTERIAL');
INSERT INTO type_status VALUES('PARATYPE','paratype','PARATYPE',1,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('ORIGINAL_MATERIAL','original material','ORIGINAL_MATERIAL',1,'BOTANICAL');
INSERT INTO type_status VALUES('NEOTYPE','neotype','NEOTYPE',1,'BACTERIAL,BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('LECTOTYPE','lectotype','LECTOTYPE',1,'BACTERIAL,BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('ICONOTYPE','iconotype','ICONOTYPE',0,'BOTANICAL');
INSERT INTO type_status VALUES('HOLOTYPE','holotype','HOLOTYPE',1,'BACTERIAL,BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('HAPANTOTYPE','hapantotype','HAPANTOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('EX_TYPE','ex type','EX_TYPE',0,'BOTANICAL,ZOOLOGICAL');
INSERT INTO type_status VALUES('ERGATOTYPE','ergatotype','ERGATOTYPE',0,'ZOOLOGICAL');
INSERT INTO type_status VALUES('EPITYPE','epitype','EPITYPE',0,'BOTANICAL');
CREATE TABLE nom_rel_type (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO nom_rel_type VALUES('');
INSERT INTO nom_rel_type VALUES('SPELLING_CORRECTION');
INSERT INTO nom_rel_type VALUES('BASIONYM');
INSERT INTO nom_rel_type VALUES('BASEDON');
INSERT INTO nom_rel_type VALUES('REPLACEMENT_NAME');
INSERT INTO nom_rel_type VALUES('CONSERVED');
INSERT INTO nom_rel_type VALUES('LATER_HOMONYM');
INSERT INTO nom_rel_type VALUES('SUPERFLUOUS');
INSERT INTO nom_rel_type VALUES('HOMOTYPIC');
INSERT INTO nom_rel_type VALUES('TYPE');
CREATE TABLE nom_status (col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO nom_status VALUES('');
INSERT INTO nom_status VALUES('ESTABLISHED');
INSERT INTO nom_status VALUES('ACCEPTABLE');
INSERT INTO nom_status VALUES('UNACCEPTABLE');
INSERT INTO nom_status VALUES('CONSERVED');
INSERT INTO nom_status VALUES('REJECTED');
INSERT INTO nom_status VALUES('DOUBTFUL');
INSERT INTO nom_status VALUES('MANUSCRIPT');
INSERT INTO nom_status VALUES('CHRESONYM');
CREATE TABLE reference_type(col__id TEXT PRIMARY KEY) STRICT;
INSERT INTO reference_type VALUES('');
INSERT INTO reference_type VALUES('ARTICLE');
INSERT INTO reference_type VALUES('ARTICLE_JOURNAL');
INSERT INTO reference_type VALUES('ARTICLE_MAGAZINE');
INSERT INTO reference_type VALUES('ARTICLE_NEWSPAPER');
INSERT INTO reference_type VALUES('BILL');
INSERT INTO reference_type VALUES('BOOK');
INSERT INTO reference_type VALUES('BROADCAST');
INSERT INTO reference_type VALUES('CHAPTER');
INSERT INTO reference_type VALUES('DATASET');
INSERT INTO reference_type VALUES('ENTRY');
INSERT INTO reference_type VALUES('ENTRY_DICTIONARY');
INSERT INTO reference_type VALUES('ENTRY_ENCYCLOPEDIA');
INSERT INTO reference_type VALUES('FIGURE');
INSERT INTO reference_type VALUES('GRAPHIC');
INSERT INTO reference_type VALUES('INTERVIEW');
INSERT INTO reference_type VALUES('LEGAL_CASE');
INSERT INTO reference_type VALUES('LEGISLATION');
INSERT INTO reference_type VALUES('MANUSCRIPT');
INSERT INTO reference_type VALUES('MAP');
INSERT INTO reference_type VALUES('MOTION_PICTURE');
INSERT INTO reference_type VALUES('MUSICAL_SCORE');
INSERT INTO reference_type VALUES('PAMPHLET');
INSERT INTO reference_type VALUES('PAPER_CONFERENCE');
INSERT INTO reference_type VALUES('PATENT');
INSERT INTO reference_type VALUES('PERSONAL_COMMUNICATION');
INSERT INTO reference_type VALUES('POST');
INSERT INTO reference_type VALUES('POST_WEBLOG');
INSERT INTO reference_type VALUES('REPORT');
INSERT INTO reference_type VALUES('REVIEW');
INSERT INTO reference_type VALUES('REVIEW_BOOK');
INSERT INTO reference_type VALUES('SONG');
INSERT INTO reference_type VALUES('SPEECH');
INSERT INTO reference_type VALUES('THESIS');
INSERT INTO reference_type VALUES('TREATY');
INSERT INTO reference_type VALUES('WEBPAGE');
CREATE TABLE taxonomic_status (
  col__id TEXT PRIMARY KEY,
  col__value TEXT DEFAULT '',
  col__name TEXT DEFAULT '',
  col__bare_name INTEGER DEFAULT 0, -- bool
  col__description TEXT DEFAULT '',
  col__majorStatus TEXT DEFAULT '',
  col__synonym INTEGER DEFAULT 0, -- bool
  col__taxon INTEGER DEFAULT 0 -- bool
) STRICT;
INSERT INTO taxonomic_status VALUES('','','',0,'','',0,0);
INSERT INTO taxonomic_status VALUES('ACCEPTED','','accepted',0,'A taxonomically accepted, current name','ACCEPTED',0,1);
INSERT INTO taxonomic_status VALUES('PROVISIONALLY_ACCEPTED','','provisionally accepted',0,'Treated as accepted, but doubtful whether this is correct.','ACCEPTED',0,1);
INSERT INTO taxonomic_status VALUES('VALID','','valid',0,'A taxonomically valid (ICZN), current name','VALID',0,1);
INSERT INTO taxonomic_status VALUES('PROVISIONALLY_VALID','','provisionally valid',0,'Treated as valid (ICZN), but doubtful whether this is correct.','VALID',0,1);
INSERT INTO taxonomic_status VALUES('SYNONYM','','synonym',0,'Names which point unambiguously at one species (not specifying whether homo- or heterotypic).Synonyms, in the CoL sense, include also orthographic variants and published misspellings.','SYNONYM',1,0);
INSERT INTO taxonomic_status VALUES('AMBIGUOUS_SYNONYM','','ambiguous synonym',0,'Names which are ambiguous because they point at the current species and one or more others e.g. homonyms, pro-parte synonyms (in other words, names which appear more than in one place in the Catalogue).','SYNONYM',1,0);
INSERT INTO taxonomic_status VALUES('MISAPPLIED','','misapplied',0,'A misapplied name. Usually accompanied with an accordingTo on the synonym to indicate the source the misapplication can be found in.','SYNONYM',1,0);
INSERT INTO taxonomic_status VALUES('BARE_NAME','','bare name',1,'A name alone without any usage, neither a synonym nor a taxon.','BARE_NAME',0,0);
CREATE TABLE species_interaction_type (
  col__id TEXT PRIMARY KEY,
  col__name TEXT NOT NULL,
  col__inverse TEXT REFERENCES species_interaction_type(col__id),
  col__superTypes TEXT DEFAULT '', -- ids sep ','
  col__obo TEXT DEFAULT '',
  col__symmetrical INTEGER DEFAULT 0, -- bool
  col__description TEXT DEFAULT ''
);
INSERT INTO species_interaction_type VALUES('','','','','',0,'');
INSERT INTO species_interaction_type VALUES('MUTUALIST_OF','mutualist of','MUTUALIST_OF','SYMBIONT_OF','http://purl.obolibrary.org/obo/RO_0002442',1,'An interaction relationship between two organisms living together in more or less intimate association in a relationship in which both organisms benefit from each other (GO).');
INSERT INTO species_interaction_type VALUES('COMMENSALIST_OF','commensalist of','COMMENSALIST_OF','SYMBIONT_OF','http://purl.obolibrary.org/obo/RO_0002441',1,'An interaction relationship between two organisms living together in more or less intimate association in a relationship in which one benefits and the other is unaffected (GO).');
INSERT INTO species_interaction_type VALUES('HAS_EPIPHYTE','has epiphyte','EPIPHYTE_OF','SYMBIONT_OF','http://purl.obolibrary.org/obo/RO_0008502',0,'Inverse of epiphyte of');
INSERT INTO species_interaction_type VALUES('EPIPHYTE_OF','epiphyte of','HAS_EPIPHYTE','SYMBIONT_OF','http://purl.obolibrary.org/obo/RO_0008501',0,'An interaction relationship wherein a plant or algae is living on the outside surface of another plant.');
INSERT INTO species_interaction_type VALUES('HAS_EGGS_LAYED_ON_BY','has eggs layed on by','LAYS_EGGS_ON','HOST_OF','http://purl.obolibrary.org/obo/RO_0008508',0,'Inverse of lays eggs on');
INSERT INTO species_interaction_type VALUES('LAYS_EGGS_ON','lays eggs on','HAS_EGGS_LAYED_ON_BY','HAS_HOST','http://purl.obolibrary.org/obo/RO_0008507',0,'An interaction relationship in which organism a lays eggs on the outside surface of organism b. Organism b is neither helped nor harmed in the process of egg laying or incubation.');
INSERT INTO species_interaction_type VALUES('POLLINATED_BY','pollinated by','POLLINATES','FLOWERS_VISITED_BY','http://purl.obolibrary.org/obo/RO_0002456',0,'Inverse of pollinates');
INSERT INTO species_interaction_type VALUES('POLLINATES','pollinates','POLLINATED_BY','VISITS_FLOWERS_OF','http://purl.obolibrary.org/obo/RO_0002455',0,'This relation is intended to be used for biotic pollination - e.g. a bee pollinating a flowering plant. ');
INSERT INTO species_interaction_type VALUES('FLOWERS_VISITED_BY','flowers visited by','VISITS_FLOWERS_OF','VISITED_BY','http://purl.obolibrary.org/obo/RO_0002623',0,'Inverse of visits flowers of');
INSERT INTO species_interaction_type VALUES('VISITS_FLOWERS_OF','visits flowers of','FLOWERS_VISITED_BY','VISITS','http://purl.obolibrary.org/obo/RO_0002622',0,'');
INSERT INTO species_interaction_type VALUES('VISITED_BY','visited by','VISITS','HOST_OF','http://purl.obolibrary.org/obo/RO_0002619',0,'Inverse of visits');
INSERT INTO species_interaction_type VALUES('VISITS','visits','VISITED_BY','HAS_HOST','http://purl.obolibrary.org/obo/RO_0002618',0,'');
INSERT INTO species_interaction_type VALUES('HAS_HYPERPARASITOID','has hyperparasitoid','HYPERPARASITOID_OF','HAS_PARASITOID','http://purl.obolibrary.org/obo/RO_0002554',0,'Inverse of hyperparasitoid of');
INSERT INTO species_interaction_type VALUES('HYPERPARASITOID_OF','hyperparasitoid of','HAS_HYPERPARASITOID','PARASITOID_OF','http://purl.obolibrary.org/obo/RO_0002553',0,'X is a hyperparasite of y if x is a parasite of a parasite of the target organism y');
INSERT INTO species_interaction_type VALUES('HAS_PARASITOID','has parasitoid','PARASITOID_OF','HAS_PARASITE','http://purl.obolibrary.org/obo/RO_0002209',0,'Inverse of parasitoid of');
INSERT INTO species_interaction_type VALUES('PARASITOID_OF','parasitoid of','HAS_PARASITOID','PARASITE_OF','http://purl.obolibrary.org/obo/RO_0002208',0,'A parasite that kills or sterilizes its host');
INSERT INTO species_interaction_type VALUES('HAS_KLEPTOPARASITE','has kleptoparasite','KLEPTOPARASITE_OF','HAS_PARASITE','http://purl.obolibrary.org/obo/RO_0008503',0,'Inverse of kleptoparasite of');
INSERT INTO species_interaction_type VALUES('KLEPTOPARASITE_OF','kleptoparasite of','HAS_KLEPTOPARASITE','PARASITE_OF','http://purl.obolibrary.org/obo/RO_0008503',0,'A sub-relation of parasite of in which a parasite steals resources from another organism, usually food or nest material');
INSERT INTO species_interaction_type VALUES('HAS_HYPERPARASITE','has hyperparasite','HYPERPARASITE_OF','HAS_PARASITE','http://purl.obolibrary.org/obo/RO_0002554',0,'Inverse of hyperparasite of');
INSERT INTO species_interaction_type VALUES('HYPERPARASITE_OF','hyperparasite of','HAS_HYPERPARASITE','PARASITE_OF','http://purl.obolibrary.org/obo/RO_0002553',0,'X is a hyperparasite of y iff x is a parasite of a parasite of the target organism y');
INSERT INTO species_interaction_type VALUES('HAS_ECTOPARASITE','has ectoparasite','ECTOPARASITE_OF','HAS_PARASITE','http://purl.obolibrary.org/obo/RO_0002633',0,'Inverse of ectoparasite of');
INSERT INTO species_interaction_type VALUES('ECTOPARASITE_OF','ectoparasite of','HAS_ECTOPARASITE','PARASITE_OF','http://purl.obolibrary.org/obo/RO_0002632',0,'A sub-relation of parasite-of in which the parasite lives on or in the integumental system of the host');
INSERT INTO species_interaction_type VALUES('HAS_ENDOPARASITE','has endoparasite','ENDOPARASITE_OF','HAS_PARASITE','http://purl.obolibrary.org/obo/RO_0002635',0,'Inverse of endoparasite of');
INSERT INTO species_interaction_type VALUES('ENDOPARASITE_OF','endoparasite of','HAS_ENDOPARASITE','PARASITE_OF','http://purl.obolibrary.org/obo/RO_0002634',0,'A sub-relation of parasite-of in which the parasite lives inside the host, beneath the integumental system');
INSERT INTO species_interaction_type VALUES('HAS_VECTOR','has vector','VECTOR_OF','HAS_HOST','http://purl.obolibrary.org/obo/RO_0002460',0,'Inverse of vector of');
INSERT INTO species_interaction_type VALUES('VECTOR_OF','vector of','HAS_VECTOR','HOST_OF','http://purl.obolibrary.org/obo/RO_0002459',0,'a is a vector for b if a carries and transmits an infectious pathogen b into another living organism');
INSERT INTO species_interaction_type VALUES('HAS_PATHOGEN','has pathogen','PATHOGEN_OF','HAS_PARASITE','http://purl.obolibrary.org/obo/RO_0002557',0,'Inverse of pathogen of');
INSERT INTO species_interaction_type VALUES('PATHOGEN_OF','pathogen of','HAS_PATHOGEN','PARASITE_OF','http://purl.obolibrary.org/obo/RO_0002556',0,'');
INSERT INTO species_interaction_type VALUES('HAS_PARASITE','has parasite','PARASITE_OF','EATEN_BY,HOST_OF','http://purl.obolibrary.org/obo/RO_0002445',0,'Inverse of parasite of');
INSERT INTO species_interaction_type VALUES('PARASITE_OF','parasite of','HAS_PARASITE','EATS,HAS_HOST','http://purl.obolibrary.org/obo/RO_0002444',0,'');
INSERT INTO species_interaction_type VALUES('HAS_HOST','has host','HOST_OF','SYMBIONT_OF','http://purl.obolibrary.org/obo/RO_0002454',0,'Inverse of host of');
INSERT INTO species_interaction_type VALUES('HOST_OF','host of','HAS_HOST','SYMBIONT_OF','http://purl.obolibrary.org/obo/RO_0002453',0,'The term host is usually used for the larger (macro) of the two members of a symbiosis');
INSERT INTO species_interaction_type VALUES('PREYED_UPON_BY','preyed upon by','PREYS_UPON','EATEN_BY,KILLED_BY','http://purl.obolibrary.org/obo/RO_0002458',0,'Inverse of preys upon');
INSERT INTO species_interaction_type VALUES('PREYS_UPON','preys upon','PREYED_UPON_BY','EATS,KILLS','http://purl.obolibrary.org/obo/RO_0002439',0,'An interaction relationship involving a predation process, where the subject kills the object in order to eat it or to feed to siblings, offspring or group members');
INSERT INTO species_interaction_type VALUES('KILLED_BY','killed by','KILLS','INTERACTS_WITH','http://purl.obolibrary.org/obo/RO_0002627',0,'Inverse of kills');
INSERT INTO species_interaction_type VALUES('KILLS','kills','KILLED_BY','INTERACTS_WITH','http://purl.obolibrary.org/obo/RO_0002626',0,'');
INSERT INTO species_interaction_type VALUES('EATEN_BY','eaten by','EATS','INTERACTS_WITH','http://purl.obolibrary.org/obo/RO_0002471',0,'Inverse of eats');
INSERT INTO species_interaction_type VALUES('EATS','eats','EATEN_BY','INTERACTS_WITH','http://purl.obolibrary.org/obo/RO_0002470',0,'Herbivores, fungivores, predators or other forms of organims eating or feeding on the related taxon.');
INSERT INTO species_interaction_type VALUES('SYMBIONT_OF','symbiont of','SYMBIONT_OF','INTERACTS_WITH','http://purl.obolibrary.org/obo/RO_0002440',1,'A symbiotic relationship, a more or less intimate association, with another organism. The various forms of symbiosis include parasitism, in which the association is disadvantageous or destructive to one of the organisms; mutualism, in which the association is advantageous, or often necessary to one or both and not harmful to either; and commensalism, in which one member of the association benefits while the other is not affected. However, mutualism, parasitism, and commensalism are often not discrete categories of interactions and should rather be perceived as a continuum of interaction ranging from parasitism to mutualism. In fact, the direction of a symbiotic interaction can change during the lifetime of the symbionts due to developmental changes as well as changes in the biotic/abiotic environment in which the interaction occurs. ');
INSERT INTO species_interaction_type VALUES('ADJACENT_TO','adjacent to','ADJACENT_TO','CO_OCCURS_WITH','http://purl.obolibrary.org/obo/RO_0002220',1,'X adjacent to y if and only if x and y share a boundary.');
INSERT INTO species_interaction_type VALUES('INTERACTS_WITH','interacts with','INTERACTS_WITH','CO_OCCURS_WITH','http://purl.obolibrary.org/obo/RO_0002437',1,'An interaction relationship in which at least one of the partners is an organism and the other is either an organism or an abiotic entity with which the organism interacts.');
INSERT INTO species_interaction_type VALUES('CO_OCCURS_WITH','co occurs with','CO_OCCURS_WITH','RELATED_TO','http://purl.obolibrary.org/obo/RO_0008506',1,'An interaction relationship describing organisms that often occur together at the same time and space or in the same environment.');
INSERT INTO species_interaction_type VALUES('RELATED_TO','related to','RELATED_TO','','http://purl.obolibrary.org/obo/RO_0002321',1,'Ecologically related to');
CREATE TABLE taxon_concept_rel_type (
  col__id TEXT PRIMARY KEY,
  col__name TEXT DEFAULT '',
  col__rcc5 TEXT DEFAULT '',
  col__description TEXT
) STRICT;
INSERT INTO taxon_concept_rel_type VALUES('','','','');
INSERT INTO taxon_concept_rel_type VALUES('EQUALS','equals','equal (EQ)','The circumscription of this taxon is (essentially) identical to the related taxon.');
INSERT INTO taxon_concept_rel_type VALUES('INCLUDES','includes','proper part inverse (PPi)','The related taxon concept is a subset of this taxon concept.');
INSERT INTO taxon_concept_rel_type VALUES('INCLUDED_IN','included in','proper part (PP)','This taxon concept is a subset of the related taxon concept.');
INSERT INTO taxon_concept_rel_type VALUES('OVERLAPS','overlaps','partially overlapping (PO)','Both taxon concepts share some members/children in common, and each contain some members not shared with the other.');
INSERT INTO taxon_concept_rel_type VALUES('EXCLUDES','excludes','disjoint (DR)','The related taxon concept is not a subset of this concept.');
CREATE TABLE gazetteer(
  col__id TEXT PRIMARY KEY,
  col__name TEXT,
  col__title TEXT,
  col__link TEXT,
  col__areaLinkTemplate TEXT,
  col__description TEXT
) STRICT;
INSERT INTO gazetteer VALUES('','','','','','');
INSERT INTO gazetteer VALUES('TDWG','tdwg','World Geographical Scheme for Recording Plant Distributions','http://www.tdwg.org/standards/109','','World Geographical Scheme for Recording Plant Distributions published by TDWG at level 1, 2, 3 or 4.  Level 1 = Continents, Level 2 = Regions, Level 3 = Botanical countries, Level 4 = Basic recording units.');
INSERT INTO gazetteer VALUES('ISO','iso','ISO 3166 Country Codes','https://en.wikipedia.org/wiki/ISO_3166','https://www.iso.org/obp/ui/#iso:code:3166:','ISO 3166 codes for the representation of names of countries and their subdivisions. Codes for current countries (ISO 3166-1), country subdivisions (ISO 3166-2) and formerly used names of countries (ISO 3166-3). Country codes can be given either as alpha-2, alpha-3 or numeric codes.');
INSERT INTO gazetteer VALUES('FAO','fao','FAO Major Fishing Areas','http://www.fao.org/fishery/cwp/handbook/H/en','https://www.fao.org/fishery/en/area/','FAO Major Fishing Areas');
INSERT INTO gazetteer VALUES('LONGHURST','longhurst','Longhurst Biogeographical Provinces','http://www.marineregions.org/sources.php#longhurst','','Longhurst Biogeographical Provinces, a partition of the world oceans into provinces as defined by Longhurst, A.R. (2006). Ecological Geography of the Sea. 2nd Edition.');
INSERT INTO gazetteer VALUES('TEOW','teow','Terrestrial Ecoregions of the World','https://www.worldwildlife.org/publications/terrestrial-ecoregions-of-the-world','','Terrestrial Ecoregions of the World is a biogeographic regionalization of the Earth''s terrestrial biodiversity. See Olson et al. 2001. Terrestrial ecoregions of the world: a new map of life on Earth. Bioscience 51(11):933-938.');
INSERT INTO gazetteer VALUES('IHO','iho','International Hydrographic Organization See Areas','','','Sea areas published by the International Hydrographic Organization as boundaries of the major oceans and seas of the world. See Limits of Oceans & Seas, Special Publication No. 23 published by the International Hydrographic Organization in 1953.');
INSERT INTO gazetteer VALUES('MRGID','mrgid','Marine Regions Geographic Identifier','https://www.marineregions.org/gazetteer.php','http://marineregions.org/mrgid/','Standard, relational list of geographic names developed by VLIZ covering mainly marine names such as seas, sandbanks, ridges, bays or even standard sampling stations used in marine research.The geographic cover is global; however the gazetteer is focused on the Belgian Continental Shelf, the Scheldt Estuary and the Southern Bight of the North Sea.');
INSERT INTO gazetteer VALUES('TEXT','text','Free Text','','','Free text not following any standard');
CREATE TABLE rank(
  col__id TEXT PRIMARY KEY,
  col__name TEXT DEFAULT '',
  col__plural TEXT DEFAULT '',
  col__marker TEXT DEFAULT '',
  col__major_rank_id TEXT REFERENCES rank(col__id),
  col__ambiguous_marker INTEGER DEFAULT 0, -- bool
  col__family_group INTEGER DEFAULT 0, -- bool
  col__genus_group INTEGER DEFAULT 0, -- bool
  col__infraspecific INTEGER DEFAULT 0, -- bool
  col__legacy INTEGER DEFAULT 0, -- bool
  col__linnean INTEGER DEFAULT 0, -- bool
  col__suprageneric INTEGER DEFAULT 0, -- bool
  col__supraspecific INTEGER DEFAULT 0, -- bool
  col__uncomparable INTEGER DEFAULT 0 -- bool
) STRICT;
INSERT INTO rank VALUES('','','','','',0,0,0,0,0,0,0,0,0);
INSERT INTO rank VALUES('SUPERDOMAIN','superdomain','superdomains','superdom.','DOMAIN',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('DOMAIN','domain','domains','dom.','DOMAIN',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBDOMAIN','subdomain','subdomains','subdom.','DOMAIN',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRADOMAIN','infradomain','infradomains','infradom.','DOMAIN',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('EMPIRE','empire','empires','imp.','EMPIRE',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('REALM','realm','realms','realm','REALM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBREALM','subrealm','subrealms','subrealm','REALM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERKINGDOM','superkingdom','superkingdoms','superreg.','KINGDOM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('KINGDOM','kingdom','kingdoms','regn.','KINGDOM',0,0,0,0,0,1,1,1,0);
INSERT INTO rank VALUES('SUBKINGDOM','subkingdom','subkingdoms','subreg.','KINGDOM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRAKINGDOM','infrakingdom','infrakingdoms','infrareg.','KINGDOM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERPHYLUM','superphylum','superphyla','superphyl.','PHYLUM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('PHYLUM','phylum','phyla','phyl.','PHYLUM',0,0,0,0,0,1,1,1,0);
INSERT INTO rank VALUES('SUBPHYLUM','subphylum','subphyla','subphyl.','PHYLUM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRAPHYLUM','infraphylum','infraphyla','infraphyl.','PHYLUM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('PARVPHYLUM','parvphylum','parvphyla','parvphyl.','PHYLUM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MICROPHYLUM','microphylum','microphyla','microphyl.','PHYLUM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('NANOPHYLUM','nanophylum','nanophyla','nanophyl.','PHYLUM',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('CLAUDIUS','claudius','claudius','claud.','CLAUDIUS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('GIGACLASS','gigaclass','gigaclasses','gigacl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MEGACLASS','megaclass','megaclasses','megacl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERCLASS','superclass','superclasses','supercl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('CLASS','class','classes','cl.','CLASS',0,0,0,0,0,1,1,1,0);
INSERT INTO rank VALUES('SUBCLASS','subclass','subclasses','subcl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRACLASS','infraclass','infraclasses','infracl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBTERCLASS','subterclass','subterclasses','subtercl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('PARVCLASS','parvclass','parvclasses','parvcl.','CLASS',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERDIVISION','superdivision','superdivisions','superdiv.','DIVISION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('DIVISION','division','divisions','div.','DIVISION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBDIVISION','subdivision','subdivisions','subdiv.','DIVISION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRADIVISION','infradivision','infradivisions','infradiv.','DIVISION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERLEGION','superlegion','superlegions','superleg.','LEGION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('LEGION','legion','legions','leg.','LEGION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBLEGION','sublegion','sublegions','subleg.','LEGION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRALEGION','infralegion','infralegions','infraleg.','LEGION',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MEGACOHORT','megacohort','megacohorts','megacohort','COHORT',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERCOHORT','supercohort','supercohorts','supercohort','COHORT',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('COHORT','cohort','cohorts','cohort','COHORT',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBCOHORT','subcohort','subcohorts','subcohort','COHORT',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRACOHORT','infracohort','infracohorts','infracohort','COHORT',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('GIGAORDER','gigaorder','gigaorders','gigaord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MAGNORDER','magnorder','magnorders','magnord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('GRANDORDER','grandorder','grandorders','grandord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MIRORDER','mirorder','mirorders','mirord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERORDER','superorder','superorders','superord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('ORDER','order','orders','ord.','ORDER',0,0,0,0,0,1,1,1,0);
INSERT INTO rank VALUES('NANORDER','nanorder','nanorders','nanord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('HYPOORDER','hypoorder','hypoorders','hypoord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MINORDER','minorder','minorders','minord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBORDER','suborder','suborders','subord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRAORDER','infraorder','infraorders','infraord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('PARVORDER','parvorder','parvorders','parvord.','ORDER',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERSECTION_ZOOLOGY','supersection zoology','supersection_zoologys','supersect.','SECTION_ZOOLOGY',1,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SECTION_ZOOLOGY','section zoology','section_zoologys','sect.','SECTION_ZOOLOGY',1,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBSECTION_ZOOLOGY','subsection zoology','subsection_zoologys','subsect.','SECTION_ZOOLOGY',1,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('FALANX','falanx','falanges','falanx','FALANX',0,0,0,0,1,0,1,1,0);
INSERT INTO rank VALUES('GIGAFAMILY','gigafamily','gigafamilies','gigafam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('MEGAFAMILY','megafamily','megafamilies','megafam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('GRANDFAMILY','grandfamily','grandfamilies','grandfam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERFAMILY','superfamily','superfamilies','superfam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('EPIFAMILY','epifamily','epifamilies','epifam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('FAMILY','family','families','fam.','FAMILY',0,0,0,0,0,1,1,1,0);
INSERT INTO rank VALUES('SUBFAMILY','subfamily','subfamilies','subfam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRAFAMILY','infrafamily','infrafamilies','infrafam.','FAMILY',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPERTRIBE','supertribe','supertribes','supertrib.','TRIBE',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('TRIBE','tribe','tribes','trib.','TRIBE',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUBTRIBE','subtribe','subtribes','subtrib.','TRIBE',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('INFRATRIBE','infratribe','infratribes','infratrib.','TRIBE',0,0,0,0,0,0,1,1,0);
INSERT INTO rank VALUES('SUPRAGENERIC_NAME','suprageneric name','suprageneric_names','supragen.','SUPRAGENERIC_NAME',0,0,0,0,0,0,1,1,1);
INSERT INTO rank VALUES('SUPERGENUS','supergenus','supergenera','supergen.','GENUS',0,0,1,0,0,0,1,1,0);
INSERT INTO rank VALUES('GENUS','genus','genera','gen.','GENUS',0,0,1,0,0,1,0,1,0);
INSERT INTO rank VALUES('SUBGENUS','subgenus','subgenera','subgen.','GENUS',0,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('INFRAGENUS','infragenus','infragenera','infrag.','GENUS',0,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('SUPERSECTION_BOTANY','supersection botany','supersection_botanys','supersect.','SECTION_BOTANY',1,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('SECTION_BOTANY','section botany','section_botanys','sect.','SECTION_BOTANY',1,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('SUBSECTION_BOTANY','subsection botany','subsection_botanys','subsect.','SECTION_BOTANY',1,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('SUPERSERIES','superseries','superseries','superser.','SERIES',0,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('SERIES','series','series','ser.','SERIES',0,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('SUBSERIES','subseries','subseries','subser.','SERIES',0,0,1,0,0,0,0,1,0);
INSERT INTO rank VALUES('INFRAGENERIC_NAME','infrageneric name','infrageneric_names','infragen.','GENUS',0,0,1,0,0,0,0,1,1);
INSERT INTO rank VALUES('SPECIES_AGGREGATE','species aggregate','species_aggregates','agg.','SPECIES',0,0,0,0,0,0,0,0,0);
INSERT INTO rank VALUES('SPECIES','species','species','sp.','SPECIES',0,0,0,0,0,1,0,0,0);
INSERT INTO rank VALUES('INFRASPECIFIC_NAME','infraspecific name','infraspecific_names','infrasp.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,1);
INSERT INTO rank VALUES('GREX','grex','grexs','gx','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('KLEPTON','klepton','kleptons','klepton','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('SUBSPECIES','subspecies','subspecies','subsp.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('CULTIVAR_GROUP','cultivar group','','','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('CONVARIETY','convariety','convarieties','convar.','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('INFRASUBSPECIFIC_NAME','infrasubspecific name','infrasubspecific_names','infrasubsp.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,1);
INSERT INTO rank VALUES('PROLES','proles','proles','prol.','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('NATIO','natio','natios','natio','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('ABERRATION','aberration','aberrations','ab.','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('MORPH','morph','morphs','morph','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('SUPERVARIETY','supervariety','supervarieties','supervar.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('VARIETY','variety','varieties','var.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('SUBVARIETY','subvariety','subvarieties','subvar.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('SUPERFORM','superform','superforms','superf.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('FORM','form','forms','f.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('SUBFORM','subform','subforms','subf.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('PATHOVAR','pathovar','pathovars','pv.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('BIOVAR','biovar','biovars','biovar','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('CHEMOVAR','chemovar','chemovars','chemovar','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('MORPHOVAR','morphovar','morphovars','morphovar','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('PHAGOVAR','phagovar','phagovars','phagovar','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('SEROVAR','serovar','serovars','serovar','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('CHEMOFORM','chemoform','chemoforms','chemoform','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('FORMA_SPECIALIS','forma specialis','forma_specialiss','f.sp.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('LUSUS','lusus','lusi','lusus','INFRASPECIFIC_NAME',0,0,0,1,1,0,0,0,0);
INSERT INTO rank VALUES('CULTIVAR','cultivar','cultivars','cv.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('MUTATIO','mutatio','mutatios','mut.','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('STRAIN','strain','strains','strain','INFRASPECIFIC_NAME',0,0,0,1,0,0,0,0,0);
INSERT INTO rank VALUES('OTHER','other','','','OTHER',0,0,0,0,0,0,0,0,1);
INSERT INTO rank VALUES('UNRANKED','unranked','','','UNRANKED',0,0,0,0,0,0,0,0,1);
CREATE TABLE geo_time (
  col__id TEXT PRIMARY KEY,
  col__parent_id TEXT REFERENCES geo_time(col__id),
  col__name TEXT DEFAULT '',
  col__type TEXT DEFAULT '',
  col__start REAL DEFAULT 0,
  col__end REAL) STRICT;
INSERT INTO geo_time VALUES('','','','',0.0,0.0);
INSERT INTO geo_time VALUES('HADEAN','PRECAMBRIAN','Hadean','eon',4567.0,4000.0);
INSERT INTO geo_time VALUES('PRECAMBRIAN','','Precambrian','supereon',4567.0,541.0);
INSERT INTO geo_time VALUES('ARCHEAN','PRECAMBRIAN','Archean','eon',4000.0,2500.0);
INSERT INTO geo_time VALUES('EOARCHEAN','ARCHEAN','Eoarchean','era',4000.0,3600.0);
INSERT INTO geo_time VALUES('PALEOARCHEAN','ARCHEAN','Paleoarchean','era',3600.0,3200.0);
INSERT INTO geo_time VALUES('MESOARCHEAN','ARCHEAN','Mesoarchean','era',3200.0,2800.0);
INSERT INTO geo_time VALUES('NEOARCHEAN','ARCHEAN','Neoarchean','era',2800.0,2500.0);
INSERT INTO geo_time VALUES('PROTEROZOIC','PRECAMBRIAN','Proterozoic','eon',2500.0,541.0);
INSERT INTO geo_time VALUES('PALEOPROTEROZOIC','PROTEROZOIC','Paleoproterozoic','era',2500.0,1600.0);
INSERT INTO geo_time VALUES('SIDERIAN','PALEOPROTEROZOIC','Siderian','period',2500.0,2300.0);
INSERT INTO geo_time VALUES('RHYACIAN','PALEOPROTEROZOIC','Rhyacian','period',2300.0,2050.0);
INSERT INTO geo_time VALUES('OROSIRIAN','PALEOPROTEROZOIC','Orosirian','period',2050.0,1800.0);
INSERT INTO geo_time VALUES('STATHERIAN','PALEOPROTEROZOIC','Statherian','period',1800.0,1600.0);
INSERT INTO geo_time VALUES('MESOPROTEROZOIC','PROTEROZOIC','Mesoproterozoic','era',1600.0,1000.0);
INSERT INTO geo_time VALUES('CALYMMIAN','MESOPROTEROZOIC','Calymmian','period',1600.0,1400.0);
INSERT INTO geo_time VALUES('ECTASIAN','MESOPROTEROZOIC','Ectasian','period',1400.0,1200.0);
INSERT INTO geo_time VALUES('STENIAN','MESOPROTEROZOIC','Stenian','period',1200.0,1000.0);
INSERT INTO geo_time VALUES('TONIAN','NEOPROTEROZOIC','Tonian','period',1000.0,720.0);
INSERT INTO geo_time VALUES('NEOPROTEROZOIC','PROTEROZOIC','Neoproterozoic','era',1000.0,541.0);
INSERT INTO geo_time VALUES('CRYOGENIAN','NEOPROTEROZOIC','Cryogenian','period',720.0,635.0);
INSERT INTO geo_time VALUES('EDIACARAN','NEOPROTEROZOIC','Ediacaran','period',635.0,541.0);
INSERT INTO geo_time VALUES('CAMBRIAN','PALEOZOIC','Cambrian','period',541.0,485.4);
INSERT INTO geo_time VALUES('FORTUNIAN','TERRENEUVIAN','Fortunian','age',541.0,529.0);
INSERT INTO geo_time VALUES('PALEOZOIC','PHANEROZOIC','Paleozoic','era',541.0,251.902);
INSERT INTO geo_time VALUES('PHANEROZOIC','','Phanerozoic','eon',541.0,0.0);
INSERT INTO geo_time VALUES('TERRENEUVIAN','CAMBRIAN','Terreneuvian','epoch',541.0,521.0);
INSERT INTO geo_time VALUES('CAMBRIANSTAGE2','TERRENEUVIAN','CambrianStage2','age',529.0,521.0);
INSERT INTO geo_time VALUES('CAMBRIANSERIES2','CAMBRIAN','CambrianSeries2','epoch',521.0,509.0);
INSERT INTO geo_time VALUES('CAMBRIANSTAGE3','CAMBRIANSERIES2','CambrianStage3','age',521.0,514.0);
INSERT INTO geo_time VALUES('CAMBRIANSTAGE4','CAMBRIANSERIES2','CambrianStage4','age',514.0,509.0);
INSERT INTO geo_time VALUES('WULIUAN','MIAOLINGIAN','Wuliuan','age',509.0,504.5);
INSERT INTO geo_time VALUES('MIAOLINGIAN','CAMBRIAN','Miaolingian','epoch',509.0,497.0);
INSERT INTO geo_time VALUES('DRUMIAN','MIAOLINGIAN','Drumian','age',504.5,500.5);
INSERT INTO geo_time VALUES('GUZHANGIAN','MIAOLINGIAN','Guzhangian','age',500.5,497.0);
INSERT INTO geo_time VALUES('FURONGIAN','CAMBRIAN','Furongian','epoch',497.0,485.4);
INSERT INTO geo_time VALUES('PAIBIAN','FURONGIAN','Paibian','age',497.0,494.0);
INSERT INTO geo_time VALUES('JIANGSHANIAN','FURONGIAN','Jiangshanian','age',494.0,489.5);
INSERT INTO geo_time VALUES('CAMBRIANSTAGE10','FURONGIAN','CambrianStage10','age',489.5,485.4);
INSERT INTO geo_time VALUES('TREMADOCIAN','LOWER_ORDOVICIAN','Tremadocian','age',485.4,477.7);
INSERT INTO geo_time VALUES('LOWER_ORDOVICIAN','ORDOVICIAN','LowerOrdovician','epoch',485.4,470.0);
INSERT INTO geo_time VALUES('ORDOVICIAN','PALEOZOIC','Ordovician','period',485.4,443.8);
INSERT INTO geo_time VALUES('FLOIAN','LOWER_ORDOVICIAN','Floian','age',477.7,470.0);
INSERT INTO geo_time VALUES('DAPINGIAN','MIDDLE_ORDOVICIAN','Dapingian','age',470.0,467.3);
INSERT INTO geo_time VALUES('MIDDLE_ORDOVICIAN','ORDOVICIAN','MiddleOrdovician','epoch',470.0,458.4);
INSERT INTO geo_time VALUES('DARRIWILIAN','MIDDLE_ORDOVICIAN','Darriwilian','age',467.3,458.4);
INSERT INTO geo_time VALUES('SANDBIAN','UPPER_ORDOVICIAN','Sandbian','age',458.4,453.0);
INSERT INTO geo_time VALUES('UPPER_ORDOVICIAN','ORDOVICIAN','UpperOrdovician','epoch',458.4,443.8);
INSERT INTO geo_time VALUES('KATIAN','UPPER_ORDOVICIAN','Katian','age',453.0,445.2);
INSERT INTO geo_time VALUES('HIRNANTIAN','UPPER_ORDOVICIAN','Hirnantian','age',445.2,443.8);
INSERT INTO geo_time VALUES('LLANDOVERY','SILURIAN','Llandovery','epoch',443.8,433.4);
INSERT INTO geo_time VALUES('RHUDDANIAN','LLANDOVERY','Rhuddanian','age',443.8,440.8);
INSERT INTO geo_time VALUES('SILURIAN','PALEOZOIC','Silurian','period',443.8,419.2);
INSERT INTO geo_time VALUES('AERONIAN','LLANDOVERY','Aeronian','age',440.8,438.5);
INSERT INTO geo_time VALUES('TELYCHIAN','LLANDOVERY','Telychian','age',438.5,433.4);
INSERT INTO geo_time VALUES('SHEINWOODIAN','WENLOCK','Sheinwoodian','age',433.4,430.5);
INSERT INTO geo_time VALUES('WENLOCK','SILURIAN','Wenlock','epoch',433.4,427.4);
INSERT INTO geo_time VALUES('HOMERIAN','WENLOCK','Homerian','age',430.5,427.4);
INSERT INTO geo_time VALUES('LUDLOW','SILURIAN','Ludlow','epoch',427.4,423.0);
INSERT INTO geo_time VALUES('GORSTIAN','LUDLOW','Gorstian','age',427.4,425.6);
INSERT INTO geo_time VALUES('LUDFORDIAN','LUDLOW','Ludfordian','age',425.6,423.0);
INSERT INTO geo_time VALUES('PRIDOLI','SILURIAN','Pridoli','age',423.0,419.2);
INSERT INTO geo_time VALUES('DEVONIAN','PALEOZOIC','Devonian','period',419.2,358.9);
INSERT INTO geo_time VALUES('LOWER_DEVONIAN','DEVONIAN','LowerDevonian','epoch',419.2,393.3);
INSERT INTO geo_time VALUES('LOCHKOVIAN','LOWER_DEVONIAN','Lochkovian','age',419.2,410.8);
INSERT INTO geo_time VALUES('PRAGIAN','LOWER_DEVONIAN','Pragian','age',410.8,407.6);
INSERT INTO geo_time VALUES('EMSIAN','LOWER_DEVONIAN','Emsian','age',407.6,393.3);
INSERT INTO geo_time VALUES('EIFELIAN','MIDDLE_DEVONIAN','Eifelian','age',393.3,387.7);
INSERT INTO geo_time VALUES('MIDDLE_DEVONIAN','DEVONIAN','MiddleDevonian','epoch',393.3,382.7);
INSERT INTO geo_time VALUES('GIVETIAN','MIDDLE_DEVONIAN','Givetian','age',387.7,382.7);
INSERT INTO geo_time VALUES('UPPER_DEVONIAN','DEVONIAN','UpperDevonian','epoch',382.7,358.9);
INSERT INTO geo_time VALUES('FRASNIAN','UPPER_DEVONIAN','Frasnian','age',382.7,372.2);
INSERT INTO geo_time VALUES('FAMENNIAN','UPPER_DEVONIAN','Famennian','age',372.2,358.9);
INSERT INTO geo_time VALUES('LOWER_MISSISSIPPIAN','MISSISSIPPIAN','LowerMississippian','epoch',358.9,346.7);
INSERT INTO geo_time VALUES('TOURNAISIAN','LOWER_MISSISSIPPIAN','Tournaisian','age',358.9,346.7);
INSERT INTO geo_time VALUES('MISSISSIPPIAN','CARBONIFEROUS','Mississippian','subperiod',358.9,323.2);
INSERT INTO geo_time VALUES('CARBONIFEROUS','PALEOZOIC','Carboniferous','period',358.9,298.9);
INSERT INTO geo_time VALUES('MIDDLE_MISSISSIPPIAN','MISSISSIPPIAN','MiddleMississippian','epoch',346.7,330.9);
INSERT INTO geo_time VALUES('VISEAN','MIDDLE_MISSISSIPPIAN','Visean','age',346.7,330.9);
INSERT INTO geo_time VALUES('SERPUKHOVIAN','UPPER_MISSISSIPPIAN','Serpukhovian','age',330.9,323.2);
INSERT INTO geo_time VALUES('UPPER_MISSISSIPPIAN','MISSISSIPPIAN','UpperMississippian','epoch',330.9,298.9);
INSERT INTO geo_time VALUES('BASHKIRIAN','LOWER_PENNSYLVANIAN','Bashkirian','age',323.2,315.2);
INSERT INTO geo_time VALUES('PENNSYLVANIAN','CARBONIFEROUS','Pennsylvanian','subperiod',323.2,298.9);
INSERT INTO geo_time VALUES('LOWER_PENNSYLVANIAN','PENNSYLVANIAN','LowerPennsylvanian','epoch',323.2,315.2);
INSERT INTO geo_time VALUES('MIDDLE_PENNSYLVANIAN','PENNSYLVANIAN','MiddlePennsylvanian','epoch',315.2,307.0);
INSERT INTO geo_time VALUES('MOSCOVIAN','MIDDLE_PENNSYLVANIAN','Moscovian','age',315.2,307.0);
INSERT INTO geo_time VALUES('KASIMOVIAN','UPPER_PENNSYLVANIAN','Kasimovian','age',307.0,303.7);
INSERT INTO geo_time VALUES('UPPER_PENNSYLVANIAN','PENNSYLVANIAN','UpperPennsylvanian','epoch',307.0,298.9);
INSERT INTO geo_time VALUES('GZHELIAN','UPPER_PENNSYLVANIAN','Gzhelian','age',303.7,298.9);
INSERT INTO geo_time VALUES('CISURALIAN','PERMIAN','Cisuralian','epoch',298.9,272.95);
INSERT INTO geo_time VALUES('ASSELIAN','CISURALIAN','Asselian','age',298.9,295.0);
INSERT INTO geo_time VALUES('PERMIAN','PALEOZOIC','Permian','period',298.9,251.902);
INSERT INTO geo_time VALUES('SAKMARIAN','CISURALIAN','Sakmarian','age',295.0,290.1);
INSERT INTO geo_time VALUES('ARTINSKIAN','CISURALIAN','Artinskian','age',290.1,283.5);
INSERT INTO geo_time VALUES('KUNGURIAN','CISURALIAN','Kungurian','age',283.5,272.95);
INSERT INTO geo_time VALUES('ROADIAN','GUADALUPIAN','Roadian','age',272.95,268.8);
INSERT INTO geo_time VALUES('GUADALUPIAN','PERMIAN','Guadalupian','epoch',272.95,259.1);
INSERT INTO geo_time VALUES('WORDIAN','GUADALUPIAN','Wordian','age',268.8,265.1);
INSERT INTO geo_time VALUES('CAPITANIAN','GUADALUPIAN','Capitanian','age',265.1,259.1);
INSERT INTO geo_time VALUES('LOPINGIAN','PERMIAN','Lopingian','epoch',259.1,251.902);
INSERT INTO geo_time VALUES('WUCHIAPINGIAN','LOPINGIAN','Wuchiapingian','age',259.1,254.14);
INSERT INTO geo_time VALUES('CHANGHSINGIAN','LOPINGIAN','Changhsingian','age',254.14,251.902);
INSERT INTO geo_time VALUES('INDUAN','LOWER_TRIASSIC','Induan','age',251.902,251.2);
INSERT INTO geo_time VALUES('LOWER_TRIASSIC','TRIASSIC','LowerTriassic','epoch',251.902,247.2);
INSERT INTO geo_time VALUES('MESOZOIC','PHANEROZOIC','Mesozoic','era',251.902,66.0);
INSERT INTO geo_time VALUES('TRIASSIC','MESOZOIC','Triassic','period',251.902,201.3);
INSERT INTO geo_time VALUES('OLENEKIAN','LOWER_TRIASSIC','Olenekian','age',251.2,247.2);
INSERT INTO geo_time VALUES('ANISIAN','MIDDLE_TRIASSIC','Anisian','age',247.2,242.0);
INSERT INTO geo_time VALUES('MIDDLE_TRIASSIC','TRIASSIC','MiddleTriassic','epoch',247.2,237.0);
INSERT INTO geo_time VALUES('LADINIAN','MIDDLE_TRIASSIC','Ladinian','age',242.0,237.0);
INSERT INTO geo_time VALUES('CARNIAN','UPPER_TRIASSIC','Carnian','age',237.0,227.0);
INSERT INTO geo_time VALUES('UPPER_TRIASSIC','TRIASSIC','UpperTriassic','epoch',237.0,201.3);
INSERT INTO geo_time VALUES('NORIAN','UPPER_TRIASSIC','Norian','age',227.0,208.5);
INSERT INTO geo_time VALUES('RHAETIAN','UPPER_TRIASSIC','Rhaetian','age',208.5,201.3);
INSERT INTO geo_time VALUES('JURASSIC','MESOZOIC','Jurassic','period',201.3,145.0);
INSERT INTO geo_time VALUES('HETTANGIAN','LOWER_JURASSIC','Hettangian','age',201.3,199.3);
INSERT INTO geo_time VALUES('LOWER_JURASSIC','JURASSIC','LowerJurassic','epoch',201.3,174.1);
INSERT INTO geo_time VALUES('SINEMURIAN','LOWER_JURASSIC','Sinemurian','age',199.3,190.8);
INSERT INTO geo_time VALUES('PLIENSBACHIAN','LOWER_JURASSIC','Pliensbachian','age',190.8,182.7);
INSERT INTO geo_time VALUES('TOARCIAN','LOWER_JURASSIC','Toarcian','age',182.7,174.1);
INSERT INTO geo_time VALUES('MIDDLE_JURASSIC','JURASSIC','MiddleJurassic','epoch',174.1,163.5);
INSERT INTO geo_time VALUES('AALENIAN','MIDDLE_JURASSIC','Aalenian','age',174.1,170.3);
INSERT INTO geo_time VALUES('BAJOCIAN','MIDDLE_JURASSIC','Bajocian','age',170.3,168.3);
INSERT INTO geo_time VALUES('BATHONIAN','MIDDLE_JURASSIC','Bathonian','age',168.3,166.1);
INSERT INTO geo_time VALUES('CALLOVIAN','MIDDLE_JURASSIC','Callovian','age',166.1,163.5);
INSERT INTO geo_time VALUES('OXFORDIAN','UPPER_JURASSIC','Oxfordian','age',163.5,157.3);
INSERT INTO geo_time VALUES('UPPER_JURASSIC','JURASSIC','UpperJurassic','epoch',163.5,145.0);
INSERT INTO geo_time VALUES('KIMMERIDGIAN','UPPER_JURASSIC','Kimmeridgian','age',157.3,152.1);
INSERT INTO geo_time VALUES('TITHONIAN','UPPER_JURASSIC','Tithonian','age',152.1,145.0);
INSERT INTO geo_time VALUES('LOWER_CRETACEOUS','CRETACEOUS','LowerCretaceous','epoch',145.0,100.5);
INSERT INTO geo_time VALUES('CRETACEOUS','MESOZOIC','Cretaceous','period',145.0,66.0);
INSERT INTO geo_time VALUES('BERRIASIAN','LOWER_CRETACEOUS','Berriasian','age',145.0,139.8);
INSERT INTO geo_time VALUES('VALANGINIAN','LOWER_CRETACEOUS','Valanginian','age',139.8,132.9);
INSERT INTO geo_time VALUES('HAUTERIVIAN','LOWER_CRETACEOUS','Hauterivian','age',132.9,129.4);
INSERT INTO geo_time VALUES('BARREMIAN','LOWER_CRETACEOUS','Barremian','age',129.4,125.0);
INSERT INTO geo_time VALUES('APTIAN','LOWER_CRETACEOUS','Aptian','age',125.0,113.0);
INSERT INTO geo_time VALUES('ALBIAN','LOWER_CRETACEOUS','Albian','age',113.0,100.5);
INSERT INTO geo_time VALUES('CENOMANIAN','UPPER_CRETACEOUS','Cenomanian','age',100.5,93.9);
INSERT INTO geo_time VALUES('UPPER_CRETACEOUS','CRETACEOUS','UpperCretaceous','epoch',100.5,66.0);
INSERT INTO geo_time VALUES('TURONIAN','UPPER_CRETACEOUS','Turonian','age',93.9,89.8);
INSERT INTO geo_time VALUES('CONIACIAN','UPPER_CRETACEOUS','Coniacian','age',89.8,86.3);
INSERT INTO geo_time VALUES('SANTONIAN','UPPER_CRETACEOUS','Santonian','age',86.3,83.6);
INSERT INTO geo_time VALUES('CAMPANIAN','UPPER_CRETACEOUS','Campanian','age',83.6,72.1);
INSERT INTO geo_time VALUES('MAASTRICHTIAN','UPPER_CRETACEOUS','Maastrichtian','age',72.1,66.0);
INSERT INTO geo_time VALUES('PALEOCENE','PALEOGENE','Paleocene','epoch',66.0,56.0);
INSERT INTO geo_time VALUES('PALEOGENE','CENOZOIC','Paleogene','period',66.0,23.03);
INSERT INTO geo_time VALUES('CENOZOIC','PHANEROZOIC','Cenozoic','era',66.0,0.0);
INSERT INTO geo_time VALUES('DANIAN','PALEOCENE','Danian','age',66.0,61.6);
INSERT INTO geo_time VALUES('SELANDIAN','PALEOCENE','Selandian','age',61.6,59.2);
INSERT INTO geo_time VALUES('THANETIAN','PALEOCENE','Thanetian','age',59.2,56.0);
INSERT INTO geo_time VALUES('EOCENE','PALEOGENE','Eocene','epoch',56.0,33.9);
INSERT INTO geo_time VALUES('YPRESIAN','EOCENE','Ypresian','age',56.0,47.8);
INSERT INTO geo_time VALUES('LUTETIAN','EOCENE','Lutetian','age',47.8,41.2);
INSERT INTO geo_time VALUES('BARTONIAN','EOCENE','Bartonian','age',41.2,37.8);
INSERT INTO geo_time VALUES('PRIABONIAN','EOCENE','Priabonian','age',37.8,33.9);
INSERT INTO geo_time VALUES('RUPELIAN','OLIGOCENE','Rupelian','age',33.9,28.1);
INSERT INTO geo_time VALUES('OLIGOCENE','PALEOGENE','Oligocene','epoch',33.9,23.03);
INSERT INTO geo_time VALUES('CHATTIAN','OLIGOCENE','Chattian','age',27.82,23.03);
INSERT INTO geo_time VALUES('AQUITANIAN','MIOCENE','Aquitanian','age',23.03,20.44);
INSERT INTO geo_time VALUES('NEOGENE','CENOZOIC','Neogene','period',23.03,2.58);
INSERT INTO geo_time VALUES('MIOCENE','NEOGENE','Miocene','epoch',23.03,5.333);
INSERT INTO geo_time VALUES('BURDIGALIAN','MIOCENE','Burdigalian','age',20.44,15.97);
INSERT INTO geo_time VALUES('LANGHIAN','MIOCENE','Langhian','age',15.97,13.82);
INSERT INTO geo_time VALUES('SERRAVALLIAN','MIOCENE','Serravallian','age',13.82,11.63);
INSERT INTO geo_time VALUES('TORTONIAN','MIOCENE','Tortonian','age',11.63,7.246);
INSERT INTO geo_time VALUES('MESSINIAN','MIOCENE','Messinian','age',7.246,5.333);
INSERT INTO geo_time VALUES('ZANCLEAN','PLIOCENE','Zanclean','age',5.333,3.6);
INSERT INTO geo_time VALUES('PLIOCENE','NEOGENE','Pliocene','epoch',5.333,2.58);
INSERT INTO geo_time VALUES('PIACENZIAN','PLIOCENE','Piacenzian','age',3.6,2.58);
INSERT INTO geo_time VALUES('QUATERNARY','CENOZOIC','Quaternary','period',2.58,0.0);
INSERT INTO geo_time VALUES('GELASIAN','PLEISTOCENE','Gelasian','age',2.58,1.8);
INSERT INTO geo_time VALUES('PLEISTOCENE','QUATERNARY','Pleistocene','epoch',2.58,0.0117);
INSERT INTO geo_time VALUES('CALABRIAN','PLEISTOCENE','Calabrian','age',1.8,0.781);
INSERT INTO geo_time VALUES('MIDDLE_PLEISTOCENE','PLEISTOCENE','MiddlePleistocene','age',0.781,0.126);
INSERT INTO geo_time VALUES('UPPER_PLEISTOCENE','PLEISTOCENE','UpperPleistocene','age',0.126,0.0117);
INSERT INTO geo_time VALUES('HOLOCENE','QUATERNARY','Holocene','epoch',0.0117,0.0);
INSERT INTO geo_time VALUES('GREENLANDIAN','HOLOCENE','Greenlandian','age',0.0117,0.0082);
INSERT INTO geo_time VALUES('NORTHGRIPPIAN','HOLOCENE','Northgrippian','age',0.0082,0.0042);
INSERT INTO geo_time VALUES('MEGHALAYAN','HOLOCENE','Meghalayan','age',0.0042,0.0);
CREATE INDEX idx_name_scientific_name ON name (col__scientific_name);
CREATE INDEX idx_name_canonical_simple ON name (gn__canonical_simple);
CREATE INDEX idx_synonym_id ON synonym (col__id);
CREATE INDEX idx_synonym_taxon_id ON synonym (col__taxon_id);
CREATE UNIQUE INDEX idx_synonym_id_taxon_id ON synonym (col__id, col__taxon_id) WHERE col__id != '';
CREATE INDEX idx_vernacular_taxon_id ON vernacular (col__taxon_id);
CREATE INDEX idx_type_material_id ON type_material (col__id);
COMMIT;
