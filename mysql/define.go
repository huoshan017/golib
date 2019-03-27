package mysql

// mysql engine type
const (
	MYSQL_ENGINE_MYISAM = iota
	MYSQL_ENGINE_INNODB = 1
)

var mysql_engines_name_map = map[int32]string{
	MYSQL_ENGINE_MYISAM: "MYISAM",
	MYSQL_ENGINE_INNODB: "INNODB",
}

// mysql table create flag
const (
	MYSQL_TABLE_CREATE_ZEROFILL                   = 1
	MYSQL_TABLE_CREATE_UNSIGNED                   = 2
	MYSQL_TABLE_CREATE_AUTOINCREMENT              = 4
	MYSQL_TABLE_CREATE_NULL                       = 8
	MYSQL_TABLE_CREATE_NOT_NULL                   = 16
	MYSQL_TABLE_CREATE_DEFAULT                    = 32
	MYSQL_TABLE_CREATE_CURRENTTIMESTAMP           = 64
	MYSQL_TABLE_CREATE_CURRENTTIMESTAMP_ON_UPDATE = 128
)

var mysql_table_create_flag_name_map = map[int32]string{
	MYSQL_TABLE_CREATE_ZEROFILL:      "ZEROFILL",
	MYSQL_TABLE_CREATE_UNSIGNED:      "UNSIGNED",
	MYSQL_TABLE_CREATE_AUTOINCREMENT: "AUTO_INCREMENT",
	MYSQL_TABLE_CREATE_NULL:          "NULL",
	MYSQL_TABLE_CREATE_NOT_NULL:      "NOT NULL",
	MYSQL_TABLE_CREATE_DEFAULT:       "DEFAULT",
}

// mysql field type
const (
	MYSQL_FIELD_TYPE_NONE       = iota //
	MYSQL_FIELD_TYPE_TINYINT    = 1    // TINYINT
	MYSQL_FIELD_TYPE_SMALLINT   = 2    // SMALLINT
	MYSQL_FIELD_TYPE_MEDIUMINT  = 3    // MEDIUMINT
	MYSQL_FIELD_TYPE_INT        = 4    // INT
	MYSQL_FIELD_TYPE_BIGINT     = 5    // BIGINT
	MYSQL_FIELD_TYPE_FLOAT      = 6    // FLOAT
	MYSQL_FIELD_TYPE_DOUBLE     = 7    // DOUBLE
	MYSQL_FIELD_TYPE_DATE       = 8    // DATE
	MYSQL_FIELD_TYPE_DATETIME   = 9    // DATETIME
	MYSQL_FIELD_TYPE_TIMESTAMP  = 10   // TIMESTAMP
	MYSQL_FIELD_TYPE_TIME       = 11   // TIME
	MYSQL_FIELD_TYPE_YEAR       = 12   // YEAR
	MYSQL_FIELD_TYPE_CHAR       = 13   // CHAR
	MYSQL_FIELD_TYPE_BINARY     = 14   // BINARY
	MYSQL_FIELD_TYPE_VARBINARY  = 15   // VARBINARY
	MYSQL_FIELD_TYPE_VARCHAR    = 16   // VARCHAR
	MYSQL_FIELD_TYPE_TINYBLOB   = 17   // TINYBLOB
	MYSQL_FIELD_TYPE_TINYTEXT   = 18   // TINYTEXT
	MYSQL_FIELD_TYPE_BLOB       = 19   // BLOB
	MYSQL_FIELD_TYPE_TEXT       = 20   // TEXT
	MYSQL_FIELD_TYPE_MEDIUMBLOB = 21   // MEDIUMBLOB
	MYSQL_FIELD_TYPE_MEDIUMTEXT = 22   // MEDIUMTEXT
	MYSQL_FIELD_TYPE_LONGBLOB   = 23   // LONGBLOB
	MYSQL_FIELD_TYPE_LONGTEXT   = 24   // LONGTEXT
	MYSQL_FIELD_TYPE_ENUM       = 25   // ENUM
	MYSQL_FIELD_TYPE_SET        = 26   // SET
	MYSQL_FIELD_TYPE_MAX        = 100
)

