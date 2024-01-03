/* Admin.go 管理员用户模型
------------------------
UserID :用户ID
------------------------
Password: 密码
------------------------
Level: 管理员等级
	0: 允许关闭/重启机器人，允许修改任意配置，能够浏览高级信息
	1: 只能浏览基础信息
------------------------
*/

package ModelAdmin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserID   string `gorm:"varchar(255);not null;unique;"`
	Password string `gorm:"varchar(255);not null;"`
	Level    int    `gorm:"not null;"`
}
