package main

// Virtualpg - Loadable dynamic extension for SQLite and SpatiaLite
// Homepage: https://www.gaia-gis.it/fossil/virtualpg/index

import (
	"fmt"
	
	"os/exec"
)

func installVirtualpg() {
	// Método 1: Descargar y extraer .tar.gz
	virtualpg_tar_url := "https://www.gaia-gis.it/gaia-sins/virtualpg-2.0.1.tar.gz"
	virtualpg_cmd_tar := exec.Command("curl", "-L", virtualpg_tar_url, "-o", "package.tar.gz")
	err := virtualpg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtualpg_zip_url := "https://www.gaia-gis.it/gaia-sins/virtualpg-2.0.1.zip"
	virtualpg_cmd_zip := exec.Command("curl", "-L", virtualpg_zip_url, "-o", "package.zip")
	err = virtualpg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtualpg_bin_url := "https://www.gaia-gis.it/gaia-sins/virtualpg-2.0.1.bin"
	virtualpg_cmd_bin := exec.Command("curl", "-L", virtualpg_bin_url, "-o", "binary.bin")
	err = virtualpg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtualpg_src_url := "https://www.gaia-gis.it/gaia-sins/virtualpg-2.0.1.src.tar.gz"
	virtualpg_cmd_src := exec.Command("curl", "-L", virtualpg_src_url, "-o", "source.tar.gz")
	err = virtualpg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtualpg_cmd_direct := exec.Command("./binary")
	err = virtualpg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libspatialite")
	exec.Command("latte", "install", "libspatialite").Run()
	fmt.Println("Instalando dependencia: postgis")
	exec.Command("latte", "install", "postgis").Run()
	fmt.Println("Instalando dependencia: postgresql@14")
	exec.Command("latte", "install", "postgresql@14").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
}