var mysql_field_type_name_map = map[int32]string{
	MYSQL_FIELD_TYPE_TINYINT:    "TINYINT",
	MYSQL_FIELD_TYPE_SMALLINT:   "SMALLINT",
	MYSQL_FIELD_TYPE_MEDIUMINT:  "MEDIUMINT",
	MYSQL_FIELD_TYPE_INT:        "INT",
	MYSQL_FIELD_TYPE_BIGINT:     "BIGINT",
	MYSQL_FIELD_TYPE_FLOAT:      "FLOAT",
	MYSQL_FIELD_TYPE_DOUBLE:     "DOUBLE",
	MYSQL_FIELD_TYPE_DATE:       "",
	MYSQL_FIELD_TYPE_DATETIME:   "",
	MYSQL_FIELD_TYPE_TIMESTAMP:  "TIMESTAMP",
	MYSQL_FIELD_TYPE_TIME:       "TIME",
	MYSQL_FIELD_TYPE_YEAR:       "YEAR",
	MYSQL_FIELD_TYPE_CHAR:       "CHAR",
	MYSQL_FIELD_TYPE_VARCHAR:    "VARCHAR",
	MYSQL_FIELD_TYPE_BINARY:     "BINARY",
	MYSQL_FIELD_TYPE_VARBINARY:  "VARBINARY",
	MYSQL_FIELD_TYPE_TINYBLOB:   "TINYBLOB",
	MYSQL_FIELD_TYPE_TINYTEXT:   "TINYTEXT",
	MYSQL_FIELD_TYPE_BLOB:       "BLOB",
	MYSQL_FIELD_TYPE_TEXT:       "TEXT",
	MYSQL_FIELD_TYPE_MEDIUMBLOB: "MEDIUMBLOB",
	MYSQL_FIELD_TYPE_MEDIUMTEXT: "MEDIUMTEXT",
	MYSQL_FIELD_TYPE_LONGBLOB:   "LONGBLOB",
	MYSQL_FIELD_TYPE_LONGTEXT:   "LONGTEXT",
}

// mysql default field length
const (
	MYSQL_FIELD_DEFAULT_LENGTH            = iota
	MYSQL_FIELD_DEFAULT_LENGTH_TINYINT    = 4
	MYSQL_FIELD_DEFAULT_LENGTH_SMALLINT   = 6
	MYSQL_FIELD_DEFAULT_LENGTH_MEDIUMINT  = 8
	MYSQL_FIELD_DEFAULT_LENGTH_INT        = 11
	MYSQL_FIELD_DEFAULT_LENGTH_BIGINT     = 20
	MYSQL_FIELD_DEFAULT_LENGTH_FLOAT      = 11
	MYSQL_FIELD_DEFAULT_LENGTH_DOUBLE     = 20
	MYSQL_FIELD_DEFAULT_LENGTH_DATE       = 10
	MYSQL_FIELD_DEFAULT_LENGTH_DATETIME   = 19
	MYSQL_FIELD_DEFAULT_LENGTH_TIMESTAMP  = 6
	MYSQL_FIELD_DEFAULT_LENGTH_TIME       = 8
	MYSQL_FIELD_DEFAULT_LENGTH_YEAR       = 4
	MYSQL_FIELD_DEFAULT_LENGTH_CHAR       = 255
	MYSQL_FIELD_DEFAULT_LENGTH_VARCHAR    = 65530
	MYSQL_FIELD_DEFAULT_LENGTH_BINARY     = 8000
	MYSQL_FIELD_DEFAULT_LENGTH_VARBINARY  = 8000
	MYSQL_FIELD_DEFAULT_LENGTH_TINYBLOB   = 255
	MYSQL_FIELD_DEFAULT_LENGTH_TINYTEXT   = 255
	MYSQL_FIELD_DEFAULT_LENGTH_BLOB       = 65535
	MYSQL_FIELD_DEFAULT_LENGTH_TEXT       = 65535
	MYSQL_FIELD_DEFAULT_LENGTH_MEDIUMBLOB = 16777215
	MYSQL_FIELD_DEFAULT_LENGTH_MEDIUMTEXT = 16777215
	MYSQL_FIELD_DEFAULT_LENGTH_LONGBLOB   = 4294967295
	MYSQL_FIELD_DEFAULT_LENGTH_LONGTEXT   = 4294967295
)

