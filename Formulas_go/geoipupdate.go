package main

// Geoipupdate - Automatic updates of GeoIP2 and GeoIP Legacy databases
// Homepage: https://github.com/maxmind/geoipupdate

import (
	"fmt"
	
	"os/exec"
)

func installGeoipupdate() {
	// Método 1: Descargar y extraer .tar.gz
	geoipupdate_tar_url := "https://github.com/maxmind/geoipupdate/archive/refs/tags/v7.0.1.tar.gz"
	geoipupdate_cmd_tar := exec.Command("curl", "-L", geoipupdate_tar_url, "-o", "package.tar.gz")
	err := geoipupdate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geoipupdate_zip_url := "https://github.com/maxmind/geoipupdate/archive/refs/tags/v7.0.1.zip"
	geoipupdate_cmd_zip := exec.Command("curl", "-L", geoipupdate_zip_url, "-o", "package.zip")
	err = geoipupdate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geoipupdate_bin_url := "https://github.com/maxmind/geoipupdate/archive/refs/tags/v7.0.1.bin"
	geoipupdate_cmd_bin := exec.Command("curl", "-L", geoipupdate_bin_url, "-o", "binary.bin")
	err = geoipupdate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geoipupdate_src_url := "https://github.com/maxmind/geoipupdate/archive/refs/tags/v7.0.1.src.tar.gz"
	geoipupdate_cmd_src := exec.Command("curl", "-L", geoipupdate_src_url, "-o", "source.tar.gz")
	err = geoipupdate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geoipupdate_cmd_direct := exec.Command("./binary")
	err = geoipupdate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
}
