package metrics

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"k8s.io/klog"

	"github.com/openbsi/kubesim/pkg/metrics/config"
)

// MysqlSink implements metrics.Interface
type MysqlSink struct {
	sink     *config.SinkConfig
	connect  bool
	database *sql.DB
	node     string
}

// NewMysqlSink create a new MysqlSink
func NewMysqlSink(sink *config.SinkConfig, node string) Interface {
	return &MysqlSink{
		sink:     sink,
		connect:  false,
		database: nil,
		node:     node,
	}
}

// Initialization is used to initialize MysqlSink
func (ms *MysqlSink) Initialization() error {
	err := ms.initDB()
	if err != nil {
		return err
	}

	ms.connect = true
	return nil
}

// LogNodeMetrics is used to insert node metrics to Mysql
func (ms *MysqlSink) LogNodeMetrics(nm *NodeMetric) error {
	if !ms.connect {
		return fmt.Errorf("LogNodeMetrics, not connect to mysql server")
	}

	switch nm.MetricType {
	case "static":
		return ms.insertStaticNodeMetrics(nm)
	case "real":
		return ms.insertRealNodeMetrics(nm)
	default:
		klog.Errorf("unsupported node metrics type")
	}
	return nil
}

// LogVolcanlJobMetrics is used to insert volcano job metrics to Mysql
func (ms *MysqlSink) LogVolcanlJobMetrics() error {
	if !ms.connect {
		return fmt.Errorf("LogNodeMetrics, not connect to mysql server")
	}
	return nil
}

// LogPodMetrics is used to insert pod metrics to Mysql
func (ms *MysqlSink) LogPodMetrics() error {
	if !ms.connect {
		return fmt.Errorf("LogPodMetrics, not connect to mysql server")
	}
	return nil
}

func (ms *MysqlSink) initDB() error {
	userName := ms.sink.Parameter["user"]
	password := ms.sink.Parameter["password"]
	ip := ms.sink.Parameter["ip"]
	port := ms.sink.Parameter["port"]
	dbName := ms.sink.Parameter["database"]

	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	DB, err := sql.Open("mysql", path)
	if err != nil {
		return err
	}

	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("open database fail, %+v", err)
	}

	ms.database = DB
	return nil
}

func (ms *MysqlSink) insertStaticNodeMetrics(nm *NodeMetric) error {
	tx, err := ms.database.Begin()
	if err != nil {
		return fmt.Errorf("insertStaticNodeMetrics, tx fail, %+v", err)
	}

	stmt, err := tx.Prepare("INSERT INTO Node_Static_Info (`name`, `capacity_cpu`, `capacity_memory`) VALUES (?, ?, ?)")
	if err != nil {
		return fmt.Errorf("insertStaticNodeMetrics, prepare fail, %+v", err)
	}

	res, err := stmt.Exec(ms.node, nm.Capacity["cpu"], nm.Capacity["memory"])
	if err != nil {
		return fmt.Errorf("insertStaticNodeMetrics, exec fail, %+v", err)
	}

	tx.Commit()

	fmt.Println(res.LastInsertId())
	return nil
}

func (ms *MysqlSink) insertRealNodeMetrics(nm *NodeMetric) error {
	tx, err := ms.database.Begin()
	if err != nil {
		return fmt.Errorf("insertRealNodeMetrics, tx fail, %+v", err)
	}

	stmt, err := tx.Prepare("INSERT INTO Node_Real_Info (`name`, `used_cpu`, `used_memory`, `timestamp`) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("insertRealNodeMetrics, prepare fail, %+v", err)
	}

	res, err := stmt.Exec(ms.node, nm.Capacity["cpu"], nm.Capacity["memory"], nm.SampleTime.Unix())
	if err != nil {
		return fmt.Errorf("insertRealNodeMetrics, exec fail, %+v", err)
	}

	tx.Commit()

	fmt.Println(res.LastInsertId())
	return nil
}
