package main

// Lmdb - Lightning memory-mapped database: key-value data store
// Homepage: https://www.symas.com/symas-embedded-database-lmdb

import (
	"fmt"
	
	"os/exec"
)

func installLmdb() {
	// Método 1: Descargar y extraer .tar.gz
	lmdb_tar_url := "https://git.openldap.org/openldap/openldap/-/archive/LMDB_0.9.33/openldap-LMDB_0.9.33.tar.bz2"
	lmdb_cmd_tar := exec.Command("curl", "-L", lmdb_tar_url, "-o", "package.tar.gz")
	err := lmdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lmdb_zip_url := "https://git.openldap.org/openldap/openldap/-/archive/LMDB_0.9.33/openldap-LMDB_0.9.33.tar.bz2"
	lmdb_cmd_zip := exec.Command("curl", "-L", lmdb_zip_url, "-o", "package.zip")
	err = lmdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lmdb_bin_url := "https://git.openldap.org/openldap/openldap/-/archive/LMDB_0.9.33/openldap-LMDB_0.9.33.tar.bz2"
	lmdb_cmd_bin := exec.Command("curl", "-L", lmdb_bin_url, "-o", "binary.bin")
	err = lmdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lmdb_src_url := "https://git.openldap.org/openldap/openldap/-/archive/LMDB_0.9.33/openldap-LMDB_0.9.33.tar.bz2"
	lmdb_cmd_src := exec.Command("curl", "-L", lmdb_src_url, "-o", "source.tar.gz")
	err = lmdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lmdb_cmd_direct := exec.Command("./binary")
	err = lmdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
