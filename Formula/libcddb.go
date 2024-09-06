package main

// Libcddb - CDDB server access library
// Homepage: https://libcddb.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibcddb() {
	// Método 1: Descargar y extraer .tar.gz
	libcddb_tar_url := "https://downloads.sourceforge.net/project/libcddb/libcddb/1.3.2/libcddb-1.3.2.tar.bz2"
	libcddb_cmd_tar := exec.Command("curl", "-L", libcddb_tar_url, "-o", "package.tar.gz")
	err := libcddb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcddb_zip_url := "https://downloads.sourceforge.net/project/libcddb/libcddb/1.3.2/libcddb-1.3.2.tar.bz2"
	libcddb_cmd_zip := exec.Command("curl", "-L", libcddb_zip_url, "-o", "package.zip")
	err = libcddb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcddb_bin_url := "https://downloads.sourceforge.net/project/libcddb/libcddb/1.3.2/libcddb-1.3.2.tar.bz2"
	libcddb_cmd_bin := exec.Command("curl", "-L", libcddb_bin_url, "-o", "binary.bin")
	err = libcddb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcddb_src_url := "https://downloads.sourceforge.net/project/libcddb/libcddb/1.3.2/libcddb-1.3.2.tar.bz2"
	libcddb_cmd_src := exec.Command("curl", "-L", libcddb_src_url, "-o", "source.tar.gz")
	err = libcddb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcddb_cmd_direct := exec.Command("./binary")
	err = libcddb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libcdio")
	exec.Command("latte", "install", "libcdio").Run()
}
