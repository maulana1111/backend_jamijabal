package Admin

import (
	"backend_jamijabal/config"
	"backend_jamijabal/entities"
)

func GetAdmins(username string) entities.Admin {
	res := config.DB.QueryRow(`
				SELECT 
				    admin.id,
				    admin.name,
					admin.username,
					admin.password,
					mst_role_admin.name as role_admin_name,
					admin.photo,
					admin.status
				FROM admin
				JOIN mst_role_admin ON admin.role_admin_id = mst_role_admin.id
				WHERE admin.username = ?
			`, username)

	var admin entities.Admin
	if err := res.Scan(
		&admin.Id,
		&admin.Name,
		&admin.Username,
		&admin.Password,
		&admin.RoleAdmin.Name,
		&admin.Photo,
		&admin.Status,
	); err != nil {
		panic(err.Error())
	}
	return admin
}

func postAdmin(admin entities.Admin) bool {
	res, err := config.DB.Exec(`
					INSERT INTO admin (name, username, password, role_admin_id, photo, status)
					VALUES (?,?,?,?,?,?)
				`, admin.Name, admin.Username, admin.Password, admin.RoleAdmin.Id, admin.Photo, admin.Status)

	if err != nil {
		panic(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastId > 0
}
