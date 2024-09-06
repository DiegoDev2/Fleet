package main

// Sdb - Ondisk/memory hashtable based on CDB
// Homepage: https://github.com/radareorg/sdb

import (
	"fmt"
	
	"os/exec"
)

func installSdb() {
	// Método 1: Descargar y extraer .tar.gz
	sdb_tar_url := "https://github.com/radareorg/sdb/archive/refs/tags/2.0.1.tar.gz"
	sdb_cmd_tar := exec.Command("curl", "-L", sdb_tar_url, "-o", "package.tar.gz")
	err := sdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdb_zip_url := "https://github.com/radareorg/sdb/archive/refs/tags/2.0.1.zip"
	sdb_cmd_zip := exec.Command("curl", "-L", sdb_zip_url, "-o", "package.zip")
	err = sdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdb_bin_url := "https://github.com/radareorg/sdb/archive/refs/tags/2.0.1.bin"
	sdb_cmd_bin := exec.Command("curl", "-L", sdb_bin_url, "-o", "binary.bin")
	err = sdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdb_src_url := "https://github.com/radareorg/sdb/archive/refs/tags/2.0.1.src.tar.gz"
	sdb_cmd_src := exec.Command("curl", "-L", sdb_src_url, "-o", "source.tar.gz")
	err = sdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdb_cmd_direct := exec.Command("./binary")
	err = sdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