var mysql_field_type_default_length_map = map[int32]int64{
	MYSQL_FIELD_TYPE_TINYINT:    MYSQL_FIELD_DEFAULT_LENGTH_TINYINT,
	MYSQL_FIELD_TYPE_SMALLINT:   MYSQL_FIELD_DEFAULT_LENGTH_SMALLINT,
	MYSQL_FIELD_TYPE_MEDIUMINT:  MYSQL_FIELD_DEFAULT_LENGTH_MEDIUMINT,
	MYSQL_FIELD_TYPE_INT:        MYSQL_FIELD_DEFAULT_LENGTH_INT,
	MYSQL_FIELD_TYPE_BIGINT:     MYSQL_FIELD_DEFAULT_LENGTH_BIGINT,
	MYSQL_FIELD_TYPE_FLOAT:      MYSQL_FIELD_DEFAULT_LENGTH_FLOAT,
	MYSQL_FIELD_TYPE_DOUBLE:     MYSQL_FIELD_DEFAULT_LENGTH_DOUBLE,
	MYSQL_FIELD_TYPE_DATE:       MYSQL_FIELD_DEFAULT_LENGTH_DATE,
	MYSQL_FIELD_TYPE_DATETIME:   MYSQL_FIELD_DEFAULT_LENGTH_DATETIME,
	MYSQL_FIELD_TYPE_TIMESTAMP:  MYSQL_FIELD_DEFAULT_LENGTH_TIMESTAMP,
	MYSQL_FIELD_TYPE_TIME:       MYSQL_FIELD_DEFAULT_LENGTH_TIME,
	MYSQL_FIELD_TYPE_YEAR:       MYSQL_FIELD_DEFAULT_LENGTH_YEAR,
	MYSQL_FIELD_TYPE_CHAR:       MYSQL_FIELD_DEFAULT_LENGTH_CHAR,
	MYSQL_FIELD_TYPE_VARCHAR:    MYSQL_FIELD_DEFAULT_LENGTH_VARCHAR,
	MYSQL_FIELD_TYPE_BINARY:     MYSQL_FIELD_DEFAULT_LENGTH_BINARY,
	MYSQL_FIELD_TYPE_VARBINARY:  MYSQL_FIELD_DEFAULT_LENGTH_VARBINARY,
	MYSQL_FIELD_TYPE_TINYBLOB:   MYSQL_FIELD_DEFAULT_LENGTH_TINYBLOB,
	MYSQL_FIELD_TYPE_TINYTEXT:   MYSQL_FIELD_DEFAULT_LENGTH_TINYTEXT,
	MYSQL_FIELD_TYPE_BLOB:       MYSQL_FIELD_DEFAULT_LENGTH_BLOB,
	MYSQL_FIELD_TYPE_TEXT:       MYSQL_FIELD_DEFAULT_LENGTH_TEXT,
	MYSQL_FIELD_TYPE_MEDIUMBLOB: MYSQL_FIELD_DEFAULT_LENGTH_MEDIUMBLOB,
	MYSQL_FIELD_TYPE_MEDIUMTEXT: MYSQL_FIELD_DEFAULT_LENGTH_MEDIUMTEXT,
	MYSQL_FIELD_TYPE_LONGBLOB:   MYSQL_FIELD_DEFAULT_LENGTH_LONGBLOB,
	MYSQL_FIELD_TYPE_LONGTEXT:   MYSQL_FIELD_DEFAULT_LENGTH_LONGTEXT,
}

// mysql index type
const (
	MYSQL_INDEX_TYPE_NONE     = iota
	MYSQL_INDEX_TYPE_NORMAL   = 1
	MYSQL_INDEX_TYPE_UNIQUE   = 2
	MYSQL_INDEX_TYPE_FULLTEXT = 3
)

var mysql_index_type_name_map = map[int32]string{
	MYSQL_INDEX_TYPE_NONE:     "",
	MYSQL_INDEX_TYPE_NORMAL:   "index",
	MYSQL_INDEX_TYPE_UNIQUE:   "unique",
	MYSQL_INDEX_TYPE_FULLTEXT: "fulltext",
}
