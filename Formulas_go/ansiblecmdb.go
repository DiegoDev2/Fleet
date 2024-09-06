package main

// AnsibleCmdb - Generates static HTML overview page from Ansible facts
// Homepage: https://github.com/fboender/ansible-cmdb

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleCmdb() {
	// Método 1: Descargar y extraer .tar.gz
	ansiblecmdb_tar_url := "https://github.com/fboender/ansible-cmdb/releases/download/1.31/ansible-cmdb-1.31.tar.gz"
	ansiblecmdb_cmd_tar := exec.Command("curl", "-L", ansiblecmdb_tar_url, "-o", "package.tar.gz")
	err := ansiblecmdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansiblecmdb_zip_url := "https://github.com/fboender/ansible-cmdb/releases/download/1.31/ansible-cmdb-1.31.zip"
	ansiblecmdb_cmd_zip := exec.Command("curl", "-L", ansiblecmdb_zip_url, "-o", "package.zip")
	err = ansiblecmdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansiblecmdb_bin_url := "https://github.com/fboender/ansible-cmdb/releases/download/1.31/ansible-cmdb-1.31.bin"
	ansiblecmdb_cmd_bin := exec.Command("curl", "-L", ansiblecmdb_bin_url, "-o", "binary.bin")
	err = ansiblecmdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansiblecmdb_src_url := "https://github.com/fboender/ansible-cmdb/releases/download/1.31/ansible-cmdb-1.31.src.tar.gz"
	ansiblecmdb_cmd_src := exec.Command("curl", "-L", ansiblecmdb_src_url, "-o", "source.tar.gz")
	err = ansiblecmdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansiblecmdb_cmd_direct := exec.Command("./binary")
	err = ansiblecmdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
}
