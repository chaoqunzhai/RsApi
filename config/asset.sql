BEGIN;

-- -- 1.资产详情表
-- DROP TABLE IF EXISTS `asset`;
-- CREATE TABLE `asset` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `asset_code` varchar(128) NOT NULL COMMENT '资产编号',
--     `sn_code` varchar(255) DEFAULT NULL COMMENT 'SN编码',
--     `category_id` int(11) NOT NULL COMMENT '资产类别',
--     `specification` varchar(100) NULL DEFAULT NULL COMMENT '规格型号',
--     `brand` varchar(100) NULL DEFAULT NULL COMMENT '品牌',
--     `unit` varchar(50) NOT NULL COMMENT '计量单位',
--     `unit_price` decimal(10, 2) NOT NULL COMMENT '单价',
--     `status` tinyint(1) NOT NULL DEFAULT 0 NOT NULL DEFAULT 0 COMMENT '状态(0=在库, 1=出库, 2=在用, 3=处置)',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_deleted_at` (`deleted_at`) USING BTREE,
--     UNIQUE INDEX unique_asset_code (`asset_code`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产详情';


-- -- 2.资产类别表
-- DROP TABLE IF EXISTS `asset_category`;
-- CREATE TABLE `asset_category` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `category_name` varchar(100) NOT NULL COMMENT '类别名称',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_category_deleted_at` (`deleted_at`) USING BTREE,
--     UNIQUE INDEX unique_category_name (`category_name`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产类别';


-- -- 3.资产组合表
-- DROP TABLE IF EXISTS `asset_group`;
-- CREATE TABLE `asset_group` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `group_name` varchar(128) NOT NULL COMMENT '资产组合名称',
--     `main_asset_id` int(11) NOT NULL COMMENT '主资产编码',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_group_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_group_main_asset_id` (`main_asset_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产组合';


-- -- 4.资产组合成员表
-- DROP TABLE IF EXISTS `asset_group_member`;
-- CREATE TABLE `asset_group_member` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `asset_group_id` int(11) NOT NULL COMMENT '资产组合编码',
--     `asset_id` int(11) NOT NULL COMMENT '资产编码',
--     `is_main` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否为主资产(1=是,0=否)',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_group_member_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_group_member_group_id` (`asset_group_id`),
--     KEY `idx_asset_group_member_asset_id` (`asset_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产组合成员';



-- 5.资产入库记录表
DROP TABLE IF EXISTS `asset_inbound`;
CREATE TABLE `asset_inbound` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `inbound_code` varchar(100) NOT NULL COMMENT '入库单号',
    `warehouse_id` int(11) NOT NULL COMMENT '库房编码',
    `inbound_from` tinyint(1) NOT NULL DEFAULT 0 COMMENT '来源(1=采购、0=直接入库)',
    `from_code` varchar(100) NOT NULL COMMENT '来源凭证编码(采购编码)',
    `inbound_by` int(11) NOT NULL COMMENT '入库人编码',
    `inbound_at` timestamp NOT NULL COMMENT '入库时间',
    `remark` text NULL DEFAULT NULL COMMENT '备注',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    KEY `idx_asset_inbound_deleted_at` (`deleted_at`) USING BTREE,
    KEY `idx_asset_inbound_code` (`inbound_code`),
    KEY `idx_asset_inbound_warehouse_id` (`warehouse_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产入库记录';


-- 6.资产入库成员表
DROP TABLE IF EXISTS `asset_inbound_member`;
CREATE TABLE `asset_inbound_member` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `asset_inbound_id` int(11) NOT NULL COMMENT '资产入库编码',
    `asset_inbound_code` varchar(100) NOT NULL COMMENT '资产入库单号',
    `asset_id` int(11) NOT NULL COMMENT '资产编码',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    KEY `idx_asset_group_member_deleted_at` (`deleted_at`) USING BTREE,
    KEY `idx_asset_inbound_id` (`asset_inbound_id`),
    KEY `idx_asset_inbound_code` (`asset_inbound_code`),
    KEY `idx_asset_inbound_member_asset_id` (`asset_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产入库成员表';


-- 7.资产出库记录表
DROP TABLE IF EXISTS `asset_outbound`;
CREATE TABLE `asset_outbound` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `outbound_code` varchar(100) NOT NULL COMMENT '出库单号',
    `warehouse_id` int(11) NOT NULL COMMENT '库房编码',
    `outbound_to` int(11) NOT NULL COMMENT '出库去向(客户编码)',
    `outbound_by` int(11) NOT NULL COMMENT '出库人编码',
    `outbound_at` timestamp NOT NULL COMMENT '出库时间',
    `remark` text NULL DEFAULT NULL COMMENT '备注',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    KEY `idx_asset_outbound_deleted_at` (`deleted_at`) USING BTREE,
    KEY `idx_asset_outbound_code` (`outbound_code`),
    KEY `idx_asset_outbound_warehouse_id` (`warehouse_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产出库记录';


-- 8.资产出库成员表
DROP TABLE IF EXISTS `asset_outbound_member`;
CREATE TABLE `asset_outbound_member` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
    `asset_outbound_id` int(11) NOT NULL COMMENT '资产出库编码',
    `asset_outbound_code` varchar(100) NOT NULL COMMENT '资产出库单号',
    `asset_id` int(11) NOT NULL COMMENT '资产编码',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `create_by` int(11) unsigned DEFAULT NULL,
    `update_by` int(11) unsigned DEFAULT NULL,
    KEY `idx_asset_group_member_deleted_at` (`deleted_at`) USING BTREE,
    KEY `idx_asset_outbound_id` (`asset_outbound_id`),
    KEY `idx_asset_outbound_code` (`asset_outbound_code`),
    KEY `idx_asset_outbound_member_asset_id` (`asset_id`)
) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产出库成员表';


-- -- 9.资产库房表
-- DROP TABLE IF EXISTS `asset_warehouse`;
-- CREATE TABLE `asset_warehouse` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `warehouse_name` VARCHAR(100) NOT NULL COMMENT '库房名称',
--     `administrator_id` int(11) NOT NULL COMMENT '管理员编码',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_warehouse_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_warehouse_administrator_id` (`administrator_id`),
--     UNIQUE INDEX unique_category_name (`warehouse_name`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产库房';


-- -- 10.资产库存表
-- DROP TABLE IF EXISTS `asset_stock`;
-- CREATE TABLE `asset_stock` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `warehouse_id` int(11) NOT NULL COMMENT '库房编码',
--     `category_id` int(11) NOT NULL COMMENT '资产类别编码',
--     `quantity` int(11) NOT NULL COMMENT '资产库存数量',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_stock_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_stock_warehouse_id` (`warehouse_id`),
--     KEY `idx_asset_stock_category_id` (`category_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产库存';


-- -- 11.资产退还记录表
-- DROP TABLE IF EXISTS `asset_return`;
-- CREATE TABLE `asset_return` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `asset_id` int(11) NOT NULL COMMENT '资产编码',
--     `return_person` int(11) NOT NULL COMMENT '退还人编码',
--     `reason` varchar(50) NOT NULL COMMENT '退还原因',
--     `return_at` timestamp NOT NULL COMMENT '退还时间',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_return_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_return_asset_id` (`asset_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产退还记录';


-- -- 12.资产处置记录表
-- DROP TABLE IF EXISTS `asset_disposal`;
-- CREATE TABLE `asset_disposal` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `asset_id` int(11) NOT NULL COMMENT '资产编码',
--     `disposal_person` int(11) NOT NULL COMMENT '处置人编码',
--     `reason` varchar(50) NOT NULL COMMENT '处置原因',
--     `disposal_way` tinyint(1) NOT NULL DEFAULT 0 COMMENT '处置方式(0=报废, 1=出售, 2=出租, 3=退租, 4=捐赠, 5=其它)',
--     `disposal_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '处置地点类型(0=机房, 1=库房)',
--     `location_id` int(11) NOT NULL COMMENT '处置地点编码(机房编码/库房编码)',
--     `amount` decimal(10, 2) NOT NULL COMMENT '处置金额',
--     `disposal_at` timestamp NOT NULL COMMENT '处置时间',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_disposal_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_disposal_asset_id` (`asset_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产处置记录';


-- -- 13.资产采购申请表
-- DROP TABLE IF EXISTS `asset_purchase_apply`;
-- CREATE TABLE `asset_purchase_apply` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `apply_code` varchar(100) NOT NULL COMMENT '申请单编号',
--     `category_id` int(11) NOT NULL COMMENT '资产类型编码',
--     `supplier_id` int(11) NOT NULL COMMENT '供应商编码',
--     `apply_user` int(11) NOT NULL COMMENT '申购人编码',
--     `quantity` int(11) NOT NULL COMMENT '申购数量',
--     `specification` varchar(100) NULL DEFAULT NULL COMMENT '规格型号',
--     `brand` varchar(100) NULL DEFAULT NULL COMMENT '品牌',
--     `unit` varchar(50) NOT NULL COMMENT '计量单位',
--     `unit_price` decimal(10, 2) NOT NULL COMMENT '预估单价',
--     `total_amount` decimal(10, 2) NOT NULL COMMENT '预估金额',
--     `apply_reason` text NOT NULL COMMENT '申购理由',
--     `apply_at` date NOT NULL COMMENT '申购日期',
--     `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '申购状态(0=待审批, 1=已审批, 2=已驳回, 3=已取消)',
--     `approver` int(11) NULL DEFAULT NULL COMMENT '审批人编码',
--     `approve_at` date NULL DEFAULT NULL COMMENT '审批时间',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_purchase_apply_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_purchase_apply_category_id` (`category_id`),
--     KEY `idx_asset_purchase_apply_supplier_id` (`supplier_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产采购申请';


-- -- 14.资产采购记录表
-- DROP TABLE IF EXISTS `asset_purchase`;
-- CREATE TABLE `asset_purchase` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `purchase_code` varchar(100) NOT NULL COMMENT '采购单编号',
--     `category_id` int(11) NOT NULL COMMENT '资产类型编码',
--     `supplier_id` int(11) NOT NULL COMMENT '供应商编码',
--     `purchase_user` int(11) NOT NULL COMMENT '采购人编码',
--     `specification` varchar(100) NULL DEFAULT NULL COMMENT '规格型号',
--     `brand` varchar(100) NULL DEFAULT NULL COMMENT '品牌',
--     `quantity` int(11) NOT NULL COMMENT '采购数量',
--     `unit` varchar(50) NOT NULL COMMENT '计量单位',
--     `unit_price` decimal(10, 2) NOT NULL COMMENT '采购单价',
--     `total_amount` decimal(10, 2) NOT NULL COMMENT '采购金额',
--     `purchase_at` date NOT NULL COMMENT '采购日期',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_purchase_deleted_at` (`deleted_at`) USING BTREE,
--     KEY `idx_asset_purchase_category_id` (`category_id`),
--     KEY `idx_asset_purchase_supplier_id` (`supplier_id`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产采购记录';


-- -- 15.资产供应商表
-- DROP TABLE IF EXISTS `asset_supplier`;
-- CREATE TABLE `asset_supplier` (
--     `id` int(11) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
--     `supplier_name` varchar(100) NOT NULL COMMENT '供应商名称',
--     `contact_person` varchar(100) NOT NULL COMMENT '联系人',
--     `phone_number` varchar(20) NOT NULL COMMENT '联系电话',
--     `email` varchar(100) NOT NULL COMMENT '邮箱',
--     `address` varchar(200) NOT NULL COMMENT '地址',
--     `remark` text NULL DEFAULT NULL COMMENT '备注',
--     `created_at` timestamp NULL DEFAULT NULL,
--     `updated_at` timestamp NULL DEFAULT NULL,
--     `deleted_at` timestamp NULL DEFAULT NULL,
--     `create_by` int(11) unsigned DEFAULT NULL,
--     `update_by` int(11) unsigned DEFAULT NULL,
--     KEY `idx_asset_supplier_deleted_at` (`deleted_at`) USING BTREE,
--     UNIQUE KEY `idx_asset_supplier_supplier_name` (`supplier_name`)
-- ) ENGINE = InnoDB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8mb4 COMMENT = '资产供应商';


COMMIT;
