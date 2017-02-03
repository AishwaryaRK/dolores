package dolores_drives

import dolores_corecode "dolores/corecode"

func GiveSshAccess(parameters ...string) (err error) {
	err = dolores_corecode.Exec("gcloud-ssh-access", parameters[0:]...)
	return
}
