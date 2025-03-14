package version

import (
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"runtime"

	"go-admin/cmd/migrate/migration"
	"go-admin/cmd/migrate/migration/models"
	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.Migrate.SetVersion(migration.GetFilename(fileName), _1599190683659Tables)
}

func _1599190683659Tables(db *gorm.DB, version string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if config.DatabaseConfig.Driver == "mysql" {
			tx = tx.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		}
		err := tx.Migrator().AutoMigrate(
			new(models.SysDept),
			new(models.SysConfig),
			new(models.SysTables),
			new(models.SysColumns),
			new(models.SysMenu),
			new(models.SysLoginLog),
			new(models.SysOperaLog),
			new(models.SysRoleDept),
			new(models.SysUser),
			new(models.SysRole),
			new(models.SysPost),
			new(models.DictData),
			new(models.DictType),
			new(models.SysJob),
			new(models.SysConfig),
			new(models.SysApi),
			new(models.TbDemo),
			// RS
			new(models.ChinaData),
			new(models.Business),
			new(models.Tag),
			new(models.BusinessCostCnf),
			new(models.Idc),
			new(models.Host),
			new(models.HostSystem),
			//计费
			new(models.HostIncome),
			new(models.HostIncomeMonth),

			new(models.HostSoftware),
			new(models.HostSwitchLog),
			new(models.HostNetDevice),
			new(models.RsHostSuspendLog),
			new(models.DataBurningHost),
			new(models.Dial),
			new(models.Contract),
			//new(models.HostChargingDay),
			new(models.Custom),
			new(models.CustomUser),
			new(models.BandwidthFees),
			new(models.HostExecLog),
			new(models.OperationLog),

			//资产
			new(models.AssetWarehouse),
			new(models.AssetSupplier),
			new(models.AdditionsOrder),
			new(models.AdditionsWarehousing),
			new(models.Combination),
			new(models.OutboundOrder),
			new(models.AssetRecording),
		)
		if err != nil {
			return err
		}
		if err := models.InitDb(tx); err != nil {
			return err
		}
		return nil
	})
}
